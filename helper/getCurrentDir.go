package helper

import (
	"os"
	"path/filepath"
)

func GetCurrentDir(debug bool) string {
	var directory string

	if debug {
		directory, _ = os.Getwd()
	} else {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		directory = filepath.Dir(ex)
	}
	return directory
}
