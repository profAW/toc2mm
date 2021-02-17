package infrastructure

import (
	"fmt"
	"os/exec"
)

func GeneratePlantumlDiagramm(filename string){
	path, _ := exec.LookPath("java")
	fmt.Println(path)

	cmd := exec.Command(path, "--jar plantuml.jar " + filename ).Run()

	fmt.Println(cmd)
}