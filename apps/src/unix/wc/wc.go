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

// countFn defines a function which returns the counts for a given rune in the input stream.
type countFn func(r rune) counts

// inWord is a countFn which returns the count of elements for a given rune when we are
// currently in a word counting.
func inWord(r rune) (counts counts) {
	counts.chars = 1
	switch classify(r) {
	case charWord: // still in a word, keep counting
	case charNewline:
		counts.lines = 1
		fallthrough
	case charSpace: // done with a word, count it
		counts.words = 1
	}
	return
}

func inSpace(r rune) (counts counts) {
	counts.chars = 1
	switch classify(r) {
	case charWord, charSpace: // nothing else to count
	case charNewline:
		counts.lines = 1
	}
	return
}

func scan(reader *bufio.Reader) counts {
	var s counts
	countFn := inSpace
	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			break
		}
		s.Add(countFn(r))
		// set the countFn function for each new char. We are either in a word
		// counting, or in whitespace counting. This will enable us to determine
		// when we are transistioning from a word to whitespace and visa-versa.
		if unicode.IsSpace(r) {
			countFn = inSpace
		} else {
			countFn = inWord
		}
	}
	return s
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
