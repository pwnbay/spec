package program

import (
	"encoding/json"
	"fmt"
)

type Challenge struct {
	Type     ChallengeType          `json:"type"`
	Metadata map[string]interface{} `json:"metadata"`
}

type ChallengeType int

const (
	Flag ChallengeType = iota
	RandomFlag
	TestSuite
)

var ChallengeTypeNames = map[ChallengeType]string{
	Flag:       "flag",
	RandomFlag: "random-flag",
	TestSuite:  "test-suite",
}

func (t ChallengeType) String() string {
	return ChallengeTypeNames[t]
}

type FlagChallengeMetadata struct {
	Flag string `json:"flag"`
}

type RandomFlagChallengeMetadata struct {
	Seed int `json:"seed"`
}

type TestSuiteChallengeMetadata struct {
	Path string `json:"path"`
}

// MarshalJSON implements custom JSON marshaling for Challenge
func (c *Challenge) MarshalJSON() ([]byte, error) {
	// Create a temporary struct for marshaling
	type challengeAlias Challenge

	// Validate the challenge type
	if _, exists := ChallengeTypeNames[c.Type]; !exists {
		return nil, fmt.Errorf("invalid challenge type: %d", c.Type)
	}

	// Validate metadata based on type
	if err := c.validateMetadata(); err != nil {
		return nil, fmt.Errorf("invalid metadata for type %s: %w", c.Type.String(), err)
	}

	return json.Marshal((*challengeAlias)(c))
}

// UnmarshalJSON implements custom JSON unmarshaling for Challenge
func (c *Challenge) UnmarshalJSON(data []byte) error {
	// Create a temporary struct for unmarshaling
	type challengeAlias Challenge

	var temp challengeAlias
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Validate the challenge type
	if _, exists := ChallengeTypeNames[temp.Type]; !exists {
		return fmt.Errorf("invalid challenge type: %d", temp.Type)
	}

	// Convert metadata to proper struct based on type
	challenge := Challenge(temp)
	if err := challenge.convertMetadata(); err != nil {
		return fmt.Errorf("failed to convert metadata for type %s: %w", challenge.Type.String(), err)
	}

	*c = challenge
	return nil
}

// validateMetadata validates that the metadata matches the expected structure for the challenge type
func (c *Challenge) validateMetadata() error {
	switch c.Type {
	case Flag:
		if _, ok := c.Metadata["flag"]; !ok {
			return fmt.Errorf("flag challenge requires 'flag' field in metadata")
		}
		if flag, ok := c.Metadata["flag"].(string); !ok || flag == "" {
			return fmt.Errorf("flag challenge requires non-empty string 'flag' field")
		}

	case RandomFlag:
		if _, ok := c.Metadata["seed"]; !ok {
			return fmt.Errorf("random-flag challenge requires 'seed' field in metadata")
		}
		if _, ok := c.Metadata["seed"].(float64); !ok {
			return fmt.Errorf("random-flag challenge requires numeric 'seed' field")
		}

	case TestSuite:
		if _, ok := c.Metadata["path"]; !ok {
			return fmt.Errorf("test-suite challenge requires 'path' field in metadata")
		}
		if path, ok := c.Metadata["path"].(string); !ok || path == "" {
			return fmt.Errorf("test-suite challenge requires non-empty string 'path' field")
		}

	default:
		return fmt.Errorf("unknown challenge type: %d", c.Type)
	}

	return nil
}

// convertMetadata converts the generic metadata map to the appropriate typed struct
func (c *Challenge) convertMetadata() error {
	switch c.Type {
	case Flag:
		if flag, ok := c.Metadata["flag"].(string); ok {
			// Metadata is already in the correct format
			_ = flag // Use flag to avoid unused variable warning
		} else {
			return fmt.Errorf("invalid flag metadata format")
		}

	case RandomFlag:
		if seedFloat, ok := c.Metadata["seed"].(float64); ok {
			// Convert float64 to int (JSON numbers are unmarshaled as float64)
			c.Metadata["seed"] = int(seedFloat)
		} else {
			return fmt.Errorf("invalid random-flag metadata format")
		}

	case TestSuite:
		if path, ok := c.Metadata["path"].(string); ok {
			// Metadata is already in the correct format
			_ = path // Use path to avoid unused variable warning
		} else {
			return fmt.Errorf("invalid test-suite metadata format")
		}

	default:
		return fmt.Errorf("unknown challenge type: %d", c.Type)
	}

	return nil
}

// GetFlagMetadata returns the metadata as FlagChallengeMetadata if the type is Flag
func (c *Challenge) GetFlagMetadata() (*FlagChallengeMetadata, error) {
	if c.Type != Flag {
		return nil, fmt.Errorf("challenge is not of type Flag")
	}

	if err := c.validateMetadata(); err != nil {
		return nil, err
	}

	flag, _ := c.Metadata["flag"].(string)
	return &FlagChallengeMetadata{Flag: flag}, nil
}

// GetRandomFlagMetadata returns the metadata as RandomFlagChallengeMetadata if the type is RandomFlag
func (c *Challenge) GetRandomFlagMetadata() (*RandomFlagChallengeMetadata, error) {
	if c.Type != RandomFlag {
		return nil, fmt.Errorf("challenge is not of type RandomFlag")
	}

	if err := c.validateMetadata(); err != nil {
		return nil, err
	}

	seed, _ := c.Metadata["seed"].(int)
	return &RandomFlagChallengeMetadata{Seed: seed}, nil
}

// GetTestSuiteMetadata returns the metadata as TestSuiteChallengeMetadata if the type is TestSuite
func (c *Challenge) GetTestSuiteMetadata() (*TestSuiteChallengeMetadata, error) {
	if c.Type != TestSuite {
		return nil, fmt.Errorf("challenge is not of type TestSuite")
	}

	if err := c.validateMetadata(); err != nil {
		return nil, err
	}

	path, _ := c.Metadata["path"].(string)
	return &TestSuiteChallengeMetadata{Path: path}, nil
}
