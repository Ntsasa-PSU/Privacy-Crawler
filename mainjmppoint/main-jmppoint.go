package main

import (
	"fmt"
	"log"
	"net/http"
	"privcrawler/internal/jmppoint"
)

func main() {
	fmt.Println("Listening on : ",  jmppoint.PORT)
	log.Fatal(http.ListenAndServe(":22", nil))
}

