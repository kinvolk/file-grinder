package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readafile(path string) string {
	file, err := os.Open(path)
	check(err)
	str := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str += scanner.Text()
	}
	file.Close()
	return string(str)
}

func main() {
	pathtofile := flag.String("path", "", "please declare the path of your file.")
	flag.Parse()
	data := readafile(*pathtofile)
	fmt.Println(string(data))
}
