package main

import (
	"fmt"
	lexer "go-lexer/parser"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

const (
	debugT1 = false
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
	lex := lexer.NewCalcLexerLexer(input)

	//Remove o listener padrão para poder controlar o que será impresso no caso de nao encontrar token
	//E adiciona o listener próprio.
	//Note que ele contem um arquivo no qual será escrito no caso de erro
	lex.RemoveErrorListeners()
	listener := &MyErrorListener{
		output: output,
	}
	lex.AddErrorListener(listener)

	erroLexico := false

	//Armazenar os tokens p o parser
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	//Itera pelos tokens (lexer)
	i := 0
	for {
		// Pega o token (força carregamento no buffer)
		token := tokens.LT(i + 1)

		// Para se chegar no fim do arquivo
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}

		//Acaba a execução caso seja um token invalido
		//O erro surgido é jogado para o listener, que exibe para o usuário
		if token.GetTokenType() == antlr.TokenInvalidType {
			break
		}

		//Caso o erro seja um comentário não fechado, ele será capturado pela
		//Gramática e retornará um "ERRO_COMENTARIO"
		if lex.SymbolicNames[token.GetTokenType()] == "ERRO_COMENTARIO" {
			fmt.Fprintf(output, "Linha %d: comentario nao fechado\n", token.GetLine())
			erroLexico = true
			break
		}

		//Caso o erro seja uma cadeia não fechada, ele será capturado pela
		//Gramática e retornará um "ERRO_CADEIA"
		if lex.SymbolicNames[token.GetTokenType()] == "ERRO_CADEIA" {
			fmt.Fprintf(output, "Linha %d: cadeia literal nao fechada\n", token.GetLine())
			erroLexico = true
			break
		}

		if lex.SymbolicNames[token.GetTokenType()] == "ERRO_DEMAIS" {
			fmt.Fprintf(output, "Linha %d: %s - simbolo nao identificado\n", token.GetLine(), token.GetText())
			erroLexico = true
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
		if debugT1 {
			fmt.Fprintf(output, "<'%s',%s>\n", token.GetText(), name)
		}
		i++
	}

	// Parser: só é executado se n tem erros lexicos
	if !erroLexico {
		//volta para o início dos tokens
		tokens.Seek(0)

		p := lexer.NewCalcLexerParser(tokens)

		// Adiciona o listener de erro no parser
		p.RemoveErrorListeners()
		p.AddErrorListener(listener)

		//Chama a primeira regra do arquivo .g4
		p.Programa()
	} else {
		fmt.Fprintf(output, "Fim da compilacao\n")
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

func (l *MyErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {

	// Recupera token que causou o erro
	token := offendingSymbol.(antlr.Token)
	tokenText := token.GetText()

	//Formatando EOF p impressao
	if tokenText == "<EOF>" {
		tokenText = "EOF"
	}

	// Linha X: erro sintatico proximo a TOKEN
	fmt.Fprintf(l.output, "Linha %d: erro sintatico proximo a %s\n", line, tokenText)
	fmt.Fprintf(l.output, "Fim da compilacao\n")

	// Para no primeiro erro sintático
	l.output.Close()
	os.Exit(0)
}
