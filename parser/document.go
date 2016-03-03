package parser

type Node interface {
}

// Document can have three types of subnodes.
//
type Document struct {
	Filename string
	Children []Node
}

type Block struct {
	Type    string
	Headers map[string]string
	Body    string // Body is everything till end », we need to keep track of matching «»
}

type LineElement struct {
	Type     ElementType
	Line     string
	Children []Node
}

type InLineElement struct {
	Type     string
	elements []Node
}

type ParagraphElement struct {
	lines    []string
	elements []Node
}
