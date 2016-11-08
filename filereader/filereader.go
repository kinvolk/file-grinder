package main

import(
	"flag"
	"fmt"
	"io/ioutil"
	
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func readafile (path string) string{
	dat, err := ioutil.ReadFile(path)
    	check(err)
    	return string(dat)
}

func main (){
	pathtofile := flag.String("p","","please declare the path of your file.")
	flag.Parse()
	data := readafile(*pathtofile)
	fmt.Println(string(data))	
}
