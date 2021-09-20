package domain

import (
	"strings"
)

func DoMindnodeConversion(lines []string) []string {

	lines = convertPlantuml2mindnode(lines)
	lines = trimTrailingWhitespaces(lines)
	lines = removeEmptyLines(lines) // reuse from plantumlexport.go

	return lines
}

func trimTrailingWhitespaces(lines []string) []string {

	for key, line := range lines {
		line = strings.Trim(line, "  ")
		lines[key] = line
	}
	return lines
}

func convertPlantuml2mindnode(lines []string) []string {
	for key, line := range lines {
		line = strings.ReplaceAll(line, "* TOC", "")
		line = strings.ReplaceAll(line, "**[#Orange]", "")
		line = strings.ReplaceAll(line, "***[#lightblue]", "\t")
		line = strings.ReplaceAll(line, "****[#lightgreen] ", "\t\t")
		line = strings.ReplaceAll(line, "*****[#lightyellow] ", "\t\t\t")
		line = strings.ReplaceAll(line, "******[#lightcyan] ", "\t\t\t\t")
		line = strings.ReplaceAll(line, "*******[#Ivory] ", "\t\t\t\t\t")
		line = strings.ReplaceAll(line, "********[#MintCream] ", "\t\t\t\t\t\t")
		line = strings.ReplaceAll(line, "@startmindmap", "")
		line = strings.ReplaceAll(line, "@endmindmap", "")
		line = strings.ReplaceAll(line, "left side", "")
		line = strings.TrimPrefix(line, " ")
		lines[key] = line
	}
	return lines
}
