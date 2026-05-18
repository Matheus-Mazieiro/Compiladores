// LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
package main

import (
	"fmt"
	parser "go-lexer/parser"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go entrada.txt saida.txt")
		return
	}

	input, _ := antlr.NewFileStream(os.Args[1])
	output, _ := os.Create(os.Args[2])
	defer output.Close()

	// LEXER
	lex := parser.NewCalcLexerLexer(input)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(NewErrorListener(output))

	// TOKENS
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	// PARSER
	p := parser.NewCalcLexerParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(NewErrorListener(output))

	// PARSE TREE
	tree := p.Programa()

	// SEMÂNTICO
	sem := NewJanderSemantico()
	tree.Accept(sem)

	// SAÍDA
	if len(ErrosSemanticos) > 0 {
		for _, e := range ErrosSemanticos {
			fmt.Fprintln(output, e)
		}
	} else {
		fmt.Fprintln(output, "Nenhum erro semantico encontrado.")
	}

	fmt.Fprintln(output, "Fim da compilacao")
}