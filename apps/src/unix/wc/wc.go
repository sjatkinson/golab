// wc counts the chars, words, and lines in a file.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

var listChars, listWords, listLines bool

type options struct {
	ShowChars bool
	ShowWords bool
	ShowLines bool
}

func (o *options) init() {
	var chars, words, lines bool
	flag.BoolVar(&chars, "c", false, "list the chars")
	flag.BoolVar(&words, "w", false, "list the words")
	flag.BoolVar(&lines, "l", false, "list the lines")

	o.ShowChars, o.ShowWords, o.ShowLines = true, true, true

	flag.Parse()
	if flag.NFlag() > 0 {
		o.ShowChars = chars
		o.ShowWords = words
		o.ShowLines = lines
	}
}

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
func (s *counts) Report(title string, opts *options) {
	if opts.ShowLines {
		fmt.Printf("%8v", s.lines)
	}
	if opts.ShowWords {
		fmt.Printf("%8v", s.words)
	}
	if opts.ShowChars {
		fmt.Printf("%8v", s.chars)
	}
	fmt.Printf(" %8v\n", title)
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
			lines++
		}
		prev = c
	}
	return counts{chars, words, lines, 1}
}

func main() {
	var anyErrs bool
	var opts options
	opts.init()
	if flag.NArg() == 0 {
		rdr := bufio.NewReader(os.Stdin)
		s := scan(rdr)
		s.Report("", &opts)
	} else {
		var totals counts
		for _, arg := range flag.Args() {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				anyErrs = true
				continue
			}
			defer file.Close()
			rdr := bufio.NewReader(file)
			s := scan(rdr)
			s.Report(arg, &opts)
			totals.Add(s)
		}
		if totals.Added > 1 {
			totals.Report("total", &opts)
		}
		if anyErrs {
			os.Exit(1)
		}
	}
}
