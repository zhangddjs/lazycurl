package main

import (
	"fmt"
	"os"

	sw "github.com/mattn/go-shellwords"
)

func main() {
	analyzeFile()
}

func analyze() {
	s := "curl -X POST -H \"Content-Type: application/json\" -d '{\"key\":\"value\"}' https://example.com/api"
	cmdParts, err := sw.Parse(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cmdParts[4])
	// TODO: need import go-flags to parse cmd args to struct
	// args, err := flags.ParseCmd()...
}

func analyzeFile() {
	fileContent, _ := os.ReadFile("shopping.curl")
	cmdParts, err := sw.Parse(string(fileContent))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("16:", cmdParts[16])
}
