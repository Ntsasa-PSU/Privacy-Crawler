package crawler

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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

	// Iterate through map of browsers
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

// Next Step 1: Handle HTTP Headers
//    1. Connect to a URL
//    2. Get HTTP Headers
//    3. Parse HTTP Headers
//    4. Return HTTP Headers

// Next Step 2: Create Server Go agents can connect to
//    1. Create a server that can analyze the connection
//    2. Create a server that can analyze the HTTP headers
//    3. Process the HTTP headers data
//    4. Return the HTTP headers analysis results in JSON format
// PROMETHESUS PACKAGE - METRIC SERVER: Use? YES
