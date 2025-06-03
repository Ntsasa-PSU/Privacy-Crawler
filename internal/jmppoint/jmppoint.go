package jmppoint

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

// ---- DATA STRUCTURES ---- //

// Options: Represents tags for the main method.
type ProcessOptions struct{
	
	// Functionality.
	browser string
	url string
	hidden bool

	// Testing & Logs.
	verbose bool
	test bool
}

//Process: Holds the option for the given process.
type Process struct{

	options ProcessOptions

	// Port that this porcess will run off.
	port int
}

//Data type that will act as a warpper for intializing with functions
type ProcessOptionsFunc func (*ProcessOptions)


func defaultProcessOptions() ProcessOptions{

	return &ProcessOptions{
		browser: chrome,
		url: "https://www.google.com",
		hidden: false,

		verbose: false,
		test: false
	}
	
}

func newProcess(opts ...ProcessOptionsFunc) *Process {

	o := defaultProcessOptions()
	
	for _, fn := range opts{

		fn((&o))
	}

	return &Process{

		options: o, 
	}


}






















































// TESTING: This is a test function to check if the code is working as expected.
// Current understanding is a little shakey. This code was used to demo a wrapper for the main crawler code.
func UserInputHandler(w http.ResponseWriter, r *http.Request) {
	// Serve HTML form for GET requests
	if r.Method == "GET" {
		html := `
        <!DOCTYPE html>
        <html>
        <body>
            <form action="/test" method="POST">
                <label>Name:</label><br>
                <input type="text" name="name"><br>
                <label>Age:</label><br>
                <input type="number" name="age"><br>
                <label>Browser:</label><br>
                <select name="browser">
                    <option value="chrome">Chrome</option>
                    <option value="firefox">Firefox</option>
                    <option value="safari">Safari</option>
                </select><br>
                <label>Verbose:</label>
                <input type="checkbox" name="verbose" value="v"><br><br>
                <input type="submit" value="Run Crawler">
            </form>
        </body>
        </html>`
		fmt.Fprint(w, html)
		return
	}

	// Process form submission for POST requests
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get form values
		name := r.FormValue("name")
		age := r.FormValue("age")
		browser := r.FormValue("browser")
		verbose := r.FormValue("verbose")

		// Build command args
		args := []string{"-t"}

		if verbose == "v" {
			args = append(args, "-v")
		}
		if browser != "" {
			args = append(args, "-b", browser)
		}

		// Execute crawler with input
		crawlerPath := "./maincrawler/main-crawler.go"
		cmdArgs := append([]string{"run", crawlerPath}, args...)
		cmd := exec.Command("go", cmdArgs...)

		// Set the input values as environment variables
		cmd.Env = append(os.Environ(),
			fmt.Sprintf("USER_NAME=%s", name),
			fmt.Sprintf("USER_AGE=%s", age),
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v\nOutput: %s", err, output), 500)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write(output)
	}
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
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
