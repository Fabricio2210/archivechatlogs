package elastic

import (
	"context"
	"strings"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type MyDocument struct {
	Name       string `json:"name"`
	Text       string `json:"text"`
	Hour       string `json:"hour"`
	Date 	   string `json:"date"`
	Subject    string `json:"subject"`
}

type Client struct {
	esClient *elasticsearch.Client
}

// NewClient creates a new Elasticsearch client and initializes the connection.
func NewClient(addresses []string) (*Client, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		esClient: esClient,
	}, nil
}

// IndexDocument indexes the provided document in Elasticsearch with an auto-generated document ID.
func (c *Client) IndexDocument(indexName string, doc MyDocument) error {
	// Serialize the document to JSON
	docJSON, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	// Create the index request
	req := esapi.IndexRequest{
		Index:   indexName,
		Refresh: "true",
		Body:    strings.NewReader(string(docJSON)), // Use a strings.NewReader to wrap the JSON data.
	}

	// Perform the request
	res, err := req.Do(context.Background(), c.esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}

	return nil
}
