package services

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpont Hit")
}

func HandleRequests() {
	var databasePass string
	databasePass = os.Getenv("TEST_ENV_VAR")
	fmt.Printf("Database Password: %s\n", databasePass)
	http.HandleFunc("/", homePage)
	log.Fatalln(http.ListenAndServe(":8081", nil))
}
