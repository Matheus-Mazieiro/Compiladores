package main

import (
	"fmt"
	"go-lexer/lexer" // altere conforme seu módulo
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type MyErrorListener struct {
	*antlr.DefaultErrorListener
	output *os.File
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong number of parameters\nPlease run \"go run inputFile outputFile\"")
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de entrada:", err)
		return
	}

	output, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de saida:", err)
		return
	}
	defer output.Close()

	lex := lexer.NewCalcLexer(input)
	lex.RemoveErrorListeners()
	listener := &MyErrorListener{
		output: output,
	}
	lex.AddErrorListener(listener)

	for token := lex.NextToken(); token.GetTokenType() != antlr.TokenEOF; token = lex.NextToken() {

		if token.GetTokenType() == antlr.TokenInvalidType {
			break
		}

		if lex.SymbolicNames[token.GetTokenType()] == "ERRO_COMENTARIO" {
			fmt.Fprintf(output, "Linha %d: comentario nao fechado\n", token.GetLine())
			break
		}

		if lex.SymbolicNames[token.GetTokenType()] == "ERRO_CADEIA" {
			fmt.Fprintf(output, "Linha %d: cadeia literal nao fechada\n", token.GetLine())
			break
		}

		if lex.SymbolicNames[token.GetTokenType()] == "RELAC" ||
			lex.SymbolicNames[token.GetTokenType()] == "ARIT" ||
			lex.SymbolicNames[token.GetTokenType()] == "LOGIC" {
			fmt.Fprintf(output, "<'%s','%s'>\n", token.GetText(), token.GetText())
			continue
		}

		name := getLiteralName(token.GetTokenType(), lex.LiteralNames, lex.SymbolicNames)
		fmt.Fprintf(output, "<'%s',%s>\n", token.GetText(), name)

	}
}

func getLiteralName(ttype int, literalNames, symbolicNames []string) string {
	if ttype >= 0 && ttype < len(literalNames) && literalNames[ttype] != "" {
		return literalNames[ttype]
	}

	if ttype >= 0 && ttype < len(symbolicNames) && symbolicNames[ttype] != "" {
		return symbolicNames[ttype]
	}

	return "UNKNOWN"
}

func (l *MyErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	const prefix = "token recognition error at: "

	if strings.HasPrefix(msg, prefix) {
		invalidChar := strings.Trim(strings.TrimPrefix(msg, prefix), "'")
		fmt.Fprintf(l.output, "Linha %d: %s - simbolo nao identificado\n", line, invalidChar)
	} else {
		// Caso seja erro sintático ou outro erro
		fmt.Fprintf(l.output, "Erro na linha %d, coluna %d: %s\n", line, column, msg)
	}
	os.Exit(0)
}
