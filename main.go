package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func readData() ([]byte, string) {
	fileName := "sample.csv"
	var f io.Reader
	var err error
	var bytes []byte
	if fileName != "" {
		f, err = os.OpenFile(fileName, os.O_RDONLY, 0)
		if err != nil {
			fmt.Println("unable to open file: ", err)
			os.Exit(1)
		}
	} else {
		f = os.Stdin
	}
	bytes, err = io.ReadAll(f)
	if err != nil {
		fmt.Println("some error reading data: ", err)
		os.Exit(1)
	}
	return bytes, fileName
}

func count(bytes []byte) (counts counter) {
	counts.bytes = len(bytes)
	start := 0
	countingWord := false
	for i := 0; i < counts.bytes; i++ {
		if bytes[i] == '\n' {
			counts.lines++
			counts.chars = counts.chars + utf8.RuneCount(bytes[start:i])
			start = i
		}
		if unicode.IsSpace(rune(bytes[i])) {
			if countingWord {
				counts.words++
			}
			countingWord = false
		} else {
			countingWord = true
		}
	}
	counts.chars = counts.chars + utf8.RuneCount(bytes[start:])
	if countingWord {
		counts.words++
	}
	return
}

type counter struct {
	bytes int
	lines int
	chars int
	words int
}

func main() {
	bytes, fileName := readData()
	c := count(bytes)
	fmt.Printf("%d %d %d %d %s", c.bytes, c.chars, c.lines, c.words, fileName)
}
