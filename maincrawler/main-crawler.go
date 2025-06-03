package main

import (
	"flag"
	"fmt"
	"privcrawler/internal/crawler"
)

func main() {

	// -- COMMNAD LINE ARGUMENTS -- //
	verbose := flag.Bool("v", false, "Enable verbose output.")
	browser := flag.String("b", "chrome", "Other browser option selected.")
	isHidden := flag.Bool("i", false, "Hides browser")
	url := flag.String("u", "https://www.amazon.com", "URL for website to analyze")
	duration := flag.Int("d", 2000, "Duration for the browser to run in milliseconds (default: 20000)")


	// Parse command line flags
	flag.Parse()

	browserList := crawler.GetBrowsers(verbose)
	crawler.VerifyTargetBrowser(browserList, *browser, verbose)

	// --- TESTING COOKIES WITH MULTIPLE URL's AND TESTING SAFE AND LESS SAFE URL's ---

	// Declare structure for privacy metrics
	safePrivacyMetric := crawler.PrivacyMetric{}

	// Fetch cookies from amazon
	cookie1 := crawler.FetchCookies(*browser, *isHidden, *url, &safePrivacyMetric, verbose, *duration)

	// Print cookies from amazon
	crawler.PrintCookies(cookie1, *url, verbose)

	urltarget := *url + ": Cookies"
	// Print metrics from amazon
	data := crawler.GetMetricsReport(safePrivacyMetric, urltarget)


	err := crawler.AppendDataToFile(data, *url, *browser)
	if err != nil {
		fmt.Printf("Error appending report to file: %v\n", err)
	}

	// Print out the report gathered 
	fmt.Println(data)
}
