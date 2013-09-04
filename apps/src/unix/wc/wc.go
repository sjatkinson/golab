/*
   wc counts the chars, words, and lines in a file.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Stats struct {
	chars int
	words int
	lines int
}

func (s *Stats) Accum(other *Stats) {
   s.chars += other.chars
   s.words += other.words
   s.lines += other.lines
}

func (s *Stats)Report(filename string) {
	fmt.Printf("%8v%8v%8v %v\n", s.lines, s.words, s.chars, filename)
}

func scanFile(reader *bufio.Reader) (Stats) {
	var chars, words, lines int
	var inWord bool
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}
		chars++
		switch c {
		case ' ', '\t':
			if inWord {
				inWord = false
			}
		case '\n':
			lines++
			inWord = false
		default:
			if !inWord {
				words++
			}
			inWord = true
		}
	}
	return Stats{chars, words, lines}
}

func main() {
	if len(os.Args) == 1 {
		rdr := bufio.NewReader(os.Stdin)
		s := scanFile(rdr)
		s.Report("")
	} else {
                var totals Stats
                var files int
		for _, arg := range os.Args[1:] {
                        files++
			file, err := os.Open(arg)
			if err != nil {
				logger := log.New(os.Stderr, "", 0)
				logger.Println(err)
			}
			rdr := bufio.NewReader(file)
			s := scanFile(rdr)
			s.Report(arg)
                        totals.Accum(&s)
		}
                if files > 1 {
                   totals.Report("total")
                }
	}

}
