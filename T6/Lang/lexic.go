package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/matheus-mazieiro/pl-compiler/parser"
)

type LexicAnal struct {
	el     *MyErrorListener
	output *os.File
	lex    *parser.GrammarLexer
}

type LexicErr struct {
	Msg    string
	Ln     int
	Simbol string
}

type MyErrorListener struct {
	*antlr.DefaultErrorListener
	output *os.File
	errors []LexicErr
}

func (l *MyErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.errors = append(l.errors, LexicErr{
		Msg: msg,
		Ln:  line,
	})

	fmt.Fprintf(l.output, "Linha %d:%d %s\n", line, column, msg)
}

func NewLexicAnal(output *os.File) *LexicAnal {
	return &LexicAnal{
		output: output,
	}
}

// Realiza a análise léxica e retorna os erros encontrados.
func (la *LexicAnal) Anal(input *antlr.FileStream) []LexicErr {
	la.lex = parser.NewGrammarLexer(input)

	// Remove o listener padrão
	la.lex.RemoveErrorListeners()

	// Cria o listener personalizado
	la.el = &MyErrorListener{
		output: la.output,
	}

	la.lex.AddErrorListener(la.el)

	// Consome todos os tokens
	for token := la.lex.NextToken(); token.GetTokenType() != antlr.TokenEOF; token = la.lex.NextToken() {

		// Se o lexer produziu um token inválido, interrompe
		if token.GetTokenType() == antlr.TokenInvalidType {
			break
		}

		if lexer_print {
			name := getLiteralName(
				token.GetTokenType(),
				la.lex.LiteralNames,
				la.lex.SymbolicNames,
			)

			fmt.Fprintf(la.output, "<'%s', %s>\n", token.GetText(), name)
		}
	}
	return la.el.errors
}

// Realiza a análise sintática.
func (la *LexicAnal) Parser(input *antlr.FileStream) []LexicErr {
	lexer := parser.NewGrammarLexer(input)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGrammarParser(tokens)

	p.RemoveErrorListeners()
	p.AddErrorListener(la.el)

	// Regra inicial
	p.Program()

	return la.el.errors
}
