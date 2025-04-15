package crawler

import (
	"encoding/json"
	"fmt"
	"os"
	"io"
	"time"
	"net/http"
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
// Return: String(Browser found)
func VerifyTargetBrowser(browsers map[string]string, selectedBrowser string, verbose *bool) string {

	// - Verbose Output - //
	if *verbose {

		fmt.Println("\n-- Verifying Browser Selection -- ")
	}

	// Iterate through map of browsers.
	for browserName := range browsers {

		if browserName == selectedBrowser {

			// - Verbose Output - //
			if *verbose {
				fmt.Printf("Found: %s\n", browserName)
			}

			return browserName
		}
	}

	fmt.Printf("Error: Failure to find browser.\n")
	return "None"
}

// Fucntion: Fetch URL
// Operation: Connect and returns header from selected URL.
// Return: Header (map), Status Code (int), Error

// Note: browser is just a string in field for "user-agent"
func FetchHeaders(url string, browser string, verbose *bool) (map[string][]string, int, error) {
    
	// Create a new HTTP client.
	client := &http.Client{
		// Timeout if cannot connect.
        Timeout: time.Second * 10,
    }
    
    // Create a new request.
	request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, 0, fmt.Errorf("error creating request: %v", err)
    }
    
    // Set the User-Agent header based on browser.
	request.Header.Set("User-Agent", browser)

    // Make the request.
	response, err := client.Do(request)
    if err != nil {
        return nil, 0, fmt.Errorf("error making request: %v", err)
    }
	
	// - Verbose output - //
	if *verbose {
		// Read the response body
		bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
    		return response.Header, response.StatusCode, fmt.Errorf("error reading response body: %v", err)
		}
		// Print the response body
		fmt.Printf("\n-- Response body -- \n%s\n", bodyBytes)
	}
	// Close the response.
	defer response.Body.Close()

   
	
   // - Verbose output - //
	if *verbose {
		fmt.Printf("\n-- HTTP Headers for %s --\n", url)
		fmt.Printf("Status: %s\n", response.Status)
	}

	// Return the headers and status code.
	return response.Header, response.StatusCode, nil
}



// Next Step 1: Handle HTTP Headers
//    3. Parse HTTP Headers

// Next Step 2: Create Server Go agents can connect to
//    1. Create a server that can analyze the connection
//    2. Create a server that can analyze the HTTP headers
//    3. Process the HTTP headers data
//    4. Return the HTTP headers analysis results in JSON format