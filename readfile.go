package main

import (
	"fmt"
	"bufio"
    "log"
    "os"
)

func main() {
	file, err := os.Open("file.txt")
	if err!=nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
	file.Close()
}