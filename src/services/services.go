package services

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpont Hit")
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatalln(http.ListenAndServe(":8081", nil))
}
