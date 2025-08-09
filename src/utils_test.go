package src

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPrintToStderr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Send a random error to Stderr",
			args: args{
				err: fmt.Errorf("This is a random error"),
			},
		},
	}

	for _, tt := range tests {
		oldStderr := os.Stderr
		defer func() { os.Stderr = oldStderr }()
		r, w, _ := os.Pipe()
		os.Stderr = w
		t.Run(tt.name, func(t *testing.T) {
			PrintToStderr(tt.args.err)
			w.Close()
		})

		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)

		assert.Equal(t, tt.args.err.Error(), buf.String())
	}
}

func TestPrintToStdout(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Check if a string passed is printed in Stdout",
			args: args{
				s: "This is a test string",
			},
		},
	}

	for _, tt := range tests {
		oldStdout := os.Stdout
		defer func() { os.Stdout = oldStdout }()
		r, w, _ := os.Pipe()
		os.Stdout = w
		t.Run(tt.name, func(t *testing.T) {
			printToStdout(tt.args.s)
			w.Close()
		})

		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)

		assert.Equal(t, tt.args.s, buf.String())
	}

}
