package src

import (
	"strings"
)

func count(lines <-chan string, errChan <-chan error) result {
	var res result

	for {
		select {
		case err := <-errChan:
			if err != nil {
				res.err = err
				errChan = nil
				return res
			}
		case line, ok := <-lines:
			if !ok {
				return res
			}
			res.lineCount++
			words := strings.Fields(line)
			res.wordCount += len(words)
			res.charCount += len(line) + 1
		}
	}
}
