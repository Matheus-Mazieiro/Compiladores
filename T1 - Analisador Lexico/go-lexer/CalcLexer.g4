lexer grammar CalcLexer;

ALG: 'algoritmo';
FIM_ALG: 'fim_algoritmo';
DECLARE: 'declare';
LEIA: 'leia';
ESCREVA: 'escreva';
LITERAL: 'literal';
INTEIRO: 'inteiro';

ABREPAR: '(';
FECHAPAR: ')';
ABRECOL: '[';
FECHACOL: ']';
DOIS_PONTOS: ':';
VIRG: ',';
PONTO: '.';
ARIT: '+' | '-' | '*' | '/' | '%';
RELAC: '>' | '>=' | '<' | '<=' | '<>' | '=';
LOGIC: 'e' | 'ou' | 'nao';

INT: [0-9]+;
FLOAT: [0-9]+ '.' [0-9]+;

IDENT: [a-zA-Z_][a-zA-Z_0-9]*;
CADEIA 	: '"' ( ESC_SEQ | ~('\''|'\\'|'"'|'\n') )* '"';
fragment
ESC_SEQ	: '\\\'';
ERRO_CADEIA:  '"' ( ESC_SEQ | ~('\''|'\\'|'"'| '\n') )* '\n';

WS: [ \t\r\n]+ -> skip;
COMENTARIO: '{' ( ~'\n' )*? '}' -> skip;
ERRO_COMENTARIO: '{' ( ~'}' )*? '\n';

