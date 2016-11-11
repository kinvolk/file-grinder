package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pathToFile := flag.String("path", "", "Declare the path of your file.")
	lineNumbers := flag.Bool("show-line-numbers", false, "Show the linenumbers")
	odd := flag.Bool("odd", false, "Show odd lines only")
	even := flag.Bool("even", false, "Show even lines only")
	info := flag.Bool("info", false, "Show Information about text in file")
	grind := flag.String("grind", "", "--grind l for print lines in random order")
	flag.Parse()

	data := readFile(*pathToFile)
	switch {
	case *even && !*odd:
		data = getEven(data)
	case *odd && !*even:
		data = getOdd(data)
	}

	if *grind == "l" {
		shuffleLines(data)
	}

	if *lineNumbers {
		printWithLineNumbers(data)
	} else {
		printPlainText(data)
	}

	if *info {
		printInfo(data)
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

func shuffleLines(a []Line) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
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

func getTotalCharacters(data []Line) int {
	c := 0
	for _, l := range data {
		c += strings.Count(l.content, "") - 1
	}
	return c
}
func getTotalWhiteCharacters(data []Line) int {
	c := 0
	for _, l := range data {
		c += strings.Count(l.content, " ")
		c += strings.Count(l.content, string(9)) //9 is ASCII for tabulator
	}
	return c
}
func getSymbolAmount(data []Line) int {
	s := 0
	for _, l := range data {
		for _, c := range l.content {
			if !(c == 32 || c == 9 || (c >= 48 && c <= 57) || //spaces, tabs, numbers
				(c >= 65 && c <= 90) || (c >= 97 && c <= 122)) { //a-z, A-Z
				s++
			}
		}
	}
	return s
}

func printInfo(data []Line) {
	fmt.Println()
	fmt.Println("information:")
	fmt.Println("-lines:\t\t\t", len(data))
	fmt.Println("-characters:\t\t", getTotalCharacters(data))
	fmt.Println("-white characters:\t", getTotalWhiteCharacters(data))
	fmt.Println("-symbols:\t\t", getSymbolAmount(data))
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
