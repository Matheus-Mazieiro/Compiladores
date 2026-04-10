grammar CalcLexer;

//Transcrevendo gramática da linguagem LA
//Regras sintáticas

programa : declaracoes ALG corpo FIM_ALG ;
declaracoes : decl_local_global* ;
decl_local_global: declaracao_local | declaracao_global;
declaracao_local: DECLARE variavel | CONSTANTE IDENT DOIS_PONTOS tipo_basico '=' valor_constante | TIPO IDENT DOIS_PONTOS tipo;
variavel : identificador (VIRG identificador)* DOIS_PONTOS tipo;
identificador: IDENT (PONTO IDENT)* dimensao;
dimensao : (ABRECOL exp_aritmetica FECHACOL)* ;
tipo: registro | tipo_estendido;
tipo_basico: LITERAL | INTEIRO | REAL | LOGICO;tipo_basico_ident : tipo_basico | IDENT;
tipo_estendido : '^'? tipo_basico_ident ;
valor_constante: CADEIA | NUM_INT | NUM_REAL | VERDADEIRO | FALSO;
registro : REGISTRO variavel* FIM_REGISTRO ;
declaracao_global : PROCEDIMENTO IDENT ABREPAR parametros? FECHAPAR declaracao_local* cmd* FIM_PROCEDIMENTO | FUNCAO IDENT ABREPAR parametros? FECHAPAR DOIS_PONTOS tipo_estendido declaracao_local* cmd* FIM_FUNCAO;
parametros: parametro (VIRG parametro)*;
corpo : declaracao_local* cmd* ;
parametro : VAR? identificador (VIRG identificador)* DOIS_PONTOS tipo_estendido ;
cmd : cmdLeia | cmdEscreva | cmdSe | cmdCaso | cmdPara | cmdEnquanto | cmdFaca | cmdAtribuicao | cmdChamada | cmdRetorne;
cmdLeia : LEIA ABREPAR '^'? identificador (VIRG '^'? identificador)* FECHAPAR ;
cmdEscreva : ESCREVA ABREPAR expressao (VIRG expressao)* FECHAPAR ;
cmdSe : SE expressao ENTAO cmd* (SENAO cmd*)? FIM_SE ;
cmdCaso : CASO exp_aritmetica SEJA selecao (SENAO cmd*)? FIM_CASO ;
cmdPara : PARA IDENT ATRIB exp_aritmetica ATE exp_aritmetica FACA cmd* FIM_PARA ;
cmdEnquanto : ENQUANTO expressao FACA cmd* FIM_ENQUANTO ;
cmdFaca : FACA cmd* ATE expressao ;
cmdAtribuicao: '^'? identificador ATRIB expressao;
cmdChamada: IDENT ABREPAR expressao (VIRG expressao)* FECHAPAR;
cmdRetorne: RETORNE expressao;
selecao: item_selecao*;
item_selecao: constantes DOIS_PONTOS cmd*;
constantes: numero_intervalo (VIRG numero_intervalo)*;
numero_intervalo : op_unario? NUM_INT (PONTOPONTO op_unario? NUM_INT)? ;
op_unario: '-';
exp_aritmetica: termo (op1 termo)*;
termo: fator (op2 fator)*;
fator: parcela (op3 parcela)*;
op1 : '+' | '-';
op2: '*' | '/';
op3: '%';
parcela: op_unario? parcela_unario | parcela_nao_unario;
parcela_unario: '^'? identificador | IDENT ABREPAR expressao (VIRG expressao)* FECHAPAR | NUM_INT | NUM_REAL | ABREPAR expressao FECHAPAR;
parcela_nao_unario: ENDERECO identificador | CADEIA;
exp_relacional : exp_aritmetica (op_relacional exp_aritmetica)? ;
op_relacional: '=' | '<>' | '>=' | '<=' | '>' | '<';
expressao: termo_logico (op_logico_1 termo_logico)*;
termo_logico: fator_logico (op_logico_2 fator_logico)*;
fator_logico: 'nao'? parcela_logica;
parcela_logica: (VERDADEIRO | FALSO) | exp_relacional;
op_logico_1: 'ou';
op_logico_2: 'e';

//Define uma regra para cada tipo de token (palavras reservadas, tipos de variaveis, etc) : regras léxicas
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