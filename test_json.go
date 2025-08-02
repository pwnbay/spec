package main

import (
	"encoding/json"
	"log"
	"log/slog"

	"github.com/pwnbay/spec/program"
)

// Example usage of the JSON marshaller and unmarshaller
func main() {

	fromProg := program.Program{
		Name: "test",
		Assets: []program.Asset{
			{
				Type:     program.File,
				Metadata: program.FileAssetMetadata{Path: "test.txt"},
			},
		},
		Challenges: []program.Challenge{
			{
				Type:     program.Flag,
				Metadata: program.FlagChallengeMetadata{Flag: "pwnbay{this_is_a_test_flag}"},
			},
		},
	}

	jsonData, err := json.Marshal(fromProg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonData))

	var toProg program.Program
	err = json.Unmarshal(jsonData, &toProg)
	if err != nil {
		slog.Error(err.Error())
	}

	log.Println(toProg)
}
