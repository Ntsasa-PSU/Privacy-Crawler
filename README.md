# Privacy-Crawler


## ***-- Crawler --*** 

### Compile: Create 'main' executable.
    - go build \maincrawler\main-crawler.go

### Run: Execute application with 'main' executable.
    - .\maincrawler\main-crawler.go

### Run: Execute without executable {Supported through Go}.
    - go run .\maincrawler\main-crawler.go


### Tags: Toggle options.
    - '-v' Verbose Mode
    - '-b' Browser Selection
    - '-i' true hides browser, false reveals browser
    - '-u' Choose desired url, e.g https://www.example.com
    - '-t' Prototype test flag

## ***-- Jump Point --***

### Run: Execute application with 'main' executable.
    - go run .\mainjmppoint\main-jmppoint.go

### URLs: Commands ran through http.
- Run: Standard Process.
    - http://localhost:8080/run

- Run w/ Tags: Allows tags to be executed along with process.
    - http://localhost:8080/run?verbose=v
    - Valid arguments: {broswer, verbose}

- Run: Testing user input mode.
    - http://localhost:8080/test

## ***How to***:
    
    - Run the Jump Point application. Once it is listening on port 8080, connect to
      URL and run process from jmppoint instead of local.

## **_-- How to run with PlayWright --_**

### Install Playwright and browser binaries

- go get github.com/playwright-community/playwright-go
- go run github.com/playwright-community/playwright-go/cmd/playwright install
