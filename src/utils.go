package src

import (
	"fmt"
	"os"
)

func printToStderr(err error) {
	fmt.Fprint(os.Stderr, err.Error())
}

func printToStdout(s string) {
	fmt.Fprint(os.Stdout, s)
}
