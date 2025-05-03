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

	// Declare structure for safe sites
	safePrivacyMetric := crawler.PrivacyMetric{}

	// Fetch Cookies With Login Feature
	cookies1 := crawler.FetchCookies(*browser, *isHidden, "https://www.amazon.com", &safePrivacyMetric, verbose)
	cookies2 := crawler.FetchCookies(*browser, *isHidden, "https://www.youtube.com", &safePrivacyMetric, verbose)
	cookies3 := crawler.FetchCookies(*browser, *isHidden, "https://www.instagram.com", &safePrivacyMetric, verbose)
	crawler.PrintCookies(cookies1, "https://www.amazon.com", verbose)
	crawler.PrintCookies(cookies2, "https://www.youtube.com", verbose)
	crawler.PrintCookies(cookies3, "https://www.instagram.com", verbose)

	// Declare structure for safe sites
	sketchyPrivacyMetric := crawler.PrivacyMetric{}

	// Fetch Cookies with Analytics/Ads
	cookies4 := crawler.FetchCookies(*browser, *isHidden, "https://freerobux.en.uptodown.com/android", &sketchyPrivacyMetric, verbose)
	cookies5 := crawler.FetchCookies(*browser, *isHidden, "https://moviekeen.com/soap-2day/", &sketchyPrivacyMetric, verbose)
	crawler.PrintCookies(cookies4, "https://freerobux.en.uptodown.com/android", verbose)
	crawler.PrintCookies(cookies5, "https://moviekeen.com/soap-2day/", verbose)

	// Print Metrics
	crawler.PrintMetrics(safePrivacyMetric, "Secure")
	crawler.PrintMetrics(sketchyPrivacyMetric, "Less Secure")
}
