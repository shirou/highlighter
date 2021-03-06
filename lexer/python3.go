package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	python3 "github.com/shirou/antlr-grammars-v4-go/python3"

	"github.com/shirou/golightan"
)

type Python3Lexer struct {
	lexer    antlr.Lexer
	tokenMap TokenMap
}

func (l Python3Lexer) Tokenize(input antlr.CharStream) (golightan.Tokens, error) {
	le := python3.NewPython3Lexer(input)
	stream := NewAllTokenStream(le)
	p := python3.NewPython3Parser(stream)

	// TODO: error handling
	//	p.SetErrorHandler(golightan.NewNullErrorStrategy())
	//p.RemoveErrorListeners()

	listener := NewCommonParseTreeListener(l.tokenMap)
	tree := p.File_input()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetTokens(), nil
}

func NewPython3Lexer() Lexer {
	return Python3Lexer{
		tokenMap: NewPython3TokenMap(),
	}
}
