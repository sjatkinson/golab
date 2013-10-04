// head displays the first lines of a file
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

// HeadFunction defines a function type we can pass around for output
type HeadFunction func(*bufio.Reader, int)

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

// headFile outputs a file to stdout using HeadFunction and count
func headFile(fname string, count int, printHead HeadFunction) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	printHead(r, count)
}

// getOutputFuncton determines if we are outputting bytes or lines, returns a function and count
func getOutputFunction(bytes, lines int) (outputFunc HeadFunction, count int) {
	if bytes > 0 {
		outputFunc = printBytes
		count = bytes
	} else {
		outputFunc = printLines
		count = lines
	}
	return
}

func main() {
	flag.Parse()
	if flag.NFlag() > 1 {
		fmt.Println("Can not combine line and byte counts")
		os.Exit(1)
	}
	printHead, count := getOutputFunction(bytes, lines)
	if flag.NArg() == 0 {
		rdr := bufio.NewReader(os.Stdin)
		printHead(rdr, count)
	} else {
		if flag.NArg() > 1 {
			for _, arg := range flag.Args() {
				fmt.Printf("==> %v <==\n", arg)
				headFile(arg, count, printHead)
				fmt.Println("")
			}
		} else {
			headFile(flag.Arg(0), count, printHead)
		}
	}
}
