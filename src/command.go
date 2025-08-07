package src

import (
	"sync"
)

const maxOpenFileLimit = 10

func Command(flagSet FlagOptions, args []string) {
	var totalCounter totalCounter

	if len(args) == 0 {
		args = []string{"-"}
	}
	var wg sync.WaitGroup
	maxOpenFilesLimitBuffer := make(chan int, maxOpenFileLimit)

	for _, arg := range args {
		wg.Add(1)
		go worker(arg, flagSet, &totalCounter, &wg, maxOpenFilesLimitBuffer)
	}

	wg.Wait()

	if len(args) > 1 {
		totalResult, err := result{
			lineCount: totalCounter.lineCount,
			wordCount: totalCounter.wordCount,
			charCount: totalCounter.charCount,
			fileName:  "total",
		}.generateOutput(flagSet)
		if err != nil {
			PrintToStderr(err)
		}
		printToStdout(totalResult)
	}
}

