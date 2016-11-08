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
	pathPtr := flag.String("path", "", "path of the file to read")
	linePtr := flag.Bool("linenumbers", false, "Show Line Numbers")
	odd_Ptr := flag.Bool("odd", false, "Only show odd Lines")
	evenPtr := flag.Bool("even", false, "Only show even Lines")
	flag.Parse()
	
	file, err := os.Open(*pathPtr)
	if err!=nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i:=0
    for scanner.Scan() {
		i++
		if (Even(i)&&*evenPtr)||(!Even(i)&&*odd_Ptr)||(!*odd_Ptr&&!*evenPtr){
			if *linePtr {
				fmt.Println(strconv.Itoa(i)+") "+scanner.Text())
			} else {
				fmt.Println(scanner.Text())
			}
		}
    }
	file.Close()
}

func Even(number int) bool {
    return number%2 == 0
}