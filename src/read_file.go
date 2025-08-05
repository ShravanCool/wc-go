package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validateFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("File is protected or not accessible: %w", err)
		}

		return fmt.Errorf("%w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Could not stat file: %w", err)
	}

	mode := info.Mode()
	if mode.IsDir() {
		return fmt.Errorf("Is a directory")
	}

	return nil
}

func readFile(filePath string, lines chan<- string, errChan chan<- error) {
	const chunkSize = 1024 * 1024
	defer close(lines)
	defer close(errChan)

	err := validateFile(filePath)
	if err != nil {
		errChan <- err
	}

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, chunkSize), chunkSize)

	for scanner.Scan() {
		line := scanner.Text()
		lines <- line
	}

	if err := scanner.Err(); err != nil {
		err = fmt.Errorf(
			"wc-go: " + strings.Replace(err.Error(), "read ", "", 1) + "\n",
		)
		errChan <- err
	}
}
