package jmppoint

import (
	"privcrawler/internal/crawler"
	"sync"
)

var PORT int = 22

// ---- DATA STRUCTURES ---- //

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
	var wg sync.WaitGroup

	// Count total processes: 9 websites × 4 browsers × 6 durations = 216 processes
	wg.Add(216)

	// ---- Start all processes immediately ---- //

	// -- WEBSITE: AMAZON -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_amazon_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(0))
		chrome_amazon_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_amazon_0 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(0))
		safari_amazon_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_amazon_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(0))
		firefox_amazon_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_amazon_0 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(0))
		edge_amazon_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_amazon_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(5000))
		chrome_amazon_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_amazon_5 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(5000))
		safari_amazon_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_amazon_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(5000))
		firefox_amazon_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_amazon_5 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(5000))
		edge_amazon_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_amazon_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(10000))
		chrome_amazon_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_amazon_10 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(10000))
		safari_amazon_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_amazon_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(10000))
		firefox_amazon_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_amazon_10 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(10000))
		edge_amazon_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_amazon_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(15000))
		chrome_amazon_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_amazon_15 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(15000))
		safari_amazon_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_amazon_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(15000))
		firefox_amazon_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_amazon_15 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(15000))
		edge_amazon_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_amazon_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(20000))
		chrome_amazon_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_amazon_20 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(20000))
		safari_amazon_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_amazon_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(20000))
		firefox_amazon_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_amazon_20 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(20000))
		edge_amazon_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_amazon_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(25000))
		chrome_amazon_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_amazon_25 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(25000))
		safari_amazon_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_amazon_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(25000))
		firefox_amazon_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_amazon_25 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(25000))
		edge_amazon_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_amazon_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.amazon.com"), WithDuration(30000))
		chrome_amazon_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_amazon_30 := NewProcess(WithBrowser(safari), WithURL("https://www.amazon.com"), WithDuration(30000))
		safari_amazon_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_amazon_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.amazon.com"), WithDuration(30000))
		firefox_amazon_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_amazon_30 := NewProcess(WithBrowser(edge), WithURL("https://www.amazon.com"), WithDuration(30000))
		edge_amazon_30.Run()
	}()

	// -- WEBSITE: YAHOO -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_yahoo_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(0))
		chrome_yahoo_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_yahoo_0 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(0))
		safari_yahoo_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_yahoo_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(0))
		firefox_yahoo_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_yahoo_0 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(0))
		edge_yahoo_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_yahoo_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(5000))
		chrome_yahoo_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_yahoo_5 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(5000))
		safari_yahoo_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_yahoo_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(5000))
		firefox_yahoo_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_yahoo_5 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(5000))
		edge_yahoo_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_yahoo_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(10000))
		chrome_yahoo_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_yahoo_10 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(10000))
		safari_yahoo_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_yahoo_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(10000))
		firefox_yahoo_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_yahoo_10 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(10000))
		edge_yahoo_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_yahoo_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(15000))
		chrome_yahoo_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_yahoo_15 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(15000))
		safari_yahoo_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_yahoo_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(15000))
		firefox_yahoo_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_yahoo_15 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(15000))
		edge_yahoo_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_yahoo_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(20000))
		chrome_yahoo_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_yahoo_20 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(20000))
		safari_yahoo_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_yahoo_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(20000))
		firefox_yahoo_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_yahoo_20 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(20000))
		edge_yahoo_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_yahoo_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(25000))
		chrome_yahoo_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_yahoo_25 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(25000))
		safari_yahoo_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_yahoo_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(25000))
		firefox_yahoo_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_yahoo_25 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(25000))
		edge_yahoo_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_yahoo_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.yahoo.com"), WithDuration(30000))
		chrome_yahoo_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_yahoo_30 := NewProcess(WithBrowser(safari), WithURL("https://www.yahoo.com"), WithDuration(30000))
		safari_yahoo_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_yahoo_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.yahoo.com"), WithDuration(30000))
		firefox_yahoo_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_yahoo_30 := NewProcess(WithBrowser(edge), WithURL("https://www.yahoo.com"), WithDuration(30000))
		edge_yahoo_30.Run()
	}()

	// -- WEBSITE: REDDIT -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_reddit_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(0))
		chrome_reddit_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_reddit_0 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(0))
		safari_reddit_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_reddit_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(0))
		firefox_reddit_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_reddit_0 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(0))
		edge_reddit_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_reddit_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(5000))
		chrome_reddit_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_reddit_5 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(5000))
		safari_reddit_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_reddit_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(5000))
		firefox_reddit_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_reddit_5 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(5000))
		edge_reddit_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_reddit_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(10000))
		chrome_reddit_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_reddit_10 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(10000))
		safari_reddit_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_reddit_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(10000))
		firefox_reddit_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_reddit_10 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(10000))
		edge_reddit_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_reddit_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(15000))
		chrome_reddit_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_reddit_15 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(15000))
		safari_reddit_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_reddit_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(15000))
		firefox_reddit_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_reddit_15 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(15000))
		edge_reddit_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_reddit_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(20000))
		chrome_reddit_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_reddit_20 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(20000))
		safari_reddit_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_reddit_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(20000))
		firefox_reddit_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_reddit_20 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(20000))
		edge_reddit_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_reddit_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(25000))
		chrome_reddit_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_reddit_25 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(25000))
		safari_reddit_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_reddit_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(25000))
		firefox_reddit_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_reddit_25 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(25000))
		edge_reddit_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_reddit_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.reddit.com"), WithDuration(30000))
		chrome_reddit_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_reddit_30 := NewProcess(WithBrowser(safari), WithURL("https://www.reddit.com"), WithDuration(30000))
		safari_reddit_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_reddit_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.reddit.com"), WithDuration(30000))
		firefox_reddit_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_reddit_30 := NewProcess(WithBrowser(edge), WithURL("https://www.reddit.com"), WithDuration(30000))
		edge_reddit_30.Run()
	}()

	// -- WEBSITE: PINTEREST -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_pinterest_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(0))
		chrome_pinterest_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pinterest_0 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(0))
		safari_pinterest_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pinterest_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(0))
		firefox_pinterest_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pinterest_0 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(0))
		edge_pinterest_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_pinterest_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(5000))
		chrome_pinterest_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pinterest_5 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(5000))
		safari_pinterest_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pinterest_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(5000))
		firefox_pinterest_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pinterest_5 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(5000))
		edge_pinterest_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_pinterest_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(10000))
		chrome_pinterest_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pinterest_10 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(10000))
		safari_pinterest_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pinterest_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(10000))
		firefox_pinterest_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pinterest_10 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(10000))
		edge_pinterest_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_pinterest_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(15000))
		chrome_pinterest_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pinterest_15 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(15000))
		safari_pinterest_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pinterest_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(15000))
		firefox_pinterest_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pinterest_15 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(15000))
		edge_pinterest_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_pinterest_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(20000))
		chrome_pinterest_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pinterest_20 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(20000))
		safari_pinterest_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pinterest_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(20000))
		firefox_pinterest_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pinterest_20 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(20000))
		edge_pinterest_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_pinterest_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(25000))
		chrome_pinterest_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pinterest_25 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(25000))
		safari_pinterest_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pinterest_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(25000))
		firefox_pinterest_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pinterest_25 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(25000))
		edge_pinterest_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_pinterest_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.pinterest.com"), WithDuration(30000))
		chrome_pinterest_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pinterest_30 := NewProcess(WithBrowser(safari), WithURL("https://www.pinterest.com"), WithDuration(30000))
		safari_pinterest_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pinterest_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.pinterest.com"), WithDuration(30000))
		firefox_pinterest_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pinterest_30 := NewProcess(WithBrowser(edge), WithURL("https://www.pinterest.com"), WithDuration(30000))
		edge_pinterest_30.Run()
	}()

	// -- WEBSITE: FANDOM -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_fandom_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(0))
		chrome_fandom_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fandom_0 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(0))
		safari_fandom_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fandom_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(0))
		firefox_fandom_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fandom_0 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(0))
		edge_fandom_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_fandom_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(5000))
		chrome_fandom_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fandom_5 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(5000))
		safari_fandom_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fandom_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(5000))
		firefox_fandom_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fandom_5 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(5000))
		edge_fandom_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_fandom_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(10000))
		chrome_fandom_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fandom_10 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(10000))
		safari_fandom_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fandom_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(10000))
		firefox_fandom_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fandom_10 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(10000))
		edge_fandom_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_fandom_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(15000))
		chrome_fandom_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fandom_15 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(15000))
		safari_fandom_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fandom_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(15000))
		firefox_fandom_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fandom_15 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(15000))
		edge_fandom_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_fandom_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(20000))
		chrome_fandom_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fandom_20 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(20000))
		safari_fandom_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fandom_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(20000))
		firefox_fandom_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fandom_20 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(20000))
		edge_fandom_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_fandom_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(25000))
		chrome_fandom_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fandom_25 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(25000))
		safari_fandom_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fandom_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(25000))
		firefox_fandom_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fandom_25 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(25000))
		edge_fandom_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_fandom_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.fandom.com"), WithDuration(30000))
		chrome_fandom_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fandom_30 := NewProcess(WithBrowser(safari), WithURL("https://www.fandom.com"), WithDuration(30000))
		safari_fandom_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fandom_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.fandom.com"), WithDuration(30000))
		firefox_fandom_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fandom_30 := NewProcess(WithBrowser(edge), WithURL("https://www.fandom.com"), WithDuration(30000))
		edge_fandom_30.Run()
	}()

	// -- WEBSITE: PORNHUB -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_pornhub_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(0))
		chrome_pornhub_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pornhub_0 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(0))
		safari_pornhub_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pornhub_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(0))
		firefox_pornhub_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pornhub_0 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(0))
		edge_pornhub_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_pornhub_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(5000))
		chrome_pornhub_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pornhub_5 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(5000))
		safari_pornhub_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pornhub_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(5000))
		firefox_pornhub_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pornhub_5 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(5000))
		edge_pornhub_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_pornhub_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(10000))
		chrome_pornhub_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pornhub_10 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(10000))
		safari_pornhub_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pornhub_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(10000))
		firefox_pornhub_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pornhub_10 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(10000))
		edge_pornhub_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_pornhub_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(15000))
		chrome_pornhub_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pornhub_15 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(15000))
		safari_pornhub_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pornhub_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(15000))
		firefox_pornhub_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pornhub_15 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(15000))
		edge_pornhub_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_pornhub_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(20000))
		chrome_pornhub_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pornhub_20 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(20000))
		safari_pornhub_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pornhub_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(20000))
		firefox_pornhub_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pornhub_20 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(20000))
		edge_pornhub_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_pornhub_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(25000))
		chrome_pornhub_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pornhub_25 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(25000))
		safari_pornhub_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pornhub_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(25000))
		firefox_pornhub_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pornhub_25 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(25000))
		edge_pornhub_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_pornhub_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.pornhub.com"), WithDuration(30000))
		chrome_pornhub_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_pornhub_30 := NewProcess(WithBrowser(safari), WithURL("https://www.pornhub.com"), WithDuration(30000))
		safari_pornhub_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_pornhub_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.pornhub.com"), WithDuration(30000))
		firefox_pornhub_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_pornhub_30 := NewProcess(WithBrowser(edge), WithURL("https://www.pornhub.com"), WithDuration(30000))
		edge_pornhub_30.Run()
	}()

	// -- WEBSITE: SOAP2DAY -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_soap2day_0 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(0))
		chrome_soap2day_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_soap2day_0 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(0))
		safari_soap2day_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_soap2day_0 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(0))
		firefox_soap2day_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_soap2day_0 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(0))
		edge_soap2day_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_soap2day_5 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(5000))
		chrome_soap2day_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_soap2day_5 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(5000))
		safari_soap2day_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_soap2day_5 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(5000))
		firefox_soap2day_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_soap2day_5 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(5000))
		edge_soap2day_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_soap2day_10 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(10000))
		chrome_soap2day_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_soap2day_10 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(10000))
		safari_soap2day_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_soap2day_10 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(10000))
		firefox_soap2day_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_soap2day_10 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(10000))
		edge_soap2day_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_soap2day_15 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(15000))
		chrome_soap2day_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_soap2day_15 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(15000))
		safari_soap2day_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_soap2day_15 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(15000))
		firefox_soap2day_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_soap2day_15 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(15000))
		edge_soap2day_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_soap2day_20 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(20000))
		chrome_soap2day_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_soap2day_20 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(20000))
		safari_soap2day_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_soap2day_20 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(20000))
		firefox_soap2day_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_soap2day_20 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(20000))
		edge_soap2day_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_soap2day_25 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(25000))
		chrome_soap2day_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_soap2day_25 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(25000))
		safari_soap2day_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_soap2day_25 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(25000))
		firefox_soap2day_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_soap2day_25 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(25000))
		edge_soap2day_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_soap2day_30 := NewProcess(WithBrowser(chrome), WithURL("https://soap2day.to"), WithDuration(30000))
		chrome_soap2day_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_soap2day_30 := NewProcess(WithBrowser(safari), WithURL("https://soap2day.to"), WithDuration(30000))
		safari_soap2day_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_soap2day_30 := NewProcess(WithBrowser(firefox), WithURL("https://soap2day.to"), WithDuration(30000))
		firefox_soap2day_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_soap2day_30 := NewProcess(WithBrowser(edge), WithURL("https://soap2day.to"), WithDuration(30000))
		edge_soap2day_30.Run()
	}()

	// -- WEBSITE: XVIDEOS -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_xvideos_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(0))
		chrome_xvideos_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_xvideos_0 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(0))
		safari_xvideos_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_xvideos_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(0))
		firefox_xvideos_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_xvideos_0 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(0))
		edge_xvideos_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_xvideos_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(5000))
		chrome_xvideos_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_xvideos_5 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(5000))
		safari_xvideos_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_xvideos_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(5000))
		firefox_xvideos_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_xvideos_5 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(5000))
		edge_xvideos_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_xvideos_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(10000))
		chrome_xvideos_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_xvideos_10 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(10000))
		safari_xvideos_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_xvideos_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(10000))
		firefox_xvideos_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_xvideos_10 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(10000))
		edge_xvideos_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_xvideos_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(15000))
		chrome_xvideos_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_xvideos_15 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(15000))
		safari_xvideos_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_xvideos_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(15000))
		firefox_xvideos_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_xvideos_15 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(15000))
		edge_xvideos_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_xvideos_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(20000))
		chrome_xvideos_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_xvideos_20 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(20000))
		safari_xvideos_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_xvideos_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(20000))
		firefox_xvideos_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_xvideos_20 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(20000))
		edge_xvideos_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_xvideos_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(25000))
		chrome_xvideos_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_xvideos_25 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(25000))
		safari_xvideos_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_xvideos_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(25000))
		firefox_xvideos_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_xvideos_25 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(25000))
		edge_xvideos_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_xvideos_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.xvideos.com"), WithDuration(30000))
		chrome_xvideos_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_xvideos_30 := NewProcess(WithBrowser(safari), WithURL("https://www.xvideos.com"), WithDuration(30000))
		safari_xvideos_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_xvideos_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.xvideos.com"), WithDuration(30000))
		firefox_xvideos_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_xvideos_30 := NewProcess(WithBrowser(edge), WithURL("https://www.xvideos.com"), WithDuration(30000))
		edge_xvideos_30.Run()
	}()

	// -- WEBSITE: FASTDOWNLOAD -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_fastdownload_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(0))
		chrome_fastdownload_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fastdownload_0 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(0))
		safari_fastdownload_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fastdownload_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(0))
		firefox_fastdownload_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fastdownload_0 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(0))
		edge_fastdownload_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_fastdownload_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		chrome_fastdownload_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fastdownload_5 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		safari_fastdownload_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fastdownload_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		firefox_fastdownload_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fastdownload_5 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(5000))
		edge_fastdownload_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_fastdownload_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		chrome_fastdownload_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fastdownload_10 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		safari_fastdownload_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fastdownload_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		firefox_fastdownload_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fastdownload_10 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(10000))
		edge_fastdownload_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_fastdownload_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		chrome_fastdownload_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fastdownload_15 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		safari_fastdownload_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fastdownload_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		firefox_fastdownload_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fastdownload_15 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(15000))
		edge_fastdownload_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_fastdownload_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		chrome_fastdownload_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fastdownload_20 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		safari_fastdownload_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fastdownload_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		firefox_fastdownload_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fastdownload_20 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(20000))
		edge_fastdownload_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_fastdownload_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		chrome_fastdownload_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fastdownload_25 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		safari_fastdownload_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fastdownload_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		firefox_fastdownload_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fastdownload_25 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(25000))
		edge_fastdownload_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_fastdownload_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.fastdownload.com"), WithDuration(30000))
		chrome_fastdownload_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_fastdownload_30 := NewProcess(WithBrowser(safari), WithURL("https://www.fastdownload.com"), WithDuration(30000))
		safari_fastdownload_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_fastdownload_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.fastdownload.com"), WithDuration(30000))
		firefox_fastdownload_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_fastdownload_30 := NewProcess(WithBrowser(edge), WithURL("https://www.fastdownload.com"), WithDuration(30000))
		edge_fastdownload_30.Run()
	}()

	// -- WEBSITE: ENDACE -- //
	// - Instant - //
	go func() {
		defer wg.Done()
		chrome_endace_0 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(0))
		chrome_endace_0.Run()
	}()
	go func() {
		defer wg.Done()
		safari_endace_0 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(0))
		safari_endace_0.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_endace_0 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(0))
		firefox_endace_0.Run()
	}()
	go func() {
		defer wg.Done()
		edge_endace_0 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(0))
		edge_endace_0.Run()
	}()

	// - Delayed: 5 seconds - //
	go func() {
		defer wg.Done()
		chrome_endace_5 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(5000))
		chrome_endace_5.Run()
	}()
	go func() {
		defer wg.Done()
		safari_endace_5 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(5000))
		safari_endace_5.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_endace_5 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(5000))
		firefox_endace_5.Run()
	}()
	go func() {
		defer wg.Done()
		edge_endace_5 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(5000))
		edge_endace_5.Run()
	}()

	// - Delayed: 10 seconds - //
	go func() {
		defer wg.Done()
		chrome_endace_10 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(10000))
		chrome_endace_10.Run()
	}()
	go func() {
		defer wg.Done()
		safari_endace_10 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(10000))
		safari_endace_10.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_endace_10 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(10000))
		firefox_endace_10.Run()
	}()
	go func() {
		defer wg.Done()
		edge_endace_10 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(10000))
		edge_endace_10.Run()
	}()

	// - Delayed: 15 seconds - //
	go func() {
		defer wg.Done()
		chrome_endace_15 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(15000))
		chrome_endace_15.Run()
	}()
	go func() {
		defer wg.Done()
		safari_endace_15 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(15000))
		safari_endace_15.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_endace_15 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(15000))
		firefox_endace_15.Run()
	}()
	go func() {
		defer wg.Done()
		edge_endace_15 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(15000))
		edge_endace_15.Run()
	}()

	// - Delayed: 20 seconds - //
	go func() {
		defer wg.Done()
		chrome_endace_20 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(20000))
		chrome_endace_20.Run()
	}()
	go func() {
		defer wg.Done()
		safari_endace_20 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(20000))
		safari_endace_20.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_endace_20 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(20000))
		firefox_endace_20.Run()
	}()
	go func() {
		defer wg.Done()
		edge_endace_20 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(20000))
		edge_endace_20.Run()
	}()

	// - Delayed: 25 seconds - //
	go func() {
		defer wg.Done()
		chrome_endace_25 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(25000))
		chrome_endace_25.Run()
	}()
	go func() {
		defer wg.Done()
		safari_endace_25 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(25000))
		safari_endace_25.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_endace_25 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(25000))
		firefox_endace_25.Run()
	}()
	go func() {
		defer wg.Done()
		edge_endace_25 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(25000))
		edge_endace_25.Run()
	}()

	// - Delayed: 30 seconds - //
	go func() {
		defer wg.Done()
		chrome_endace_30 := NewProcess(WithBrowser(chrome), WithURL("https://www.endace.com"), WithDuration(30000))
		chrome_endace_30.Run()
	}()
	go func() {
		defer wg.Done()
		safari_endace_30 := NewProcess(WithBrowser(safari), WithURL("https://www.endace.com"), WithDuration(30000))
		safari_endace_30.Run()
	}()
	go func() {
		defer wg.Done()
		firefox_endace_30 := NewProcess(WithBrowser(firefox), WithURL("https://www.endace.com"), WithDuration(30000))
		firefox_endace_30.Run()
	}()
	go func() {
		defer wg.Done()
		edge_endace_30 := NewProcess(WithBrowser(edge), WithURL("https://www.endace.com"), WithDuration(30000))
		edge_endace_30.Run()
	}()

	// Wait for all 216 processes to complete
	wg.Wait()

	return nil
}
