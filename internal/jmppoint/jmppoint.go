package jmppoint
import (
	"privcrawler/internal/crawler"
)
var PORT int = 22

// ---- DATA STRUCTURES ---- //

// Options: Represents tags for the main method.
type ProcessOptions struct {
	// Functionality.
	browser string
	url     string
	hidden  bool
	duration int
	verbose bool
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
	safari = "safari"
)

func defaultProcessOptions() ProcessOptions {
	return ProcessOptions{
		browser: chrome,
		url:     "https://www.google.com",
		duration: 2000,
		hidden:  false,
		verbose: false,
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

