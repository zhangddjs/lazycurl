package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	sw "github.com/mattn/go-shellwords"
)

func main() {
	// analyzeFile()
	analyze()
}

type Curl struct {
	Header     []string `short:"H" long:"header" description:"curl headers"`
	Method     string   `short:"X" long:"--request" description:"request command"`
	Body       string   `short:"d" long:"data" description:"curl request body"`
	Verbose    []bool   `short:"v" long:"verbose" description:"make the operation more talkative"`
	Compressed []bool   `long:"compressed" description:"Request compressed response"`
	Rawcurl    string
}

func analyze() {
	s := "curl -X POST -H \"Content-Type: application/json\" -d '{\"key\":\"value\"}' https://example.com/api"
	cmdParts, err := sw.Parse(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := Curl{}
	args, _ := flags.ParseArgs(&res, cmdParts)
	fmt.Println("Curl:", res)
	fmt.Println("arg:", args)
}

func analyzeFile() {
	res := Curl{}
	fileContent, _ := os.ReadFile("shopping.curl")
	cmdParts, err := sw.Parse(string(fileContent))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = flags.ParseArgs(&res, cmdParts)
	fmt.Println("curl:", cmdParts)
	fmt.Println("headers:", res.Header[0])
}
