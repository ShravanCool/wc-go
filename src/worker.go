package src

import (
	"sync"
)

func worker(filePath string, flagSet FlagOptions, totalCounter *totalCounter, wg *sync.WaitGroup, maxOpenFilesLimitBuffer chan int) {
	lines := make(chan string)
	errChan := make(chan error)
	maxOpenFilesLimitBuffer <- 1

	defer func() {
		wg.Done()
		<-maxOpenFilesLimitBuffer
	}()

	go readFile(filePath, lines, errChan)

	result := count(lines, errChan)
	totalCounter.lineCount += result.lineCount
	totalCounter.wordCount += result.wordCount
	totalCounter.charCount += result.charCount

	result.fileName = filePath

	output, err := result.generateOutput(flagSet)
	if err != nil {
		PrintToStderr(err)
	}

	printToStdout(output)
}
