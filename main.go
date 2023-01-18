package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ncmprbll/hack-asm/parser"
)

func main() {
	file, err := os.Open("Add.asm")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()

		index := strings.Index(line, "//")

		if index != -1 {
			line = line[:index]
		}

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		p := parser.NewParser(line)
		bin, err := p.ToBinary()

		if err != nil {
			continue
		}

		fmt.Println(lineNum, line, bin)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
