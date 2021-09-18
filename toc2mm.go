package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"toc2mm/helper"
)

var version = "0.0.9"

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

func doConversion(debug bool) {

	var directory string = helper.GetCurrentDir(debug)
	fmt.Println(directory)

	files, _ := getTocFilesInFolders(directory)

	fmt.Println(files)
	for _, file := range files {
		var lines = readBasicFileData(file)
		lines = removeOverheadLines(lines)
		lines = convertToPlantUmlSyntax(lines)
		file2 := strings.TrimSuffix(file, ".toc")
		s := []string{file2, "-toc-mm.plantuml"}
		outputfileName := strings.Join(s, "")
		createMindMapFile(lines, outputfileName)
		// infrastructure.GeneratePlantumlDiagramm(outputfileName)
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
	const delString3 = "\\babel@toc"

	for key, element := range lines {
		if strings.Contains(element, delString2) {
			lines[key] = ""
		}
		if element == (delString1) {
			lines[key] = ""
		}
		if element == (delString3) {
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
	line = strings.ReplaceAll(line, "subsubsectionintoc", "****[#lightgreen]")
	line = strings.ReplaceAll(line, "subsectionintoc", "***[#lightblue]")
	line = strings.ReplaceAll(line, "sectionintoc", "**[#Orange]")
	return line
}

func removeArticlePrefixes(line string) string {
	line = strings.ReplaceAll(line, "}", " ")
	//	fmt.Println(line)
	line = strings.ReplaceAll(line, "subsubsection", "****[#lightgreen]")
	//	fmt.Println(line)
	line = strings.ReplaceAll(line, "subsection", "***[#lightblue]")
	//	fmt.Println(line)
	line = strings.ReplaceAll(line, "section", "**[#Orange]")
	//	fmt.Println(line)
	var lineArray = strings.Split(line, "{")
	line = lineArray[1] + " " + lineArray[3] //+ lineArray[4]
	fmt.Println(line)
	return line
}

func convertToPlantUmlSyntax(lines []string) []string {

	for key, element := range lines {
		var line string
		line = element
		if strings.Contains(line, "contentsline") {
			line = removeArticlePrefixes(line)
		} else {
			entr := strings.Split(element, "{")
			entr = removeLastThreeEntries(entr)
			line = strings.Join(entr, "")
			line = removeAndReplaceToc2Plant(line)
		}
		lines[key] = line
	}

	breakPoint := addLeftSectionsInToc(lines) //+ 1
	result := []string{"@startmindmap", "* TOC"}
	lines = append(result, lines...)
	lines = append(lines, "@endmindmap")

	lines = append(lines, "")
	copy(lines[breakPoint+1:], lines[breakPoint:])
	lines[breakPoint+1] = "left side"

	return lines
}
func countSectionsInToc(lines []string) int {
	//Create a   dictionary of values for each element
	var ret = 0
	for _, line := range lines {
		entry := strings.Split(line, ".")
		entry[0] = replaceTocString(entry[0])
		i64, _ := strconv.ParseInt(entry[0], 10, 64)
		ret = int(i64)
	}
	return ret
}

func replaceTocString(entry string) string {
	entry = strings.ReplaceAll(entry, "*", "")
	entry = strings.ReplaceAll(entry, " ", "")
	entry = strings.ReplaceAll(entry, "[#lightgreen]", "")
	entry = strings.ReplaceAll(entry, "[#lightblue]", "")
	entry = strings.ReplaceAll(entry, "[#Orange]", "")
	return entry
}

func addLeftSectionsInToc(lines []string) int {
	//Create a   dictionary of values for each element
	numberOfSection := countSectionsInToc(lines)
	d := float64(numberOfSection) / float64(2)
	breakPoint := int(math.Ceil(d)) + 1

	var ret = 0
	var postion = 0
	for key, line := range lines {
		entry := strings.Split(line, ".")
		entry[0] = replaceTocString(entry[0])

		i64, _ := strconv.ParseInt(entry[0], 10, 64)
		ret = int(i64)
		if ret == breakPoint {
			postion = key
			break
		}
	}
	return postion

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

var debug bool = true

func main() {

	log.Info("### Welcome and remember 'never forget your towel' ###")
	log.Info("------------------------------------------------------")
	log.Info("toc2mm-Version: " + version)

	doConversion(debug)

	log.Info("Press enter key to exit...")
	helper.CloseApplicationWithOutError()
}
