package mylogger

import "log"

func ListenForLog(ch chan string) {

	for {
		msg := <-ch //the message is assigned to whatever is in the channel,
		//which is the input...
		log.Println(msg) //then it prints the message
	}

}
