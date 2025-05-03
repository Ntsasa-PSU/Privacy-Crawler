package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
)

// ---- DATA STRUCTURES ---- //

// URLList: Represents the structure of JSON file.
// Tag 'json:"urls"' envokes JSON parser to  map "urls" field
// from JSON to URLS slice.
type URLList struct {
	URLs []string `json:"urls"`
}

// Result: Represent the outcome from fetching URL.
type Result struct {
	URL      string
	Status   int
	Size     int
	Duration time.Duration
	Error    error
}

// CookieList: Represents the List of cookies collected.
// Uses type Cookie for each entry.
type CookiesList struct {
	List map[string][]Cookie
}

// Cookie: Represents the privacy characteristics of the collected cookies
type Cookie struct {
	Name         string
	Value        string
	Domain       string
	Path         string
	Expires      float64
	HttpOnly     bool
	Secure       bool
	SameSite     string
	IsFirstParty bool
}

// Privacy Metric: Represents the privacy fields to consider
type PrivacyMetric struct {
	TotalCookies    int
	TotalFirstParty int
	TotalThirdParty int
	TotalSecure     int
	TotalNotSecure  int
}

// ---- Global Definitions ---- //

// URL File: Location fo pre-configed JSON file.
const URLFILE string = "internal/config/urls.json"

// ---- Functions ---- //

// Function: Read JSON
// Operation: This will read a pre-configured JSON file with URLs.
// Return: *URL List
func ReadJSON(verbose *bool) *URLList {

	// Read JSON file into data variable.
	data, err := os.ReadFile(URLFILE)

	// Check error from JSON file.
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return nil
	}

	// Create URL data structure.
	var urlList URLList

	// Parse JSON into data structure.
	err = json.Unmarshal(data, &urlList)

	if err != nil {
		fmt.Printf("Error parsing JSON file: %v\n", err)
		return nil
	}

	// - Verbose Output - //
	if *verbose {

		fmt.Println("\n-- Parsed JSON --")

		for i, url := range urlList.URLs {
			fmt.Printf("%d. %s\n", i+1, url)
		}
	}

	return &urlList
}

// Function: Get Browser
// Operation: Returns map of browers and agents for them.
// Return: map[string]string
func GetBrowsers(verbose *bool) map[string]string {

	browsers := map[string]string{
		"chrome":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
		"firefox": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/112.0",
		"safari":  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Safari/605.1.15",
		"edge":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.48",
		"bot":     "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"curl":    "curl/7.84.0",
	}

	// - Verbose Output - //
	if *verbose {

		fmt.Println("\n-- Browsers/Agent --")

		index := 1
		for browserName, userAgent := range browsers {
			fmt.Printf("[%s]: %s\n", browserName, userAgent)
			index++
		}
	}

	return browsers
}

// Fucntion: Verify Browser
// Operation: Verify we have an agent for selected browser.
// Return: String(Browser found), String (User-Agent), Error
func VerifyTargetBrowser(browsers map[string]string, selectedBrowser string, verbose *bool) (string, string, error) {

	// - Verbose Output - //
	if *verbose {

		fmt.Println("\n-- Verifying Browser Selection -- ")
	}

	// Iterate through map of browsers.
	for browserName, browserUserAgent := range browsers {

		if browserName == selectedBrowser {

			// - Verbose Output - //
			if *verbose {

				fmt.Printf("Found: %s\n", browserName)
			}

			return browserName, browserUserAgent, nil
		}
	}

	return "None", "None", fmt.Errorf("browser not found")
}

// Fucntion: Fetch URL
// Operation: Connect and returns header from selected URL.
// Return: Header (map), Body (map), Status Code (int), Error
func FetchPacket(url string, userAgent string, verbose *bool) (map[string][]string, []byte, int, error) {

	// Create a new HTTP client.
	client := &http.Client{
		// Timeout if cannot connect.
		Timeout: time.Second * 10,
	}

	// Create a new request.
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("error creating request: %v", err)
	}

	// Set the User-Agent header based on browser.
	request.Header.Set("User-Agent", userAgent)

	// Make the request.
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("error making request: %v", err)
	}

	// Cleanup on response.
	defer response.Body.Close()

	// Read the body.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("error reading response body: %v", err)
	}

	// Cleanup on body.
	response.Body.Close()

	// - Verbose output - //
	if *verbose {
		fmt.Printf("\n-- HTTP Packet: Header for %s --\n", url)
		fmt.Printf("Status: %s\n", response.Status)

		fmt.Printf("\n-- HTTP Packet: Body for %s --\n\n", url)
		fmt.Printf("Body: %s\n", string(body))
	}

	// Return the headers and status code.
	return response.Header, body, response.StatusCode, nil
}

// Function: Fetch Cookies
// Operation: Connect to url with browser, creates cookies using playwright
// to fully generate all cookies due to Javascript delys.
// Return: A list of cookies collected and stored in a struct (*CookiesList)
func FetchCookies(browser string, isHidden bool, url string, privacyMetrics *PrivacyMetric, verbose *bool) *CookiesList {
	// - Run Playwright - //
	pw, err := playwright.Run()
	if err != nil {
		fmt.Printf("could not lauch playwright: %v", err)
		return nil
	}
	defer pw.Stop()

	// Declare launcher ahead of time
	var launcher playwright.Browser

	// launching specific browser (edge is missing)
	if browser == "chrome" {
		launcher, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(isHidden),
		})
	} else if browser == "firefox" {
		launcher, err = pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(isHidden),
		})
	} else if browser == "safari" {
		launcher, err = pw.WebKit.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(isHidden),
		})
	} else if browser == "edge" {
		launcher, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Channel:  playwright.String("msedge"),
			Headless: playwright.Bool(isHidden),
		})
	} else {
		fmt.Printf("Browser %s is not compativle\n", browser)
		return nil
	}
	if err != nil {
		fmt.Printf("could not launch browser: %s\n", browser)
		return nil
	}

	if *verbose {
		fmt.Printf("Launching %s with %s...\n", browser, url)
		if isHidden {
			fmt.Println("\t- Browser is hidden.")
		} else {
			fmt.Println("\t- Browser is visible.")
		}
	}

	// Create the context for the broswer
	context, err := launcher.NewContext() // creates an isolated browsers contents
	if err != nil {
		fmt.Printf("could not create context: %v", err)
		return nil
	}

	// Open up a new tab from the context
	page, err := context.NewPage() // Add a new Tab
	if err != nil {
		fmt.Printf("could not create a new Tab: %v", err)
		return nil
	}

	// Navigate to the desired URL
	_, err = page.Goto(url)
	if err != nil {
		fmt.Printf("could not go to url page: %v", err)
	}

	// Wait for a few seconds for JS to run and set cookies
	page.WaitForTimeout(5000)

	// Get cookies
	cookies, err := context.Cookies()
	if err != nil {
		fmt.Printf("could not get cookies: %v", err)
		return nil
	}
	if len(cookies) == 0 {
		fmt.Println("No cookies were returned.")
		return nil
	}

	// Store cookies in a struc, organized
	collectedCookies := CookiesList{
		List: make(map[string][]Cookie),
	}

	collectedCount := 0
	for _, c := range cookies {
		// Check for first or third part
		isFirst := isFirstParty(c.Domain, url) // check if url and cookie Domain are the same

		// Store fields inside of a cookie
		cookie := Cookie{
			Name:         c.Name,
			Value:        c.Value,
			Domain:       c.Domain,
			Path:         c.Path,
			Expires:      float64(c.Expires),
			HttpOnly:     c.HttpOnly,
			Secure:       c.Secure,
			SameSite:     string(*c.SameSite),
			IsFirstParty: isFirst,
		}
		// Map to Cookie Key
		collectedCookies.List[c.Domain] = append(collectedCookies.List[c.Domain], cookie)
		collectedCount++

		addToMetric(privacyMetrics, cookie, url)

	}

	if *verbose {
		fmt.Printf("*** Cookies Collected: %d ***\n", collectedCount)
	}

	return &collectedCookies
}

// Function: Add to Privacy Metric
// Operation: Adds specific data from cookies to use for later analysis
// Return: None
func addToMetric(privacyMetrics *PrivacyMetric, cookie Cookie, url string) {
	// Add 1 to cookie total
	privacyMetrics.TotalCookies++

	// Check for cookie party
	if isFirstParty(cookie.Domain, url) {
		privacyMetrics.TotalFirstParty++
	} else {
		privacyMetrics.TotalThirdParty++
	}

	// Check for Security
	if cookie.Secure {
		privacyMetrics.TotalSecure++
	} else {
		privacyMetrics.TotalNotSecure++
	}
}

// Function: Print Privacy Metrics
// Operation: Prints the privacy metrics in a readable formate
// Return: None
func PrintMetrics(privacyMetrics PrivacyMetric, metricName string) {
	fmt.Printf("#----- Printing %s Privacy Metrics ------#\n", metricName)
	fmt.Printf("Total Cookies: %d\n", privacyMetrics.TotalCookies)
	fmt.Printf("Total First-Party Cookies: %d\n", privacyMetrics.TotalFirstParty)
	fmt.Printf("Total Third-Party Cookies: %d\n", privacyMetrics.TotalThirdParty)
	fmt.Printf("Total Secure Domains: %d\n", privacyMetrics.TotalSecure)
	fmt.Printf("Total Unsecure Domains: %d\n", privacyMetrics.TotalNotSecure)
	fmt.Println("#--------------------------------------------#")
}

// Function: Print Cookies
// Operation: Print the cookies in a readable format
// Return: None
func PrintCookies(cookieList *CookiesList, url string, verbose *bool) {
	if cookieList == nil || len(cookieList.List) == 0 {
		fmt.Println("No cookies to display.")
		return
	}

	fmt.Println("\n-- Cookies Displayed --")

	cookieCount := 0
	domainCount := 1
	for domain, cookies := range cookieList.List {
		fmt.Printf("Domain %d: %s\n", domainCount, domain)

		for i, cookie := range cookies {
			partyType := "third-party"
			if cookie.IsFirstParty {
				partyType = "first-party"
			}
			fmt.Printf("\tCookie %d.[%s] \n", i+1, partyType)

			fmt.Printf("\t\tFrom Page: %s\n", url)
			fmt.Printf("\t\tDomain: %s\n", cookie.Domain)
			fmt.Printf("\t\tName: %s\n", cookie.Name)
			fmt.Printf("\t\tValue: %s\n", cookie.Value)
			fmt.Printf("\t\tPath: %s\n", cookie.Path)

			// Check if Expires is greater than 0 to avoid negative time
			// Expires is a float64, so we convert it to int64 for Unix time
			if cookie.Expires > 0 {
				expTime := time.Unix(int64(cookie.Expires), 0)
				fmt.Printf("\t\tExpires: %s\n", expTime.Format("Mon, 02 Jan 2006"))
			} else {
				fmt.Println("\t\tExpires: No Expiration")
			}

			fmt.Printf("\t\tHttpOnly: %t\n", cookie.HttpOnly)
			fmt.Printf("\t\tSecure: %t\n", cookie.Secure)
			fmt.Printf("\t\tSameSite: %s\n", cookie.SameSite)
			cookieCount++
		}
		domainCount++
		fmt.Println("#-------------------------------------#")
	}

	if *verbose {
		fmt.Printf("Total Cookies Collected: %d\n", cookieCount)
	}
}

// Function: isFirstParty
// Operation: Check if the cookie domain is first-party or third-party
// Return: True if first-party, false if third-party
func isFirstParty(cookieDomain, fullURL string) bool {
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		fmt.Printf("could not parse URL: %v", err)
		return false
	}

	topLevelDomain := parsedURL.Host                           // Get the domain from the URL (e.g., "www.amazon.com")
	cleanCookieDomain := strings.TrimPrefix(cookieDomain, ".") // TrimPrefix: removes "." from the start of cookieDomain

	// HasPrefix: checks if the cleanCookieDomain STARTS with the topLevelDomain.
	isPrefix := strings.HasPrefix(cleanCookieDomain, topLevelDomain)
	// HasSuffix: checks if the topLevelDomain ENDS with the cleanCookieDomain.
	isSuffix := strings.HasSuffix(topLevelDomain, cleanCookieDomain)

	return isSuffix || isPrefix
}

// Make PrivacyMetric Struct
// Make function to import packet data into struct.

//Then can make application do all borwsers+ urls
// Host on server
// Can figure out where to host

//Once data gathered, analyze it? and make a report

// Next Step 2: Create Server Go agents can connect to
//    1. Create a server that can analyze the connection
//    2. Create a server that can analyze the HTTP headers
//    3. Process the HTTP headers data
//    4. Return the HTTP headers analysis results in JSON format
