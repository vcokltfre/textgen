package textgen

import "math/rand"

type Weights struct {
	weights map[string]map[string]int
}

func NewWeights() *Weights {
	return &Weights{weights: map[string]map[string]int{}}
}

func (w *Weights) Add(word, next string) {
	if _, ok := w.weights[word]; !ok {
		w.weights[word] = map[string]int{}
	}
	w.weights[word][next]++
}

func (w *Weights) Predict(word string) string {
	if word == "" {
		for w := range w.weights {
			return w
		}
	}

	if _, ok := w.weights[word]; !ok {
		return ""
	}

	weights := w.weights[word]
	total := 0

	for _, weight := range weights {
		total += weight
	}

	r := rand.Intn(total)

	for next, weight := range weights {
		r -= weight
		if r < 0 {
			return next
		}
	}

	return ""
}

func (w *Weights) Train(text string) {
	previous := ""

	for _, word := range SplitText(text) {
		w.Add(previous, word)
		previous = word
	}
}
