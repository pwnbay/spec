package program

import (
	"encoding/json"
	"fmt"
	"log"
)

// Example usage of the JSON marshaller and unmarshaller
func ExampleJSONMarshaling() {
	// Create a valid flag challenge
	flagChallenge := &Challenge{
		Type: Flag,
		Metadata: map[string]interface{}{
			"flag": "pwn{test_flag_123}",
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(flagChallenge)
	if err != nil {
		log.Fatalf("Failed to marshal flag challenge: %v", err)
	}
	fmt.Printf("Flag challenge JSON: %s\n", jsonData)

	// Unmarshal back
	var unmarshaledChallenge Challenge
	err = json.Unmarshal(jsonData, &unmarshaledChallenge)
	if err != nil {
		log.Fatalf("Failed to unmarshal flag challenge: %v", err)
	}

	// Get typed metadata
	flagMetadata, err := unmarshaledChallenge.GetFlagMetadata()
	if err != nil {
		log.Fatalf("Failed to get flag metadata: %v", err)
	}
	fmt.Printf("Flag: %s\n", flagMetadata.Flag)

	// Create a valid random flag challenge
	randomFlagChallenge := &Challenge{
		Type: RandomFlag,
		Metadata: map[string]interface{}{
			"seed": 42,
		},
	}

	// Marshal to JSON
	jsonData, err = json.Marshal(randomFlagChallenge)
	if err != nil {
		log.Fatalf("Failed to marshal random flag challenge: %v", err)
	}
	fmt.Printf("Random flag challenge JSON: %s\n", jsonData)

	// Create a valid file asset
	fileAsset := &Asset{
		Type: File,
		Metadata: map[string]interface{}{
			"path": "/path/to/file.txt",
		},
	}

	// Marshal to JSON
	jsonData, err = json.Marshal(fileAsset)
	if err != nil {
		log.Fatalf("Failed to marshal file asset: %v", err)
	}
	fmt.Printf("File asset JSON: %s\n", jsonData)

	// Example of invalid challenge (missing required metadata)
	invalidChallenge := &Challenge{
		Type:     Flag,
		Metadata: map[string]interface{}{
			// Missing "flag" field
		},
	}

	_, err = json.Marshal(invalidChallenge)
	if err != nil {
		fmt.Printf("Expected error for invalid challenge: %v\n", err)
	}

	// Example of invalid asset type
	invalidAsset := &Asset{
		Type: 999, // Invalid type
		Metadata: map[string]interface{}{
			"path": "/path/to/file.txt",
		},
	}

	_, err = json.Marshal(invalidAsset)
	if err != nil {
		fmt.Printf("Expected error for invalid asset type: %v\n", err)
	}
}

// TestJSONRoundTrip tests that marshaling and unmarshaling work correctly
func TestJSONRoundTrip() {
	// Test challenge round trip
	originalChallenge := &Challenge{
		Type: TestSuite,
		Metadata: map[string]interface{}{
			"path": "/tests/suite1",
		},
	}

	jsonData, err := json.Marshal(originalChallenge)
	if err != nil {
		log.Fatalf("Failed to marshal challenge: %v", err)
	}

	var roundTripChallenge Challenge
	err = json.Unmarshal(jsonData, &roundTripChallenge)
	if err != nil {
		log.Fatalf("Failed to unmarshal challenge: %v", err)
	}

	if roundTripChallenge.Type != originalChallenge.Type {
		log.Fatalf("Type mismatch: expected %d, got %d", originalChallenge.Type, roundTripChallenge.Type)
	}

	// Test asset round trip
	originalAsset := &Asset{
		Type: DockerCompose,
		Metadata: map[string]interface{}{
			"path": "/docker-compose.yml",
		},
	}

	jsonData, err = json.Marshal(originalAsset)
	if err != nil {
		log.Fatalf("Failed to marshal asset: %v", err)
	}

	var roundTripAsset Asset
	err = json.Unmarshal(jsonData, &roundTripAsset)
	if err != nil {
		log.Fatalf("Failed to unmarshal asset: %v", err)
	}

	if roundTripAsset.Type != originalAsset.Type {
		log.Fatalf("Type mismatch: expected %d, got %d", originalAsset.Type, roundTripAsset.Type)
	}

	fmt.Println("JSON round trip tests passed!")
}
