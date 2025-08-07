package src

import (
	"fmt"
	"os"
)

func PrintToStderr(err error) {
	fmt.Fprint(os.Stderr, err.Error())
}

func printToStdout(s string) {
	fmt.Fprint(os.Stdout, s)
}
