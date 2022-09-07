package textgen

import "strings"

func SplitText(text string) []string {
	words := []string{}
	word := ""
	prev := ' '

	for i, char := range strings.ToLower(text) {
		if (char >= 'a' && char <= 'z') || char == '\'' || char == '’' {
			word += string(char)
		}

		if char == '-' && (prev >= 'a' && prev <= 'z') && i < len(text)-1 && (text[i+1] >= 'a' && text[i+1] <= 'z') {
			word += string(char)
		}

		if char == ' ' || char == '\n' || char == '\t' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}

		prev = char
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}
