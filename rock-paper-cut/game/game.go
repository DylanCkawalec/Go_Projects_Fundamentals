package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	//use select to process inout in channels
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)

			g.DisplayChan <- ""

		}

	}

}

func (g *Game) ClearScreen() { //this is cool
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {

	g.DisplayChan <- fmt.Sprintf(`
	Rock, Paper & Scissors
	----------------------
	Game is played for three rounds, and best of three wins the game, PLAY!
	
	`)
	<-g.DisplayChan
}

func (g *Game) PlayerRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	g.DisplayChan <- fmt.Sprintf(`
	Round %d
	--------
	`, g.Round.RoundNumber)
	<-g.DisplayChan

	fmt.Print("PLease enter rock, paper, or scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer Chose Rock"
		<-g.DisplayChan
	case PAPER:
		g.DisplayChan <- "Computer Chose Paper"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer Chose Scissors"
		<-g.DisplayChan
	default:
	}
	//-----------------------------------------------------------------------------------
	g.DisplayChan <- ""
	<-g.DisplayChan

	if playerValue == computerValue {

		g.DisplayChan <- "It's a draw!"
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break

		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break

		default:
			g.DisplayChan <- "Invalid Choice"
			<-g.DisplayChan
			return false
		}
	}

	return true

}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer Wins!"
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player Wins!"
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	fmt.Println("Final Score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	fmt.Println()
	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player Wins the game!"
		<-g.DisplayChan
	} else {
		g.DisplayChan <- "Computer Wins The Game :( !!!"
		<-g.DisplayChan
	}
}
