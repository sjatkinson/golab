// echo writes all command line arguments to stdout
package main

import (
	"fmt"
        "os"
	"strings"
)


func main() {
      fmt.Printf("%s\n", strings.Join(os.Args[1:], " "))
}
