package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func LineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\r')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		r := regexp.MustCompile("[^\\s]+")

		words := r.FindAllString(line, -1)

		for i := range words {
			fmt.Print(words[i], " ")
		}
	}

	return nil
}

func main() {

	LineByLine("file.txt")

}
