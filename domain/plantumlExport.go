package domain

import (
	"math"
	"strings"
)

func DoPlantumlConversion(lines []string) []string {
	lines = removeOverheadLines(lines)
	lines = removeEmptyLines(lines)
	lines = convertToPlantUmlSyntax(lines)
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

func changeAndRemoveBeamerContent(element string) string {
	var line string
	entr := strings.Split(element, "{")
	entr = removeLastThreeEntries(entr)
	line = strings.Join(entr, "")

	line = strings.TrimPrefix(line, "\\beamer@")
	line = strings.TrimSuffix(line, "}")
	line = strings.ReplaceAll(line, "}", ".")

	// Do not change order of Replacements
	line = strings.ReplaceAll(line, "subsubsectionintoc", "****[#lightgreen]")
	line = strings.ReplaceAll(line, "subsectionintoc", "***[#lightblue]")
	line = strings.ReplaceAll(line, "sectionintoc", "**[#Orange]")
	return line
}

func changeAndRemoveArticleContent(element string) string {
	var line string
	line = element
	line = strings.ReplaceAll(line, "}", " ")
	line = strings.ReplaceAll(line, "subsubsection", "****[#lightgreen]")
	line = strings.ReplaceAll(line, "subsection", "***[#lightblue]")
	line = strings.ReplaceAll(line, "section", "**[#Orange]")
	var lineArray = strings.Split(line, "{")
	line = lineArray[1] + " " + lineArray[3]
	return line
}

func tocIsFromArticle(line string) bool {
	return strings.Contains(line, "contentsline")
}

func convertToPlantUmlSyntax(lines []string) []string {

	for key, element := range lines {
		var line string
		line = element
		if tocIsFromArticle(line) {
			line = changeAndRemoveArticleContent(element)
		} else {
			line = changeAndRemoveBeamerContent(element)
		}
		lines[key] = line
	}

	breakPoint := calculateNumberOfLeftSectionsInToc(lines)

	startlines := []string{"@startmindmap", "* TOC"}
	lines = append(startlines, lines...)
	lines = append(lines, "@endmindmap")

	lines = append(lines, "")
	copy(lines[breakPoint+1:], lines[breakPoint:])
	lines[breakPoint+1] = "left side"

	return lines
}

func countSectionsInToc(lines []string) int {
	var value int
	for _, line := range lines {
		if strings.Contains(line, "#Orange") {
			value += 1
		}
	}
	return value
}

func calculateNumberOfLeftSectionsInToc(lines []string) int {
	numberOfSection := countSectionsInToc(lines)

	d := float64(numberOfSection) / float64(2)
	breakPoint := int(math.Ceil(d)) + 1

	var postion = 0
	var value = 0
	for key, line := range lines {
		if strings.Contains(line, "#Orange") {
			value += 1
		}
		if value == breakPoint {
			postion = key + 1 // cause of index stuff
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
