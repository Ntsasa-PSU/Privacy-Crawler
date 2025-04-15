package main

import (
	"fmt"
	"log"
	"net/http"
	"privcrawler/internal/jmppoint"
)

func main() {
	http.HandleFunc("/test", jmppoint.UserInputHandler)
	http.HandleFunc("/run", jmppoint.GeneralHandler)
	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

