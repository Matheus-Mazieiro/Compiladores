lexer grammar CalcLexer;

//Define uma regra para cada tipo de token (palavras reservadas, tipos de variaveis, etc)
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

//Regras para operadores
ABREPAR: '(';
FECHAPAR: ')';
ABRECOL: '[';
FECHACOL: ']';
DOIS_PONTOS: ':';
VIRG: ',';
PONTO: '.';
PONTOPONTO: '..';
ARIT: '+' | '-' | '*' | '/' | '%' | '^'; //Define operadores artimeticos
RELAC: '>' | '>=' | '<' | '<=' | '<>' | '='; //Define operadores relacionais
LOGIC: 'e' | 'ou' | 'nao'; //Define operadores lógicos
ATRIB: '<-';
ENDERECO: '&'; 

// Expressão regular para números inteiros e reais 
NUM_INT: [0-9]+;
NUM_REAL: [0-9]+ '.' [0-9]+;

//Expressão regular para capturar identificadores 
IDENT: [a-zA-Z_][a-zA-Z_0-9]*;

//Expressão regular para cadeia de caracteres
CADEIA 	: '"' ( ESC_SEQ | ~('\''|'\\'|'"'|'\n') )* '"';
fragment
ESC_SEQ	: '\\\'';

//Define regra para capturar erro de cadeia : não encontra aspas de fechamento, mas sim um \n
ERRO_CADEIA:  '"' ( ESC_SEQ | ~('\''|'\\'|'"'| '\n') )* '\n';

// Ignora espaços e comentários
WS: [ \t\r\n]+ -> skip;
COMENTARIO: '{' ( ~'\n' )*? '}' -> skip;

// Regra para capturar comentário não fechado
ERRO_COMENTARIO: '{' ( ~'}' )*? '\n';

// Captura os demais casos de erro, porque . é um caractere coringa
ERRO_DEMAIS: .;
