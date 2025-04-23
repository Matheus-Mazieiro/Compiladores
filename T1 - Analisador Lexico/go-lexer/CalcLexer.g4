lexer grammar CalcLexer;

ALG: 'algoritmo';
FIM_ALG: 'fim_algoritmo';
DECLARE: 'declare';
LEIA: 'leia';
ESCREVA: 'escreva';
LITERAL: 'literal';
INTEIRO: 'inteiro';
CASO: 'caso';
SEJA: 'seja';
SENAO: 'senao';
FIM_CASO: 'fim_caso';
REAL: 'real';
PARA: 'para';
ATE: 'ate';
FACA : 'faca';
FIM_PARA: 'fim_para';
ENQUANTO: 'enquanto';
FIM_ENQUANTO: 'fim_enquanto';
SE: 'se';
ENTAO: 'entao';
FIM_SE: 'fim_se';
REGISTRO: 'registro';
FIM_REGISTRO: 'fim_registro';
TIPO: 'tipo';
PROCEDIMENTO: 'procedimento';
VAR: 'var';
FIM_PROCEDIMENTO: 'fim_procedimento';
FUNCAO: 'funcao';
RETORNE: 'retorne';
FIM_FUNCAO: 'fim_funcao';
CONSTANTE: 'constante';
LOGICO: 'logico';
FALSO: 'falso';
VERDADEIRO: 'verdadeiro';

ABREPAR: '(';
FECHAPAR: ')';
ABRECOL: '[';
FECHACOL: ']';
DOIS_PONTOS: ':';
VIRG: ',';
PONTO: '.';
PONTOPONTO: '..';
ARIT: '+' | '-' | '*' | '/' | '%' | '^';
RELAC: '>' | '>=' | '<' | '<=' | '<>' | '=';
LOGIC: 'e' | 'ou' | 'nao';
ATRIB: '<-';
ENDERECO: '&'; 

NUM_INT: [0-9]+;
NUM_REAL: [0-9]+ '.' [0-9]+;

IDENT: [a-zA-Z_][a-zA-Z_0-9]*;
CADEIA 	: '"' ( ESC_SEQ | ~('\''|'\\'|'"'|'\n') )* '"';
fragment
ESC_SEQ	: '\\\'';
ERRO_CADEIA:  '"' ( ESC_SEQ | ~('\''|'\\'|'"'| '\n') )* '\n';

WS: [ \t\r\n]+ -> skip;
COMENTARIO: '{' ( ~'\n' )*? '}' -> skip;
ERRO_COMENTARIO: '{' ( ~'}' )*? '\n';

