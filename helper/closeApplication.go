package helper

import "os"

func CloseApplicationWithOutError() {
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
	os.Exit(1)
}

func CloseApplicationWithError() {
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
	os.Exit(3)
}
