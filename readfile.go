package main

import (
	"fmt"
	"flag"
	"bufio"
    "log"
    "os"
	"strconv"
)

func main() {
	striPtr := flag.String("path", "", "path of the file to read")
	boolPtr := flag.Bool("linenumbers", false, "Show Line Numbers")
	
	flag.Parse()
	file, err := os.Open(*striPtr)
	if err!=nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i:=0
    for scanner.Scan() {
		i++
		if (*boolPtr) {
			fmt.Println(strconv.Itoa(i)+") "+scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
    }
	file.Close()
}