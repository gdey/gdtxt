package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r       *bufio.Reader
	lastTok Token
	lastLit string
}

// NewScanner returns a new instance of Scanner
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

/*
«front-matter
| author : Gautam Dey <gautam.dey77@gmail.com>
| date : 16 Dec 2015
;
»

   This is a paragraph it is ended by a newline [&#13&] charater.

§ This is a section.

• This is a bullet point
*/

/*
We have six main types of tokens.
Documents contain Paragraphs, Line Elements and Blocks
Paragraphs contain Lines
Lines contain inline elements.
LineElements contain a type and a line.
BlockElements contain a type, headers, and body
Type is a set of characters.
Headers contain a headerline element.
Headerline element contains Key and a Value
Key is a set of characters not containing ":"
Value is a set of characters till the first new line.
Body is all characters till the matching close '»' character, there may be sets
of “«” and “»” in the body.
*/

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	tok, lit = s.lastTok, s.lastLit
	if tok == uninit_ {
		tok, lit = s.scan()
	}
	if tok == EOF {
		return
	}
	s.lastTok, s.lastLit = s.scan()
	if tok == Word {
		for s.lastTok == Word {
			lit += s.lastLit
			s.lastTok, s.lastLit = s.scan()
		}
	}
	if tok == ILLEGAL {
	}
	return
}

// scan returns the next token and literal value.
func (s *Scanner) scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()
	if isNewLine(ch) {
		s.unread()
		return s.scanNewLine()
	}
	if isWhiteSpace(ch) {
		s.unread()
		return s.scanWhitespace()
	}
	if isWord(ch) {
		s.unread()
		return s.scanWord()
	}
	switch ch {
	case '\\':
		return Word, string(s.read())
	case eof:
		return EOF, ""
	case '«':
		return BlockElementStart, string(ch)
	case '»':
		return BlockElementEnd, string(ch)
	case '[':
		s.unread()
		return s.scanStartInline()
	case '*', '-', '_', ']', '@':
		s.unread()
		return s.scanEndInLine()
	case '|':
		return Pipe, string(ch)
	case ';':
		return Semicolon, string(ch)
	case ':':
		return Colon, string(ch)
	case '§':
		return LineElementSection, string(ch)
	case '•':
		return LineElementUnorderedList, string(ch)
	case '#':
		return LineElementOrderedList, string(ch)
	}
	return Word, string(ch)

}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// scanNewLine consume the current rune and all runes that would make up a single new line.
func (s *Scanner) scanNewLine() (tok Token, lit string) {
	ch := s.read()
	var buff bytes.Buffer
	buff.WriteRune(ch)
	if ch == '\r' {
		ch := s.read()
		if ch == '\n' {
			buff.WriteRune(ch)
		} else {
			s.unread()
		}
	}
	return NL, buff.String()

}

// scanStartInline consumes the current rune and all contigious start inline runes. The start of an inline element is always “[” but this alone may be just a “[” character.
func (s *Scanner) scanStartInline() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read()) // Assume we are on “[”
	ch := s.read()
	buf.WriteRune(ch)
	switch ch {
	case '*':
		ch = s.read()
		if ch == '*' {
			buf.WriteRune(ch)
			return InLineStrongStart, buf.String()
		}
		s.unread()
		return InLineEmphasizeStart, buf.String()
	case '-':
		return InLineStrikeThroughStart, buf.String()
	case '_':
		return InLineUnderlineStart, buf.String()
	case '@':
		return InLineInsertStart, buf.String()
	case '[':
		return InLineReferenceStart, buf.String()
	case '#':
		if ch == '[' {
			buf.WriteRune(ch)
			return InLineAnchorStart, buf.String()
		}
		s.unread()
		return Word, buf.String()
	}
	s.unread()
	return Word, buf.String()
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buff bytes.Buffer
	for ch := s.read(); isWhiteSpace(ch); ch = s.read() {
		buff.WriteRune(ch)
	}
	s.unread()
	return WS, buff.String()
}

// scanEndInline consumes the current rune and all contigious end inline runes.
func (s *Scanner) scanEndInLine() (tok Token, lit string) {
	var buff bytes.Buffer
	ch := s.read()
	buff.WriteRune(ch) // Assume we are on a valid end inline marker.
	switch ch {
	case '*':
		ch = s.read()
		if ch == '*' {
			buff.WriteRune(ch)
			tok = InLineStrongEnd
		} else {
			s.unread()
			tok = InLineEmphasizeEnd
		}
	case '-':
		tok = InLineStrikeThroughEnd
	case '_':
		tok = InLineUnderlineEnd
	case ']':
		ch = s.read()
		if ch == '[' {
			buff.WriteRune(ch)
			return InLineReferenseMiddle, buff.String()
		}
		s.unread()
		tok = InLineReferenceEnd
	default:
		return Word, buff.String()
	}
	ch = s.read()
	if ch == ']' {
		buff.WriteRune(ch)
		lit = buff.String()
		return
	}
	s.unread()
	if tok == InLineReferenceEnd {
		return CLoseBracket, buff.String()
	}
	return Word, buff.String()
}

// scanEndInline consumes the current rune and all contigious end inline runes.
func (s *Scanner) scanWord() (tok Token, lit string) {
	var buff bytes.Buffer
	for ch := s.read(); isWord(ch); ch = s.read() {
		buff.WriteRune(ch)
	}
	s.unread()
	return Word, buff.String()
}

// scanWord consumes the current rune and all contigious rune that make up a word, which is anything that does not make up other things.

// isWord will return true if the rune is not a special character, and can only be a word character.
func isWord(ch rune) bool {
	return !isNewLine(ch) &&
		!isWhiteSpace(ch) &&
		!isInLine(ch) &&
		!isLineStart(ch) &&
		!isBlock(ch) &&
		!unicode.IsControl(ch) &&
		!strings.ContainsRune("\\|];:", ch)
}

// isInLine return true if the rune could be an inline start or end char.
func isInLine(ch rune) bool      { return isInLineStart(ch) || isInLineEnd(ch) }
func isInLineStart(ch rune) bool { return ch == inLineStart }
func isInLineEnd(ch rune) bool   { return strings.ContainsRune(inLineEnd, ch) }

// isLineStart return true if the rune could be an line start char.
func isLineStart(ch rune) bool { return strings.ContainsRune(lineElementStart, ch) }

func isBlock(ch rune) bool { return ch == '«' || ch == '»' }

// isNewLine returns true if the rune is considered a newline.
func isNewLine(ch rune) bool { return ch == '\n' || ch == '\r' }

// isSpace returns true if the rune is consider a whitespace as
// described by the unicode package. But not '\n' and '\r'
func isSpace(ch rune) bool      { return !isNewLine(ch) && unicode.IsSpace(ch) }
func isWhiteSpace(ch rune) bool { return unicode.IsSpace(ch) }
