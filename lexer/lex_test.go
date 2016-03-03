package lexer

import (
	"fmt"
	"strings"
	"testing"
)

type expectedCase struct {
	ExpectedTok Token
	ExpectedLit string
}

type TestCase struct {
	Test     string
	Expected []expectedCase
}

func TestScan(t *testing.T) {
	var testCases = []TestCase{
		{
			Test: "This is a valid doc.",
			Expected: []expectedCase{
				{Word, "This"},
				{WS, " "},
				{Word, "is"},
				{WS, " "},
				{Word, "a"},
				{WS, " "},
				{Word, "valid"},
				{WS, " "},
				{Word, "doc."},
				{WS, " "},
				{EOF, ""},
			},
		},
		{
			Test: "This is a  [*  valid  *]\n  doc.",
			Expected: []expectedCase{
				{Word, "This"},
				{WS, " "},
				{Word, "is"},
				{WS, " "},
				{Word, "a"},
				{WS, " "},
				{Word, "valid"},
				{WS, " "},
				{Word, "doc."},
				{WS, " "},
				{EOF, ""},
			},
		},
		{
			Test: `
«front-matter 
| Author : Gautam Dey <gautam.dey77@gmail.com>
| Date   : 2 Jan 2016
;»

§ Here is a section

Hello this is something to talking to you about something.

§§ Section 2

This is section [_  two  _]. Sections are delimited with a \§ character.
`,
		},
	}

	for _, c := range testCases {
		s := NewScanner(strings.NewReader(c.Test))
		for tok, lit := s.Scan(); tok != EOF; tok, lit = s.Scan() {
			fmt.Printf("Tok: %v, lit: “%v”\n", tok, lit)
		}
	}
}
