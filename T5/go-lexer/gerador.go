package main

import (
	parser "go-lexer/parser"
	"strings"
)

type JanderGerador struct {
	parser.BaseCalcLexerVisitor
	tabela *TabelaDeSimbolos
	saida  strings.Builder
}

func NewJanderGerador(
	tabela *TabelaDeSimbolos,
) *JanderGerador {

	return &JanderGerador{
		tabela: tabela,
	}
}

// Gerar inicia a geração do código C
func (g *JanderGerador) Gerar(
	ctx *parser.ProgramaContext,
) string {

	g.saida.WriteString("#include <stdio.h>\n")
	g.saida.WriteString("#include <stdlib.h>\n")
	g.saida.WriteString("\n")

	g.saida.WriteString("int main() {\n")

	// declarações globais/localizadas no corpo
	if ctx.Corpo() != nil {
		ctx.Corpo().Accept(g)
	}

	g.saida.WriteString("\treturn 0;\n")
	g.saida.WriteString("}\n")

	return g.saida.String()
}

// =========================
// CORPO
// =========================

func (g *JanderGerador) VisitCorpo(
	ctx *parser.CorpoContext,
) interface{} {

	// declarações
	for _, d := range ctx.AllDeclaracao_local() {
		d.Accept(g)
	}

	// comandos
	for _, c := range ctx.AllCmd() {
		c.Accept(g)
	}

	return nil
}

// =========================
// DECLARAÇÕES
// =========================

func (g *JanderGerador) VisitDeclaracao_local(
	ctx *parser.Declaracao_localContext,
) interface{} {

	if ctx.Variavel() != nil {
		ctx.Variavel().Accept(g)
	}

	return nil
}

func (g *JanderGerador) VisitVariavel(
	ctx *parser.VariavelContext,
) interface{} {

	tipo := VerificarTipo(g.tabela, ctx.Tipo())

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.GetText()

		switch tipo {

		case INTEIRO:

			g.saida.WriteString(
				"\tint " + nome + ";\n",
			)

		case REAL:

			g.saida.WriteString(
				"\tfloat " + nome + ";\n",
			)

		case LITERAL:

			g.saida.WriteString(
				"\tchar " + nome + "[80];\n",
			)

		case LOGICO:

			g.saida.WriteString(
				"\tint " + nome + ";\n",
			)
		}
	}

	return nil
}

// =========================
// COMANDOS
// =========================

func (g *JanderGerador) VisitCmd(
	ctx *parser.CmdContext,
) interface{} {

	if ctx.CmdLeia() != nil {
		return ctx.CmdLeia().Accept(g)
	}

	if ctx.CmdEscreva() != nil {
		return ctx.CmdEscreva().Accept(g)
	}

	if ctx.CmdAtribuicao() != nil {
		return ctx.CmdAtribuicao().Accept(g)
	}

	if ctx.CmdSe() != nil {
		return ctx.CmdSe().Accept(g)
	}

	if ctx.CmdEnquanto() != nil {
		return ctx.CmdEnquanto().Accept(g)
	}

	if ctx.CmdFaca() != nil {
		return ctx.CmdFaca().Accept(g)
	}

	return nil
}

// =========================
// LEIA
// =========================

func (g *JanderGerador) VisitCmdLeia(
	ctx *parser.CmdLeiaContext,
) interface{} {

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.GetText()

		tipo := g.tabela.Verificar(nome)

		switch tipo {

		case INTEIRO:

			g.saida.WriteString(
				"\tscanf(\"%d\", &" + nome + ");\n",
			)

		case REAL:

			g.saida.WriteString(
				"\tscanf(\"%f\", &" + nome + ");\n",
			)

		case LITERAL:

			g.saida.WriteString(
				"\tgets(" + nome + ");\n",
			)

		case LOGICO:

			g.saida.WriteString(
				"\tscanf(\"%d\", &" + nome + ");\n",
			)
		}
	}

	return nil
}

// =========================
// ESCREVA
// =========================

func (g *JanderGerador) VisitCmdEscreva(
	ctx *parser.CmdEscrevaContext,
) interface{} {

	for _, expr := range ctx.AllExpressao() {

		texto := TraduzExpressao(
			expr.GetText(),
		)

		tipo := VerificarExpressao(
			g.tabela,
			expr,
		)

		switch tipo {

		case INTEIRO:

			g.saida.WriteString(
				"\tprintf(\"%d\", " + texto + ");\n",
			)

		case REAL:

			g.saida.WriteString(
				"\tprintf(\"%f\", " + texto + ");\n",
			)

		case LITERAL:

			g.saida.WriteString(
				"\tprintf(\"%s\", " + texto + ");\n",
			)

		case LOGICO:

			g.saida.WriteString(
				"\tprintf(\"%d\", " + texto + ");\n",
			)
		}
	}

	return nil
}

// =========================
// ATRIBUIÇÃO
// =========================

func (g *JanderGerador) VisitCmdAtribuicao(
	ctx *parser.CmdAtribuicaoContext,
) interface{} {

	nome := ctx.Identificador().GetText()

	expr := TraduzExpressao(
		ctx.Expressao().GetText(),
	)

	g.saida.WriteString(
		"\t" + nome + " = " + expr + ";\n",
	)

	return nil
}

// =========================
// SE
// =========================

func (g *JanderGerador) VisitCmdSe(
	ctx *parser.CmdSeContext,
) interface{} {

	condicao := TraduzExpressao(
		ctx.Expressao().GetText(),
	)

	g.saida.WriteString(
		"\tif (" + condicao + ") {\n",
	)

	cmds := ctx.AllCmd()

	temSenao := ctx.SENAO() != nil

	if temSenao {

		meio := len(cmds) / 2

		for i := 0; i < meio; i++ {
			cmds[i].Accept(g)
		}

		g.saida.WriteString("\t}\n")
		g.saida.WriteString("\telse {\n")

		for i := meio; i < len(cmds); i++ {
			cmds[i].Accept(g)
		}

		g.saida.WriteString("\t}\n")

	} else {

		for _, c := range cmds {
			c.Accept(g)
		}

		g.saida.WriteString("\t}\n")
	}

	return nil
}

// =========================
// ENQUANTO
// =========================

func (g *JanderGerador) VisitCmdEnquanto(
	ctx *parser.CmdEnquantoContext,
) interface{} {

	condicao := TraduzExpressao(
		ctx.Expressao().GetText(),
	)

	g.saida.WriteString(
		"\twhile (" + condicao + ") {\n",
	)

	for _, c := range ctx.AllCmd() {
		c.Accept(g)
	}

	g.saida.WriteString("\t}\n")

	return nil
}

// =========================
// FAÇA
// =========================

func (g *JanderGerador) VisitCmdFaca(
	ctx *parser.CmdFacaContext,
) interface{} {

	g.saida.WriteString("\tdo {\n")

	for _, c := range ctx.AllCmd() {
		c.Accept(g)
	}

	condicao := TraduzExpressao(
		ctx.Expressao().GetText(),
	)

	g.saida.WriteString(
		"\t} while (" + condicao + ");\n",
	)

	return nil
}

// =========================
// EXPRESSÕES
// =========================

func TraduzExpressao(expr string) string {

	// operadores lógicos
	expr = strings.ReplaceAll(
		expr,
		" e ",
		" && ",
	)

	expr = strings.ReplaceAll(
		expr,
		" ou ",
		" || ",
	)

	expr = strings.ReplaceAll(
		expr,
		"nao",
		"!",
	)

	// booleanos
	expr = strings.ReplaceAll(
		expr,
		"verdadeiro",
		"1",
	)

	expr = strings.ReplaceAll(
		expr,
		"falso",
		"0",
	)

	// atribuição
	expr = strings.ReplaceAll(
		expr,
		"<-",
		"=",
	)

	return expr
}
