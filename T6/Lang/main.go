package main

import (
	"fmt"
	"os"

	"github.com/matheus-mazieiro/pl-compiler/parser"

	"github.com/antlr4-go/antlr/v4"
)

// Estrutura que captará o erro, necessária para pegar o 'token not found'
type MyErrorListener struct {
	*antlr.DefaultErrorListener
	output *os.File
}

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

	//Objeto do lexer
	lex := parser.NewGrammarLexer(input)

	//Remove o listener padrão para poder controlar o que será impresso no caso de nao encontrar token
	//E adiciona o listener próprio.
	//Note que ele contem um arquivo no qual será escrito no caso de erro
	lex.RemoveErrorListeners()
	listener := &MyErrorListener{
		output: output,
	}
	lex.AddErrorListener(listener)

	//Itera pelos tokens
	for token := lex.NextToken(); token.GetTokenType() != antlr.TokenEOF; token = lex.NextToken() {

		//Acaba a execução caso seja um token invalido
		//O erro surgido é jogado para o listener, que exibe para o usuário
		if token.GetTokenType() == antlr.TokenInvalidType {
			break
		}

		//Caso seja algum operador, derá ser impresso o prórprio operador
		if lex.SymbolicNames[token.GetTokenType()] == "RELAC" ||
			lex.SymbolicNames[token.GetTokenType()] == "ARIT" ||
			lex.SymbolicNames[token.GetTokenType()] == "LOGIC" {
			fmt.Fprintf(output, "<'%s','%s'>\n", token.GetText(), token.GetText())
			continue
		}

		//Demais tokens
		name := getLiteralName(token.GetTokenType(), lex.LiteralNames, lex.SymbolicNames)
		fmt.Fprintf(output, "<'%s',%s>\n", token.GetText(), name)

	}
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
