package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// ArcDescription is a struct representing a single arc
// in a CYOA story.
type ArcDescription struct {
	Title   string
	Story   []string
	Options []Option
}

// Option is a possible choice to be made in an arc
type Option struct {
	Text string
	Arc  string
}

// ParseJSONAdventureFile receives a file path representing
// CYOA and parses it to a map, mapping between an arc name
// and ArcDescription
func ParseJSONAdventureFile(path string) (map[string]ArcDescription, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	result := make(map[string]ArcDescription)
	err = json.Unmarshal(data, &result)

	if _, ok := result["intro"]; !ok {
		return result, errors.New("can't find intro arc")
	}

	return result, err
}
