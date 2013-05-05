// echo writes all command line arguments to stdout
package main

import (
	"fmt"
	"os"
)

func main() {
	last := len(os.Args) - 1
	for _, arg := range os.Args[1:last] {
		fmt.Printf("%s ", arg)
	}
	fmt.Printf("%s\n", os.Args[last])
}
