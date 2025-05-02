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
	isHidden := flag.Bool("i", false, "Hides browser")
	url := flag.String("u", "https://www.amazon.com", "URL for website to analyze")
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
	crawler.VerifyTargetBrowser(browserList, *browser, verbose)

	// Fetch Cookies
	cookies := crawler.FetchCookies(*browser, *isHidden, *url, verbose)
	crawler.PrintCookies(cookies, *url, verbose)
}
