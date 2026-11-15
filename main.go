package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"searAI/internal/aggregate"
	"searAI/internal/config"
)

func main() {
	// Load configuration from TOML file
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	timeout := flag.Duration("timeout", 5*time.Second, "Search timeout duration")
	asJSON := true
	flag.Parse()

	// Load enabled engines based on configuration
	engs := cfg.LoadEnabledEngines()
	if len(engs) == 0 {
		log.Fatal("No search engines enabled in configuration")
	}

	for {
		fmt.Print("Enter your search query: ")
		reader := bufio.NewReader(os.Stdin)

		// Read everything until the user presses Enter ('\n')
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Clean up whitespace and hidden characters
		query := strings.TrimSpace(input)
		//timeout context
		ctx, cancel := context.WithTimeout(context.Background(), *timeout)
		defer cancel()

		fmt.Printf("🔍 Searching for %q...\n\n", query)

		// Run the aggregator directly
		// Parameters: ctx, engines, query, timeout, page, pageSize
		results, timings := aggregate.Run(ctx, engs, query, *timeout, 1, 30)

		//Output Formatting
		fmt.Printf("===============================================================================\n")
		if asJSON {
			// Pretty-print as JSON if -json flag is passed
			jsonData, err := json.MarshalIndent(results, "", "  ")
			if err != nil {
				fmt.Printf("Failed to marshal results: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(jsonData))
			fmt.Printf("------------------------------------------------------------------------------\n")

			// Optional: Print timing stats
			fmt.Println("--- Performance Timings ---")
			for _, t := range timings {
				fmt.Printf("Engine: %s | Took: %v | Error: %v\n", t.Engine, t.Ms, t.Error)
			}
			fmt.Printf("\n===============================================================================\n\n")
		}
	}
}
