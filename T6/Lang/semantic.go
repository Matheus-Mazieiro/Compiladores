package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/matheus-mazieiro/pl-compiler/parser"
)

const DEBUG = false

type ErroSemantico struct {
	Msg   string
	Linha int
}

type PlSemantic struct {
	parser.BaseGrammarVisitor
	tabela *TabelaDeSimbolos
	erros  []ErroSemantico
}

func NewPlSemantic() *PlSemantic {
	return &PlSemantic{
		tabela: nil,
	}
}

func debug(s string) {
	if DEBUG {
		fmt.Println("Visit", s)
	}
}

func (s *PlSemantic) erro(msg string, linha int) {
	s.erros = append(s.erros,
		ErroSemantico{
			msg,
			linha,
		})
}

func (s *PlSemantic) declarar(
	nome string,
	tipo TipoSimbolo,
	dimensoes int,
	dominios []string,
	linha int,
) {

	if !s.tabela.Adicionar(nome, tipo, dimensoes, dominios) {
		s.erro(fmt.Sprintf("'%s' já foi declarado", nome), linha)
	}
}

func (s *PlSemantic) VisitProgram(ctx *parser.ProgramContext) interface{} {
	debug("Program")

	s.tabela = NewTabela()

	if ctx.Declarations() != nil {
		ctx.Declarations().Accept(s)
	}

	if ctx.Objective() != nil {
		ctx.Objective().Accept(s)
	}

	if ctx.Constraints() != nil {
		ctx.Constraints().Accept(s)
	}

	if DEBUG {
		fmt.Println("End semantic")
	}
	return nil
}

func (s *PlSemantic) VisitDeclarations(ctx *parser.DeclarationsContext) interface{} {
	debug("Declarations")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}

	return nil
}

func (s *PlSemantic) VisitSetDecl(ctx *parser.SetDeclContext) interface{} {
	debug("SetDecl")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#setDefinition.
func (s *PlSemantic) VisitSetDefinition(ctx *parser.SetDefinitionContext) interface{} {
	debug("SetDefinition")

	nome := ctx.IDENT().GetText()

	s.declarar(
		nome,
		SET,
		0,
		nil,
		ctx.GetStart().GetLine(),
	)

	return nil
}

// Visit a parse tree produced by GrammarParser#paramDecl.
func (s *PlSemantic) VisitParamDecl(ctx *parser.ParamDeclContext) interface{} {
	debug("ParamDecl")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#parameter.
func (s *PlSemantic) VisitParameter(ctx *parser.ParameterContext) interface{} {
	debug("Parameter")

	for _, ref := range ctx.IdentifierList().AllReference() {

		nome := ref.IDENT().GetText()

		var dominios []string

		if ref.ExpressionList() != nil {

			for _, expr := range ref.ExpressionList().AllExpression() {

				// Como estamos numa DECLARAÇÃO,
				// esperamos apenas nomes de conjuntos.
				txt := expr.GetText()

				if s.tabela.Verificar(txt) != SET {
					s.erro(fmt.Sprintf(
						"'%s' não é um conjunto declarado",
						txt,
					),
						ctx.GetStart().GetLine())
				}

				dominios = append(dominios, txt)
			}
		}

		s.declarar(
			nome,
			PARAM,
			len(dominios),
			dominios,
			ctx.GetStart().GetLine(),
		)
	}

	return nil
}

// Visit a parse tree produced by GrammarParser#varDecl.
func (s *PlSemantic) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	debug("VarDecl")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#variable.
func (s *PlSemantic) VisitVariable(ctx *parser.VariableContext) interface{} {

	debug("Variable")

	var tipo TipoSimbolo

	switch ctx.Domain().GetText() {

	case "binary":
		tipo = VAR_BINARY

	case "integer":
		tipo = VAR_INTEGER

	default:
		tipo = VAR_CONTINUOUS
	}

	for _, ref := range ctx.IdentifierList().AllReference() {

		nome := ref.IDENT().GetText()

		var dominios []string

		if ref.ExpressionList() != nil {

			for _, expr := range ref.ExpressionList().AllExpression() {

				txt := expr.GetText()

				if s.tabela.Verificar(txt) != SET {

					s.erro(fmt.Sprintf(
						"'%s' não é um conjunto declarado",
						txt,
					),
						ctx.GetStart().GetLine())
				}

				dominios = append(dominios, txt)
			}
		}

		s.declarar(
			nome,
			tipo,
			len(dominios),
			dominios,
			ctx.GetStart().GetLine(),
		)
	}

	return nil
}

// Visit a parse tree produced by GrammarParser#identifierList.
func (s *PlSemantic) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	debug("IdentifierList")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#reference.
// Visit a parse tree produced by GrammarParser#reference.
func (s *PlSemantic) VisitReference(ctx *parser.ReferenceContext) interface{} {
	debug("Reference")

	nome := ctx.IDENT().GetText()

	// Verifica se o identificador existe
	entrada, ok := s.tabela.ObterEntrada(nome)
	if !ok {
		s.erro(fmt.Sprintf(
			"identificador '%s' não declarado",
			nome,
		),
			ctx.GetStart().GetLine())
		return nil
	}

	// Quantidade de índices utilizados
	indicesUsados := 0

	if ctx.ExpressionList() != nil {

		indicesUsados = len(ctx.ExpressionList().AllExpression())

		// Visita as expressões dos índices
		for _, expr := range ctx.ExpressionList().AllExpression() {
			expr.Accept(s)
		}
	}

	// Verifica a dimensionalidade
	if entrada.Indices != indicesUsados {
		s.erro(fmt.Sprintf(
			"'%s' foi declarado com %d dimensão(ões), mas utilizado com %d",
			nome,
			entrada.Indices,
			indicesUsados,
		),
			ctx.GetStart().GetLine())
	}

	return nil
}

// Visit a parse tree produced by GrammarParser#domain.
func (s *PlSemantic) VisitDomain(ctx *parser.DomainContext) interface{} {
	debug("Domain")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#objective.
func (s *PlSemantic) VisitObjective(ctx *parser.ObjectiveContext) interface{} {
	debug("Objective")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#constraints.
func (s *PlSemantic) VisitConstraints(ctx *parser.ConstraintsContext) interface{} {
	debug("Constraints")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#constraint.
func (s *PlSemantic) VisitConstraint(ctx *parser.ConstraintContext) interface{} {
	debug("Constraint")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#simpleConstraint.
func (s *PlSemantic) VisitSimpleConstraint(ctx *parser.SimpleConstraintContext) interface{} {
	debug("SimpleConstraint")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#quantifiedConstraint.
func (s *PlSemantic) VisitQuantifiedConstraint(ctx *parser.QuantifiedConstraintContext) interface{} {
	debug("QuantifiedConstraint")

	s.tabela.NovoEscopo()

	ctx.IteratorList().Accept(s)

	ctx.LinearConstraint().Accept(s)

	s.tabela.AbandonarEscopo()

	return nil
}

// Visit a parse tree produced by GrammarParser#subsetConstraint.
func (s *PlSemantic) VisitSubsetConstraint(ctx *parser.SubsetConstraintContext) interface{} {
	debug("SubsetConstraint")

	nomeSubconjunto := ctx.IDENT(0).GetText()

	// O conjunto após "subsetof"
	ctx.Expression().Accept(s)

	s.tabela.NovoEscopo()

	s.tabela.Adicionar(
		nomeSubconjunto,
		SET,
		0,
		nil,
	)

	if ctx.WhereClause() != nil {
		ctx.WhereClause().Accept(s)
	}

	ctx.LinearConstraint().Accept(s)

	s.tabela.AbandonarEscopo()

	return nil
}

// Visit a parse tree produced by GrammarParser#whereClause.
func (s *PlSemantic) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	debug("WhereClause")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#iteratorList.
func (s *PlSemantic) VisitIteratorList(ctx *parser.IteratorListContext) interface{} {
	debug("IteratorList")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#iterator.
func (s *PlSemantic) VisitIterator(ctx *parser.IteratorContext) interface{} {
	debug("Iterator")

	nome := ctx.IDENT().GetText()

	// Visita a expressão após "in"
	ctx.Expression().Accept(s)

	// Não permite reutilizar nomes no mesmo escopo
	if s.tabela.ExisteNoEscopoAtual(nome) {
		s.erro(fmt.Sprintf(
			"iterador '%s' já declarado neste escopo",
			nome,
		),
			ctx.GetStart().GetLine())
		return nil
	}

	s.tabela.Adicionar(
		nome,
		ITERATOR,
		0,
		nil,
	)

	return nil
}

// Visit a parse tree produced by GrammarParser#linearConstraint.
func (s *PlSemantic) VisitLinearConstraint(ctx *parser.LinearConstraintContext) interface{} {
	debug("LinearConstraint")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#expression.
func (s *PlSemantic) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	debug("Expression")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#additiveExpr.
func (s *PlSemantic) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) interface{} {
	debug("AdditiveExpr")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#multiplicativeExpr.
func (s *PlSemantic) VisitMultiplicativeExpr(ctx *parser.MultiplicativeExprContext) interface{} {
	debug("MultiplicativeExpr")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#unaryExpr.
func (s *PlSemantic) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	debug("UnaryExpr")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#summation.
func (s *PlSemantic) VisitSummation(ctx *parser.SummationContext) interface{} {
	debug("Summation")

	s.tabela.NovoEscopo()

	ctx.IteratorList().Accept(s)

	ctx.Expression().Accept(s)

	s.tabela.AbandonarEscopo()

	return nil
}

// Visit a parse tree produced by GrammarParser#expressionList.
func (s *PlSemantic) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	debug("ExpressionList")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}

// Visit a parse tree produced by GrammarParser#comparator.
func (s *PlSemantic) VisitComparator(ctx *parser.ComparatorContext) interface{} {
	debug("Comparator")

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(s)
		}
	}
	return nil
}
