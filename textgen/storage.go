package textgen

import (
	"encoding/json"
	"os"
)

func LoadWeights(filename string) (*Weights, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var weights map[string]map[string]int
	err = json.Unmarshal(data, &weights)
	if err != nil {
		return nil, err
	}

	return &Weights{weights}, nil
}

func (w *Weights) Save(filename string) error {
	data, err := json.Marshal(w.weights)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
