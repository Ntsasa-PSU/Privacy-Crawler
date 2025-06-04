package main

import (
	"privcrawler/internal/jmppoint"
)

func main() {
	//jmppoint.RunServer()
	jmppoint.GenerateTotalsFile()
	jmppoint.BrowserRanking()
}

