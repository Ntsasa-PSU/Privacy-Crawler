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
	TotalCookies int

	TotalFirstParty int
	TotalThirdParty int

	TotalSecure    int
	TotalNotSecure int

	SuspiciousPaths []Cookie // [SuspiciousCookieName {..., ..., Path, ...}]

	TotalHttpOnly    int
	TotalNotHttpOnly int

	SameSiteStrict int
	SameSiteNone   int
	SameSiteLax    int
	SameSiteUnset  int

	TotalSessionCookies    int
	TotalPersistentCookies int
}

// Possible additions to PrivacyMetric
// - track cookie paths (if suspicious, like not "/" path)
// - HttpOnly
// - SameSite (to see how strict(Strict, Lax, None) passing to third-party sites is like)
// - Expiration date? (maybe we can tell if it is a session cookie meaning it expires when you close the tab)

// ---- Global Definitions ---- //

// URL File: Location fo pre-configed JSON file.
const URLFILE string = "internal/config/urls.json"

// Security Ratio of cookie metrics
// These values are the percentage thresholds used to determine the minimum security level
// There are no constant value that makes these fields considered true. but these fields are
// stepping stones to determining the security of a website based on the cookies collected.
const SECURE_THRESHOLD float64 = 80.0         // 80%
const FIRST_THRESHOLD float64 = 50.0          // 50%
const THIRD_THRESHOLD float64 = 50.0          // 50%
const HTTP_THRESHOLD float64 = 70.0           // 70%
const SAMESITESTRICT_THRESHOLD float64 = 50.0 // 50%
const SAMESITELAX_THRESHOLD float64 = 50.0    // 50%
const SAMESITENONE_THRESHOLD float64 = 50.0   // 50%
const SAMESITEUNSET_THRESHOLD float64 = 50.0  // 50%
const SESSION_THRESHOLD float64 = 50.0        // 50%
const PERSISTENT_THRESHOLD float64 = 50.0     // 50%

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
func FetchCookies(browser string, isHidden bool, url string, privacyMetrics *PrivacyMetric, verbose *bool, duration int) *CookiesList {

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
	page.WaitForTimeout(float64(duration))

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

	// Check for Path
	if cookie.Path != "/" {
		privacyMetrics.SuspiciousPaths = append(privacyMetrics.SuspiciousPaths, cookie)
	}

	// Check for HttpOnly
	if cookie.HttpOnly {
		privacyMetrics.TotalHttpOnly++
	} else {
		privacyMetrics.TotalNotHttpOnly++
	}

	// Check for SameSite
	if cookie.SameSite == "Strict" {
		privacyMetrics.SameSiteStrict++
	} else if cookie.SameSite == "Lax" {
		privacyMetrics.SameSiteLax++
	} else {
		privacyMetrics.SameSiteNone++
	}

	// Check for Session/Persistent
	if cookie.Expires < 0 {
		privacyMetrics.TotalSessionCookies++
	} else {
		privacyMetrics.TotalPersistentCookies++
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

	if len(privacyMetrics.SuspiciousPaths) > 0 {
		fmt.Println("All Suspicious Paths")
		for i := 0; i < len(privacyMetrics.SuspiciousPaths); i++ {
			cookie := privacyMetrics.SuspiciousPaths[i]

			partyType := "third-party"
			if cookie.IsFirstParty {
				partyType = "first-party"
			}
			fmt.Printf("\tCookie %d. [%s]\n", i, partyType)
			fmt.Printf("\t\tDomain: %s\n", cookie.Domain)
			fmt.Printf("\t\tName: %s\n", cookie.Name)
			fmt.Printf("\t\tPath: %s\n", cookie.Path)
		}
	} else {
		fmt.Println("No Suspicious Paths")
	}

	fmt.Printf("Total HttpOnly: %d\n", privacyMetrics.TotalHttpOnly)
	fmt.Printf("Total Not HttpOnly: %d\n", privacyMetrics.TotalNotHttpOnly)

	fmt.Printf("Total SameSite with Strict: %d\n", privacyMetrics.SameSiteStrict)
	fmt.Printf("Total SameSite with Lax: %d\n", privacyMetrics.SameSiteLax)
	fmt.Printf("Total SameSite with None: %d\n", privacyMetrics.SameSiteNone)
	fmt.Printf("Total SameSite with Unset: %d\n", privacyMetrics.SameSiteUnset)

	fmt.Printf("Total Session Cookies: %d\n", privacyMetrics.TotalSessionCookies)
	fmt.Printf("Total Persistent Cookies: %d\n", privacyMetrics.TotalPersistentCookies)

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
				fmt.Printf("\t\tExpires: %s\n", expTime.Format("Jan/02/2006 15:04:05"))
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

// Function: Analyze Metrics
// Operation: analyzes the privacy metrics collected from the privacyMetric struct
// Return: A map of analysis results
func AnalyzeMetrics(privacyMetric PrivacyMetric) map[string]float64 {
	analysis := map[string]float64{
		"secure":            0.0,
		"firstParty":        0.0,
		"thirdParty":        0.0,
		"httpOnly":          0.0,
		"sameSiteStrict":    0.0,
		"sameSiteLax":       0.0,
		"sameSiteNone":      0.0,
		"sameSiteUnset":     0.0,
		"sessionCookies":    0.0,
		"persistentCookies": 0.0,
	}

	// --- Calculate Ratios of privacy metrics --- //
	if privacyMetric.TotalCookies > 0 {
		// Secure
		if privacyMetric.TotalSecure > 0 {
			analysis["secure"] = (float64(privacyMetric.TotalSecure) / float64(privacyMetric.TotalCookies)) * 100
		}

		// Party
		if privacyMetric.TotalFirstParty > 0 {
			analysis["firstParty"] = (float64(privacyMetric.TotalFirstParty) / float64(privacyMetric.TotalCookies)) * 100
		}
		if privacyMetric.TotalThirdParty > 0 {
			analysis["thirdParty"] = (float64(privacyMetric.TotalThirdParty) / float64(privacyMetric.TotalCookies)) * 100
		}

		// HttpOnly
		if privacyMetric.TotalHttpOnly > 0 {
			analysis["httpOnly"] = (float64(privacyMetric.TotalHttpOnly) / float64(privacyMetric.TotalCookies)) * 100
		}

		// SameSite
		if privacyMetric.SameSiteNone > 0 {
			analysis["sameSiteStrict"] = (float64(privacyMetric.SameSiteStrict) / float64(privacyMetric.TotalCookies)) * 100
		}
		if privacyMetric.SameSiteLax > 0 {
			analysis["sameSiteLax"] = (float64(privacyMetric.SameSiteLax) / float64(privacyMetric.TotalCookies)) * 100
		}
		if privacyMetric.SameSiteNone > 0 {
			analysis["sameSiteNone"] = (float64(privacyMetric.SameSiteNone) / float64(privacyMetric.TotalCookies)) * 100
		}
		if privacyMetric.SameSiteUnset > 0 {
			analysis["sameSiteUnset"] = (float64(privacyMetric.SameSiteUnset) / float64(privacyMetric.TotalCookies)) * 100
		}

		// Sessions
		if privacyMetric.TotalSessionCookies > 0 {
			analysis["sessionCookies"] = (float64(privacyMetric.TotalSessionCookies) / float64(privacyMetric.TotalCookies)) * 100
		}
		if privacyMetric.TotalPersistentCookies > 0 {
			analysis["persistentCookies"] = (float64(privacyMetric.TotalPersistentCookies) / float64(privacyMetric.TotalCookies)) * 100
		}
	} else {
		fmt.Println("No cookies collected, cannot analyze metrics.")
		return nil
	}

	return analysis
}

// Make constants of ranges for the report to use
// What defines a secure website with ratio of secure cookies?
// how many is to many third-party cookies?

// Function: Create Report
// Operation: Creates a report based on the analysis of the privacy metrics.
// Return: A string which contains the report summarizing the analysis
func CreateReport(analysis map[string]float64) string {

	var report string

	// --- Create report based on analysis --- //
	// It may be better if we used match cases for these reports?

	// ### SECURE METRIC ###
	// is it secure?
	if analysis["secure"] >= SECURE_THRESHOLD {
		report += fmt.Sprintf("The website uses the Secure flag in %.2f%% of its cookies, which "+
			"helps prevent man-in-the-middle attacks by ensuring cookies are only sent over HTTPS. ",
			analysis["secure"])
	} else {
		report += fmt.Sprintf("The website lacks sufficient use of the Secure flag with only %.2f%% "+
			" of cookies being secure, meaning some cookies may be sent over unencrypted HTTP, "+
			"increasing the risk of interception. ",
			analysis["secure"])
	}

	report += "\n\n"

	// ### PARTY METRIC ###
	// What about first-party? to many? to few? enough? does it make sense?
	if analysis["firstParty"] > 0 || analysis["thirdParty"] > 0 {

		if analysis["firstParty"] > analysis["thirdParty"] {
			report += fmt.Sprintf("There are more first-party cookies than third-party cookies (%.2f%% "+
				"first-party), which is typically a good sign as it indicates the site is primarily "+
				"using its own cookies for functionality and personalization. ",
				analysis["firstParty"])
		} else {
			report += fmt.Sprintf("There are more third-party cookies than first-party cookies (%.2f%% "+
				"third-party), which may indicate that the site relies heavily on external services for "+
				"functionality, advertising, or analytics. ",
				analysis["thirdParty"])
		}
	} else {
		report += "There was an issue in detecting the sites first-party cookies and third-party " +
			"cookies. "
	}

	report += "Below are the details of the first-party and third-party cookies:\n"
	report += fmt.Sprintf("\t- First-Party: %.2f%%\n", analysis["firstParty"])
	report += fmt.Sprintf("\t- Third-Party: %.2f%%\n", analysis["thirdParty"])
	report += "\n"

	// ### HTTPONLY METRIC ###
	// is the ratio of HttpOnly valid?
	if analysis["httpOnly"] >= HTTP_THRESHOLD {
		report += fmt.Sprintf("%.2f%% of cookies use the HttpOnly flag, which helps protect session "+
			"data from client-side scripts and reduces the risk of XSS attacks. ",
			analysis["httpOnly"])
		if analysis["secure"] < SECURE_THRESHOLD {
			report += "However, the lack of Secure cookies limits the overall protection HttpOnly " +
				"provides. "
		}
	} else if analysis["httpOnly"] == 0.0 {
		report += "No HttpOnly cookies were found, which is a strong indicator of vulnerability to " +
			"cross-site scripting (XSS) attacks. "
	} else {
		report += fmt.Sprintf("Only %.2f%% of cookies use the HttpOnly flag, which is below the recommended "+
			"threshold. This may leave some session data vulnerable to client-side scripts. ",
			analysis["httpOnly"])
	}
	report += "Below are the details of the HttpOnly attributes:\n"
	report += fmt.Sprintf("\t- HttpOnly: %.2f%%\n", analysis["httpOnly"])
	report += fmt.Sprintf("\t- Not HttpOnly: %.2f%%\n", analysis["httpOnly"])
	report += "\n"

	// ### SAMESITE METRIC ###
	if analysis["sameSiteStrict"] >= SAMESITESTRICT_THRESHOLD {
		report += fmt.Sprintf("Most cookies use SameSite=Strict (%.2f%% Strict), which prevents them from "+
			"being sent in cross-site requests. This is good for defending against CSRF attacks, though it "+
			"may reduce compatibility with some cross-site features. ",
			analysis["sameSiteStrict"])
	}
	if analysis["sameSiteLax"] >= SAMESITELAX_THRESHOLD {
		report += fmt.Sprintf("Most cookies use SameSite=Lax (%.2f%% Lax), a balanced setting that "+
			"permits top-level navigation (e.g., links) while still protecting against most CSRF attacks. ",
			analysis["sameSiteLax"])
	}
	if analysis["sameSiteNone"] >= SAMESITENONE_THRESHOLD {
		report += fmt.Sprintf("Most cookies use SameSite=None (%.2f%% None), allowing them to be sent "+
			"with all cross-site requests. This setting is commonly required for third-party cookies but "+
			"must be paired with Secure to reduce risk. ",
			analysis["sameSiteNone"])
	}
	if analysis["sameSiteUnset"] >= SAMESITEUNSET_THRESHOLD {
		report += fmt.Sprintf("Most cookies have no SameSite attribute set (%.2f%% Unset), which could leave them "+
			"vulnerable to CSRF or privacy leaks. ",
			analysis["sameSiteUnset"])
	}

	report += "Below are the details of the SameSite attributes:\n"
	report += fmt.Sprintf("\t- Strict: %.2f%%\n", analysis["sameSiteStrict"])
	report += fmt.Sprintf("\t- Lax: %.2f%%\n", analysis["sameSiteLax"])
	report += fmt.Sprintf("\t- None: %.2f%%\n", analysis["sameSiteNone"])
	report += fmt.Sprintf("\t- Unset: %.2f%%\n", analysis["sameSiteUnset"])
	report += "\n"

	// ### SESSION METRIC ###
	if analysis["sessionCookies"] >= 0 {
		report += "there is at least one session cookie, which means you have logged in to at least one " +
			"site. Session cookies are temporary and are deleted when the browser is closed. "
	}

	if analysis["sessionCookies"] >= SESSION_THRESHOLD {
		report += "However, A significant number of cookies are session-based. "
	}
	if analysis["persistentCookies"] >= PERSISTENT_THRESHOLD {
		report += "However, there are more Persistent cookies than session cookies. "
	}

	report += "Below are the details of the session and persistent cookies:\n"
	report += fmt.Sprintf("\t- Session Cookies: %.2f%%\n", analysis["sessionCookies"])
	report += fmt.Sprintf("\t- Persistent Cookies: %.2f%%\n", analysis["persistentCookies"])
	report += "\n"

	return report

}

// appendReportToFile appends the report to DATA.txt with timestamp and metadata
func AppendDataToFile(report, url, browser string) error {
	// Open file in append mode, create if it doesn't exist
	file, err := os.OpenFile("DATA.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open DATA.txt: %v", err)
	}
	defer file.Close()

	// Create a formatted entry with timestamp and metadata
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("\n=== Privacy Analysis Report ===\n")
	entry += fmt.Sprintf("Timestamp: %s\n", timestamp)
	entry += fmt.Sprintf("URL: %s\n", url)
	entry += fmt.Sprintf("Browser: %s\n", browser)
	entry += fmt.Sprintf("Report:\n%s\n", report)
	entry += fmt.Sprintf("=== End Report ===\n\n")

	// Write the entry to the file
	_, err = file.WriteString(entry)
	if err != nil {
		return fmt.Errorf("failed to write to DATA.txt: %v", err)
	}

	return nil
}

// Function: Get Metrics Report
// Operation: Generates a report of the privacy metrics in a structured format
// Return: A string containing the formatted metrics report
func GetMetricsReport(privacyMetrics PrivacyMetric, metricName string) string {
	var report strings.Builder

	report.WriteString(fmt.Sprintf("#----- Printing %s Privacy Metrics ------#\n", metricName))

	report.WriteString(fmt.Sprintf("Total Cookies: %d\n", privacyMetrics.TotalCookies))

	report.WriteString(fmt.Sprintf("Total First-Party Cookies: %d\n", privacyMetrics.TotalFirstParty))
	report.WriteString(fmt.Sprintf("Total Third-Party Cookies: %d\n", privacyMetrics.TotalThirdParty))

	report.WriteString(fmt.Sprintf("Total Secure Domains: %d\n", privacyMetrics.TotalSecure))
	report.WriteString(fmt.Sprintf("Total Unsecure Domains: %d\n", privacyMetrics.TotalNotSecure))

	if len(privacyMetrics.SuspiciousPaths) > 0 {
		report.WriteString("All Suspicious Paths\n")
		for i := 0; i < len(privacyMetrics.SuspiciousPaths); i++ {
			cookie := privacyMetrics.SuspiciousPaths[i]

			partyType := "third-party"
			if cookie.IsFirstParty {
				partyType = "first-party"
			}
			report.WriteString(fmt.Sprintf("\tCookie %d. [%s]\n", i, partyType))
			report.WriteString(fmt.Sprintf("\t\tDomain: %s\n", cookie.Domain))
			report.WriteString(fmt.Sprintf("\t\tName: %s\n", cookie.Name))
			report.WriteString(fmt.Sprintf("\t\tPath: %s\n", cookie.Path))
		}
	} else {
		report.WriteString("No Suspicious Paths\n")
	}

	report.WriteString(fmt.Sprintf("Total HttpOnly: %d\n", privacyMetrics.TotalHttpOnly))
	report.WriteString(fmt.Sprintf("Total Not HttpOnly: %d\n", privacyMetrics.TotalNotHttpOnly))

	report.WriteString(fmt.Sprintf("Total SameSite with Strict: %d\n", privacyMetrics.SameSiteStrict))
	report.WriteString(fmt.Sprintf("Total SameSite with Lax: %d\n", privacyMetrics.SameSiteLax))
	report.WriteString(fmt.Sprintf("Total SameSite with None: %d\n", privacyMetrics.SameSiteNone))
	report.WriteString(fmt.Sprintf("Total SameSite with Unset: %d\n", privacyMetrics.SameSiteUnset))

	report.WriteString(fmt.Sprintf("Total Session Cookies: %d\n", privacyMetrics.TotalSessionCookies))
	report.WriteString(fmt.Sprintf("Total Persistent Cookies: %d\n", privacyMetrics.TotalPersistentCookies))

	report.WriteString("#--------------------------------------------#\n")

	return report.String()
}

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
