/*
   cat writes files to the stdout. 

   If no arguments are given, cat copies the stdin to stdout. Otherwise each file
   is writtne to stdout in order. Any errors are written to stderr, but processing is
   not halted.
*/
package main

import (
	"io"
	"log"
	"os"
)

// catFile copies a file to Stdout
func catFile(f *os.File) error {
	_, err := io.Copy(os.Stdout, f)
	if err != nil {
		return err
	}
	err = os.Stdout.Sync()
	return err
}

func main() {
	if len(os.Args) == 1 {
		err := catFile(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			file, err := os.Open(arg)
			if err != nil {
                                 logger := log.New(os.Stderr, "", 0)
				logger.Println(err)
			}
			catFile(file)
		}
	}
}
