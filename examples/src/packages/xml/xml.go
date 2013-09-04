// Example of parsing xml in Go without marshalling into structures.
package main

import (
	"encoding/xml"
	"fmt"
	"os"
        "io"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "file")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	depth := 0
        inside := false
	parser := xml.NewDecoder(file)
	for {
		token, err := parser.Token()
		if token == nil && err == io.EOF {
                fmt.Println("done!!", err)
			break
		}
                if err != nil {
                  fmt.Println("Error parsing xml:", err)
                  os.Exit(2)
                }
		switch t := token.(type) {
		case xml.StartElement:
			name := t.Name.Local
			print(name, depth)
			for _, v := range t.Attr {
				print("name", depth+1)
				print(v.Name.Local, depth+2)
				print("value", depth+1)
				print(v.Value, depth+2)
			}
                        inside = true
			depth++
		case xml.EndElement:
			depth--
                        inside = false
		case xml.CharData:
			if inside {
				print(string(t), depth)
			}
		case xml.Comment:
			print("Comment", depth)
			print(string(t), depth+1)
		case xml.ProcInst:
			print("ProcInst", depth)
			print("Target", depth+1)
			print(t.Target, depth+2)
			print("Inst", depth+1)
			print(string(t.Inst), depth+2)
		case xml.Directive:
			print("Directive", depth)
			print(string(t), depth+1)
		default:
			fmt.Println("Unknown xml element")
		}
	}
}

func print(s interface{}, depth int) {
	fmt.Printf("%*v%v\n", depth*3, "", s)
}
