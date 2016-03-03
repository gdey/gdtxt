//go:generate stringer -type=Token
package lexer

type Token uint8

const (

	// Special Tokens
	uninit_ Token = iota
	ILLEGAL
	EOF
	WS
	NL
	Pipe
	CLoseBracket
	Colon
	Semicolon

	DocumentStartToken

	Word
	EmptyLine

	InLineElementStart

	InLineEmphasizeStart     // “[* ”
	InLineEmphasizeEnd       // “ *]”
	InLineStrongStart        // “[** ”
	InLineStrongEnd          // “ **]”
	InLineStrikeThroughStart // “[- ”
	InLineStrikeThroughEnd   // “ -]”
	InLineUnderlineStart     // “[_ ”
	InLineUnderlineEnd       // “ _]”
	InLineInsertStart        // “[@ ”
	InLineInsertEnd          // “ @]”
	InLineAnchorStart        // “[#[ ”
	InLineReferenceStart     // “[[ ”
	InLineReferenseMiddle    // “ ][ ”
	InLineReferenceEnd       // “ ]]”

	InLineElementEnd

	LineElementStart

	LineElementSection       // “§”
	LineElementUnorderedList // “• ”
	LineElementOrderedList   // “#”
	LineElementList          // “#”|“•”
	LineElementOrderedCheck  // LineElementList,“[”,WS|[“x”|“X”],“]”
	LineElementHorzontalLine

	LineElementEnd

	BlockElementStart // “«”
	BlockElementEnd   // “»”
)

var inLineStart = '['
var inLineEnd = "*-_@]"
var blockStart = '«'
var blockEnd = '»'
var lineElementStart = "§•#"
var eof = rune(0)
