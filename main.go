package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func main() {
	pathtofile := flag.String("path", "", "please declare the path of your file.")
	linenumbers := flag.Bool("linenumbers", false, "Show the linenumbers")
	flag.Parse()
	data := readafile(*pathtofile)
	print(data, *linenumbers)
}

func print(data []string, numb bool) {
	for s := range data {
		if numb {
			fmt.Println(strconv.Itoa(s+1) + ") " + data[s])
		} else {
			fmt.Println(data[s])
		}
	}
}
