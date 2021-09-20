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

func changeAndRemoveBeamerContent(element string, tocLevel int) string {
	var line string
	entr := strings.Split(element, "{")
	entr = removeLastThreeEntries(entr)
	line = strings.Join(entr, "")

	line = strings.TrimPrefix(line, "\\beamer@")
	line = strings.TrimSuffix(line, "}")
	line = strings.ReplaceAll(line, "}", ".")

	// Do not change order of Replacements
	line = replaceTocWithPlantuml(line, tocLevel)

	return line
}

func changeAndRemoveArticleContent(element string, tocLevel int) string {
	var line string
	line = element
	line = strings.ReplaceAll(line, "}", " ")
	line = strings.ReplaceAll(line, "1em", "")
	line = replaceTocWithPlantuml(line, tocLevel)
	var lineArray = strings.Split(line, "{")
	line = lineArray[1] + " " + lineArray[3]
	return line
}

func replaceTocWithPlantuml(line string, tocLevel int) string {

	switch tocLevel {
	case 1:
		line = strings.ReplaceAll(line, "subparagraph", "********[#MintCream]")
		line = strings.ReplaceAll(line, "paragraph", "*******[#Ivory]")
		line = strings.ReplaceAll(line, "subsubsection", "******[#lightcyan]")
		line = strings.ReplaceAll(line, "subsection", "*****[#lightyellow]")
		line = strings.ReplaceAll(line, "section", "****[#lightgreen]")
		line = strings.ReplaceAll(line, "chapter", "***[#lightblue]")
		line = strings.ReplaceAll(line, "part", "**[#Orange]")

	case 2:
		line = strings.ReplaceAll(line, "subparagraph", "*******[#Ivory]")
		line = strings.ReplaceAll(line, "paragraph", "******[#lightcyan]")
		line = strings.ReplaceAll(line, "subsubsection", "*****[#lightyellow]")
		line = strings.ReplaceAll(line, "subsection", "****[#lightgreen]")
		line = strings.ReplaceAll(line, "section", "***[#lightblue]")
		line = strings.ReplaceAll(line, "chapter", "**[#Orange]")

	default:
		line = strings.ReplaceAll(line, "subparagraph", "******[#lightcyan]")
		line = strings.ReplaceAll(line, "paragraph", "*****[#lightyellow]")
		line = strings.ReplaceAll(line, "subsubsection", "****[#lightgreen]")
		line = strings.ReplaceAll(line, "subsection", "***[#lightblue]")
		line = strings.ReplaceAll(line, "section", "**[#Orange]")
	}
	// if beamer
	line = strings.ReplaceAll(line, "intoc", "]") // remove "intoc"

	return line
}

// Source: https://freshman.tech/snippets/go/check-if-slice-contains-element/
// plus anpassung
func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(v, str) {
			return true
		}
	}

	return false
}

func tocIsFromArticle(lines []string) bool {
	return contains(lines, "contentsline")
}

func getTocLevel(lines []string) int {
	if contains(lines, "{part}") {
		return 1
	}
	if contains(lines, "{chapter}") {
		return 2
	}
	return 3

}

func convertToPlantUmlSyntax(lines []string) []string {

	var isAnArticle = tocIsFromArticle(lines)
	var tocLevel = getTocLevel(lines)
	for key, element := range lines {
		var line string
		line = element
		if isAnArticle {
			line = changeAndRemoveArticleContent(element, tocLevel)
		} else {
			line = changeAndRemoveBeamerContent(element, tocLevel)
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
