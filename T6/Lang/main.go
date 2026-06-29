package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

const (
	lexer_print = false
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong number of parameters\nPlease run \"go run inputFile outputFile\"")
	}

	//Abre arquivo de entrada
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de entrada:", err)
		return
	}

	//Arquivo de saida
	output, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de saida:", err)
		return
	}
	defer output.Close()

	lex := NewLexicAnal(output)

	if errs := lex.Anal(input); len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e.Msg)
		}
		return
	}

	// Reabre arquivo para parser
	input, err = antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de entrada:", err)
		return
	}

	tree, errs := lex.Parser(input)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e.Msg)
		}
		return
	}

	// Analise Semantica
	sem := NewPlSemantic()
	tree.Accept(sem)
	if len(sem.erros) > 0 {
		for _, e := range sem.erros {
			fmt.Fprintf(output, "Linha %d: %s\n", e.Linha, e.Msg)
		}
		return
	}

	fmt.Println("END\n\n")
}

// Necessário para evitar index out of bounds ao recuperar os nomes
func getLiteralName(ttype int, literalNames, symbolicNames []string) string {
	if ttype >= 0 && ttype < len(literalNames) && literalNames[ttype] != "" {
		return literalNames[ttype]
	}

	if ttype >= 0 && ttype < len(symbolicNames) && symbolicNames[ttype] != "" {
		return symbolicNames[ttype]
	}

	return "UNKNOWN"
}
