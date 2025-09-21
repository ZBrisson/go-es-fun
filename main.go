package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	// Get configuration from environment variables
	cloudID := os.Getenv("ELASTIC_CLOUD_ID")
	apiKey := os.Getenv("ELASTIC_API_KEY")

	if cloudID == "" || apiKey == "" {
		log.Fatal("ELASTIC_CLOUD_ID and ELASTIC_API_KEY environment variables must be set")
	}

	// Create Elasticsearch client configuration
	cfg := elasticsearch.Config{
		CloudID: cloudID,
		APIKey:  apiKey,
	}

	// Create the client
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	// Check cluster health
	healthData, status, err := checkClusterHealth(client)
	if err != nil {
		log.Fatalf("Error checking cluster health: %v", err)
	}

	// Display results
	displayHealthStatus(healthData, status)
}

func checkClusterHealth(client *elasticsearch.Client) (map[string]interface{}, string, error) {
	ctx := context.Background()

	// Make cluster health API request
	res, err := client.Cluster.Health(
		client.Cluster.Health.WithContext(ctx),
		client.Cluster.Health.WithPretty(),
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get cluster health: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, "", fmt.Errorf("cluster health API returned error: %s", res.Status())
	}

	// Parse response into generic map
	var healthData map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&healthData); err != nil {
		return nil, "", fmt.Errorf("failed to decode response: %w", err)
	}

	// Extract status for special handling
	status, ok := healthData["status"].(string)
	if !ok {
		status = "unknown"
	}

	return healthData, status, nil
}

func displayHealthStatus(healthData map[string]interface{}, status string) {
	fmt.Printf("=== Elasticsearch Cluster Health ===\n")

	// Handle status specially with color coding
	fmt.Printf("Status: %s\n", getStatusWithColor(status))

	// Get all keys except status and sort them for consistent output
	var keys []string
	for key := range healthData {
		if key != "status" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	// Display all other fields dynamically
	for _, key := range keys {
		value := healthData[key]
		displayField(key, value)
	}

	// Health status interpretation
	fmt.Printf("\n=== Health Status Interpretation ===\n")
	switch status {
	case "green":
		fmt.Printf("✅ All shards are allocated and the cluster is fully operational\n")
	case "yellow":
		fmt.Printf("⚠️  All primary shards are allocated, but some replica shards are not\n")
	case "red":
		fmt.Printf("❌ Some primary shards are not allocated\n")
	default:
		fmt.Printf("❓ Unknown status: %s\n", status)
	}

	// Check for unassigned shards warning
	if unassigned, exists := healthData["unassigned_shards"]; exists {
		if unassignedCount, ok := unassigned.(float64); ok && unassignedCount > 0 {
			fmt.Printf("⚠️  Warning: %.0f unassigned shards detected\n", unassignedCount)
		}
	}
}

func displayField(key string, value interface{}) {
	// Convert snake_case to Title Case for display
	displayKey := formatFieldName(key)

	// Handle different value types
	switch v := value.(type) {
	case string:
		fmt.Printf("%s: %s\n", displayKey, v)
	case float64:
		// Check if it's actually an integer
		if v == float64(int64(v)) {
			fmt.Printf("%s: %.0f\n", displayKey, v)
		} else {
			fmt.Printf("%s: %.2f\n", displayKey, v)
		}
	case bool:
		fmt.Printf("%s: %t\n", displayKey, v)
	case nil:
		fmt.Printf("%s: <null>\n", displayKey)
	default:
		// Handle complex types (arrays, objects) by converting to JSON
		if jsonBytes, err := json.Marshal(v); err == nil {
			fmt.Printf("%s: %s\n", displayKey, string(jsonBytes))
		} else {
			fmt.Printf("%s: %v\n", displayKey, v)
		}
	}
}

func formatFieldName(fieldName string) string {
	// Convert snake_case to Title Case
	words := strings.Split(fieldName, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func getStatusWithColor(status string) string {
	switch status {
	case "green":
		return "🟢 " + status
	case "yellow":
		return "🟡 " + status
	case "red":
		return "🔴 " + status
	default:
		return status
	}
}
