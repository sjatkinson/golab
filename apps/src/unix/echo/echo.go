package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	count := flag.NArg()
	last := count - 1
	for i := 0; i < last; i++ {
		fmt.Printf("%s ", flag.Arg(i))
	}
	fmt.Printf("%s\n", flag.Arg(last))
}
