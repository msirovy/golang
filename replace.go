package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"flag"
)


func replace(file string, match string, replace string) error {
	// try to read given file
	src, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Processing ", file)

	// replace all matching strings and save result to dest
	dest := strings.Replace(string(src), match, replace, -1)

	// print dest
	fmt.Println(dest)

	return nil
}

func main() {

	// define command flags
	aFile := flag.String("file", "", "a string")
	aMatch := flag.String("match", "", "a string")
	aReplace := flag.String("replace", "", "a string")

	// parse input
	flag.Parse()

	// process file
	fmt.Println("Parse file ", *aFile)
	replace(*aFile, *aMatch, *aReplace)

}
