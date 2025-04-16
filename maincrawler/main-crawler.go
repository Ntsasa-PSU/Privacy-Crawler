package main

import (
	"flag"
	"fmt"
	"os"
	"privcrawler/internal/crawler"
)

func main() {

	// -- COMMNAD LINE ARGUMENTS -- //
	verbose := flag.Bool("v", false, "Enable verbose output.")
	browser := flag.String("b", "chrome", "Other browser option selected.")
	test := flag.Bool("t", false, "Prototype test flag.")

	if *test {
		// Get input from environment variables
		name := os.Getenv("USER_NAME")
		age := os.Getenv("USER_AGE")

		fmt.Printf("Your name is %s and you are %s years old.\n", name, age)
	}

	// Parse command line flags
	flag.Parse()

	browserList := crawler.GetBrowsers(verbose)
	_, userAgent, err := crawler.VerifyTargetBrowser(browserList, *browser, verbose)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	crawler.ReadJSON(verbose)

	//Website currently returns user-agent of user.
	crawler.FetchPacket("https://google.com", userAgent, verbose)

}
