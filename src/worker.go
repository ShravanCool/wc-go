package src

func Worker(filePath string) {
	lines := make(chan string)
	errChan := make(chan error)

	go readFile(filePath, lines, errChan)

	result := count(lines, errChan)
	result.fileName = filePath

	output, err := result.generateOutput()
	if err != nil {
		printToStderr(err)
	}

	printToStdout(output)

}
