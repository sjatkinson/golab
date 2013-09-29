// wc counts the chars, words, and lines in a file.
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// counts keeps track of the chars, words, and lines in a files or set of files
type counts struct {
	chars int
	words int
	lines int
	Added int // the number of structs added to ours
}

// Add takes in another counts struct and adds them to our current counts
func (s *counts) Add(other counts) {
	s.chars += other.chars
	s.words += other.words
	s.lines += other.lines
	s.Added++
}

// Report outputs the current counts to stdout
func (s *counts) Report(title string) {
	fmt.Printf("%8v%8v%8v %v\n", s.lines, s.words, s.chars, title)
}

// typeChar represents the type of chars contained in the elements we count
type typeChar int

// the list of possible typeChar values
const (
	charNewline typeChar = iota // a newline char
	charSpace                   // a whitespace char other than newline
	charWord                    // any other char is considered part of a word
)

// clasify determine what type of typeChar a given rune is
func classify(r rune) typeChar {
	if r == '\n' {
		return charNewline
	}
	if unicode.IsSpace(r) {
		return charSpace
	}
	return charWord
}


// scan counts all the chars, words, and lines in the given file
func scan(reader *bufio.Reader) counts {
        var prev typeChar = charSpace
        var chars, words, lines int
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break // TODO: what errors might we get here, and how to handle?
		}
                chars++
                c := classify(r)
                switch c {
                  case charWord:
                     if prev != c {
                        words++
                     }
                     case charSpace: // nothing
                     case charNewline:
                        lines++;
               }
               prev = c
	}
	return counts {chars, words, lines, 1}
}

func main() {
	var anyErrs bool
	if len(os.Args) == 1 {
		rdr := bufio.NewReader(os.Stdin)
		s := scan(rdr)
		s.Report("")
	} else {
		var totals counts
		for _, arg := range os.Args[1:] {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				anyErrs = true
				continue
			}
                        defer file.Close()
			rdr := bufio.NewReader(file)
			s := scan(rdr)
			s.Report(arg)
			totals.Add(s)
		}
		if totals.Added > 1 {
			totals.Report("total")
		}
		if anyErrs {
			os.Exit(1)
		}
	}
}
