package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"toc2mm/helper"
)

var version = "0.0.8"

func getTocFilesInFolders(root string) ([]string, error) {
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

func doConversion() {
	//dir, _ := os.Getwd()
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	fmt.Println(exePath)
	files, _ := getTocFilesInFolders(exePath)
	fmt.Println(files)
	for _, file := range files {
		var lines = readBasicFileData(file)
		lines = removeOverheadLines(lines)
		lines = convertToPlantUmlSyntax(lines)
		file2 := strings.TrimSuffix(file, ".toc")
		s := []string{file2, "-toc-mm.plantuml"}
		file3 := strings.Join(s, "")
		createMindMapFile(lines, file3)
	}
}

func readBasicFileData(file string) []string {
	var lines []string
	bytesRead, _ := ioutil.ReadFile(file)
	fileContent := string(bytesRead)
	lines = strings.Split(fileContent, "\n")
	return lines
}

func removeOverheadLines(lines []string) []string {
	const delString1 = "\\babel@toc {german}{}"
	const delString2 = "relax"

	for key, element := range lines {
		if strings.Contains(element, delString2) {
			lines[key] = ""
		}
		if element == (delString1) {
			lines[key] = ""
		}
	}
	lines = removeEmptyLines(lines)
	return lines
}

func removeEmptyLines(lines []string) []string {
	var result []string
	for _, str := range lines {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func removeAndReplaceToc2Plant(line string) string {
	line = strings.TrimPrefix(line, "\\beamer@")
	line = strings.TrimSuffix(line, "}")
	line = strings.ReplaceAll(line, "}", ".")

	// Do not change order of Replacements
	line = strings.ReplaceAll(line, "subsubsectionintoc", "****")
	line = strings.ReplaceAll(line, "subsectionintoc", "***")
	line = strings.ReplaceAll(line, "sectionintoc", "**")
	return line
}

func convertToPlantUmlSyntax(lines []string) []string {
	for key, element := range lines {
		entr := strings.Split(element, "{")
		entr = removeLastThreeEntries(entr)
		var line = strings.Join(entr, "")
		line = removeAndReplaceToc2Plant(line)
		lines[key] = line
	}

	result := []string{"@startmindmap", "* TOC"}
	lines = append(result, lines...)
	lines = append(lines, "@endmindmap")

	return lines
}

func removeLastThreeEntries(entry []string) []string {
	entry = entry[:len(entry)-1]
	entry = entry[:len(entry)-1]
	entry = entry[:len(entry)-1]
	return entry
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func createMindMapFile(lines []string, path string) {
	if err := writeLines(lines, path); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}

func main() {

	log.Info("### Welcome and remember 'never forget your towel' ###")
	log.Info("------------------------------------------------------")
	log.Info("toc2mm-Version: " + version)

	doConversion()

	log.Info("Press enter key to exit...")
	helper.CloseApplicationWithOutError()

}