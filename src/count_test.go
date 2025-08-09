package src

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	// Test empty input
	emptyLines := make(chan string)
	emptyErrChan := make(chan error)
	go func() {
		close(emptyLines)
	}()

	res := count(emptyLines, emptyErrChan)
	if res.lineCount != 0 || res.wordCount != 0 || res.charCount != 0 || res.err != nil {
		t.Errorf("count(emptyLines, emptyErrChan) = %v; expected lineCount=0, wordCount=0, charCount=0, err=nil", res)
	}

	// Test input with one line
	oneLine := make(chan string)
	oneErrChan := make(chan error)
	go func() {
		oneLine <- "Hello, World!"
		close(oneLine)
	}()

	res = count(oneLine, oneErrChan)
	if res.lineCount != 1 || res.wordCount != 2 || res.charCount != 14 || res.err != nil {
		t.Errorf("count(oneLine, oneErrChan) = %v; expected lineCount=1, wordCount=2, charCount=14, err=nil", res)
	}

	// Test input with multiple lines
	multiLine := make(chan string)
	multiErrChan := make(chan error)
	go func() {
		multiLine <- "The quick brown fox"
		multiLine <- "jumps over the lazy dog."
		close(multiLine)
	}()

	res = count(multiLine, multiErrChan)
	if res.lineCount != 2 || res.wordCount != 9 || res.charCount != 45 || res.err != nil {
		t.Errorf("count(multiLine, multiErrChan) = %v; expected lineCount=2, wordCount=9, charCount=45, err=nil", res)
	}

	// Test input with error
	errorLine := make(chan string)
	errorErrChan := make(chan error)
	expectedErr := fmt.Errorf("test error")
	go func() {
		errorErrChan <- expectedErr
		close(errorLine)
	}()
	res = count(errorLine, errorErrChan)
	if res.lineCount != 0 || res.wordCount != 0 || res.charCount != 0 || res.err != expectedErr {
		t.Errorf(
			"count(errorLine, errorErrChan) = %v; expected lineCount=0, wordCount=0, charCount=0, err=%v",
			res,
			expectedErr,
		)
	}

}
