package src

import (
	"fmt"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestValidateFile(t *testing.T) {
	t.Run("Case when file path exists", func(t *testing.T) {
		tmpFile, e := os.CreateTemp("", "testFile")
		check(e)
		defer os.Remove(tmpFile.Name())
		tmpFile.Close()

		if err := validateFile(tmpFile.Name()); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Case when file path does not exist", func(t *testing.T) {
		err := validateFile("testFile")
		errorText := "open testFile: no such file or directory"

		if err == nil || err.Error() != errorText {
			t.Errorf("Expected 'No such file or directory' error, got %v", err.Error())
		}
	})

	t.Run("Case when file path passed is a directory", func(t *testing.T) {
		e := os.Mkdir("testDir", 0755)
		check(e)
		defer os.RemoveAll("testDir")

		err := validateFile("testDir")
		errorText := "testDir is a directory\n"
		if err == nil || err.Error() != errorText {
			t.Errorf("Expected 'is a directory' error, got %v", err)
		}
	})

	t.Run("Case when file path provided is protected file", func(t *testing.T) {
		tmpFile, e := os.CreateTemp("", "testFile")
		check(e)
		defer os.Remove(tmpFile.Name())
		tmpFile.Close()

		e = os.Chmod(tmpFile.Name(), 0000)
		check(e)
		defer os.Chmod(tmpFile.Name(), 0644)

		err := validateFile(tmpFile.Name())
		errorText := fmt.Sprintf("File is protected or not accessible: open %s: permission denied\n", tmpFile.Name())

		if err == nil || err.Error() != errorText {
			t.Errorf("Expected 'Permission denied' error, got %v", err)
		}
	})
}
