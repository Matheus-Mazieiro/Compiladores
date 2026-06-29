package main

import (
	"fmt"
	"strings"

	"github.com/matheus-mazieiro/pl-compiler/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Gerador struct {
	parser.BaseGrammarVisitor

	tabela *TabelaDeSimbolos

	code strings.Builder

	indent int
}

func NewGerador() *Gerador {
	return &Gerador{
		tabela: NewTabela(),
	}
}

func (g *Gerador) Codigo() string {
	return g.code.String()
}

func (g *Gerador) tab() string {
	return strings.Repeat("\t", g.indent)
}

func (g *Gerador) writeln(format string, args ...interface{}) {
	g.code.WriteString(g.tab())
	g.code.WriteString(fmt.Sprintf(format, args...))
	g.code.WriteByte('\n')
}

func (g *Gerador) write(format string, args ...interface{}) {
	g.code.WriteString(fmt.Sprintf(format, args...))
}

// ----------------------------------------------------
// Program
// ----------------------------------------------------

func (g *Gerador) VisitProgram(ctx *parser.ProgramContext) interface{} {

	g.writeln("#include <ilcplex/ilocplex.h>")
	g.writeln("#include <vector>")
	g.writeln("#include <iostream>")
	g.writeln("")
	g.writeln("ILOSTLBEGIN")
	g.writeln("")
	g.writeln("int main(){")

	g.indent++

	g.writeln("IloEnv env;")
	g.writeln("")

	g.writeln("try {")

	g.indent++

	g.writeln("IloModel model(env);")
	g.writeln("")

	ctx.Declarations().Accept(g)
	ctx.Objective().Accept(g)
	ctx.Constraints().Accept(g)

	g.writeln("")
	g.writeln("IloCplex cplex(model);")
	g.writeln("")

	g.writeln("if(cplex.solve()){")
	g.indent++
	g.writeln(`std::cout << "Objective = " << cplex.getObjValue() << std::endl;`)
	g.indent--
	g.writeln("}")

	g.indent--

	g.writeln("}")
	g.writeln("catch(IloException& e){")
	g.indent++
	g.writeln("std::cerr << e << std::endl;")
	g.indent--
	g.writeln("}")

	g.writeln("")
	g.writeln("env.end();")
	g.writeln("return 0;")

	g.indent--

	g.writeln("}")

	return nil
}

// ----------------------------------------------------
// Visitors que apenas propagam a visita
// ----------------------------------------------------

func (g *Gerador) VisitDeclarations(ctx *parser.DeclarationsContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitSetDecl(ctx *parser.SetDeclContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitParamDecl(ctx *parser.ParamDeclContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitObjective(ctx *parser.ObjectiveContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitConstraints(ctx *parser.ConstraintsContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitConstraint(ctx *parser.ConstraintContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitSimpleConstraint(ctx *parser.SimpleConstraintContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitQuantifiedConstraint(ctx *parser.QuantifiedConstraintContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitSubsetConstraint(ctx *parser.SubsetConstraintContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitIteratorList(ctx *parser.IteratorListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitIterator(ctx *parser.IteratorContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitLinearConstraint(ctx *parser.LinearConstraintContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitMultiplicativeExpr(ctx *parser.MultiplicativeExprContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitSummation(ctx *parser.SummationContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitComparator(ctx *parser.ComparatorContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitReference(ctx *parser.ReferenceContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(g)
		}
	}
	return nil
}

func (g *Gerador) VisitSetDefinition(ctx *parser.SetDefinitionContext) interface{} {
	return nil
}

func (g *Gerador) VisitParameter(ctx *parser.ParameterContext) interface{} {
	return nil
}

func (g *Gerador) VisitVariable(ctx *parser.VariableContext) interface{} {
	return nil
}

func (g *Gerador) VisitDomain(ctx *parser.DomainContext) interface{} {
	return nil
}
