//
// Simple ls example using go
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

// TODO: command line handling
// -- -l
// -- -a
// -- -r
func main() {
	flag.Parse()
	var dir string = "."
	if flag.NArg() > 0 {
		dir = flag.Args()[0]
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("Error reading dir: ", err)
	}
	for _, f := range files {
		fmt.Printf("%s\n", f.Name())
	}
}
