package spec

import (
	"encoding/json"
	"fmt"
)

type ChallengeMetadata interface {
	Type() ChallengeType
}

type Challenge struct {
	Type     ChallengeType     `json:"type"`
	Metadata ChallengeMetadata `json:"metadata"`
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

func (t ChallengeType) toString() (string, error) {
	if t < 0 || int(t) >= len(ChallengeTypeNames) {
		return "", fmt.Errorf("invalid challenge type: %d", t)
	}

	return ChallengeTypeNames[t], nil
}

func toChallengeType(s string) (ChallengeType, error) {
	for chalType, typeName := range ChallengeTypeNames {
		if s == typeName {
			return chalType, nil
		}
	}

	return -1, fmt.Errorf("invalid challenge type: %s", s)
}

type FlagChallengeMetadata struct {
	Flag string `json:"flag"`
}

func (f FlagChallengeMetadata) Type() ChallengeType {
	return Flag
}

type RandomFlagChallengeMetadata struct {
	Seed int `json:"seed"`
}

func (r RandomFlagChallengeMetadata) Type() ChallengeType {
	return RandomFlag
}

type TestSuiteChallengeMetadata struct {
	Path string `json:"path"`
}

func (t TestSuiteChallengeMetadata) Type() ChallengeType {
	return TestSuite
}

func (c *Challenge) MarshalJSON() ([]byte, error) {
	type MarshalledChallenge struct {
		Type     string      `json:"type"`
		Metadata interface{} `json:"metadata"`
	}

	chalType, err := c.Type.toString()
	if err != nil {
		return nil, err
	}

	temp := MarshalledChallenge{
		Type:     chalType,
		Metadata: c.Metadata,
	}

	return json.Marshal(temp)
}

func (c *Challenge) UnmarshalJSON(data []byte) error {
	type UnmarshalledChallenge struct {
		Type     string          `json:"type"`
		Metadata json.RawMessage `json:"metadata"`
	}

	var temp UnmarshalledChallenge

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	chalType, err := toChallengeType(temp.Type)
	if err != nil {
		return fmt.Errorf("invalid challenge type: %s", temp.Type)
	}

	c.Type = chalType

	switch chalType {
	case Flag:
		var metadata FlagChallengeMetadata
		if err := json.Unmarshal(temp.Metadata, &metadata); err != nil {
			return fmt.Errorf("failed to unmarshal challenge metadata: %w", err)
		}
		c.Metadata = metadata
	case RandomFlag:
		var metadata RandomFlagChallengeMetadata
		if err := json.Unmarshal(temp.Metadata, &metadata); err != nil {
			return fmt.Errorf("failed to unmarshal challenge metadata: %w", err)
		}
		c.Metadata = metadata
	case TestSuite:
		var metadata TestSuiteChallengeMetadata
		if err := json.Unmarshal(temp.Metadata, &metadata); err != nil {
			return fmt.Errorf("failed to unmarshal challenge metadata: %w", err)
		}
		c.Metadata = metadata
	default:
		return fmt.Errorf("unimplemented challenge type unmarshaller: %s", temp.Type)
	}

	return nil
}
