package savedata

import (
	"context"
	"fmt"
	"github.com/Fabricio2210/dateFormat"
	"github.com/Fabricio2210/elastic"
	"github.com/Fabricio2210/parseString"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

func Savedata(fileName string, subject string) error {
	// Replace the Elasticsearch server addresses with your actual Elasticsearch server configuration.
	elasticAddresses := []string{"http://localhost:9200"}

	// Create a new Elasticsearch client.
	client, err := elastic.NewClient(elasticAddresses)
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch client: %s", err)
	}

	// Read the text file
	data, err := ioutil.ReadFile(fmt.Sprintf("./%s/%s", subject, fileName))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Split the data into lines
	lines := strings.Split(string(data), "\n")

	// Set the maximum number of concurrent writes
	maxConcurrency := 100
	// Create a channel to control the concurrency
	concurrencyCh := make(chan struct{}, maxConcurrency)

	// Create a wait group to track the completion of all writes
	var wg sync.WaitGroup

	// Iterate over the lines and save them concurrently with controlled concurrency
	for _, line := range lines {
		// If the line is empty skip it
		if line == "" {
			continue
		}
		wg.Add(1)
		concurrencyCh <- struct{}{} // Acquire a concurrency slot
		fmt.Println(line)
		go func(l string) {
			defer func() {
				<-concurrencyCh // Release the concurrency slot
				wg.Done()
			}()
			// Parse the string data into a structured format
			data := parsestring.Parsestring(l)
			// Change the date format of the parsed data
			parsedDate := changedateformat.ChangeDateFormat(data.Date)
			// Create a custom document for Elasticsearch indexing
			doc := elastic.MyDocument{
				Name:    data.Name,
				Text:    data.Message,
				Hour:    data.Hour,
				Date:    parsedDate,
				Subject: subject,
			}
			// Index the document in Elasticsearch
			err = client.IndexDocument("logschemas", doc)
			if err != nil {
				log.Fatalf("Failed to index document: %s", err)
			}
			// Create a context with a timeout of 30 seconds for each write operation
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if ctx.Err() != nil {
				fmt.Println("Context canceled:", ctx.Err())
				return
			}
		}(line)
	}

	// Wait for all writes to complete
	wg.Wait()

	fmt.Println("All lines saved.")

	return nil
}

