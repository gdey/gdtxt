// Code generated by "stringer -type=Token"; DO NOT EDIT

package lexer

import "fmt"

const _Token_name = "uninit_ILLEGALEOFWSNLPipeCLoseBracketColonSemicolonDocumentStartTokenWordEmptyLineInLineElementStartInLineEmphasizeStartInLineEmphasizeEndInLineStrongStartInLineStrongEndInLineStrikeThroughStartInLineStrikeThroughEndInLineUnderlineStartInLineUnderlineEndInLineInsertStartInLineInsertEndInLineAnchorStartInLineReferenceStartInLineReferenseMiddleInLineReferenceEndInLineElementEndLineElementStartLineElementSectionLineElementUnorderedListLineElementOrderedListLineElementListLineElementOrderedCheckLineElementHorzontalLineLineElementEndBlockElementStartBlockElementEnd"

var _Token_index = [...]uint16{0, 7, 14, 17, 19, 21, 25, 37, 42, 51, 69, 73, 82, 100, 120, 138, 155, 170, 194, 216, 236, 254, 271, 286, 303, 323, 344, 362, 378, 394, 412, 436, 458, 473, 496, 520, 534, 551, 566}

func (i Token) String() string {
	if i >= Token(len(_Token_index)-1) {
		return fmt.Sprintf("Token(%d)", i)
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}