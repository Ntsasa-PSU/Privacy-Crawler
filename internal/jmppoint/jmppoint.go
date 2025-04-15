package jmppoint

import (
	"fmt"
	"net/http"
	"os/exec"
)

// TESTING: This is a test function to check if the code is working as expected.
// Current understanding is a little shakey. This code was used to demo a wrapper for the main crawler code.
// Further working on the handler will 
func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	verbose := query.Get("verbose")
	browser := query.Get("browser")

	fmt.Printf("[DEBUG] Request received\n")
	fmt.Printf("[DEBUG] Query parameters: verbose=%q, browser=%q\n", verbose, browser)

	args := []string{}
	if verbose != "" && verbose != "v" {
		errMsg := fmt.Sprintf("Invalid verbose value: %q. Expected 'v'", verbose)
		fmt.Printf("[ERROR] %s\n", errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	if verbose == "v" {
		args = append(args, "-v")
		fmt.Printf("[DEBUG] Added verbose flag\n")
	}

	if browser != "" {
		args = append(args, "-b", browser)
		fmt.Printf("[DEBUG] Added browser flag: %s\n", browser)
	}

	crawlerPath := "./maincrawler/main-crawler.go"
	fmt.Printf("[DEBUG] Executing crawler at: %s\n", crawlerPath)
	fmt.Printf("[DEBUG] Command args: %v\n", args)

	cmdArgs := append([]string{"run", crawlerPath}, args...)
	cmd := exec.Command("go", cmdArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		errMsg := fmt.Sprintf("Crawler execution failed: %v\nOutput: %s", err, output)
		fmt.Printf("[ERROR] %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}
