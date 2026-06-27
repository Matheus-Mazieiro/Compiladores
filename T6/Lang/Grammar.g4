grammar Grammar;

program
    : declarations objective constraints EOF
    ;

declarations
    : (setDecl | paramDecl | varDecl)*
    ;

setDecl
    : SETS setDefinition+
    ;

setDefinition
    : IDENT
    ;

paramDecl
    : PARAMS parameter+
    ;

parameter
    : identifierList
    ;

varDecl
    : VARS variable+
    ;

variable
    : identifierList domain
    ;

identifierList
    : reference (',' reference)*
    ;

reference
    : IDENT ('[' expressionList ']')?
    ;

domain
    : BINARY
    | INTEGER
    | CONTINUOUS
    ;

objective
    : MAXIMIZE expression
    | MINIMIZE expression
    ;

constraints
    : SUBJECT TO constraint*
    ;

constraint
    : simpleConstraint
    | quantifiedConstraint
    | subsetConstraint
    ;

simpleConstraint
    : IDENT ':' linearConstraint
    ;

quantifiedConstraint
    : IDENT '(' iteratorList ')' ':' linearConstraint
    ;

subsetConstraint
    : IDENT '(' IDENT SUBSETOF expression ')' whereClause? ':' linearConstraint
    ;

whereClause
    : WHERE expression (AND expression)*
    ;

iteratorList
    : iterator (',' iterator)*
    ;

iterator
    : IDENT IN expression
    ;

linearConstraint
    : expression comparator expression
    ;

expression
    : additiveExpr
    ;

additiveExpr
    : multiplicativeExpr (('+'|'-') multiplicativeExpr)*
    ;

multiplicativeExpr
    : unaryExpr ('*' unaryExpr)*
    ;

unaryExpr
    : reference
    | NUMBER
    | summation
    | '(' expression ')'
    ;

summation
    : SUM '(' iteratorList ')' expression
    ;

expressionList
    : expression (',' expression)*
    ;

comparator
    : '<='
    | '>='
    | '='
    ;

SETS:'sets';
PARAMS:'params';
VARS:'vars';
MAXIMIZE:'maximize';
MINIMIZE:'minimize';
SUBJECT:'subject';
TO:'to';
SUM:'sum';
IN:'in';
SUBSETOF:'subsetof';
WHERE:'where';
AND:'and';
BINARY:'binary';
INTEGER:'integer';
CONTINUOUS:'continuous';

IDENT: [a-zA-Z_][a-zA-Z_0-9]*;
NUMBER: [0-9]+('.'[0-9]+)?;
WS: [ \t\r\n]+ -> skip;
