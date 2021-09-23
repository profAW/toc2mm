package main

import (
	"fmt"
	"os"
	"strconv"
	"toc2mm/domain"
	"toc2mm/helper"
	"toc2mm/infrastructure"
)

func FolderIsValidDir(folder string) bool {
	dir, err := os.Stat(folder)
	if err != nil {
		fmt.Println("failed to open directory, error: %w", err)
		return false
	}
	if !dir.IsDir() {
		fmt.Println("%q is not a directory", dir.Name())
		return false
	}
	return true
}
func getConversionFoler(args []string, debug bool) string {

	folder := helper.GetCurrentDir(debug)
	if len(args) > 0 {
		folder = args[0]
	}

	isAValidFolder := FolderIsValidDir(folder)

	if !isAValidFolder {
		fmt.Println("---------------- Conversion failed, no valid folder or files -----------------------")
		fmt.Println("Press enter key to exit...")
		helper.CloseApplicationWithError()
	}

	fmt.Println("Working director is : " + folder)
	return folder
}

func doConversion(directory string) {

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
var version = "0.2.0"

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

	argsWithoutProg := os.Args[1:]
	directory := getConversionFoler(argsWithoutProg, debug)

	doConversion(directory)

	fmt.Println("---------------- Conversion finished -----------------------")
	fmt.Println("Press enter key to exit...")
	helper.CloseApplicationWithOutError()
}
