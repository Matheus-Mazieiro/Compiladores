package main

import (
	"fmt"
	parser "go-lexer/parser"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {

	// =========================
	// VALIDA ARGUMENTOS
	// =========================

	if len(os.Args) < 3 {

		fmt.Println(
			"Uso: go run main.go entrada.txt saida.txt",
		)

		return
	}

	arquivoEntrada := os.Args[1]
	arquivoSaida := os.Args[2]

	// =========================
	// ABRE ENTRADA
	// =========================

	input, err := antlr.NewFileStream(
		arquivoEntrada,
	)

	if err != nil {

		fmt.Println(
			"Erro ao abrir arquivo de entrada",
		)

		return
	}

	// =========================
	// CRIA SAÍDA
	// =========================

	output, err := os.Create(
		arquivoSaida,
	)

	if err != nil {

		fmt.Println(
			"Erro ao criar arquivo de saida",
		)

		return
	}

	defer output.Close()

	// =========================
	// LIMPA ERROS
	// =========================

	ResetarErros()

	// =========================
	// LEXER
	// =========================

	lex := parser.NewCalcLexerLexer(input)

	lex.RemoveErrorListeners()

	lex.AddErrorListener(
		NewErrorListener(output),
	)

	// =========================
	// TOKENS
	// =========================

	tokens := antlr.NewCommonTokenStream(
		lex,
		antlr.TokenDefaultChannel,
	)

	// =========================
	// PARSER
	// =========================

	p := parser.NewCalcLexerParser(tokens)

	p.RemoveErrorListeners()

	p.AddErrorListener(
		NewErrorListener(output),
	)

	// =========================
	// ÁRVORE SINTÁTICA
	// =========================

	tree := p.Programa()

	// =========================
	// ANÁLISE SEMÂNTICA
	// =========================

	sem := NewJanderSemantico()

	tree.Accept(sem)

	// =========================
	// SE TIVER ERROS:
	// imprime erros e encerra
	// =========================

	if len(ErrosSemanticos) > 0 {

		for _, e := range ErrosSemanticos {

			fmt.Fprintln(output, e)
		}

		fmt.Fprintln(
			output,
			"Fim da compilacao",
		)

		return
	}

	// =========================
	// GERAÇÃO DE CÓDIGO C
	// =========================

	gerador := NewGerador(
		sem.tabela,
	)

	codigoC := tree.Accept(gerador)
	fmt.Println("GERADOR EXECUTOU")
	// =========================
	// ESCREVE C NO ARQUIVO
	// =========================

	fmt.Fprint(output, codigoC)
}
