package main

import (
	"fmt"
	parser "go-lexer/parser"
)

type JanderSemantico struct {
	parser.BaseCalcLexerVisitor
	tabela *TabelaDeSimbolos
}

func NewJanderSemantico() *JanderSemantico {
	return &JanderSemantico{
		tabela: nil,
	}
}
func (j *JanderSemantico) VisitDecl_local_global(ctx *parser.Decl_local_globalContext) interface{} {
	fmt.Println("DEBUG: VisitDecl_local_global")

	if ctx.Declaracao_local() != nil {
		ctx.Declaracao_local().Accept(j)
	}
	return nil
}
func (j *JanderSemantico) VisitDeclaracao_local(ctx *parser.Declaracao_localContext) interface{} {
	fmt.Println("DEBUG: VisitDeclaracao_local")

	if ctx.Variavel() != nil {
		ctx.Variavel().Accept(j)
		return nil
	}

	if ctx.IDENT() != nil {
		nome := ctx.IDENT().GetText()
		fmt.Println("DEBUG: declarando tipo", nome)

		tipo := VerificarTipoBasico(ctx.Tipo_basico())

		if tipo == INVALIDO {
			AdicionarErroSemantico(ctx.Tipo_basico().GetStart().GetLine(),
				"tipo "+ctx.Tipo_basico().GetText()+" nao declarado")
		}

		if j.tabela.Existe(nome) {
			AdicionarErroSemantico(ctx.IDENT().GetSymbol().GetLine(),
				"identificador "+nome+" ja declarado anteriormente")
		} else {
			if tipo != INVALIDO {
				j.tabela.Adicionar(nome, tipo)
			}
		}
	}

	return nil
}
func (j *JanderSemantico) VisitDeclaracoes(ctx *parser.DeclaracoesContext) interface{} {
	fmt.Println("DEBUG: VisitDeclaracoes")

	for _, d := range ctx.AllDecl_local_global() {
		d.Accept(j)
	}

	return nil
}
func (j *JanderSemantico) VisitCorpo(ctx *parser.CorpoContext) interface{} {
	fmt.Println("DEBUG: VisitCorpo")

	for _, d := range ctx.AllDeclaracao_local() {
		d.Accept(j)
	}

	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	return nil
}
func (j *JanderSemantico) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	fmt.Println("DEBUG: VisitPrograma")
	j.tabela = NewTabelaDeSimbolos()

	if ctx.Declaracoes() != nil {
		ctx.Declaracoes().Accept(j)
	}

	// 🔥 VISITA CORPO
	if ctx.Corpo() != nil {
		ctx.Corpo().Accept(j)
	}

	return nil
}
func (j *JanderSemantico) VisitCmd(ctx *parser.CmdContext) interface{} {
	fmt.Println("DEBUG: VisitCmd")

	if ctx.CmdLeia() != nil {
		return ctx.CmdLeia().Accept(j)
	}
	if ctx.CmdEscreva() != nil {
		return ctx.CmdEscreva().Accept(j)
	}
	if ctx.CmdAtribuicao() != nil {
		return ctx.CmdAtribuicao().Accept(j)
	}

	return j.VisitChildren(ctx)
}

func (j *JanderSemantico) VisitVariavel(ctx *parser.VariavelContext) interface{} {
	fmt.Println("DEBUG: VisitVariavel")

	tipo := VerificarTipo(j.tabela, ctx.Tipo())

	for _, ident := range ctx.AllIdentificador() {
		nome := ident.GetText()
		fmt.Println("DEBUG: declarando", nome)

		if j.tabela.Existe(nome) {
			AdicionarErroSemantico(ident.GetStart().GetLine(),
				"identificador "+nome+" ja declarado anteriormente")
		} else {
			if tipo != INVALIDO {
				j.tabela.Adicionar(nome, tipo)
			}
		}
	}

	return nil
}
func (j *JanderSemantico) VisitCmdLeia(ctx *parser.CmdLeiaContext) interface{} {
	fmt.Println("DEBUG: VisitCmdLeia")

	for _, ident := range ctx.AllIdentificador() {
		nome := ident.GetText()
		fmt.Println("DEBUG: lendo", nome)

		if !j.tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(ident.GetStart().GetLine(),
					"identificador "+nome+" nao declarado")
			}
		}
	}

	return nil
}

func (j *JanderSemantico) VisitIdentificador(ctx *parser.IdentificadorContext) interface{} {

	nome := ctx.GetText()

	if variaveisComTipoInvalido[nome] {
		return nil
	}

	if !j.tabela.Existe(nome) {
		if !ErroJaReportado(nome) {
			AdicionarErroSemantico(ctx.GetStart().GetLine(),
				"identificador "+nome+" nao declarado")
		}
	}

	return nil
}