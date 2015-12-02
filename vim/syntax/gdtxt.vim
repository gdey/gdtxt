" Vim syntax file
" Language: Good (enough) Text format.
" Maintainer: Gautam Dey
" Latest Revision: 1 December 2015

if exists("b:current_syntax")
	finish
endif

" The syntax of gdtxt is made up three elements.
" Line elements, these start at the begining of a line with a sepecial set of
" characters.
"
" In-line elements these are made up of '[' + another character.
"
" Block elements thse start on a new line with the “«” and ends with a “»”
" character.
"
"
"

"runtime! syntax/html.vim
"unlet! b:current_syntax


syntax case ignore

set conceallevel=2

syntax cluster gdtxtInLineElements contains=gdtxtBold,gdtxtEscape,gdtxtStrong,gdtxtUnderline,gdtxtStrickthrough
syntax region gdtxtBold          matchgroup=boldEnds start=/\[\* /   end=/ \*\]/   contains=@gdtxtInLineElements concealends
syntax region gdtxtStrong        matchgroup=boldEnds start=/\[\*\* / end=/ \*\*\]/ contains=@gdtxtInLineElements concealends
syntax region gdtxtUnderline     matchgroup=boldEnds start=/\[_ /    end=/ _\]/    contains=@gdtxtInLineElements concealends
syntax region gdtxtStrickthrough matchgroup=boldEnds start=/\[- /    end=/ -\]/    contains=@gdtxtInLineElements concealends
syntax region gdtxtReference     start=/\[\[ /    end=/ \]\]/  keepend

syntax region gdtxtH1 matchgroup=gdtxtHeadingDelimiter start=/\s*§\s\+/      end=/\s*$/ oneline concealends
syntax region gdtxtH2 matchgroup=gdtxtHeadingDelimiter start=/\s*§§\s\+/     end=/\s*$/ oneline concealends
syntax region gdtxtH3 matchgroup=gdtxtHeadingDelimiter start=/\s*§§§\s\+/    end=/\s*$/ oneline concealends
syntax region gdtxtH4 matchgroup=gdtxtHeadingDelimiter start=/\s*§§§§\s\+/   end=/\s*$/ oneline concealends
syntax region gdtxtH5 matchgroup=gdtxtHeadingDelimiter start=/\s*§§§§§\s\+/  end=/\s*$/ oneline concealends
syntax region gdtxtH6 matchgroup=gdtxtHeadingDelimiter start=/\s*§§§§§§\s\+/ end=/\s*$/ oneline concealends

syntax region gdtxtCL1  matchgroup=gdtxtListDelimiter start=/\s*•\[\s*\]\s\+/            end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCCL1 matchgroup=gdtxtListDelimiter start=/\s*•\[\s*[Xx]\s*\]\s\+/     end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCL2  matchgroup=gdtxtListDelimiter start=/\s*••\s\+/                  end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCCL2 matchgroup=gdtxtListDelimiter start=/\s*••\[\s*[Xx]\s*\]\s\+/    end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCL3  matchgroup=gdtxtListDelimiter start=/\s*•••\s\+/                 end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCCL3 matchgroup=gdtxtListDelimiter start=/\s*•••\[\s*[Xx]\s*\]\s\+/   end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCL4  matchgroup=gdtxtListDelimiter start=/\s*••••\s\+/                end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCCL4 matchgroup=gdtxtListDelimiter start=/\s*••••\[\s*[Xx]\s*\]\s\+/  end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCL5  matchgroup=gdtxtListDelimiter start=/\s*••••\s\+/                end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCCL5 matchgroup=gdtxtListDelimiter start=/\s*••••\[\s*[Xx]\s*\]\s\+/  end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCL6  matchgroup=gdtxtListDelimiter start=/\s*•••••\s\+/               end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtCCL6 matchgroup=gdtxtListDelimiter start=/\s*•••••\[\s*[Xx]\s*\]\s\+/ end=/\s*$/ oneline contains=@gdtxtInLineElements

syntax region gdtxtUL1 matchgroup=gdtxtListDelimiter start=/\s*•\s\+/     end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtUL2 matchgroup=gdtxtListDelimiter start=/\s*••\s\+/    end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtUL3 matchgroup=gdtxtListDelimiter start=/\s*•••\s\+/   end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtUL4 matchgroup=gdtxtListDelimiter start=/\s*••••\s\+/  end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtUL5 matchgroup=gdtxtListDelimiter start=/\s*••••\s\+/  end=/\s*$/ oneline contains=@gdtxtInLineElements
syntax region gdtxtUL6 matchgroup=gdtxtListDelimiter start=/\s*•••••\s\+/ end=/\s*$/ oneline contains=@gdtxtInLineElements


syntax region gdtxtFrontMatter matchgroup=gdtxtBlockEnds start=/^\s*«front-matter/ end=/\s*»\s*$/ contains=@gdtxtBlockHeader

"syntax cluster gdtxtBlocks contains=@gdtxtNoteTitle,gdtxtBlockHead,gdtxtBlockBod
"
"syntax region gdtxtBlock start=/^\s*«/ skip=/\\«\|\\»/ end=/\s*»\s*$/ contains=@gdtxtBlocks



syntax region gdtxtComment matchgroup=gdtxtBlockEnds start=/^\s*«comment/ end=/\s*»\s*$/ contains=@gdtxtBlockHeader


syntax match gdtxtNoteTitle /note/ nextgroup=@gdtxtBlockHead contained
syntax match gdtxtBlockHead /[^;]*/ nextgroup=gdtxtBlockBod contained
syntax region gdtxtBlockBod matchgroup=gdtxtBlockEnds start=/;/ skip=/\\;\|\\»\|\\«/ end=/\s*»\s*/ contained

" syntax region gdtxtNote matchgroup=gdtxtBlockEnds start=/\s*«note/ skip=/\\;/ end=/;/ contains=@gdtxtBlockHeader nextgroup=gdtxtBlockBod contained


syn match gdtxtEscape "\\."

highlight gdtxtBold          term=bold      cterm=bold              gui=bold
highlight gdtxtStrong        term=italic    cterm=italic            gui=italic
highlight gdtxtUnderline     term=underline cterm=underline         gui=underline
highlight gdtxtStrickthrough term=underline cterm=underline,reverse gui=underline,reverse

highlight gdtxtHeading   term=bold      cterm=bold      ctermfg=DarkBlue gui=bold      guifg=Orange
highlight gdtxtReference term=underline cterm=underline ctermfg=DarkBlue gui=underline guibg=Blue

highlight gdtxtSubHeading    term=bold  cterm=bold ctermfg=Blue     gui=bold guifg=Purple
highlight gdtxtSubSubHeading term=bold  cterm=bold ctermfg=DarkCyan gui=bold guifg=DarkCyan

highlight gdtxtListDelimiter term=bold cterm=bold gui=bold
highlight gdtxtListCL        term=bold cterm=bold ctermfg=Yellow gui=bold
highlight gdtxtListCCL       term=NONE cterm=NONE ctermfg=Gray   gui=NONE



highlight gdtxtComment  term=standout cterm=NONE ctermfg=Gray gui=NONE guifg=Gray

highlight gdtxtNoteTitle  term=standout cterm=NONE ctermfg=Blue gui=NONE guifg=Blue
highlight gdtxtBlockBod  term=standout cterm=NONE ctermfg=Green gui=NONE guifg=Blue

highlight gdtxtBlockHeader term=bold cterm=bold ctermfg=Green guifg=Green gui=bold




highlight def link gdtxtH1 gdtxtHeading
highlight def link gdtxtH2 gdtxtSubHeading
highlight def link gdtxtH3 gdtxtSubSubHeading
highlight def link gdtxtH4 gdtxtSubSubHeading
highlight def link gdtxtH5 gdtxtSubSubHeading
highlight def link gdtxtH6 gdtxtSubSubHeading

highlight def link gdtxtCCL1 gdtxtListCCL
highlight def link gdtxtCCL2 gdtxtListCCL
highlight def link gdtxtCCL3 gdtxtListCCL
highlight def link gdtxtCCL4 gdtxtListCCL
highlight def link gdtxtCCL5 gdtxtListCCL
highlight def link gdtxtCCL6 gdtxtListCCL

highlight def link gdtxtCL1 gdtxtListCL
highlight def link gdtxtCL2 gdtxtListCL
highlight def link gdtxtCL3 gdtxtListCL
highlight def link gdtxtCL4 gdtxtListCL
highlight def link gdtxtCL5 gdtxtListCL
highlight def link gdtxtCL6 gdtxtListCL

highlight def link gdtxtBlockEnds gdtxtComment



let b:current_syntax="gdtxt"

