/*
The parser package will take a stream of tokens and transform them to a document.
*/
package parser

import (
	"io"

	"github.com/gdey/gdtxt/lexer"
)

type token struct {
	tok lexer.Token
	lit string
}

//
type Parse struct {
	s   *lexer.Scanner
	buf struct {
		tok     lexer.Token
		lit     string
		haveTok bool
	}
}

// New Parser returns a new instance of Parser
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok lexer.Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.haveTok {
		p.buf.haveTok = false
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buff.haveTok = true }

// scanIgnoreWhitespace scan the next non-whitesapce token. The supplied variable
// tells it to ignore NL as well.
func (p *Parser) scanIgnoreWhitespace(ignoreNL bool) (tok lexer.Token, lit string) {
	for tok, lit = p.scan(); tok == lexer.WS || (ignoreNL && tok == lexer.NL); tok, lit = p.scan() {
	}
	return
}
