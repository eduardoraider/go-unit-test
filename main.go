package main

import (
	"flag"
	"fmt"
	"github.com/eduardoraider/go-unit-test/words"
	"os"
)

var (
	file   string
	filter string
)

func init() {
	flag.StringVar(&file, "file", "", "file to read from")
	flag.StringVar(&filter, "filter", "", "filter by word")
	flag.Parse()
}

func main() {
	if file == "" {
		panic("file is required")
	}

	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	c := words.CountWords(string(content), filter)
	fmt.Println(c)
}
