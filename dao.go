package coffee_roulette

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ReadHistory reads the history from a file on the filesystem
func ReadHistory(filename string) (History, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var history History
	yaml.Unmarshal(bytes, &history)
	return history, nil
}

// WriteHistory writes the history to a file on the filesystem
func WriteHistory(filename string, history History) error {
	bytes, err := yaml.Marshal(history)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
