package main

import (
	"github.com/spf13/cobra"
	"os"
	"wc-go/src"
)

var (
	flagSet src.FlagOptions
)

var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "wc is a command-line tool to count words, lines and characters",
	Long:  `wc is a command-line tool that reads from the standard input or from a file (or multiple files) to count words, lines and characters`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = []string{"-"}
		}

		src.Worker(args[0], flagSet)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&flagSet.LineFlag, "lines", "l", false, "Count number of lines")
	rootCmd.Flags().BoolVarP(&flagSet.WordFlag, "words", "w", false, "Count number of words")
	rootCmd.Flags().BoolVarP(&flagSet.CharFlag, "chars", "c", false, "Count number of characters")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		src.PrintToStderr(err)
		os.Exit(1)
	}
}
