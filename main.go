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
	pathToFile := flag.String("path", "", "Declare the path of your file.")
	lineNumbers := flag.Bool("show-line-numbers", false, "Show the linenumbers")
	odd := flag.Bool("odd", false, "Show odd lines only")
	even := flag.Bool("even", false, "Show even lines only")
	flag.Parse()

	data := readFile(*pathToFile)
	if *even&&!*odd{
		data=getEven(data)
	}
	if *odd&&!*even{
		data=getOdd(data)
	}

	if *lineNumbers {
		printWithLineNumbers(data)
	} else {
		printPlainText(data)
	}
}

type Line struct {
	linenumber int
	content    string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readFile(path string) []Line {
	file, err := os.Open(path)
	check(err)
	i := 0
	var str []Line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		str = append(str, Line{i, scanner.Text()})
	}
	file.Close()
	return str
}

func getEven(data []Line) []Line {
	var output []Line
	for s := range data {
		if data[s].linenumber%2 == 0 {
			output = append(output, data[s])
		}
	}
	return output
}
func getOdd(data []Line) []Line {
	var output []Line
	for s := range data {
		if data[s].linenumber%2 == 1 {
			output = append(output, data[s])
		}
	}
	return output
}

func printWithLineNumbers(data []Line) {
	for s := range data {
		fmt.Println(strconv.Itoa(data[s].linenumber) + ") " + data[s].content)
	}
}

func printPlainText(data []Line) {
	for s := range data {
		fmt.Println(data[s].content)
	}
}
