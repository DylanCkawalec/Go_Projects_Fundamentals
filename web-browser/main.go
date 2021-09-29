package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

//render a homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")

}
func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/play", playRound)
	//first thing to do, is start a web server
	http.HandleFunc("/", homePage) //this is requesting the top level of the application

	log.Println("Starting web server on port 8080")

	http.ListenAndServe(":8080", nil) //listens to a port, a server, production ready web server.
}

// just make a function and then pass the same html RENDER whenever I need to.

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
