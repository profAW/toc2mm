package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"toc2mm/domain"
	"toc2mm/helper"
	"toc2mm/infrastructure"
)

func doConversion(debug bool) {

	var directory = helper.GetCurrentDir(debug)
	log.Info("Working director is : " + directory)

	files, _ := infrastructure.GetTocFilesInFolders(directory)

	for _, file := range files {
		log.Info("Do conversion for   : " + file)
		outputfileNames := infrastructure.CreateExportFileNames(file)

		var lines = infrastructure.ReadBasicFileData(file)

		lines = domain.DoPlantumlConversion(lines)
		infrastructure.CreateFile(lines, outputfileNames[0])

		lines = domain.DoMindnodeConversion(lines)
		infrastructure.CreateFile(lines, outputfileNames[1])
	}
}

var debug = true
var version = "0.0.11"

func main() {

	log.Info("████████╗░█████╗░░█████╗░██████╗░███╗░░░███╗███╗░░░███╗")
	log.Info("╚══██╔══╝██╔══██╗██╔══██╗╚════██╗████╗░████║████╗░████║")
	log.Info("░░░██║░░░██║░░██║██║░░╚═╝░░███╔═╝██╔████╔██║██╔████╔██║")
	log.Info("░░░██║░░░██║░░██║██║░░██╗██╔══╝░░██║╚██╔╝██║██║╚██╔╝██║")
	log.Info("░░░██║░░░╚█████╔╝╚█████╔╝███████╗██║░╚═╝░██║██║░╚═╝░██║")
	log.Info("░░░╚═╝░░░░╚════╝░░╚════╝░╚══════╝╚═╝░░░░░╚═╝╚═╝░░░░░╚═╝")
	log.Info("-------------------------------------------------------")
	log.Info("toc2mm-Version      : " + version)
	log.Info("Debug-Mode          : " + strconv.FormatBool(debug))

	doConversion(debug)

	log.Info("---------------- Conversion finished -----------------")
	log.Info("Press enter key to exit...")
	helper.CloseApplicationWithOutError()
}
