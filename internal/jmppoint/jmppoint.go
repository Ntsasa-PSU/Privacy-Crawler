package jmppoint

import (
	"bufio"
	"fmt"
	"os"
	"privcrawler/internal/crawler"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var PORT int = 22

// ---- DATA STRUCTURES ---- //

// Browser Statitics: Hold totaled statistics for a browser
type BrowserStats struct {
	Browser                string
	TotalReports           int
	TotalCookies           int
	TotalFirstParty        int
	TotalThirdParty        int
	TotalSecure            int
	TotalUnsecure          int
	TotalHttpOnly          int
	TotalNotHttpOnly       int
	TotalSameSiteStrict    int
	TotalSameSiteLax       int
	TotalSameSiteNone      int
	TotalSessionCookies    int
	TotalPersistentCookies int
	SuspiciousPathsCount   int
}

// Options: Represents tags for the main method.
type ProcessOptions struct {
	// Functionality.
	browser  string
	url      string
	hidden   bool
	duration int
	verbose  bool
}

// Process: Holds the option for the given process.
type Process struct {
	options ProcessOptions
	// Port that this process will run off.
	port int
}

// Data type that will act as a wrapper for initializing with functions
type ProcessOptionsFunc func(*ProcessOptions)

// Constants for browser types
const (
	chrome  = "chrome"
	firefox = "firefox"
	edge    = "edge"
	safari  = "safari"
)

func defaultProcessOptions() ProcessOptions {
	return ProcessOptions{
		browser:  chrome,
		url:      "https://www.google.com",
		duration: 2000,
		hidden:   false,
		verbose:  false,
	}
}

// ---- FUNCTIONAL OPTIONS ---- //

// WithBrowser sets the browser option
func WithBrowser(browser string) ProcessOptionsFunc {
	return func(opts *ProcessOptions) {
		opts.browser = browser
	}
}

// WithURL sets the URL option
func WithURL(url string) ProcessOptionsFunc {
	return func(opts *ProcessOptions) {
		opts.url = url
	}
}

// WithHidden sets the hidden browser option
func WithHidden(hidden bool) ProcessOptionsFunc {
	return func(opts *ProcessOptions) {
		opts.hidden = hidden
	}
}

// WithVerbose sets the verbose logging option
func WithVerbose(verbose bool) ProcessOptionsFunc {
	return func(opts *ProcessOptions) {
		opts.verbose = verbose
	}
}

// WithDuration sets the duration for the process
func WithDuration(duration int) ProcessOptionsFunc {
	return func(opts *ProcessOptions) {
		opts.duration = duration
	}
}

// ---- CONSTRUCTOR ---- //
func NewProcess(opts ...ProcessOptionsFunc) *Process {
	o := defaultProcessOptions()

	for _, fn := range opts {
		fn(&o)
	}

	return &Process{
		options: o,
		port:    PORT, // default port
	}
}

// ---- METHODS ---- //

// GetBrowser returns the browser option
func (p *Process) GetBrowser() string {
	return p.options.browser
}

// GetURL returns the URL option
func (p *Process) GetURL() string {
	return p.options.url
}

// IsHidden returns the hidden option
func (p *Process) IsHidden() bool {
	return p.options.hidden
}

// IsVerbose returns the verbose option
func (p *Process) IsVerbose() bool {
	return p.options.verbose
}

// GetPort returns the port
func (p *Process) GetPort() int {
	return p.port
}

// Run executes the privacy crawl with the given options
func (p *Process) Run() error {
	return crawler.RunPrivacyCrawl(
		p.options.browser,
		p.options.hidden,
		p.options.url,
		p.options.duration,
		p.options.verbose,
	)
}

func RunServer() error {
	fmt.Println("Starting privacy crawler with website-by-website processing...")

	// -- WEBSITE: AMAZON -- //
	fmt.Println("Processing AMAZON...")
	var amazonWg sync.WaitGroup
	amazonWg.Add(24) // 4 browsers × 6 durations = 24 processes

	// All Amazon processes
	go func() {
		defer amazonWg.Done()
		chrome_amazon_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(0))
		chrome_amazon_0.Run()
	}()
	go func() {
		defer amazonWg.Done()
		safari_amazon_0 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(0))
		safari_amazon_0.Run()
	}()
	go func() {
		defer amazonWg.Done()
		firefox_amazon_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(0))
		firefox_amazon_0.Run()
	}()
	go func() {
		defer amazonWg.Done()
		edge_amazon_0 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(0))
		edge_amazon_0.Run()
	}()

	go func() {
		defer amazonWg.Done()
		chrome_amazon_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(5000))
		chrome_amazon_5.Run()
	}()
	go func() {
		defer amazonWg.Done()
		safari_amazon_5 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(5000))
		safari_amazon_5.Run()
	}()
	go func() {
		defer amazonWg.Done()
		firefox_amazon_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(5000))
		firefox_amazon_5.Run()
	}()
	go func() {
		defer amazonWg.Done()
		edge_amazon_5 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(5000))
		edge_amazon_5.Run()
	}()

	go func() {
		defer amazonWg.Done()
		chrome_amazon_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(10000))
		chrome_amazon_10.Run()
	}()
	go func() {
		defer amazonWg.Done()
		safari_amazon_10 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(10000))
		safari_amazon_10.Run()
	}()
	go func() {
		defer amazonWg.Done()
		firefox_amazon_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(10000))
		firefox_amazon_10.Run()
	}()
	go func() {
		defer amazonWg.Done()
		edge_amazon_10 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(10000))
		edge_amazon_10.Run()
	}()

	go func() {
		defer amazonWg.Done()
		chrome_amazon_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(15000))
		chrome_amazon_15.Run()
	}()
	go func() {
		defer amazonWg.Done()
		safari_amazon_15 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(15000))
		safari_amazon_15.Run()
	}()
	go func() {
		defer amazonWg.Done()
		firefox_amazon_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(15000))
		firefox_amazon_15.Run()
	}()
	go func() {
		defer amazonWg.Done()
		edge_amazon_15 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(15000))
		edge_amazon_15.Run()
	}()

	go func() {
		defer amazonWg.Done()
		chrome_amazon_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(20000))
		chrome_amazon_20.Run()
	}()
	go func() {
		defer amazonWg.Done()
		safari_amazon_20 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(20000))
		safari_amazon_20.Run()
	}()
	go func() {
		defer amazonWg.Done()
		firefox_amazon_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(20000))
		firefox_amazon_20.Run()
	}()
	go func() {
		defer amazonWg.Done()
		edge_amazon_20 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(20000))
		edge_amazon_20.Run()
	}()

	go func() {
		defer amazonWg.Done()
		chrome_amazon_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(25000))
		chrome_amazon_25.Run()
	}()
	go func() {
		defer amazonWg.Done()
		safari_amazon_25 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(25000))
		safari_amazon_25.Run()
	}()
	go func() {
		defer amazonWg.Done()
		firefox_amazon_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(25000))
		firefox_amazon_25.Run()
	}()
	go func() {
		defer amazonWg.Done()
		edge_amazon_25 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(25000))
		edge_amazon_25.Run()
	}()

	amazonWg.Wait()
	fmt.Println("AMAZON completed!")

	// -- WEBSITE: YAHOO -- //
	fmt.Println("Processing YAHOO...")
	var yahooWg sync.WaitGroup
	yahooWg.Add(24)

	go func() {
		defer yahooWg.Done()
		chrome_yahoo_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(0))
		chrome_yahoo_0.Run()
	}()
	go func() {
		defer yahooWg.Done()
		safari_yahoo_0 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(0))
		safari_yahoo_0.Run()
	}()
	go func() {
		defer yahooWg.Done()
		firefox_yahoo_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(0))
		firefox_yahoo_0.Run()
	}()
	go func() {
		defer yahooWg.Done()
		edge_yahoo_0 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(0))
		edge_yahoo_0.Run()
	}()

	go func() {
		defer yahooWg.Done()
		chrome_yahoo_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(5000))
		chrome_yahoo_5.Run()
	}()
	go func() {
		defer yahooWg.Done()
		safari_yahoo_5 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(5000))
		safari_yahoo_5.Run()
	}()
	go func() {
		defer yahooWg.Done()
		firefox_yahoo_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(5000))
		firefox_yahoo_5.Run()
	}()
	go func() {
		defer yahooWg.Done()
		edge_yahoo_5 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(5000))
		edge_yahoo_5.Run()
	}()

	go func() {
		defer yahooWg.Done()
		chrome_yahoo_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(10000))
		chrome_yahoo_10.Run()
	}()
	go func() {
		defer yahooWg.Done()
		safari_yahoo_10 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(10000))
		safari_yahoo_10.Run()
	}()
	go func() {
		defer yahooWg.Done()
		firefox_yahoo_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(10000))
		firefox_yahoo_10.Run()
	}()
	go func() {
		defer yahooWg.Done()
		edge_yahoo_10 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(10000))
		edge_yahoo_10.Run()
	}()

	go func() {
		defer yahooWg.Done()
		chrome_yahoo_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(15000))
		chrome_yahoo_15.Run()
	}()
	go func() {
		defer yahooWg.Done()
		safari_yahoo_15 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(15000))
		safari_yahoo_15.Run()
	}()
	go func() {
		defer yahooWg.Done()
		firefox_yahoo_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(15000))
		firefox_yahoo_15.Run()
	}()
	go func() {
		defer yahooWg.Done()
		edge_yahoo_15 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(15000))
		edge_yahoo_15.Run()
	}()

	go func() {
		defer yahooWg.Done()
		chrome_yahoo_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(20000))
		chrome_yahoo_20.Run()
	}()
	go func() {
		defer yahooWg.Done()
		safari_yahoo_20 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(20000))
		safari_yahoo_20.Run()
	}()
	go func() {
		defer yahooWg.Done()
		firefox_yahoo_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(20000))
		firefox_yahoo_20.Run()
	}()
	go func() {
		defer yahooWg.Done()
		edge_yahoo_20 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(20000))
		edge_yahoo_20.Run()
	}()

	go func() {
		defer yahooWg.Done()
		chrome_yahoo_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(25000))
		chrome_yahoo_25.Run()
	}()
	go func() {
		defer yahooWg.Done()
		safari_yahoo_25 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(25000))
		safari_yahoo_25.Run()
	}()
	go func() {
		defer yahooWg.Done()
		firefox_yahoo_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(25000))
		firefox_yahoo_25.Run()
	}()
	go func() {
		defer yahooWg.Done()
		edge_yahoo_25 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(25000))
		edge_yahoo_25.Run()
	}()

	yahooWg.Wait()
	fmt.Println("YAHOO completed!")

	// -- WEBSITE: REDDIT -- //
	fmt.Println("Processing REDDIT...")
	var redditWg sync.WaitGroup
	redditWg.Add(24)

	go func() {
		defer redditWg.Done()
		chrome_reddit_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(0))
		chrome_reddit_0.Run()
	}()
	go func() {
		defer redditWg.Done()
		safari_reddit_0 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(0))
		safari_reddit_0.Run()
	}()
	go func() {
		defer redditWg.Done()
		firefox_reddit_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(0))
		firefox_reddit_0.Run()
	}()
	go func() {
		defer redditWg.Done()
		edge_reddit_0 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(0))
		edge_reddit_0.Run()
	}()

	go func() {
		defer redditWg.Done()
		chrome_reddit_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(5000))
		chrome_reddit_5.Run()
	}()
	go func() {
		defer redditWg.Done()
		safari_reddit_5 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(5000))
		safari_reddit_5.Run()
	}()
	go func() {
		defer redditWg.Done()
		firefox_reddit_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(5000))
		firefox_reddit_5.Run()
	}()
	go func() {
		defer redditWg.Done()
		edge_reddit_5 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(5000))
		edge_reddit_5.Run()
	}()

	go func() {
		defer redditWg.Done()
		chrome_reddit_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(10000))
		chrome_reddit_10.Run()
	}()
	go func() {
		defer redditWg.Done()
		safari_reddit_10 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(10000))
		safari_reddit_10.Run()
	}()
	go func() {
		defer redditWg.Done()
		firefox_reddit_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(10000))
		firefox_reddit_10.Run()
	}()
	go func() {
		defer redditWg.Done()
		edge_reddit_10 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(10000))
		edge_reddit_10.Run()
	}()

	go func() {
		defer redditWg.Done()
		chrome_reddit_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(15000))
		chrome_reddit_15.Run()
	}()
	go func() {
		defer redditWg.Done()
		safari_reddit_15 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(15000))
		safari_reddit_15.Run()
	}()
	go func() {
		defer redditWg.Done()
		firefox_reddit_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(15000))
		firefox_reddit_15.Run()
	}()
	go func() {
		defer redditWg.Done()
		edge_reddit_15 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(15000))
		edge_reddit_15.Run()
	}()

	go func() {
		defer redditWg.Done()
		chrome_reddit_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(20000))
		chrome_reddit_20.Run()
	}()
	go func() {
		defer redditWg.Done()
		safari_reddit_20 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(20000))
		safari_reddit_20.Run()
	}()
	go func() {
		defer redditWg.Done()
		firefox_reddit_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(20000))
		firefox_reddit_20.Run()
	}()
	go func() {
		defer redditWg.Done()
		edge_reddit_20 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(20000))
		edge_reddit_20.Run()
	}()

	go func() {
		defer redditWg.Done()
		chrome_reddit_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(25000))
		chrome_reddit_25.Run()
	}()
	go func() {
		defer redditWg.Done()
		safari_reddit_25 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(25000))
		safari_reddit_25.Run()
	}()
	go func() {
		defer redditWg.Done()
		firefox_reddit_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(25000))
		firefox_reddit_25.Run()
	}()
	go func() {
		defer redditWg.Done()
		edge_reddit_25 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(25000))
		edge_reddit_25.Run()
	}()

	redditWg.Wait()
	fmt.Println("REDDIT completed!")

	// -- WEBSITE: PINTEREST -- //
	fmt.Println("Processing PINTEREST...")
	var pinterestWg sync.WaitGroup
	pinterestWg.Add(24)

	go func() {
		defer pinterestWg.Done()
		chrome_pinterest_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(0))
		chrome_pinterest_0.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		safari_pinterest_0 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(0))
		safari_pinterest_0.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		firefox_pinterest_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(0))
		firefox_pinterest_0.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		edge_pinterest_0 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(0))
		edge_pinterest_0.Run()
	}()

	go func() {
		defer pinterestWg.Done()
		chrome_pinterest_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(5000))
		chrome_pinterest_5.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		safari_pinterest_5 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(5000))
		safari_pinterest_5.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		firefox_pinterest_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(5000))
		firefox_pinterest_5.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		edge_pinterest_5 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(5000))
		edge_pinterest_5.Run()
	}()

	go func() {
		defer pinterestWg.Done()
		chrome_pinterest_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(10000))
		chrome_pinterest_10.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		safari_pinterest_10 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(10000))
		safari_pinterest_10.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		firefox_pinterest_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(10000))
		firefox_pinterest_10.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		edge_pinterest_10 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(10000))
		edge_pinterest_10.Run()
	}()

	go func() {
		defer pinterestWg.Done()
		chrome_pinterest_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(15000))
		chrome_pinterest_15.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		safari_pinterest_15 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(15000))
		safari_pinterest_15.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		firefox_pinterest_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(15000))
		firefox_pinterest_15.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		edge_pinterest_15 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(15000))
		edge_pinterest_15.Run()
	}()

	go func() {
		defer pinterestWg.Done()
		chrome_pinterest_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(20000))
		chrome_pinterest_20.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		safari_pinterest_20 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(20000))
		safari_pinterest_20.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		firefox_pinterest_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(20000))
		firefox_pinterest_20.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		edge_pinterest_20 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(20000))
		edge_pinterest_20.Run()
	}()

	go func() {
		defer pinterestWg.Done()
		chrome_pinterest_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(25000))
		chrome_pinterest_25.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		safari_pinterest_25 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(25000))
		safari_pinterest_25.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		firefox_pinterest_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(25000))
		firefox_pinterest_25.Run()
	}()
	go func() {
		defer pinterestWg.Done()
		edge_pinterest_25 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(25000))
		edge_pinterest_25.Run()
	}()

	pinterestWg.Wait()
	fmt.Println("PINTEREST completed!")

	// -- WEBSITE: FANDOM -- //
	fmt.Println("Processing FANDOM...")
	var fandomWg sync.WaitGroup
	fandomWg.Add(24)

	go func() {
		defer fandomWg.Done()
		chrome_fandom_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(0))
		chrome_fandom_0.Run()
	}()
	go func() {
		defer fandomWg.Done()
		safari_fandom_0 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(0))
		safari_fandom_0.Run()
	}()
	go func() {
		defer fandomWg.Done()
		firefox_fandom_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(0))
		firefox_fandom_0.Run()
	}()
	go func() {
		defer fandomWg.Done()
		edge_fandom_0 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(0))
		edge_fandom_0.Run()
	}()

	go func() {
		defer fandomWg.Done()
		chrome_fandom_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(5000))
		chrome_fandom_5.Run()
	}()
	go func() {
		defer fandomWg.Done()
		safari_fandom_5 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(5000))
		safari_fandom_5.Run()
	}()
	go func() {
		defer fandomWg.Done()
		firefox_fandom_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(5000))
		firefox_fandom_5.Run()
	}()
	go func() {
		defer fandomWg.Done()
		edge_fandom_5 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(5000))
		edge_fandom_5.Run()
	}()

	go func() {
		defer fandomWg.Done()
		chrome_fandom_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(10000))
		chrome_fandom_10.Run()
	}()
	go func() {
		defer fandomWg.Done()
		safari_fandom_10 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(10000))
		safari_fandom_10.Run()
	}()
	go func() {
		defer fandomWg.Done()
		firefox_fandom_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(10000))
		firefox_fandom_10.Run()
	}()
	go func() {
		defer fandomWg.Done()
		edge_fandom_10 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(10000))
		edge_fandom_10.Run()
	}()

	go func() {
		defer fandomWg.Done()
		chrome_fandom_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(15000))
		chrome_fandom_15.Run()
	}()
	go func() {
		defer fandomWg.Done()
		safari_fandom_15 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(15000))
		safari_fandom_15.Run()
	}()
	go func() {
		defer fandomWg.Done()
		firefox_fandom_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(15000))
		firefox_fandom_15.Run()
	}()
	go func() {
		defer fandomWg.Done()
		edge_fandom_15 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(15000))
		edge_fandom_15.Run()
	}()

	go func() {
		defer fandomWg.Done()
		chrome_fandom_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(20000))
		chrome_fandom_20.Run()
	}()
	go func() {
		defer fandomWg.Done()
		safari_fandom_20 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(20000))
		safari_fandom_20.Run()
	}()
	go func() {
		defer fandomWg.Done()
		firefox_fandom_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(20000))
		firefox_fandom_20.Run()
	}()
	go func() {
		defer fandomWg.Done()
		edge_fandom_20 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(20000))
		edge_fandom_20.Run()
	}()

	go func() {
		defer fandomWg.Done()
		chrome_fandom_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(25000))
		chrome_fandom_25.Run()
	}()
	go func() {
		defer fandomWg.Done()
		safari_fandom_25 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(25000))
		safari_fandom_25.Run()
	}()
	go func() {
		defer fandomWg.Done()
		firefox_fandom_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(25000))
		firefox_fandom_25.Run()
	}()
	go func() {
		defer fandomWg.Done()
		edge_fandom_25 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(25000))
		edge_fandom_25.Run()
	}()

	fandomWg.Wait()
	fmt.Println("FANDOM completed!")

	// -- WEBSITE: PORNHUB -- //
	fmt.Println("Processing PORNHUB...")
	var pornhubWg sync.WaitGroup
	pornhubWg.Add(24)

	go func() {
		defer pornhubWg.Done()
		chrome_pornhub_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(0))
		chrome_pornhub_0.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		safari_pornhub_0 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(0))
		safari_pornhub_0.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		firefox_pornhub_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(0))
		firefox_pornhub_0.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		edge_pornhub_0 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(0))
		edge_pornhub_0.Run()
	}()

	go func() {
		defer pornhubWg.Done()
		chrome_pornhub_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(5000))
		chrome_pornhub_5.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		safari_pornhub_5 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(5000))
		safari_pornhub_5.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		firefox_pornhub_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(5000))
		firefox_pornhub_5.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		edge_pornhub_5 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(5000))
		edge_pornhub_5.Run()
	}()

	go func() {
		defer pornhubWg.Done()
		chrome_pornhub_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(10000))
		chrome_pornhub_10.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		safari_pornhub_10 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(10000))
		safari_pornhub_10.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		firefox_pornhub_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(10000))
		firefox_pornhub_10.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		edge_pornhub_10 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(10000))
		edge_pornhub_10.Run()
	}()

	go func() {
		defer pornhubWg.Done()
		chrome_pornhub_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(15000))
		chrome_pornhub_15.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		safari_pornhub_15 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(15000))
		safari_pornhub_15.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		firefox_pornhub_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(15000))
		firefox_pornhub_15.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		edge_pornhub_15 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(15000))
		edge_pornhub_15.Run()
	}()

	go func() {
		defer pornhubWg.Done()
		chrome_pornhub_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(20000))
		chrome_pornhub_20.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		safari_pornhub_20 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(20000))
		safari_pornhub_20.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		firefox_pornhub_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(20000))
		firefox_pornhub_20.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		edge_pornhub_20 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(20000))
		edge_pornhub_20.Run()
	}()

	go func() {
		defer pornhubWg.Done()
		chrome_pornhub_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(25000))
		chrome_pornhub_25.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		safari_pornhub_25 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(25000))
		safari_pornhub_25.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		firefox_pornhub_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(25000))
		firefox_pornhub_25.Run()
	}()
	go func() {
		defer pornhubWg.Done()
		edge_pornhub_25 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(25000))
		edge_pornhub_25.Run()
	}()

	pornhubWg.Wait()
	fmt.Println("PORNHUB completed!")

	// -- WEBSITE: SOAP2DAY -- //
	fmt.Println("Processing SOAP2DAY...")
	var soap2dayWg sync.WaitGroup
	soap2dayWg.Add(24)

	go func() {
		defer soap2dayWg.Done()
		chrome_soap2day_0 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(0))
		chrome_soap2day_0.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		safari_soap2day_0 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(0))
		safari_soap2day_0.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		firefox_soap2day_0 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(0))
		firefox_soap2day_0.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		edge_soap2day_0 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(0))
		edge_soap2day_0.Run()
	}()

	go func() {
		defer soap2dayWg.Done()
		chrome_soap2day_5 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(5000))
		chrome_soap2day_5.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		safari_soap2day_5 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(5000))
		safari_soap2day_5.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		firefox_soap2day_5 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(5000))
		firefox_soap2day_5.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		edge_soap2day_5 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(5000))
		edge_soap2day_5.Run()
	}()

	go func() {
		defer soap2dayWg.Done()
		chrome_soap2day_10 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(10000))
		chrome_soap2day_10.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		safari_soap2day_10 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(10000))
		safari_soap2day_10.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		firefox_soap2day_10 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(10000))
		firefox_soap2day_10.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		edge_soap2day_10 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(10000))
		edge_soap2day_10.Run()
	}()

	go func() {
		defer soap2dayWg.Done()
		chrome_soap2day_15 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(15000))
		chrome_soap2day_15.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		safari_soap2day_15 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(15000))
		safari_soap2day_15.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		firefox_soap2day_15 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(15000))
		firefox_soap2day_15.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		edge_soap2day_15 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(15000))
		edge_soap2day_15.Run()
	}()

	go func() {
		defer soap2dayWg.Done()
		chrome_soap2day_20 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(20000))
		chrome_soap2day_20.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		safari_soap2day_20 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(20000))
		safari_soap2day_20.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		firefox_soap2day_20 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(20000))
		firefox_soap2day_20.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		edge_soap2day_20 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(20000))
		edge_soap2day_20.Run()
	}()

	go func() {
		defer soap2dayWg.Done()
		chrome_soap2day_25 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(25000))
		chrome_soap2day_25.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		safari_soap2day_25 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(25000))
		safari_soap2day_25.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		firefox_soap2day_25 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(25000))
		firefox_soap2day_25.Run()
	}()
	go func() {
		defer soap2dayWg.Done()
		edge_soap2day_25 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(25000))
		edge_soap2day_25.Run()
	}()

	soap2dayWg.Wait()
	fmt.Println("SOAP2DAY completed!")

	// -- WEBSITE: XVIDEOS -- //
	fmt.Println("Processing XVIDEOS...")
	var xvideosWg sync.WaitGroup
	xvideosWg.Add(24)

	go func() {
		defer xvideosWg.Done()
		chrome_xvideos_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(0))
		chrome_xvideos_0.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		safari_xvideos_0 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(0))
		safari_xvideos_0.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		firefox_xvideos_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(0))
		firefox_xvideos_0.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		edge_xvideos_0 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(0))
		edge_xvideos_0.Run()
	}()

	go func() {
		defer xvideosWg.Done()
		chrome_xvideos_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(5000))
		chrome_xvideos_5.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		safari_xvideos_5 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(5000))
		safari_xvideos_5.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		firefox_xvideos_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(5000))
		firefox_xvideos_5.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		edge_xvideos_5 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(5000))
		edge_xvideos_5.Run()
	}()

	go func() {
		defer xvideosWg.Done()
		chrome_xvideos_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(10000))
		chrome_xvideos_10.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		safari_xvideos_10 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(10000))
		safari_xvideos_10.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		firefox_xvideos_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(10000))
		firefox_xvideos_10.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		edge_xvideos_10 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(10000))
		edge_xvideos_10.Run()
	}()

	go func() {
		defer xvideosWg.Done()
		chrome_xvideos_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(15000))
		chrome_xvideos_15.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		safari_xvideos_15 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(15000))
		safari_xvideos_15.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		firefox_xvideos_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(15000))
		firefox_xvideos_15.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		edge_xvideos_15 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(15000))
		edge_xvideos_15.Run()
	}()

	go func() {
		defer xvideosWg.Done()
		chrome_xvideos_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(20000))
		chrome_xvideos_20.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		safari_xvideos_20 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(20000))
		safari_xvideos_20.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		firefox_xvideos_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(20000))
		firefox_xvideos_20.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		edge_xvideos_20 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(20000))
		edge_xvideos_20.Run()
	}()

	go func() {
		defer xvideosWg.Done()
		chrome_xvideos_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(25000))
		chrome_xvideos_25.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		safari_xvideos_25 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(25000))
		safari_xvideos_25.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		firefox_xvideos_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(25000))
		firefox_xvideos_25.Run()
	}()
	go func() {
		defer xvideosWg.Done()
		edge_xvideos_25 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(25000))
		edge_xvideos_25.Run()
	}()

	xvideosWg.Wait()
	fmt.Println("XVIDEOS completed!")

	// -- WEBSITE: FASTDOWNLOAD -- //
	fmt.Println("Processing FASTDOWNLOAD...")
	var fastdownloadWg sync.WaitGroup
	fastdownloadWg.Add(24)

	go func() {
		defer fastdownloadWg.Done()
		chrome_fastdownload_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(0))
		chrome_fastdownload_0.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		safari_fastdownload_0 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(0))
		safari_fastdownload_0.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		firefox_fastdownload_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(0))
		firefox_fastdownload_0.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		edge_fastdownload_0 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(0))
		edge_fastdownload_0.Run()
	}()

	go func() {
		defer fastdownloadWg.Done()
		chrome_fastdownload_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		chrome_fastdownload_5.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		safari_fastdownload_5 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		safari_fastdownload_5.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		firefox_fastdownload_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		firefox_fastdownload_5.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		edge_fastdownload_5 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		edge_fastdownload_5.Run()
	}()

	go func() {
		defer fastdownloadWg.Done()
		chrome_fastdownload_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		chrome_fastdownload_10.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		safari_fastdownload_10 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		safari_fastdownload_10.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		firefox_fastdownload_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		firefox_fastdownload_10.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		edge_fastdownload_10 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		edge_fastdownload_10.Run()
	}()

	go func() {
		defer fastdownloadWg.Done()
		chrome_fastdownload_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		chrome_fastdownload_15.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		safari_fastdownload_15 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		safari_fastdownload_15.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		firefox_fastdownload_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		firefox_fastdownload_15.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		edge_fastdownload_15 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		edge_fastdownload_15.Run()
	}()

	go func() {
		defer fastdownloadWg.Done()
		chrome_fastdownload_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		chrome_fastdownload_20.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		safari_fastdownload_20 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		safari_fastdownload_20.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		firefox_fastdownload_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		firefox_fastdownload_20.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		edge_fastdownload_20 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		edge_fastdownload_20.Run()
	}()

	go func() {
		defer fastdownloadWg.Done()
		chrome_fastdownload_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		chrome_fastdownload_25.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		safari_fastdownload_25 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		safari_fastdownload_25.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		firefox_fastdownload_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		firefox_fastdownload_25.Run()
	}()
	go func() {
		defer fastdownloadWg.Done()
		edge_fastdownload_25 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		edge_fastdownload_25.Run()
	}()

	fastdownloadWg.Wait()
	fmt.Println("FASTDOWNLOAD completed!")

	// -- WEBSITE: ENDACE -- //
	fmt.Println("Processing ENDACE...")
	var endaceWg sync.WaitGroup
	endaceWg.Add(24)

	go func() {
		defer endaceWg.Done()
		chrome_endace_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(0))
		chrome_endace_0.Run()
	}()
	go func() {
		defer endaceWg.Done()
		safari_endace_0 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(0))
		safari_endace_0.Run()
	}()
	go func() {
		defer endaceWg.Done()
		firefox_endace_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(0))
		firefox_endace_0.Run()
	}()
	go func() {
		defer endaceWg.Done()
		edge_endace_0 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(0))
		edge_endace_0.Run()
	}()

	go func() {
		defer endaceWg.Done()
		chrome_endace_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(5000))
		chrome_endace_5.Run()
	}()
	go func() {
		defer endaceWg.Done()
		safari_endace_5 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(5000))
		safari_endace_5.Run()
	}()
	go func() {
		defer endaceWg.Done()
		firefox_endace_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(5000))
		firefox_endace_5.Run()
	}()
	go func() {
		defer endaceWg.Done()
		edge_endace_5 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(5000))
		edge_endace_5.Run()
	}()

	go func() {
		defer endaceWg.Done()
		chrome_endace_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(10000))
		chrome_endace_10.Run()
	}()
	go func() {
		defer endaceWg.Done()
		safari_endace_10 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(10000))
		safari_endace_10.Run()
	}()
	go func() {
		defer endaceWg.Done()
		firefox_endace_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(10000))
		firefox_endace_10.Run()
	}()
	go func() {
		defer endaceWg.Done()
		edge_endace_10 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(10000))
		edge_endace_10.Run()
	}()

	go func() {
		defer endaceWg.Done()
		chrome_endace_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(15000))
		chrome_endace_15.Run()
	}()
	go func() {
		defer endaceWg.Done()
		safari_endace_15 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(15000))
		safari_endace_15.Run()
	}()
	go func() {
		defer endaceWg.Done()
		firefox_endace_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(15000))
		firefox_endace_15.Run()
	}()
	go func() {
		defer endaceWg.Done()
		edge_endace_15 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(15000))
		edge_endace_15.Run()
	}()

	go func() {
		defer endaceWg.Done()
		chrome_endace_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(20000))
		chrome_endace_20.Run()
	}()
	go func() {
		defer endaceWg.Done()
		safari_endace_20 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(20000))
		safari_endace_20.Run()
	}()
	go func() {
		defer endaceWg.Done()
		firefox_endace_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(20000))
		firefox_endace_20.Run()
	}()
	go func() {
		defer endaceWg.Done()
		edge_endace_20 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(20000))
		edge_endace_20.Run()
	}()

	go func() {
		defer endaceWg.Done()
		chrome_endace_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(25000))
		chrome_endace_25.Run()
	}()
	go func() {
		defer endaceWg.Done()
		safari_endace_25 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(25000))
		safari_endace_25.Run()
	}()
	go func() {
		defer endaceWg.Done()
		firefox_endace_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(25000))
		firefox_endace_25.Run()
	}()
	go func() {
		defer endaceWg.Done()
		edge_endace_25 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(25000))
		edge_endace_25.Run()
	}()

	endaceWg.Wait()
	fmt.Println("ENDACE completed!")

	fmt.Println("All websites completed successfully!")
	return nil
}

// ---- DATA PROCESSING ---- //

func GenerateTotalsFile() {
	fmt.Println("Starting to parse DATA.txt...")

	file, err := os.Open("./DATA.txt")
	if err != nil {
		fmt.Printf("Error opening DATA.txt: %v\n", err)
		return
	}
	defer file.Close()

	// Browser totals
	chromeStats := &BrowserStats{Browser: "chrome"}
	edgeStats := &BrowserStats{Browser: "edge"}
	firefoxStats := &BrowserStats{Browser: "firefox"}
	safariStats := &BrowserStats{Browser: "safari"}

	scanner := bufio.NewScanner(file)
	var currentBrowser string
	inReport := false

	// Simple regex patterns
	browserRegex := regexp.MustCompile(`Browser: (\w+)`)
	totalCookiesRegex := regexp.MustCompile(`Total Cookies: (\d+)`)
	firstPartyRegex := regexp.MustCompile(`Total First-Party Cookies: (\d+)`)
	thirdPartyRegex := regexp.MustCompile(`Total Third-Party Cookies: (\d+)`)
	secureRegex := regexp.MustCompile(`Total Secure Domains: (\d+)`)
	unsecureRegex := regexp.MustCompile(`Total Unsecure Domains: (\d+)`)
	httpOnlyRegex := regexp.MustCompile(`Total HttpOnly: (\d+)`)
	notHttpOnlyRegex := regexp.MustCompile(`Total Not HttpOnly: (\d+)`)
	sameSiteStrictRegex := regexp.MustCompile(`Total SameSite with Strict: (\d+)`)
	sameSiteLaxRegex := regexp.MustCompile(`Total SameSite with Lax: (\d+)`)
	sameSiteNoneRegex := regexp.MustCompile(`Total SameSite with None: (\d+)`)
	sessionRegex := regexp.MustCompile(`Total Session Cookies: (\d+)`)
	persistentRegex := regexp.MustCompile(`Total Persistent Cookies: (\d+)`)

	fmt.Println("Parsing browser data...")

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "=== Privacy Analysis Report ===") {
			inReport = true
			continue
		}

		if strings.Contains(line, "=== End Report ===") {
			inReport = false
			currentBrowser = ""
			continue
		}

		if !inReport {
			continue
		}

		// Get browser name
		if matches := browserRegex.FindStringSubmatch(line); matches != nil {
			currentBrowser = matches[1]
			switch currentBrowser {
			case "chrome":
				chromeStats.TotalReports++
			case "edge":
				edgeStats.TotalReports++
			case "firefox":
				firefoxStats.TotalReports++
			case "safari":
				safariStats.TotalReports++
			}
			continue
		}

		if currentBrowser == "" {
			continue
		}

		var currentStats *BrowserStats
		switch currentBrowser {
		case "chrome":
			currentStats = chromeStats
		case "edge":
			currentStats = edgeStats
		case "firefox":
			currentStats = firefoxStats
		case "safari":
			currentStats = safariStats
		default:
			continue
		}

		// Extract and add values
		if matches := totalCookiesRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalCookies += val
			}
		}
		if matches := firstPartyRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalFirstParty += val
			}
		}
		if matches := thirdPartyRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalThirdParty += val
			}
		}
		if matches := secureRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalSecure += val
			}
		}
		if matches := unsecureRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalUnsecure += val
			}
		}
		if matches := httpOnlyRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalHttpOnly += val
			}
		}
		if matches := notHttpOnlyRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalNotHttpOnly += val
			}
		}
		if matches := sameSiteStrictRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalSameSiteStrict += val
			}
		}
		if matches := sameSiteLaxRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalSameSiteLax += val
			}
		}
		if matches := sameSiteNoneRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalSameSiteNone += val
			}
		}
		if matches := sessionRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalSessionCookies += val
			}
		}
		if matches := persistentRegex.FindStringSubmatch(line); matches != nil {
			if val, err := strconv.Atoi(matches[1]); err == nil {
				currentStats.TotalPersistentCookies += val
			}
		}
	}

	fmt.Println("Writing totals to DATA_TOTAL.txt...")

	// Write to file
	outFile, err := os.Create("DATA_TOTAL.txt")
	if err != nil {
		fmt.Printf("Error creating DATA_TOTAL.txt: %v\n", err)
		return
	}
	defer outFile.Close()

	fmt.Fprintf(outFile, "=== BROWSER TOTALS ===\n")
	fmt.Fprintf(outFile, "Generated: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	// Write Chrome totals
	fmt.Fprintf(outFile, "CHROME:\n")
	fmt.Fprintf(outFile, "Total Reports: %d\n", chromeStats.TotalReports)
	fmt.Fprintf(outFile, "Total Cookies: %d\n", chromeStats.TotalCookies)
	fmt.Fprintf(outFile, "First-Party Cookies: %d\n", chromeStats.TotalFirstParty)
	fmt.Fprintf(outFile, "Third-Party Cookies: %d\n", chromeStats.TotalThirdParty)
	fmt.Fprintf(outFile, "Secure Domains: %d\n", chromeStats.TotalSecure)
	fmt.Fprintf(outFile, "Unsecure Domains: %d\n", chromeStats.TotalUnsecure)
	fmt.Fprintf(outFile, "HttpOnly: %d\n", chromeStats.TotalHttpOnly)
	fmt.Fprintf(outFile, "Not HttpOnly: %d\n", chromeStats.TotalNotHttpOnly)
	fmt.Fprintf(outFile, "SameSite Strict: %d\n", chromeStats.TotalSameSiteStrict)
	fmt.Fprintf(outFile, "SameSite Lax: %d\n", chromeStats.TotalSameSiteLax)
	fmt.Fprintf(outFile, "SameSite None: %d\n", chromeStats.TotalSameSiteNone)
	fmt.Fprintf(outFile, "Session Cookies: %d\n", chromeStats.TotalSessionCookies)
	fmt.Fprintf(outFile, "Persistent Cookies: %d\n\n", chromeStats.TotalPersistentCookies)

	// Write Edge totals
	fmt.Fprintf(outFile, "EDGE:\n")
	fmt.Fprintf(outFile, "Total Reports: %d\n", edgeStats.TotalReports)
	fmt.Fprintf(outFile, "Total Cookies: %d\n", edgeStats.TotalCookies)
	fmt.Fprintf(outFile, "First-Party Cookies: %d\n", edgeStats.TotalFirstParty)
	fmt.Fprintf(outFile, "Third-Party Cookies: %d\n", edgeStats.TotalThirdParty)
	fmt.Fprintf(outFile, "Secure Domains: %d\n", edgeStats.TotalSecure)
	fmt.Fprintf(outFile, "Unsecure Domains: %d\n", edgeStats.TotalUnsecure)
	fmt.Fprintf(outFile, "HttpOnly: %d\n", edgeStats.TotalHttpOnly)
	fmt.Fprintf(outFile, "Not HttpOnly: %d\n", edgeStats.TotalNotHttpOnly)
	fmt.Fprintf(outFile, "SameSite Strict: %d\n", edgeStats.TotalSameSiteStrict)
	fmt.Fprintf(outFile, "SameSite Lax: %d\n", edgeStats.TotalSameSiteLax)
	fmt.Fprintf(outFile, "SameSite None: %d\n", edgeStats.TotalSameSiteNone)
	fmt.Fprintf(outFile, "Session Cookies: %d\n", edgeStats.TotalSessionCookies)
	fmt.Fprintf(outFile, "Persistent Cookies: %d\n\n", edgeStats.TotalPersistentCookies)

	// Write Firefox totals
	fmt.Fprintf(outFile, "FIREFOX:\n")
	fmt.Fprintf(outFile, "Total Reports: %d\n", firefoxStats.TotalReports)
	fmt.Fprintf(outFile, "Total Cookies: %d\n", firefoxStats.TotalCookies)
	fmt.Fprintf(outFile, "First-Party Cookies: %d\n", firefoxStats.TotalFirstParty)
	fmt.Fprintf(outFile, "Third-Party Cookies: %d\n", firefoxStats.TotalThirdParty)
	fmt.Fprintf(outFile, "Secure Domains: %d\n", firefoxStats.TotalSecure)
	fmt.Fprintf(outFile, "Unsecure Domains: %d\n", firefoxStats.TotalUnsecure)
	fmt.Fprintf(outFile, "HttpOnly: %d\n", firefoxStats.TotalHttpOnly)
	fmt.Fprintf(outFile, "Not HttpOnly: %d\n", firefoxStats.TotalNotHttpOnly)
	fmt.Fprintf(outFile, "SameSite Strict: %d\n", firefoxStats.TotalSameSiteStrict)
	fmt.Fprintf(outFile, "SameSite Lax: %d\n", firefoxStats.TotalSameSiteLax)
	fmt.Fprintf(outFile, "SameSite None: %d\n", firefoxStats.TotalSameSiteNone)
	fmt.Fprintf(outFile, "Session Cookies: %d\n", firefoxStats.TotalSessionCookies)
	fmt.Fprintf(outFile, "Persistent Cookies: %d\n\n", firefoxStats.TotalPersistentCookies)

	// Write Safari totals
	fmt.Fprintf(outFile, "SAFARI:\n")
	fmt.Fprintf(outFile, "Total Reports: %d\n", safariStats.TotalReports)
	fmt.Fprintf(outFile, "Total Cookies: %d\n", safariStats.TotalCookies)
	fmt.Fprintf(outFile, "First-Party Cookies: %d\n", safariStats.TotalFirstParty)
	fmt.Fprintf(outFile, "Third-Party Cookies: %d\n", safariStats.TotalThirdParty)
	fmt.Fprintf(outFile, "Secure Domains: %d\n", safariStats.TotalSecure)
	fmt.Fprintf(outFile, "Unsecure Domains: %d\n", safariStats.TotalUnsecure)
	fmt.Fprintf(outFile, "HttpOnly: %d\n", safariStats.TotalHttpOnly)
	fmt.Fprintf(outFile, "Not HttpOnly: %d\n", safariStats.TotalNotHttpOnly)
	fmt.Fprintf(outFile, "SameSite Strict: %d\n", safariStats.TotalSameSiteStrict)
	fmt.Fprintf(outFile, "SameSite Lax: %d\n", safariStats.TotalSameSiteLax)
	fmt.Fprintf(outFile, "SameSite None: %d\n", safariStats.TotalSameSiteNone)
	fmt.Fprintf(outFile, "Session Cookies: %d\n", safariStats.TotalSessionCookies)
	fmt.Fprintf(outFile, "Persistent Cookies: %d\n", safariStats.TotalPersistentCookies)

	fmt.Println("DATA_TOTAL.txt created successfully!")
}

func BrowserRanking() {
	fmt.Println("Analyzing browser rankings...")

	file, err := os.Open("./DATA_TOTAL.txt")
	if err != nil {
		fmt.Printf("Error opening DATA_TOTAL.txt: %v\n", err)
		return
	}
	defer file.Close()

	// Parse the existing DATA_TOTAL.txt format
	chromeStats := &BrowserStats{Browser: "chrome"}
	edgeStats := &BrowserStats{Browser: "edge"}
	firefoxStats := &BrowserStats{Browser: "firefox"}
	safariStats := &BrowserStats{Browser: "safari"}

	scanner := bufio.NewScanner(file)
	var currentStats *BrowserStats

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Identify which browser section we're in
		if line == "CHROME:" {
			currentStats = chromeStats
			continue
		} else if line == "EDGE:" {
			currentStats = edgeStats
			continue
		} else if line == "FIREFOX:" {
			currentStats = firefoxStats
			continue
		} else if line == "SAFARI:" {
			currentStats = safariStats
			continue
		}

		if currentStats == nil {
			continue
		}

		// Parse the values
		if strings.HasPrefix(line, "Total Cookies: ") {
			fmt.Sscanf(line, "Total Cookies: %d", &currentStats.TotalCookies)
		} else if strings.HasPrefix(line, "Third-Party Cookies: ") {
			fmt.Sscanf(line, "Third-Party Cookies: %d", &currentStats.TotalThirdParty)
		} else if strings.HasPrefix(line, "Secure Domains: ") {
			fmt.Sscanf(line, "Secure Domains: %d", &currentStats.TotalSecure)
		} else if strings.HasPrefix(line, "HttpOnly: ") {
			fmt.Sscanf(line, "HttpOnly: %d", &currentStats.TotalHttpOnly)
		}
	}

	// Create simple rankings file
	outFile, err := os.Create("SIMPLE_RANKINGS.txt")
	if err != nil {
		fmt.Printf("Error creating SIMPLE_RANKINGS.txt: %v\n", err)
		return
	}
	defer outFile.Close()

	fmt.Fprintf(outFile, "=== SIMPLE BROWSER RANKINGS ===\n\n")

	// 1. Fewest Total Cookies
	browsers := []*BrowserStats{chromeStats, edgeStats, firefoxStats, safariStats}

	fmt.Fprintf(outFile, "1. FEWEST COOKIES (Better for Privacy):\n")

	// Sort by total cookies
	for i := 0; i < len(browsers)-1; i++ {
		for j := i + 1; j < len(browsers); j++ {
			if browsers[i].TotalCookies > browsers[j].TotalCookies {
				browsers[i], browsers[j] = browsers[j], browsers[i]
			}
		}
	}

	for i, browser := range browsers {
		fmt.Fprintf(outFile, "   %d. %s: %d cookies\n", i+1, strings.ToUpper(browser.Browser), browser.TotalCookies)
	}

	fmt.Fprintf(outFile, "\n2. FEWEST THIRD-PARTY COOKIES:\n")

	// Reset and sort by third-party cookies
	browsers = []*BrowserStats{chromeStats, edgeStats, firefoxStats, safariStats}
	for i := 0; i < len(browsers)-1; i++ {
		for j := i + 1; j < len(browsers); j++ {
			if browsers[i].TotalThirdParty > browsers[j].TotalThirdParty {
				browsers[i], browsers[j] = browsers[j], browsers[i]
			}
		}
	}

	for i, browser := range browsers {
		fmt.Fprintf(outFile, "   %d. %s: %d third-party\n", i+1, strings.ToUpper(browser.Browser), browser.TotalThirdParty)
	}

	fmt.Fprintf(outFile, "\n3. MOST SECURE COOKIES:\n")

	// Reset and sort by secure cookies (descending)
	browsers = []*BrowserStats{chromeStats, edgeStats, firefoxStats, safariStats}
	for i := 0; i < len(browsers)-1; i++ {
		for j := i + 1; j < len(browsers); j++ {
			if browsers[i].TotalSecure < browsers[j].TotalSecure {
				browsers[i], browsers[j] = browsers[j], browsers[i]
			}
		}
	}

	for i, browser := range browsers {
		fmt.Fprintf(outFile, "   %d. %s: %d secure\n", i+1, strings.ToUpper(browser.Browser), browser.TotalSecure)
	}

	// Privacy winner (fewest total + fewest third-party)
	fmt.Fprintf(outFile, "\n=== PRIVACY WINNER ===\n")
	browsers = []*BrowserStats{chromeStats, edgeStats, firefoxStats, safariStats}

	// Scoring: total cookies + (third-party * 2) = lower is better
	minScore := 999999
	winner := ""

	for _, browser := range browsers {
		score := browser.TotalCookies + (browser.TotalThirdParty * 2)
		if score < minScore {
			minScore = score
			winner = browser.Browser
		}
		fmt.Fprintf(outFile, "%s: %d total + %d third-party = %d points\n",
			strings.ToUpper(browser.Browser), browser.TotalCookies, browser.TotalThirdParty, score)
	}

	fmt.Fprintf(outFile, "\nWINNER: %s (lowest score = best privacy)\n", strings.ToUpper(winner))

	fmt.Println("Simple rankings saved to: SIMPLE_RANKINGS.txt")
}
