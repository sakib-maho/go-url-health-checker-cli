package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sakib-maho/go-url-health-checker-cli/internal/checker"
	"github.com/sakib-maho/go-url-health-checker-cli/internal/output"
	"github.com/sakib-maho/go-url-health-checker-cli/internal/parse"
)

func main() {
	urlsRaw := flag.String("urls", "", "Comma-separated URL list")
	filePath := flag.String("file", "", "Path to newline-separated URL file")
	timeoutSec := flag.Int("timeout", 5, "Timeout in seconds")
	retries := flag.Int("retries", 1, "Retries for each URL")
	jsonOut := flag.String("json-out", "", "Optional JSON output path")
	csvOut := flag.String("csv-out", "", "Optional CSV output path")
	flag.Parse()

	var urls []string
	if *filePath != "" {
		parsed, err := parse.ParseURLsFromFile(*filePath)
		if err != nil {
			exitWithError(err)
		}
		urls = append(urls, parsed...)
	}
	if *urlsRaw != "" {
		urls = append(urls, parse.ParseURLs(*urlsRaw)...)
	}
	if len(urls) == 0 {
		exitWithError(fmt.Errorf("no URLs provided; use --urls or --file"))
	}

	results := checker.CheckURLs(urls, time.Duration(*timeoutSec)*time.Second, *retries)
	output.PrintTable(results)

	if *jsonOut != "" {
		if err := output.WriteJSON(*jsonOut, results); err != nil {
			exitWithError(err)
		}
		fmt.Printf("JSON written to %s\n", *jsonOut)
	}
	if *csvOut != "" {
		if err := output.WriteCSV(*csvOut, results); err != nil {
			exitWithError(err)
		}
		fmt.Printf("CSV written to %s\n", *csvOut)
	}
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
