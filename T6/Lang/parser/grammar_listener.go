// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterSetDecl is called when entering the setDecl production.
	EnterSetDecl(c *SetDeclContext)

	// EnterSetDefinition is called when entering the setDefinition production.
	EnterSetDefinition(c *SetDefinitionContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterObjective is called when entering the objective production.
	EnterObjective(c *ObjectiveContext)

	// EnterConstraints is called when entering the constraints production.
	EnterConstraints(c *ConstraintsContext)

	// EnterConstraint is called when entering the constraint production.
	EnterConstraint(c *ConstraintContext)

	// EnterSimpleConstraint is called when entering the simpleConstraint production.
	EnterSimpleConstraint(c *SimpleConstraintContext)

	// EnterQuantifiedConstraint is called when entering the quantifiedConstraint production.
	EnterQuantifiedConstraint(c *QuantifiedConstraintContext)

	// EnterSubsetConstraint is called when entering the subsetConstraint production.
	EnterSubsetConstraint(c *SubsetConstraintContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterIteratorList is called when entering the iteratorList production.
	EnterIteratorList(c *IteratorListContext)

	// EnterIterator is called when entering the iterator production.
	EnterIterator(c *IteratorContext)

	// EnterLinearConstraint is called when entering the linearConstraint production.
	EnterLinearConstraint(c *LinearConstraintContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterMultiplicativeExpr is called when entering the multiplicativeExpr production.
	EnterMultiplicativeExpr(c *MultiplicativeExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterSummation is called when entering the summation production.
	EnterSummation(c *SummationContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitSetDecl is called when exiting the setDecl production.
	ExitSetDecl(c *SetDeclContext)

	// ExitSetDefinition is called when exiting the setDefinition production.
	ExitSetDefinition(c *SetDefinitionContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitObjective is called when exiting the objective production.
	ExitObjective(c *ObjectiveContext)

	// ExitConstraints is called when exiting the constraints production.
	ExitConstraints(c *ConstraintsContext)

	// ExitConstraint is called when exiting the constraint production.
	ExitConstraint(c *ConstraintContext)

	// ExitSimpleConstraint is called when exiting the simpleConstraint production.
	ExitSimpleConstraint(c *SimpleConstraintContext)

	// ExitQuantifiedConstraint is called when exiting the quantifiedConstraint production.
	ExitQuantifiedConstraint(c *QuantifiedConstraintContext)

	// ExitSubsetConstraint is called when exiting the subsetConstraint production.
	ExitSubsetConstraint(c *SubsetConstraintContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitIteratorList is called when exiting the iteratorList production.
	ExitIteratorList(c *IteratorListContext)

	// ExitIterator is called when exiting the iterator production.
	ExitIterator(c *IteratorContext)

	// ExitLinearConstraint is called when exiting the linearConstraint production.
	ExitLinearConstraint(c *LinearConstraintContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitMultiplicativeExpr is called when exiting the multiplicativeExpr production.
	ExitMultiplicativeExpr(c *MultiplicativeExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitSummation is called when exiting the summation production.
	ExitSummation(c *SummationContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)
}
