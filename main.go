package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/ncmprbll/hack-asm/parser"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	path := args[1]

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	output, err := os.Create(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) + ".hack")

	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

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

		output.WriteString(bin + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
