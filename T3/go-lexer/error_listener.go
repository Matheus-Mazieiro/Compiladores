// LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

type MyCustomErrorListener struct {
	*antlr.DefaultErrorListener
	output *os.File
	flag   bool
}

func NewErrorListener(output *os.File) *MyCustomErrorListener {
	return &MyCustomErrorListener{
		output: output,
		flag:   false,
	}
}

func (l *MyCustomErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {

	if l.flag {
		return
	}

	token := offendingSymbol.(antlr.Token)
	tokenText := token.GetText()

	if tokenText == "<EOF>" {
		tokenText = "EOF"
	}

	fmt.Fprintf(l.output, "Linha %d: erro sintatico proximo a %s\n", line, tokenText)
	fmt.Fprintln(l.output, "Fim da compilacao")

	l.flag = true
	os.Exit(0)
}