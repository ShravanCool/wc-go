package src

import "fmt"

type result struct {
	lineCount int
	wordCount int
	charCount int
	fileName  string
	err       error
}

func (res result) generateOutput() (string, error) {
	var output string

	if res.err != nil {
		return "", res.err
	}

	output += fmt.Sprintf("%8d", res.lineCount)
	output += fmt.Sprintf("%8d", res.wordCount)
	output += fmt.Sprintf("%8d", res.charCount)

	return output, nil
}
