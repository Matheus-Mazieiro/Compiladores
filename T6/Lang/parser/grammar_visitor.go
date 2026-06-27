// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GrammarParser#declarations.
	VisitDeclarations(ctx *DeclarationsContext) interface{}

	// Visit a parse tree produced by GrammarParser#setDecl.
	VisitSetDecl(ctx *SetDeclContext) interface{}

	// Visit a parse tree produced by GrammarParser#setDefinition.
	VisitSetDefinition(ctx *SetDefinitionContext) interface{}

	// Visit a parse tree produced by GrammarParser#paramDecl.
	VisitParamDecl(ctx *ParamDeclContext) interface{}

	// Visit a parse tree produced by GrammarParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by GrammarParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by GrammarParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by GrammarParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by GrammarParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by GrammarParser#domain.
	VisitDomain(ctx *DomainContext) interface{}

	// Visit a parse tree produced by GrammarParser#objective.
	VisitObjective(ctx *ObjectiveContext) interface{}

	// Visit a parse tree produced by GrammarParser#constraints.
	VisitConstraints(ctx *ConstraintsContext) interface{}

	// Visit a parse tree produced by GrammarParser#constraint.
	VisitConstraint(ctx *ConstraintContext) interface{}

	// Visit a parse tree produced by GrammarParser#simpleConstraint.
	VisitSimpleConstraint(ctx *SimpleConstraintContext) interface{}

	// Visit a parse tree produced by GrammarParser#quantifiedConstraint.
	VisitQuantifiedConstraint(ctx *QuantifiedConstraintContext) interface{}

	// Visit a parse tree produced by GrammarParser#subsetConstraint.
	VisitSubsetConstraint(ctx *SubsetConstraintContext) interface{}

	// Visit a parse tree produced by GrammarParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by GrammarParser#iteratorList.
	VisitIteratorList(ctx *IteratorListContext) interface{}

	// Visit a parse tree produced by GrammarParser#iterator.
	VisitIterator(ctx *IteratorContext) interface{}

	// Visit a parse tree produced by GrammarParser#linearConstraint.
	VisitLinearConstraint(ctx *LinearConstraintContext) interface{}

	// Visit a parse tree produced by GrammarParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GrammarParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#multiplicativeExpr.
	VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#summation.
	VisitSummation(ctx *SummationContext) interface{}

	// Visit a parse tree produced by GrammarParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by GrammarParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}
}
