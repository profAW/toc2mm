package main

import (
	"fmt"
	"strconv"
	"toc2mm/domain"
	"toc2mm/helper"
	"toc2mm/infrastructure"
)

func doConversion(debug bool) {

	var directory = helper.GetCurrentDir(debug)
	fmt.Println("Working director is : " + directory)

	files, _ := infrastructure.GetTocFilesInFolders(directory)

	for _, file := range files {
		fmt.Println("Do conversion for   : " + file)
		outputfileNames := infrastructure.CreateExportFileNames(file)

		var lines = infrastructure.ReadBasicFileData(file)

		lines = domain.DoPlantumlConversion(lines)
		infrastructure.CreateFile(lines, outputfileNames[0])

		lines = domain.DoMindnodeConversion(lines)
		infrastructure.CreateFile(lines, outputfileNames[1])
	}
}

var DebugMode = "true"
var version = "0.1.1"

func main() {
	debug, _ := strconv.ParseBool(DebugMode)

	fmt.Println("------------------------------------------------------------")
	fmt.Println("████████╗░█████╗░░█████╗░██████╗░███╗░░░███╗███╗░░░███╗")
	fmt.Println("╚══██╔══╝██╔══██╗██╔══██╗╚════██╗████╗░████║████╗░████║")
	fmt.Println("░░░██║░░░██║░░██║██║░░╚═╝░░███╔═╝██╔████╔██║██╔████╔██║")
	fmt.Println("░░░██║░░░██║░░██║██║░░██╗██╔══╝░░██║╚██╔╝██║██║╚██╔╝██║")
	fmt.Println("░░░██║░░░╚█████╔╝╚█████╔╝███████╗██║░╚═╝░██║██║░╚═╝░██║")
	fmt.Println("░░░╚═╝░░░░╚════╝░░╚════╝░╚══════╝╚═╝░░░░░╚═╝╚═╝░░░░░╚═╝")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("Aim of the programm : Convert a Latex toc-file into a mindmap")
	fmt.Println("toc2mm-Version      : " + version)
	fmt.Println("Debug-Mode          : " + strconv.FormatBool(debug))
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("")

	doConversion(debug)

	fmt.Println("---------------- Conversion finished -----------------------")
	fmt.Println("Press enter key to exit...")
	helper.CloseApplicationWithOutError()
}
