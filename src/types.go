package src

import "fmt"

type result struct {
	lineCount int
	wordCount int
	charCount int
	fileName  string
	err       error
}

type FlagOptions struct {
	LineFlag bool
	WordFlag bool
	CharFlag bool
}

type totalCounter struct {
	lineCount int
	wordCount int
	charCount int
}

func (res result) generateOutput(flagSet FlagOptions) (string, error) {
	var output string

	if res.err != nil {
		return "", res.err
	}

	if flagSet.LineFlag {
		output += fmt.Sprintf("%8d", res.lineCount)
	}

	if flagSet.WordFlag {
		output += fmt.Sprintf("%8d", res.wordCount)
	}

	if flagSet.CharFlag {
		output += fmt.Sprintf("%8d", res.charCount)
	}

	if !flagSet.LineFlag && !flagSet.WordFlag && !flagSet.CharFlag {
		output += fmt.Sprintf("%8d", res.lineCount)
		output += fmt.Sprintf("%8d", res.wordCount)
		output += fmt.Sprintf("%8d", res.charCount)
	}

	if res.fileName == "-" {
		output += "-\n"
	} else {
		output += fmt.Sprintf(" %s\n", res.fileName)
	}

	return output, nil
}
