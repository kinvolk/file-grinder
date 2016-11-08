package main

import(
	"flag"
	"fmt"
	"os"
	"bufio"	
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func readafile (path string, linenumbers bool) {
	file, err := os.Open(path)
	check(err)
	i:=0
	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
		if linenumbers {
			i++
        		fmt.Println(strconv.Itoa(i)+")",scanner.Text())
		}else {
			fmt.Println(scanner.Text())		
		}
    	}
	file.Close()
}

func main (){
	pathtofile := flag.String("p","","please declare the path of your file.")
	linenumbers :=flag.Bool("line-numbers", false, "show line-numbers")
	flag.Parse()
	readafile(*pathtofile, *linenumbers)	
}
