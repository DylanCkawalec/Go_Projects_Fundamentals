package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Yo bro it's me your hommie")

	html := `<strong>This is the body</strong>`
	w.Header().Set("Content-Type", "text/tml")
	fmt.Fprint(w, html)

}

func main() {

	http.HandleFunc("/", homePage)
	log.Println("Serving to the local Host")
	//
	http.ListenAndServe(":8080", nil)

	//

}
