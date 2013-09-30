// head displays the first lines of a file
// TODO: work with multiple files
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// how many lines or bytes to display
var lines, bytes int

func init() {
	flag.IntVar(&lines, "n", 10, "number of lines to print")
	flag.IntVar(&bytes, "c", 0, "number of bytes to print")
}

// printLines displays the first specified lines in a file
func printLines(r *bufio.Reader, lines int) {
	for i := 0; i < lines; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Print(err)
			os.Exit(3)
		}
		fmt.Print(line)
	}
}

// printBytes displays the first specified bytes in a file
func printBytes(r *bufio.Reader, bytes int) {
	for i := 0; i < bytes; i++ {
		byte, err := r.ReadByte()
		if err != nil {
			fmt.Print(err)
			os.Exit(3)
		}
		fmt.Print(string(byte))
	}
}

func main() {
	flag.Parse()
	if flag.NFlag() > 1 {
		fmt.Println("Can not combine line and byte counts")
		os.Exit(1)
	}
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	if bytes > 0 {
		printBytes(r, bytes)
	} else {
		printLines(r, lines)
	}
}
