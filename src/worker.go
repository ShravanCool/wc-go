package src

func Worker(filePath string, flagSet FlagOptions) {
	lines := make(chan string)
	errChan := make(chan error)

	go readFile(filePath, lines, errChan)

	result := count(lines, errChan)
	result.fileName = filePath

	output, err := result.generateOutput(flagSet)
	if err != nil {
		PrintToStderr(err)
	}

	PrintToStdout(output)

}
