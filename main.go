package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	pathToFile := flag.String("path", "", "please declare the path of your file.")
	lineNumbers := flag.Bool("show-line-numbers", false, "Show the linenumbers")
	flag.Parse()
	data := readafile(*pathToFile)
	if *lineNumbers {
		printWithLineNumbers(data)
	} else {
		printPlainText(data)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readafile(path string) []string {
	file, err := os.Open(path)
	check(err)
	var str []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	file.Close()
	return str
}

func printWithLineNumbers(data []string) {
	for s := range data {
		fmt.Println(strconv.Itoa(s+1) + ") " + data[s])
	}
}
func printPlainText(data []string) {
	for s := range data {
		fmt.Println(data[s])
	}
}
