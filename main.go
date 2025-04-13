package main

import (
	"flag"
	"privcrawler/internal/crawler"
)


func main() {

	// -- COMMNAD LINE ARGUMENTS -- //
	verbose := flag.Bool("v", false, "Enable verbose output.")
	browser := flag.String("b", "chrome", "Other browser option selected.")
	
	// Parse command line flags
	flag.Parse()
	
	browserList := crawler.GetBrowsers(verbose)
	crawler.VerifyBrowser(browserList, *browser, verbose)

	//targetURLs := crawler.ReadJSON(verbose)

}