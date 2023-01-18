package parser

import "testing"

var isatests = []struct {
	ins string
	exp bool
}{
	{"@0", true},
	{"@1", true},
	{"@2", true},
	{"@30", true},
	{"@400", true},
	{"@i", true},
	{"@f", true},
	{"@model", true},
	{"@register", true},
	{"AMD=0;JMP", false},
	{"D=A", false},
	{"M=D", false},
	{"0;JMP", false},
}

func TestIsA(t *testing.T) {
	for _, e := range isatests {
		p := parser{e.ins}

		got := p.IsA()
		want := e.exp

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}

var isctests = []struct {
	ins string
	exp bool
}{
	{"@0", false},
	{"@1", false},
	{"@2", false},
	{"@30", false},
	{"@400", false},
	{"@i", false},
	{"@f", false},
	{"@model", false},
	{"@register", false},
	{"AMD=0;JMP", true},
	{"D=A", true},
	{"M=D", true},
	{"D=D-M", true},
	{"0;JMP", true},
	{"D;JGT", true},
}

func TestIsC(t *testing.T) {
	for _, e := range isctests {
		p := parser{e.ins}

		got := p.IsC()
		want := e.exp

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}

var getvaluetests = []struct {
	ins string
	exp string
}{
	{"@0", "0"},
	{"@1", "1"},
	{"@2", "2"},
	{"@30", "30"},
	{"@400", "400"},
	{"@5122", "5122"},
	{"@3325", "3325"},
	{"@1222", "1222"},
	{"@53322", "53322"},
}

func TestGetValue(t *testing.T) {
	for _, e := range getvaluetests {
		p := parser{e.ins}

		got := p.GetValue()
		want := e.exp

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var getvaluecodetests = []struct {
	val string
	exp string
}{
	{"@0", "000000000000000"},
	{"@1", "000000000000001"},
	{"@2", "000000000000010"},
	{"@17", "000000000010001"},
	{"@32767", "111111111111111"},
}

func TestGetValueCode(t *testing.T) {
	for _, e := range getvaluecodetests {
		p := parser{e.val}

		got, err := p.GetValueCode()
		want := e.exp

		if err != nil {
			t.Errorf("%v", err)
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var getdesttests = []struct {
	ins string
	exp string
}{
	{"0", ""},
	{"AMD=0;JMP", "AMD"},
	{"D=A", "D"},
	{"M=D", "M"},
	{"D=D-M", "D"},
	{"0;JMP", ""},
	{"D;JGT", ""},
}

func TestGetDest(t *testing.T) {
	for _, e := range getdesttests {
		p := parser{e.ins}

		got := p.GetDest()
		want := e.exp

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var getdestcodetests = []struct {
	ins string
	exp string
}{
	{"0", "000"},
	{"AMD=0;JMP", "111"},
	{"D=A", "010"},
	{"M=D", "001"},
	{"D=D-M", "010"},
	{"0;JMP", "000"},
	{"D;JGT", "000"},
}

func TestGetDestCode(t *testing.T) {
	for _, e := range getdestcodetests {
		p := parser{e.ins}

		got, err := p.GetDestCode()
		want := e.exp

		if err != nil {
			t.Errorf("%v", err)
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var getcomptests = []struct {
	ins string
	exp string
}{
	{"0", "0"},
	{"AMD=0;JMP", "0"},
	{"D=A", "A"},
	{"M=D", "D"},
	{"D=D-M", "D-M"},
	{"0;JMP", "0"},
	{"D;JGT", "D"},
}

func TestGetComp(t *testing.T) {
	for _, e := range getcomptests {
		p := parser{e.ins}

		got := p.GetComp()
		want := e.exp

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var getcompcodetests = []struct {
	ins string
	exp string
}{
	{"0", "0101010"},
	{"AMD=0;JMP", "0101010"},
	{"D=A", "0110000"},
	{"M=D", "0001100"},
	{"D=D-M", "1010011"},
	{"0;JMP", "0101010"},
	{"D;JGT", "0001100"},
}

func TestGetCompCode(t *testing.T) {
	for _, e := range getcompcodetests {
		p := parser{e.ins}

		got, err := p.GetCompCode()
		want := e.exp

		if err != nil {
			t.Errorf("%v", err)
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var getjumptests = []struct {
	ins string
	exp string
}{
	{"0", ""},
	{"AMD=0;JMP", "JMP"},
	{"D=A", ""},
	{"M=D", ""},
	{"D=D-M", ""},
	{"0;JMP", "JMP"},
	{"D;JGT", "JGT"},
	{"0;JLT", "JLT"},
	{"D;JNE", "JNE"},
	{"AMD=D-M;JEQ", "JEQ"},
}

func TestGetJump(t *testing.T) {
	for _, e := range getjumptests {
		p := parser{e.ins}

		got := p.GetJump()
		want := e.exp

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var getjumpcodetests = []struct {
	ins string
	exp string
}{
	{"0", "000"},
	{"AMD=0;JMP", "111"},
	{"D=A", "000"},
	{"M=D", "000"},
	{"D=D-M", "000"},
	{"0;JMP", "111"},
	{"D;JGT", "001"},
	{"0;JLT", "100"},
	{"D;JNE", "101"},
	{"AMD=D-M;JEQ", "010"},
}

func TestGetJumpCode(t *testing.T) {
	for _, e := range getjumpcodetests {
		p := parser{e.ins}

		got, err := p.GetJumpCode()
		want := e.exp

		if err != nil {
			t.Errorf("%v", err)
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var tobinarytests = []struct {
	ins string
	exp string
}{
	{"@0", "0000000000000000"},
	{"@1", "0000000000000001"},
	{"@9", "0000000000001001"},
	{"@50", "0000000000110010"},
	{"@4099", "0001000000000011"},
	{"@32767", "0111111111111111"},

	{"0", "1110101010000000"},
	{"1", "1110111111000000"},
	{"D", "1110001100000000"},
	{"M", "1111110000000000"},

	{"D=D-M", "1111010011010000"},
	{"MD=D-M", "1111010011011000"},

	{"D;JGT", "1110001100000001"},
	{"D;JMP", "1110001100000111"},
	{"0;JLT", "1110101010000100"},
	{"-M;JNE", "1111110011000101"},

	{"AMD=A-D;JEQ", "1110000111111010"},
}

func TestToBinary(t *testing.T) {
	for _, e := range tobinarytests {
		p := parser{e.ins}

		got, err := p.ToBinary()
		want := e.exp

		if err != nil {
			t.Errorf("%v", err)
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}
