package domain

import (
	"strings"
)

func DoMindnodeConversion(lines []string) []string {

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
		lines[key] = line
	}
	return lines
}
