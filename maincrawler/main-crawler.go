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
	// url := flag.String("u", "https://www.amazon.com", "URL for website to analyze")
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

	// --- TESTING COOKIES WITH MULTIPLE URL's AND TESTING SAFE AND LESS SAFE URL's ---

	// Declare structure for privacy metrics
	safePrivacyMetric := crawler.PrivacyMetric{}

	// Fetch cookies from amazon
	cookie1 := crawler.FetchCookies(*browser, *isHidden, "https://www.amazon.com", &safePrivacyMetric, verbose, 20000)

	// Print cookies from amazon
	crawler.PrintCookies(cookie1, "https://www.amazon.com", verbose)

	// Print metrics from amazon
	crawler.PrintMetrics(safePrivacyMetric, "Amazon Cookies")

	// Analyze the Privacy Metrics in amazon
	analysis := crawler.AnalyzeMetrics(safePrivacyMetric)

	// Generate a report for the Privacy Metrics in amazon
	report := crawler.CreateReport(analysis)

	// Print out the report gathered from amazon
	fmt.Println(report)
}
