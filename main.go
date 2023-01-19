package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"path/filepath"

	"github.com/ncmprbll/hack-asm/parser"
)

var symbols = map[string]string{
	"SP":     "0",
	"LCL":    "1",
	"ARG":    "2",
	"THIS":   "3",
	"THAT":   "4",
	"R0":     "0",
	"R1":     "1",
	"R2":     "2",
	"R3":     "3",
	"R4":     "4",
	"R5":     "5",
	"R6":     "6",
	"R7":     "7",
	"R8":     "8",
	"R9":     "9",
	"R10":    "10",
	"R11":    "11",
	"R12":    "12",
	"R13":    "13",
	"R14":    "14",
	"R15":    "15",
	"SCREEN": "16384",
	"KBD":    "24576",
}

func sanitize(str string) string {
	index := strings.Index(str, "//")

	if index != -1 {
		str = str[:index]
	}

	str = strings.TrimSpace(str)

	return str
}

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

	lineNumber := 0
	varAddress := 16

	// Resolving labels

	for scanner.Scan() {
		line := sanitize(scanner.Text())

		if line == "" {
			continue
		}

		index := strings.Index(line, "(")

		if index != -1 {
			line = strings.Trim(line, "()")

			symbols[line] = strconv.FormatInt(int64(lineNumber), 10)
		} else {
			lineNumber++
		}
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)

	// Assembling

	for scanner.Scan() {
		line := sanitize(scanner.Text())

		if line == "" {
			continue
		}

		index := strings.Index(line, "(")

		if index != -1 {
			continue
		}

		p := parser.NewParser(line)

		if p.IsA() {
			value := p.GetValue()

			if _, err := strconv.Atoi(value); err != nil {
				resolved, ok := symbols[value]

				if !ok {
					resolved = strconv.FormatInt(int64(varAddress), 10)
					symbols[value] = resolved
					varAddress++
				}

				p = parser.NewParser("@" + resolved)
			}
		}

		bin, err := p.ToBinary()

		if err != nil {
			log.Fatal(err)
		}

		output.WriteString(bin + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
