package infrastructure

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadBasicFileData(file string) []string {
	var lines []string
	bytesRead, _ := ioutil.ReadFile(file)
	fileContent := string(bytesRead)
	lines = strings.Split(fileContent, "\n")
	return lines
}

func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func CreateFile(lines []string, path string) {
	if err := WriteLines(lines, path); err != nil {
		logrus.Fatalf("WriteLines: %s", err)
	}
}

func GetTocFilesInFolders(root string) ([]string, error) {
	var matches []string
	var pattern = "*.toc"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func CreateExportFileNames(basefilename string) []string {
	var outputfileName = make([]string, 2)
	file := strings.TrimSuffix(basefilename, ".toc")
	s1 := []string{file, "-toc-mm.plantuml"}
	s2 := []string{file, "-toc-mm-4-mindnode.txt"}

	outputfileName[0] = strings.Join(s1, "")
	outputfileName[1] = strings.Join(s2, "")

	return outputfileName
}
