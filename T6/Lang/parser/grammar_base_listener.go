// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseGrammarListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseGrammarListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterSetDecl is called when production setDecl is entered.
func (s *BaseGrammarListener) EnterSetDecl(ctx *SetDeclContext) {}

// ExitSetDecl is called when production setDecl is exited.
func (s *BaseGrammarListener) ExitSetDecl(ctx *SetDeclContext) {}

// EnterSetDefinition is called when production setDefinition is entered.
func (s *BaseGrammarListener) EnterSetDefinition(ctx *SetDefinitionContext) {}

// ExitSetDefinition is called when production setDefinition is exited.
func (s *BaseGrammarListener) ExitSetDefinition(ctx *SetDefinitionContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseGrammarListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseGrammarListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseGrammarListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseGrammarListener) ExitParameter(ctx *ParameterContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseGrammarListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseGrammarListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseGrammarListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseGrammarListener) ExitVariable(ctx *VariableContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseGrammarListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseGrammarListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseGrammarListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseGrammarListener) ExitReference(ctx *ReferenceContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseGrammarListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseGrammarListener) ExitDomain(ctx *DomainContext) {}

// EnterObjective is called when production objective is entered.
func (s *BaseGrammarListener) EnterObjective(ctx *ObjectiveContext) {}

// ExitObjective is called when production objective is exited.
func (s *BaseGrammarListener) ExitObjective(ctx *ObjectiveContext) {}

// EnterConstraints is called when production constraints is entered.
func (s *BaseGrammarListener) EnterConstraints(ctx *ConstraintsContext) {}

// ExitConstraints is called when production constraints is exited.
func (s *BaseGrammarListener) ExitConstraints(ctx *ConstraintsContext) {}

// EnterConstraint is called when production constraint is entered.
func (s *BaseGrammarListener) EnterConstraint(ctx *ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseGrammarListener) ExitConstraint(ctx *ConstraintContext) {}

// EnterSimpleConstraint is called when production simpleConstraint is entered.
func (s *BaseGrammarListener) EnterSimpleConstraint(ctx *SimpleConstraintContext) {}

// ExitSimpleConstraint is called when production simpleConstraint is exited.
func (s *BaseGrammarListener) ExitSimpleConstraint(ctx *SimpleConstraintContext) {}

// EnterQuantifiedConstraint is called when production quantifiedConstraint is entered.
func (s *BaseGrammarListener) EnterQuantifiedConstraint(ctx *QuantifiedConstraintContext) {}

// ExitQuantifiedConstraint is called when production quantifiedConstraint is exited.
func (s *BaseGrammarListener) ExitQuantifiedConstraint(ctx *QuantifiedConstraintContext) {}

// EnterSubsetConstraint is called when production subsetConstraint is entered.
func (s *BaseGrammarListener) EnterSubsetConstraint(ctx *SubsetConstraintContext) {}

// ExitSubsetConstraint is called when production subsetConstraint is exited.
func (s *BaseGrammarListener) ExitSubsetConstraint(ctx *SubsetConstraintContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseGrammarListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseGrammarListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterIteratorList is called when production iteratorList is entered.
func (s *BaseGrammarListener) EnterIteratorList(ctx *IteratorListContext) {}

// ExitIteratorList is called when production iteratorList is exited.
func (s *BaseGrammarListener) ExitIteratorList(ctx *IteratorListContext) {}

// EnterIterator is called when production iterator is entered.
func (s *BaseGrammarListener) EnterIterator(ctx *IteratorContext) {}

// ExitIterator is called when production iterator is exited.
func (s *BaseGrammarListener) ExitIterator(ctx *IteratorContext) {}

// EnterLinearConstraint is called when production linearConstraint is entered.
func (s *BaseGrammarListener) EnterLinearConstraint(ctx *LinearConstraintContext) {}

// ExitLinearConstraint is called when production linearConstraint is exited.
func (s *BaseGrammarListener) ExitLinearConstraint(ctx *LinearConstraintContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGrammarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGrammarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseGrammarListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseGrammarListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterMultiplicativeExpr is called when production multiplicativeExpr is entered.
func (s *BaseGrammarListener) EnterMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// ExitMultiplicativeExpr is called when production multiplicativeExpr is exited.
func (s *BaseGrammarListener) ExitMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseGrammarListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseGrammarListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterSummation is called when production summation is entered.
func (s *BaseGrammarListener) EnterSummation(ctx *SummationContext) {}

// ExitSummation is called when production summation is exited.
func (s *BaseGrammarListener) ExitSummation(ctx *SummationContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseGrammarListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseGrammarListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseGrammarListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseGrammarListener) ExitComparator(ctx *ComparatorContext) {}
