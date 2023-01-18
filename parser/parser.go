package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type parser struct {
	str string
}

func NewParser(str string) *parser {
	p := new(parser)
	p.str = str

	return p
}

// Error-free code assumption

func (p *parser) IsA() bool {
	return p.str[0] == '@'
}

func (p *parser) IsC() bool {
	return !p.IsA()
}

func (p *parser) GetValue() string {
	return p.str[1:]
}

func (p *parser) GetValueCode() (string, error) {
	i, err := strconv.ParseInt(p.GetValue(), 10, 15)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%015b", i), nil
}

func (p *parser) GetDest() string {
	index := strings.Index(p.str, "=")

	if index == -1 {
		return ""
	}

	return p.str[:index]
}

func (p *parser) GetComp() string {
	start := 0
	end := len(p.str)

	index := strings.Index(p.str, "=")

	if index != -1 {
		start = index + 1
	}

	index = strings.Index(p.str, ";")

	if index != -1 {
		end = index
	}

	return p.str[start:end]
}

func (p *parser) GetJump() string {
	index := strings.Index(p.str, ";")

	if index == -1 {
		return ""
	}

	return p.str[index + 1:]
}

func (p *parser) GetDestCode() (string, error) {
	code, ok := destTable[p.GetDest()]

	if !ok {
		return "", errors.New("no dest code")
	}

	return code, nil
}

func (p *parser) GetCompCode() (string, error) {
	code, ok := compTable[p.GetComp()]

	if !ok {
		return "", errors.New("no comp code")
	}

	return code, nil
}

func (p *parser) GetJumpCode() (string, error) {
	code, ok := jumpTable[p.GetJump()]

	if !ok {
		return "", errors.New("no jump code")
	}

	return code, nil
}

func (p *parser) ToBinary() (string, error) {
	if p.IsA() {
		value, err := p.GetValueCode()

		if err != nil {
			return "", err
		}

		return "0" + value, nil
	} else if p.IsC() {
		dest, err := p.GetDestCode()

		if err != nil {
			return "", err
		}

		comp, err := p.GetCompCode()

		if err != nil {
			return "", err
		}

		jump, err := p.GetJumpCode()

		if err != nil {
			return "", err
		}

		return "111" + dest + comp + jump, nil
	}

	// Unreachable error since we assume the code is error-free

	return "", errors.New("not an A or C instruction")
}