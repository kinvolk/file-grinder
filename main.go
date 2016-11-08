package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readafile (path string) string{
	dat, err := ioutil.ReadFile(path)
    	check(err)
	return string(dat)
}

func main() {
	pathtofile := flag.String("path", "", "please declare the path of your file.")
	flag.Parse()
	data := readafile(*pathtofile)
	fmt.Println(string(data))
}
