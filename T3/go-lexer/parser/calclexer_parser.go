// Code generated from CalcLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcLexer

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CalcLexerParser struct {
	*antlr.BaseParser
}

var CalcLexerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerParserInit() {
	staticData := &CalcLexerParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'^'", "'-'", "'+'", "'*'", "'/'", "'%'", "'<>'", "'>='",
		"'<='", "'>'", "'<'", "'nao'", "'ou'", "'e'", "'algoritmo'", "'fim_algoritmo'",
		"'declare'", "'leia'", "'escreva'", "'literal'", "'inteiro'", "'caso'",
		"'seja'", "'senao'", "'fim_caso'", "'real'", "'para'", "'ate'", "'faca'",
		"'fim_para'", "'enquanto'", "'fim_enquanto'", "'se'", "'entao'", "'fim_se'",
		"'registro'", "'fim_registro'", "'tipo'", "'procedimento'", "'var'",
		"'fim_procedimento'", "'funcao'", "'retorne'", "'fim_funcao'", "'constante'",
		"'logico'", "'falso'", "'verdadeiro'", "'('", "')'", "'['", "']'", "':'",
		"','", "'.'", "'..'", "", "", "", "'<-'", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ALG",
		"FIM_ALG", "DECLARE", "LEIA", "ESCREVA", "LITERAL", "INTEIRO", "CASO",
		"SEJA", "SENAO", "FIM_CASO", "REAL", "PARA", "ATE", "FACA", "FIM_PARA",
		"ENQUANTO", "FIM_ENQUANTO", "SE", "ENTAO", "FIM_SE", "REGISTRO", "FIM_REGISTRO",
		"TIPO", "PROCEDIMENTO", "VAR", "FIM_PROCEDIMENTO", "FUNCAO", "RETORNE",
		"FIM_FUNCAO", "CONSTANTE", "LOGICO", "FALSO", "VERDADEIRO", "ABREPAR",
		"FECHAPAR", "ABRECOL", "FECHACOL", "DOIS_PONTOS", "VIRG", "PONTO", "PONTOPONTO",
		"ARIT", "RELAC", "LOGIC", "ATRIB", "ENDERECO", "NUM_INT", "NUM_REAL",
		"IDENT", "CADEIA", "ERRO_CADEIA", "WS", "COMENTARIO", "ERRO_COMENTARIO",
		"ERRO_DEMAIS",
	}
	staticData.RuleNames = []string{
		"programa", "declaracoes", "decl_local_global", "declaracao_local",
		"variavel", "identificador", "dimensao", "tipo", "tipo_basico", "tipo_basico_ident",
		"tipo_estendido", "valor_constante", "registro", "declaracao_global",
		"parametros", "corpo", "parametro", "cmd", "cmdLeia", "cmdEscreva",
		"cmdSe", "cmdCaso", "cmdPara", "cmdEnquanto", "cmdFaca", "cmdAtribuicao",
		"cmdChamada", "cmdRetorne", "selecao", "item_selecao", "constantes",
		"numero_intervalo", "op_unario", "exp_aritmetica", "termo", "fator",
		"op1", "op2", "op3", "parcela", "parcela_unario", "parcela_nao_unario",
		"exp_relacional", "op_relacional", "expressao", "termo_logico", "fator_logico",
		"parcela_logica", "op_logico_1", "op_logico_2",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 544, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		5, 1, 107, 8, 1, 10, 1, 12, 1, 110, 9, 1, 1, 2, 1, 2, 3, 2, 114, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 129, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 134, 8, 4, 10, 4, 12, 4,
		137, 9, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 145, 8, 5, 10, 5,
		12, 5, 148, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 156, 8, 6,
		10, 6, 12, 6, 159, 9, 6, 1, 7, 1, 7, 3, 7, 163, 8, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 3, 9, 169, 8, 9, 1, 10, 3, 10, 172, 8, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 5, 12, 180, 8, 12, 10, 12, 12, 12, 183, 9, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 191, 8, 13, 1, 13, 1, 13, 5,
		13, 195, 8, 13, 10, 13, 12, 13, 198, 9, 13, 1, 13, 5, 13, 201, 8, 13, 10,
		13, 12, 13, 204, 9, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 211,
		8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 217, 8, 13, 10, 13, 12, 13, 220,
		9, 13, 1, 13, 5, 13, 223, 8, 13, 10, 13, 12, 13, 226, 9, 13, 1, 13, 1,
		13, 3, 13, 230, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 235, 8, 14, 10, 14,
		12, 14, 238, 9, 14, 1, 15, 5, 15, 241, 8, 15, 10, 15, 12, 15, 244, 9, 15,
		1, 15, 5, 15, 247, 8, 15, 10, 15, 12, 15, 250, 9, 15, 1, 16, 3, 16, 253,
		8, 16, 1, 16, 1, 16, 1, 16, 5, 16, 258, 8, 16, 10, 16, 12, 16, 261, 9,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 276, 8, 17, 1, 18, 1, 18, 1, 18, 3, 18, 281,
		8, 18, 1, 18, 1, 18, 1, 18, 3, 18, 286, 8, 18, 1, 18, 5, 18, 289, 8, 18,
		10, 18, 12, 18, 292, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 5, 19, 301, 8, 19, 10, 19, 12, 19, 304, 9, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 312, 8, 20, 10, 20, 12, 20, 315, 9, 20, 1,
		20, 1, 20, 5, 20, 319, 8, 20, 10, 20, 12, 20, 322, 9, 20, 3, 20, 324, 8,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 334,
		8, 21, 10, 21, 12, 21, 337, 9, 21, 3, 21, 339, 8, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 351, 8, 22,
		10, 22, 12, 22, 354, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5,
		23, 362, 8, 23, 10, 23, 12, 23, 365, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		5, 24, 371, 8, 24, 10, 24, 12, 24, 374, 9, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 3, 25, 380, 8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 5, 26, 391, 8, 26, 10, 26, 12, 26, 394, 9, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 28, 5, 28, 402, 8, 28, 10, 28, 12, 28, 405,
		9, 28, 1, 29, 1, 29, 1, 29, 5, 29, 410, 8, 29, 10, 29, 12, 29, 413, 9,
		29, 1, 30, 1, 30, 1, 30, 5, 30, 418, 8, 30, 10, 30, 12, 30, 421, 9, 30,
		1, 31, 3, 31, 424, 8, 31, 1, 31, 1, 31, 1, 31, 3, 31, 429, 8, 31, 1, 31,
		3, 31, 432, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 440,
		8, 33, 10, 33, 12, 33, 443, 9, 33, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 449,
		8, 34, 10, 34, 12, 34, 452, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 458,
		8, 35, 10, 35, 12, 35, 461, 9, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1,
		38, 1, 39, 3, 39, 470, 8, 39, 1, 39, 1, 39, 3, 39, 474, 8, 39, 1, 40, 3,
		40, 477, 8, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 485, 8,
		40, 10, 40, 12, 40, 488, 9, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 3, 40, 498, 8, 40, 1, 41, 1, 41, 1, 41, 3, 41, 503, 8, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 509, 8, 42, 1, 43, 1, 43, 1, 44, 1,
		44, 1, 44, 1, 44, 5, 44, 517, 8, 44, 10, 44, 12, 44, 520, 9, 44, 1, 45,
		1, 45, 1, 45, 1, 45, 5, 45, 526, 8, 45, 10, 45, 12, 45, 529, 9, 45, 1,
		46, 3, 46, 532, 8, 46, 1, 46, 1, 46, 1, 47, 1, 47, 3, 47, 538, 8, 47, 1,
		48, 1, 48, 1, 49, 1, 49, 1, 49, 0, 0, 50, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 0, 6, 3, 0, 21, 22, 27, 27, 47, 47, 3, 0, 48, 49, 63,
		64, 66, 66, 1, 0, 3, 4, 1, 0, 5, 6, 2, 0, 1, 1, 8, 12, 1, 0, 48, 49, 562,
		0, 100, 1, 0, 0, 0, 2, 108, 1, 0, 0, 0, 4, 113, 1, 0, 0, 0, 6, 128, 1,
		0, 0, 0, 8, 130, 1, 0, 0, 0, 10, 141, 1, 0, 0, 0, 12, 157, 1, 0, 0, 0,
		14, 162, 1, 0, 0, 0, 16, 164, 1, 0, 0, 0, 18, 168, 1, 0, 0, 0, 20, 171,
		1, 0, 0, 0, 22, 175, 1, 0, 0, 0, 24, 177, 1, 0, 0, 0, 26, 229, 1, 0, 0,
		0, 28, 231, 1, 0, 0, 0, 30, 242, 1, 0, 0, 0, 32, 252, 1, 0, 0, 0, 34, 275,
		1, 0, 0, 0, 36, 277, 1, 0, 0, 0, 38, 295, 1, 0, 0, 0, 40, 307, 1, 0, 0,
		0, 42, 327, 1, 0, 0, 0, 44, 342, 1, 0, 0, 0, 46, 357, 1, 0, 0, 0, 48, 368,
		1, 0, 0, 0, 50, 379, 1, 0, 0, 0, 52, 385, 1, 0, 0, 0, 54, 397, 1, 0, 0,
		0, 56, 403, 1, 0, 0, 0, 58, 406, 1, 0, 0, 0, 60, 414, 1, 0, 0, 0, 62, 423,
		1, 0, 0, 0, 64, 433, 1, 0, 0, 0, 66, 435, 1, 0, 0, 0, 68, 444, 1, 0, 0,
		0, 70, 453, 1, 0, 0, 0, 72, 462, 1, 0, 0, 0, 74, 464, 1, 0, 0, 0, 76, 466,
		1, 0, 0, 0, 78, 473, 1, 0, 0, 0, 80, 497, 1, 0, 0, 0, 82, 502, 1, 0, 0,
		0, 84, 504, 1, 0, 0, 0, 86, 510, 1, 0, 0, 0, 88, 512, 1, 0, 0, 0, 90, 521,
		1, 0, 0, 0, 92, 531, 1, 0, 0, 0, 94, 537, 1, 0, 0, 0, 96, 539, 1, 0, 0,
		0, 98, 541, 1, 0, 0, 0, 100, 101, 3, 2, 1, 0, 101, 102, 5, 16, 0, 0, 102,
		103, 3, 30, 15, 0, 103, 104, 5, 17, 0, 0, 104, 1, 1, 0, 0, 0, 105, 107,
		3, 4, 2, 0, 106, 105, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0,
		0, 0, 108, 109, 1, 0, 0, 0, 109, 3, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111,
		114, 3, 6, 3, 0, 112, 114, 3, 26, 13, 0, 113, 111, 1, 0, 0, 0, 113, 112,
		1, 0, 0, 0, 114, 5, 1, 0, 0, 0, 115, 116, 5, 18, 0, 0, 116, 129, 3, 8,
		4, 0, 117, 118, 5, 46, 0, 0, 118, 119, 5, 65, 0, 0, 119, 120, 5, 54, 0,
		0, 120, 121, 3, 16, 8, 0, 121, 122, 5, 1, 0, 0, 122, 123, 3, 22, 11, 0,
		123, 129, 1, 0, 0, 0, 124, 125, 5, 39, 0, 0, 125, 126, 5, 65, 0, 0, 126,
		127, 5, 54, 0, 0, 127, 129, 3, 14, 7, 0, 128, 115, 1, 0, 0, 0, 128, 117,
		1, 0, 0, 0, 128, 124, 1, 0, 0, 0, 129, 7, 1, 0, 0, 0, 130, 135, 3, 10,
		5, 0, 131, 132, 5, 55, 0, 0, 132, 134, 3, 10, 5, 0, 133, 131, 1, 0, 0,
		0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136,
		138, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 5, 54, 0, 0, 139, 140,
		3, 14, 7, 0, 140, 9, 1, 0, 0, 0, 141, 146, 5, 65, 0, 0, 142, 143, 5, 56,
		0, 0, 143, 145, 5, 65, 0, 0, 144, 142, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0,
		146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148,
		146, 1, 0, 0, 0, 149, 150, 3, 12, 6, 0, 150, 11, 1, 0, 0, 0, 151, 152,
		5, 52, 0, 0, 152, 153, 3, 66, 33, 0, 153, 154, 5, 53, 0, 0, 154, 156, 1,
		0, 0, 0, 155, 151, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0,
		0, 157, 158, 1, 0, 0, 0, 158, 13, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160,
		163, 3, 24, 12, 0, 161, 163, 3, 20, 10, 0, 162, 160, 1, 0, 0, 0, 162, 161,
		1, 0, 0, 0, 163, 15, 1, 0, 0, 0, 164, 165, 7, 0, 0, 0, 165, 17, 1, 0, 0,
		0, 166, 169, 3, 16, 8, 0, 167, 169, 5, 65, 0, 0, 168, 166, 1, 0, 0, 0,
		168, 167, 1, 0, 0, 0, 169, 19, 1, 0, 0, 0, 170, 172, 5, 2, 0, 0, 171, 170,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 3, 18,
		9, 0, 174, 21, 1, 0, 0, 0, 175, 176, 7, 1, 0, 0, 176, 23, 1, 0, 0, 0, 177,
		181, 5, 37, 0, 0, 178, 180, 3, 8, 4, 0, 179, 178, 1, 0, 0, 0, 180, 183,
		1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 1, 0,
		0, 0, 183, 181, 1, 0, 0, 0, 184, 185, 5, 38, 0, 0, 185, 25, 1, 0, 0, 0,
		186, 187, 5, 40, 0, 0, 187, 188, 5, 65, 0, 0, 188, 190, 5, 50, 0, 0, 189,
		191, 3, 28, 14, 0, 190, 189, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 192,
		1, 0, 0, 0, 192, 196, 5, 51, 0, 0, 193, 195, 3, 6, 3, 0, 194, 193, 1, 0,
		0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0,
		197, 202, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 201, 3, 34, 17, 0, 200,
		199, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 230, 5, 42,
		0, 0, 206, 207, 5, 43, 0, 0, 207, 208, 5, 65, 0, 0, 208, 210, 5, 50, 0,
		0, 209, 211, 3, 28, 14, 0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 212, 1, 0, 0, 0, 212, 213, 5, 51, 0, 0, 213, 214, 5, 54, 0, 0, 214,
		218, 3, 20, 10, 0, 215, 217, 3, 6, 3, 0, 216, 215, 1, 0, 0, 0, 217, 220,
		1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 224, 1, 0,
		0, 0, 220, 218, 1, 0, 0, 0, 221, 223, 3, 34, 17, 0, 222, 221, 1, 0, 0,
		0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225,
		227, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 5, 45, 0, 0, 228, 230,
		1, 0, 0, 0, 229, 186, 1, 0, 0, 0, 229, 206, 1, 0, 0, 0, 230, 27, 1, 0,
		0, 0, 231, 236, 3, 32, 16, 0, 232, 233, 5, 55, 0, 0, 233, 235, 3, 32, 16,
		0, 234, 232, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236,
		237, 1, 0, 0, 0, 237, 29, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 241, 3,
		6, 3, 0, 240, 239, 1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0,
		0, 242, 243, 1, 0, 0, 0, 243, 248, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245,
		247, 3, 34, 17, 0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246,
		1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 31, 1, 0, 0, 0, 250, 248, 1, 0,
		0, 0, 251, 253, 5, 41, 0, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0,
		253, 254, 1, 0, 0, 0, 254, 259, 3, 10, 5, 0, 255, 256, 5, 55, 0, 0, 256,
		258, 3, 10, 5, 0, 257, 255, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257,
		1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 259, 1, 0,
		0, 0, 262, 263, 5, 54, 0, 0, 263, 264, 3, 20, 10, 0, 264, 33, 1, 0, 0,
		0, 265, 276, 3, 36, 18, 0, 266, 276, 3, 38, 19, 0, 267, 276, 3, 40, 20,
		0, 268, 276, 3, 42, 21, 0, 269, 276, 3, 44, 22, 0, 270, 276, 3, 46, 23,
		0, 271, 276, 3, 48, 24, 0, 272, 276, 3, 50, 25, 0, 273, 276, 3, 52, 26,
		0, 274, 276, 3, 54, 27, 0, 275, 265, 1, 0, 0, 0, 275, 266, 1, 0, 0, 0,
		275, 267, 1, 0, 0, 0, 275, 268, 1, 0, 0, 0, 275, 269, 1, 0, 0, 0, 275,
		270, 1, 0, 0, 0, 275, 271, 1, 0, 0, 0, 275, 272, 1, 0, 0, 0, 275, 273,
		1, 0, 0, 0, 275, 274, 1, 0, 0, 0, 276, 35, 1, 0, 0, 0, 277, 278, 5, 19,
		0, 0, 278, 280, 5, 50, 0, 0, 279, 281, 5, 2, 0, 0, 280, 279, 1, 0, 0, 0,
		280, 281, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 290, 3, 10, 5, 0, 283,
		285, 5, 55, 0, 0, 284, 286, 5, 2, 0, 0, 285, 284, 1, 0, 0, 0, 285, 286,
		1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 3, 10, 5, 0, 288, 283, 1, 0,
		0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0,
		291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294, 5, 51, 0, 0, 294,
		37, 1, 0, 0, 0, 295, 296, 5, 20, 0, 0, 296, 297, 5, 50, 0, 0, 297, 302,
		3, 88, 44, 0, 298, 299, 5, 55, 0, 0, 299, 301, 3, 88, 44, 0, 300, 298,
		1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0,
		0, 0, 303, 305, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 306, 5, 51, 0, 0,
		306, 39, 1, 0, 0, 0, 307, 308, 5, 34, 0, 0, 308, 309, 3, 88, 44, 0, 309,
		313, 5, 35, 0, 0, 310, 312, 3, 34, 17, 0, 311, 310, 1, 0, 0, 0, 312, 315,
		1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 323, 1, 0,
		0, 0, 315, 313, 1, 0, 0, 0, 316, 320, 5, 25, 0, 0, 317, 319, 3, 34, 17,
		0, 318, 317, 1, 0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320,
		321, 1, 0, 0, 0, 321, 324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 316,
		1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 5, 36,
		0, 0, 326, 41, 1, 0, 0, 0, 327, 328, 5, 23, 0, 0, 328, 329, 3, 66, 33,
		0, 329, 330, 5, 24, 0, 0, 330, 338, 3, 56, 28, 0, 331, 335, 5, 25, 0, 0,
		332, 334, 3, 34, 17, 0, 333, 332, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335,
		333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 339, 1, 0, 0, 0, 337, 335,
		1, 0, 0, 0, 338, 331, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1, 0,
		0, 0, 340, 341, 5, 26, 0, 0, 341, 43, 1, 0, 0, 0, 342, 343, 5, 28, 0, 0,
		343, 344, 5, 65, 0, 0, 344, 345, 5, 61, 0, 0, 345, 346, 3, 66, 33, 0, 346,
		347, 5, 29, 0, 0, 347, 348, 3, 66, 33, 0, 348, 352, 5, 30, 0, 0, 349, 351,
		3, 34, 17, 0, 350, 349, 1, 0, 0, 0, 351, 354, 1, 0, 0, 0, 352, 350, 1,
		0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 355, 1, 0, 0, 0, 354, 352, 1, 0, 0,
		0, 355, 356, 5, 31, 0, 0, 356, 45, 1, 0, 0, 0, 357, 358, 5, 32, 0, 0, 358,
		359, 3, 88, 44, 0, 359, 363, 5, 30, 0, 0, 360, 362, 3, 34, 17, 0, 361,
		360, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 363, 364,
		1, 0, 0, 0, 364, 366, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 366, 367, 5, 33,
		0, 0, 367, 47, 1, 0, 0, 0, 368, 372, 5, 30, 0, 0, 369, 371, 3, 34, 17,
		0, 370, 369, 1, 0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372,
		373, 1, 0, 0, 0, 373, 375, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 375, 376,
		5, 29, 0, 0, 376, 377, 3, 88, 44, 0, 377, 49, 1, 0, 0, 0, 378, 380, 5,
		2, 0, 0, 379, 378, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 1, 0, 0,
		0, 381, 382, 3, 10, 5, 0, 382, 383, 5, 61, 0, 0, 383, 384, 3, 88, 44, 0,
		384, 51, 1, 0, 0, 0, 385, 386, 5, 65, 0, 0, 386, 387, 5, 50, 0, 0, 387,
		392, 3, 88, 44, 0, 388, 389, 5, 55, 0, 0, 389, 391, 3, 88, 44, 0, 390,
		388, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 392, 393,
		1, 0, 0, 0, 393, 395, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 395, 396, 5, 51,
		0, 0, 396, 53, 1, 0, 0, 0, 397, 398, 5, 44, 0, 0, 398, 399, 3, 88, 44,
		0, 399, 55, 1, 0, 0, 0, 400, 402, 3, 58, 29, 0, 401, 400, 1, 0, 0, 0, 402,
		405, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404, 57, 1,
		0, 0, 0, 405, 403, 1, 0, 0, 0, 406, 407, 3, 60, 30, 0, 407, 411, 5, 54,
		0, 0, 408, 410, 3, 34, 17, 0, 409, 408, 1, 0, 0, 0, 410, 413, 1, 0, 0,
		0, 411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 59, 1, 0, 0, 0, 413,
		411, 1, 0, 0, 0, 414, 419, 3, 62, 31, 0, 415, 416, 5, 55, 0, 0, 416, 418,
		3, 62, 31, 0, 417, 415, 1, 0, 0, 0, 418, 421, 1, 0, 0, 0, 419, 417, 1,
		0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 61, 1, 0, 0, 0, 421, 419, 1, 0, 0,
		0, 422, 424, 3, 64, 32, 0, 423, 422, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0,
		424, 425, 1, 0, 0, 0, 425, 431, 5, 63, 0, 0, 426, 428, 5, 57, 0, 0, 427,
		429, 3, 64, 32, 0, 428, 427, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 430,
		1, 0, 0, 0, 430, 432, 5, 63, 0, 0, 431, 426, 1, 0, 0, 0, 431, 432, 1, 0,
		0, 0, 432, 63, 1, 0, 0, 0, 433, 434, 5, 3, 0, 0, 434, 65, 1, 0, 0, 0, 435,
		441, 3, 68, 34, 0, 436, 437, 3, 72, 36, 0, 437, 438, 3, 68, 34, 0, 438,
		440, 1, 0, 0, 0, 439, 436, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441, 439,
		1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 67, 1, 0, 0, 0, 443, 441, 1, 0,
		0, 0, 444, 450, 3, 70, 35, 0, 445, 446, 3, 74, 37, 0, 446, 447, 3, 70,
		35, 0, 447, 449, 1, 0, 0, 0, 448, 445, 1, 0, 0, 0, 449, 452, 1, 0, 0, 0,
		450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 69, 1, 0, 0, 0, 452, 450,
		1, 0, 0, 0, 453, 459, 3, 78, 39, 0, 454, 455, 3, 76, 38, 0, 455, 456, 3,
		78, 39, 0, 456, 458, 1, 0, 0, 0, 457, 454, 1, 0, 0, 0, 458, 461, 1, 0,
		0, 0, 459, 457, 1, 0, 0, 0, 459, 460, 1, 0, 0, 0, 460, 71, 1, 0, 0, 0,
		461, 459, 1, 0, 0, 0, 462, 463, 7, 2, 0, 0, 463, 73, 1, 0, 0, 0, 464, 465,
		7, 3, 0, 0, 465, 75, 1, 0, 0, 0, 466, 467, 5, 7, 0, 0, 467, 77, 1, 0, 0,
		0, 468, 470, 3, 64, 32, 0, 469, 468, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0,
		470, 471, 1, 0, 0, 0, 471, 474, 3, 80, 40, 0, 472, 474, 3, 82, 41, 0, 473,
		469, 1, 0, 0, 0, 473, 472, 1, 0, 0, 0, 474, 79, 1, 0, 0, 0, 475, 477, 5,
		2, 0, 0, 476, 475, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 478, 1, 0, 0,
		0, 478, 498, 3, 10, 5, 0, 479, 480, 5, 65, 0, 0, 480, 481, 5, 50, 0, 0,
		481, 486, 3, 88, 44, 0, 482, 483, 5, 55, 0, 0, 483, 485, 3, 88, 44, 0,
		484, 482, 1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 486,
		487, 1, 0, 0, 0, 487, 489, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 489, 490,
		5, 51, 0, 0, 490, 498, 1, 0, 0, 0, 491, 498, 5, 63, 0, 0, 492, 498, 5,
		64, 0, 0, 493, 494, 5, 50, 0, 0, 494, 495, 3, 88, 44, 0, 495, 496, 5, 51,
		0, 0, 496, 498, 1, 0, 0, 0, 497, 476, 1, 0, 0, 0, 497, 479, 1, 0, 0, 0,
		497, 491, 1, 0, 0, 0, 497, 492, 1, 0, 0, 0, 497, 493, 1, 0, 0, 0, 498,
		81, 1, 0, 0, 0, 499, 500, 5, 62, 0, 0, 500, 503, 3, 10, 5, 0, 501, 503,
		5, 66, 0, 0, 502, 499, 1, 0, 0, 0, 502, 501, 1, 0, 0, 0, 503, 83, 1, 0,
		0, 0, 504, 508, 3, 66, 33, 0, 505, 506, 3, 86, 43, 0, 506, 507, 3, 66,
		33, 0, 507, 509, 1, 0, 0, 0, 508, 505, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0,
		509, 85, 1, 0, 0, 0, 510, 511, 7, 4, 0, 0, 511, 87, 1, 0, 0, 0, 512, 518,
		3, 90, 45, 0, 513, 514, 3, 96, 48, 0, 514, 515, 3, 90, 45, 0, 515, 517,
		1, 0, 0, 0, 516, 513, 1, 0, 0, 0, 517, 520, 1, 0, 0, 0, 518, 516, 1, 0,
		0, 0, 518, 519, 1, 0, 0, 0, 519, 89, 1, 0, 0, 0, 520, 518, 1, 0, 0, 0,
		521, 527, 3, 92, 46, 0, 522, 523, 3, 98, 49, 0, 523, 524, 3, 92, 46, 0,
		524, 526, 1, 0, 0, 0, 525, 522, 1, 0, 0, 0, 526, 529, 1, 0, 0, 0, 527,
		525, 1, 0, 0, 0, 527, 528, 1, 0, 0, 0, 528, 91, 1, 0, 0, 0, 529, 527, 1,
		0, 0, 0, 530, 532, 5, 13, 0, 0, 531, 530, 1, 0, 0, 0, 531, 532, 1, 0, 0,
		0, 532, 533, 1, 0, 0, 0, 533, 534, 3, 94, 47, 0, 534, 93, 1, 0, 0, 0, 535,
		538, 7, 5, 0, 0, 536, 538, 3, 84, 42, 0, 537, 535, 1, 0, 0, 0, 537, 536,
		1, 0, 0, 0, 538, 95, 1, 0, 0, 0, 539, 540, 5, 14, 0, 0, 540, 97, 1, 0,
		0, 0, 541, 542, 5, 15, 0, 0, 542, 99, 1, 0, 0, 0, 57, 108, 113, 128, 135,
		146, 157, 162, 168, 171, 181, 190, 196, 202, 210, 218, 224, 229, 236, 242,
		248, 252, 259, 275, 280, 285, 290, 302, 313, 320, 323, 335, 338, 352, 363,
		372, 379, 392, 403, 411, 419, 423, 428, 431, 441, 450, 459, 469, 473, 476,
		486, 497, 502, 508, 518, 527, 531, 537,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalcLexerParserInit initializes any static state used to implement CalcLexerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalcLexerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerParserInit() {
	staticData := &CalcLexerParserStaticData
	staticData.once.Do(calclexerParserInit)
}

// NewCalcLexerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalcLexerParser(input antlr.TokenStream) *CalcLexerParser {
	CalcLexerParserInit()
	this := new(CalcLexerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalcLexerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CalcLexer.g4"

	return this
}

// CalcLexerParser tokens.
const (
	CalcLexerParserEOF              = antlr.TokenEOF
	CalcLexerParserT__0             = 1
	CalcLexerParserT__1             = 2
	CalcLexerParserT__2             = 3
	CalcLexerParserT__3             = 4
	CalcLexerParserT__4             = 5
	CalcLexerParserT__5             = 6
	CalcLexerParserT__6             = 7
	CalcLexerParserT__7             = 8
	CalcLexerParserT__8             = 9
	CalcLexerParserT__9             = 10
	CalcLexerParserT__10            = 11
	CalcLexerParserT__11            = 12
	CalcLexerParserT__12            = 13
	CalcLexerParserT__13            = 14
	CalcLexerParserT__14            = 15
	CalcLexerParserALG              = 16
	CalcLexerParserFIM_ALG          = 17
	CalcLexerParserDECLARE          = 18
	CalcLexerParserLEIA             = 19
	CalcLexerParserESCREVA          = 20
	CalcLexerParserLITERAL          = 21
	CalcLexerParserINTEIRO          = 22
	CalcLexerParserCASO             = 23
	CalcLexerParserSEJA             = 24
	CalcLexerParserSENAO            = 25
	CalcLexerParserFIM_CASO         = 26
	CalcLexerParserREAL             = 27
	CalcLexerParserPARA             = 28
	CalcLexerParserATE              = 29
	CalcLexerParserFACA             = 30
	CalcLexerParserFIM_PARA         = 31
	CalcLexerParserENQUANTO         = 32
	CalcLexerParserFIM_ENQUANTO     = 33
	CalcLexerParserSE               = 34
	CalcLexerParserENTAO            = 35
	CalcLexerParserFIM_SE           = 36
	CalcLexerParserREGISTRO         = 37
	CalcLexerParserFIM_REGISTRO     = 38
	CalcLexerParserTIPO             = 39
	CalcLexerParserPROCEDIMENTO     = 40
	CalcLexerParserVAR              = 41
	CalcLexerParserFIM_PROCEDIMENTO = 42
	CalcLexerParserFUNCAO           = 43
	CalcLexerParserRETORNE          = 44
	CalcLexerParserFIM_FUNCAO       = 45
	CalcLexerParserCONSTANTE        = 46
	CalcLexerParserLOGICO           = 47
	CalcLexerParserFALSO            = 48
	CalcLexerParserVERDADEIRO       = 49
	CalcLexerParserABREPAR          = 50
	CalcLexerParserFECHAPAR         = 51
	CalcLexerParserABRECOL          = 52
	CalcLexerParserFECHACOL         = 53
	CalcLexerParserDOIS_PONTOS      = 54
	CalcLexerParserVIRG             = 55
	CalcLexerParserPONTO            = 56
	CalcLexerParserPONTOPONTO       = 57
	CalcLexerParserARIT             = 58
	CalcLexerParserRELAC            = 59
	CalcLexerParserLOGIC            = 60
	CalcLexerParserATRIB            = 61
	CalcLexerParserENDERECO         = 62
	CalcLexerParserNUM_INT          = 63
	CalcLexerParserNUM_REAL         = 64
	CalcLexerParserIDENT            = 65
	CalcLexerParserCADEIA           = 66
	CalcLexerParserERRO_CADEIA      = 67
	CalcLexerParserWS               = 68
	CalcLexerParserCOMENTARIO       = 69
	CalcLexerParserERRO_COMENTARIO  = 70
	CalcLexerParserERRO_DEMAIS      = 71
)

// CalcLexerParser rules.
const (
	CalcLexerParserRULE_programa           = 0
	CalcLexerParserRULE_declaracoes        = 1
	CalcLexerParserRULE_decl_local_global  = 2
	CalcLexerParserRULE_declaracao_local   = 3
	CalcLexerParserRULE_variavel           = 4
	CalcLexerParserRULE_identificador      = 5
	CalcLexerParserRULE_dimensao           = 6
	CalcLexerParserRULE_tipo               = 7
	CalcLexerParserRULE_tipo_basico        = 8
	CalcLexerParserRULE_tipo_basico_ident  = 9
	CalcLexerParserRULE_tipo_estendido     = 10
	CalcLexerParserRULE_valor_constante    = 11
	CalcLexerParserRULE_registro           = 12
	CalcLexerParserRULE_declaracao_global  = 13
	CalcLexerParserRULE_parametros         = 14
	CalcLexerParserRULE_corpo              = 15
	CalcLexerParserRULE_parametro          = 16
	CalcLexerParserRULE_cmd                = 17
	CalcLexerParserRULE_cmdLeia            = 18
	CalcLexerParserRULE_cmdEscreva         = 19
	CalcLexerParserRULE_cmdSe              = 20
	CalcLexerParserRULE_cmdCaso            = 21
	CalcLexerParserRULE_cmdPara            = 22
	CalcLexerParserRULE_cmdEnquanto        = 23
	CalcLexerParserRULE_cmdFaca            = 24
	CalcLexerParserRULE_cmdAtribuicao      = 25
	CalcLexerParserRULE_cmdChamada         = 26
	CalcLexerParserRULE_cmdRetorne         = 27
	CalcLexerParserRULE_selecao            = 28
	CalcLexerParserRULE_item_selecao       = 29
	CalcLexerParserRULE_constantes         = 30
	CalcLexerParserRULE_numero_intervalo   = 31
	CalcLexerParserRULE_op_unario          = 32
	CalcLexerParserRULE_exp_aritmetica     = 33
	CalcLexerParserRULE_termo              = 34
	CalcLexerParserRULE_fator              = 35
	CalcLexerParserRULE_op1                = 36
	CalcLexerParserRULE_op2                = 37
	CalcLexerParserRULE_op3                = 38
	CalcLexerParserRULE_parcela            = 39
	CalcLexerParserRULE_parcela_unario     = 40
	CalcLexerParserRULE_parcela_nao_unario = 41
	CalcLexerParserRULE_exp_relacional     = 42
	CalcLexerParserRULE_op_relacional      = 43
	CalcLexerParserRULE_expressao          = 44
	CalcLexerParserRULE_termo_logico       = 45
	CalcLexerParserRULE_fator_logico       = 46
	CalcLexerParserRULE_parcela_logica     = 47
	CalcLexerParserRULE_op_logico_1        = 48
	CalcLexerParserRULE_op_logico_2        = 49
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracoes() IDeclaracoesContext
	ALG() antlr.TerminalNode
	Corpo() ICorpoContext
	FIM_ALG() antlr.TerminalNode

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) Declaracoes() IDeclaracoesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracoesContext)
}

func (s *ProgramaContext) ALG() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserALG, 0)
}

func (s *ProgramaContext) Corpo() ICorpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICorpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICorpoContext)
}

func (s *ProgramaContext) FIM_ALG() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_ALG, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (p *CalcLexerParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcLexerParserRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Declaracoes()
	}
	{
		p.SetState(101)
		p.Match(CalcLexerParserALG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Corpo()
	}
	{
		p.SetState(103)
		p.Match(CalcLexerParserFIM_ALG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracoesContext is an interface to support dynamic dispatch.
type IDeclaracoesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDecl_local_global() []IDecl_local_globalContext
	Decl_local_global(i int) IDecl_local_globalContext

	// IsDeclaracoesContext differentiates from other interfaces.
	IsDeclaracoesContext()
}

type DeclaracoesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracoesContext() *DeclaracoesContext {
	var p = new(DeclaracoesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_declaracoes
	return p
}

func InitEmptyDeclaracoesContext(p *DeclaracoesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_declaracoes
}

func (*DeclaracoesContext) IsDeclaracoesContext() {}

func NewDeclaracoesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracoesContext {
	var p = new(DeclaracoesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_declaracoes

	return p
}

func (s *DeclaracoesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracoesContext) AllDecl_local_global() []IDecl_local_globalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDecl_local_globalContext); ok {
			len++
		}
	}

	tst := make([]IDecl_local_globalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDecl_local_globalContext); ok {
			tst[i] = t.(IDecl_local_globalContext)
			i++
		}
	}

	return tst
}

func (s *DeclaracoesContext) Decl_local_global(i int) IDecl_local_globalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_local_globalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_local_globalContext)
}

func (s *DeclaracoesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracoesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracoesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterDeclaracoes(s)
	}
}

func (s *DeclaracoesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitDeclaracoes(s)
	}
}

func (p *CalcLexerParser) Declaracoes() (localctx IDeclaracoesContext) {
	localctx = NewDeclaracoesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalcLexerParserRULE_declaracoes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&80814104903680) != 0 {
		{
			p.SetState(105)
			p.Decl_local_global()
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecl_local_globalContext is an interface to support dynamic dispatch.
type IDecl_local_globalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracao_local() IDeclaracao_localContext
	Declaracao_global() IDeclaracao_globalContext

	// IsDecl_local_globalContext differentiates from other interfaces.
	IsDecl_local_globalContext()
}

type Decl_local_globalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_local_globalContext() *Decl_local_globalContext {
	var p = new(Decl_local_globalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_decl_local_global
	return p
}

func InitEmptyDecl_local_globalContext(p *Decl_local_globalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_decl_local_global
}

func (*Decl_local_globalContext) IsDecl_local_globalContext() {}

func NewDecl_local_globalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_local_globalContext {
	var p = new(Decl_local_globalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_decl_local_global

	return p
}

func (s *Decl_local_globalContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_local_globalContext) Declaracao_local() IDeclaracao_localContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracao_localContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracao_localContext)
}

func (s *Decl_local_globalContext) Declaracao_global() IDeclaracao_globalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracao_globalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracao_globalContext)
}

func (s *Decl_local_globalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_local_globalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decl_local_globalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterDecl_local_global(s)
	}
}

func (s *Decl_local_globalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitDecl_local_global(s)
	}
}

func (p *CalcLexerParser) Decl_local_global() (localctx IDecl_local_globalContext) {
	localctx = NewDecl_local_globalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcLexerParserRULE_decl_local_global)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserDECLARE, CalcLexerParserTIPO, CalcLexerParserCONSTANTE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Declaracao_local()
		}

	case CalcLexerParserPROCEDIMENTO, CalcLexerParserFUNCAO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Declaracao_global()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracao_localContext is an interface to support dynamic dispatch.
type IDeclaracao_localContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECLARE() antlr.TerminalNode
	Variavel() IVariavelContext
	CONSTANTE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	DOIS_PONTOS() antlr.TerminalNode
	Tipo_basico() ITipo_basicoContext
	Valor_constante() IValor_constanteContext
	TIPO() antlr.TerminalNode
	Tipo() ITipoContext

	// IsDeclaracao_localContext differentiates from other interfaces.
	IsDeclaracao_localContext()
}

type Declaracao_localContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracao_localContext() *Declaracao_localContext {
	var p = new(Declaracao_localContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_declaracao_local
	return p
}

func InitEmptyDeclaracao_localContext(p *Declaracao_localContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_declaracao_local
}

func (*Declaracao_localContext) IsDeclaracao_localContext() {}

func NewDeclaracao_localContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracao_localContext {
	var p = new(Declaracao_localContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_declaracao_local

	return p
}

func (s *Declaracao_localContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracao_localContext) DECLARE() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserDECLARE, 0)
}

func (s *Declaracao_localContext) Variavel() IVariavelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariavelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariavelContext)
}

func (s *Declaracao_localContext) CONSTANTE() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserCONSTANTE, 0)
}

func (s *Declaracao_localContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserIDENT, 0)
}

func (s *Declaracao_localContext) DOIS_PONTOS() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserDOIS_PONTOS, 0)
}

func (s *Declaracao_localContext) Tipo_basico() ITipo_basicoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_basicoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_basicoContext)
}

func (s *Declaracao_localContext) Valor_constante() IValor_constanteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValor_constanteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValor_constanteContext)
}

func (s *Declaracao_localContext) TIPO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserTIPO, 0)
}

func (s *Declaracao_localContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracao_localContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracao_localContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracao_localContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterDeclaracao_local(s)
	}
}

func (s *Declaracao_localContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitDeclaracao_local(s)
	}
}

func (p *CalcLexerParser) Declaracao_local() (localctx IDeclaracao_localContext) {
	localctx = NewDeclaracao_localContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcLexerParserRULE_declaracao_local)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserDECLARE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Match(CalcLexerParserDECLARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Variavel()
		}

	case CalcLexerParserCONSTANTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(CalcLexerParserCONSTANTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(CalcLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(CalcLexerParserDOIS_PONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Tipo_basico()
		}
		{
			p.SetState(121)
			p.Match(CalcLexerParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Valor_constante()
		}

	case CalcLexerParserTIPO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(CalcLexerParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(CalcLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(CalcLexerParserDOIS_PONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Tipo()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariavelContext is an interface to support dynamic dispatch.
type IVariavelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []IIdentificadorContext
	Identificador(i int) IIdentificadorContext
	DOIS_PONTOS() antlr.TerminalNode
	Tipo() ITipoContext
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode

	// IsVariavelContext differentiates from other interfaces.
	IsVariavelContext()
}

type VariavelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariavelContext() *VariavelContext {
	var p = new(VariavelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_variavel
	return p
}

func InitEmptyVariavelContext(p *VariavelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_variavel
}

func (*VariavelContext) IsVariavelContext() {}

func NewVariavelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariavelContext {
	var p = new(VariavelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_variavel

	return p
}

func (s *VariavelContext) GetParser() antlr.Parser { return s.parser }

func (s *VariavelContext) AllIdentificador() []IIdentificadorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentificadorContext); ok {
			len++
		}
	}

	tst := make([]IIdentificadorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentificadorContext); ok {
			tst[i] = t.(IIdentificadorContext)
			i++
		}
	}

	return tst
}

func (s *VariavelContext) Identificador(i int) IIdentificadorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *VariavelContext) DOIS_PONTOS() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserDOIS_PONTOS, 0)
}

func (s *VariavelContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VariavelContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *VariavelContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *VariavelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariavelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariavelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterVariavel(s)
	}
}

func (s *VariavelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitVariavel(s)
	}
}

func (p *CalcLexerParser) Variavel() (localctx IVariavelContext) {
	localctx = NewVariavelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcLexerParserRULE_variavel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Identificador()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserVIRG {
		{
			p.SetState(131)
			p.Match(CalcLexerParserVIRG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Identificador()
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(138)
		p.Match(CalcLexerParserDOIS_PONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Tipo()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentificadorContext is an interface to support dynamic dispatch.
type IIdentificadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	Dimensao() IDimensaoContext
	AllPONTO() []antlr.TerminalNode
	PONTO(i int) antlr.TerminalNode

	// IsIdentificadorContext differentiates from other interfaces.
	IsIdentificadorContext()
}

type IdentificadorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentificadorContext() *IdentificadorContext {
	var p = new(IdentificadorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_identificador
	return p
}

func InitEmptyIdentificadorContext(p *IdentificadorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_identificador
}

func (*IdentificadorContext) IsIdentificadorContext() {}

func NewIdentificadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentificadorContext {
	var p = new(IdentificadorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_identificador

	return p
}

func (s *IdentificadorContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentificadorContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserIDENT)
}

func (s *IdentificadorContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserIDENT, i)
}

func (s *IdentificadorContext) Dimensao() IDimensaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensaoContext)
}

func (s *IdentificadorContext) AllPONTO() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserPONTO)
}

func (s *IdentificadorContext) PONTO(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserPONTO, i)
}

func (s *IdentificadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentificadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentificadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterIdentificador(s)
	}
}

func (s *IdentificadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitIdentificador(s)
	}
}

func (p *CalcLexerParser) Identificador() (localctx IIdentificadorContext) {
	localctx = NewIdentificadorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalcLexerParserRULE_identificador)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(CalcLexerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserPONTO {
		{
			p.SetState(142)
			p.Match(CalcLexerParserPONTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(CalcLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Dimensao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDimensaoContext is an interface to support dynamic dispatch.
type IDimensaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllABRECOL() []antlr.TerminalNode
	ABRECOL(i int) antlr.TerminalNode
	AllExp_aritmetica() []IExp_aritmeticaContext
	Exp_aritmetica(i int) IExp_aritmeticaContext
	AllFECHACOL() []antlr.TerminalNode
	FECHACOL(i int) antlr.TerminalNode

	// IsDimensaoContext differentiates from other interfaces.
	IsDimensaoContext()
}

type DimensaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimensaoContext() *DimensaoContext {
	var p = new(DimensaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_dimensao
	return p
}

func InitEmptyDimensaoContext(p *DimensaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_dimensao
}

func (*DimensaoContext) IsDimensaoContext() {}

func NewDimensaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensaoContext {
	var p = new(DimensaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_dimensao

	return p
}

func (s *DimensaoContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensaoContext) AllABRECOL() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserABRECOL)
}

func (s *DimensaoContext) ABRECOL(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserABRECOL, i)
}

func (s *DimensaoContext) AllExp_aritmetica() []IExp_aritmeticaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			len++
		}
	}

	tst := make([]IExp_aritmeticaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExp_aritmeticaContext); ok {
			tst[i] = t.(IExp_aritmeticaContext)
			i++
		}
	}

	return tst
}

func (s *DimensaoContext) Exp_aritmetica(i int) IExp_aritmeticaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *DimensaoContext) AllFECHACOL() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserFECHACOL)
}

func (s *DimensaoContext) FECHACOL(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFECHACOL, i)
}

func (s *DimensaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimensaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterDimensao(s)
	}
}

func (s *DimensaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitDimensao(s)
	}
}

func (p *CalcLexerParser) Dimensao() (localctx IDimensaoContext) {
	localctx = NewDimensaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalcLexerParserRULE_dimensao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserABRECOL {
		{
			p.SetState(151)
			p.Match(CalcLexerParserABRECOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Exp_aritmetica()
		}
		{
			p.SetState(153)
			p.Match(CalcLexerParserFECHACOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Registro() IRegistroContext
	Tipo_estendido() ITipo_estendidoContext

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) Registro() IRegistroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegistroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegistroContext)
}

func (s *TipoContext) Tipo_estendido() ITipo_estendidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_estendidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_estendidoContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *CalcLexerParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalcLexerParserRULE_tipo)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserREGISTRO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Registro()
		}

	case CalcLexerParserT__1, CalcLexerParserLITERAL, CalcLexerParserINTEIRO, CalcLexerParserREAL, CalcLexerParserLOGICO, CalcLexerParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Tipo_estendido()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipo_basicoContext is an interface to support dynamic dispatch.
type ITipo_basicoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL() antlr.TerminalNode
	INTEIRO() antlr.TerminalNode
	REAL() antlr.TerminalNode
	LOGICO() antlr.TerminalNode

	// IsTipo_basicoContext differentiates from other interfaces.
	IsTipo_basicoContext()
}

type Tipo_basicoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipo_basicoContext() *Tipo_basicoContext {
	var p = new(Tipo_basicoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo_basico
	return p
}

func InitEmptyTipo_basicoContext(p *Tipo_basicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo_basico
}

func (*Tipo_basicoContext) IsTipo_basicoContext() {}

func NewTipo_basicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_basicoContext {
	var p = new(Tipo_basicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_tipo_basico

	return p
}

func (s *Tipo_basicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_basicoContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserLITERAL, 0)
}

func (s *Tipo_basicoContext) INTEIRO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserINTEIRO, 0)
}

func (s *Tipo_basicoContext) REAL() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserREAL, 0)
}

func (s *Tipo_basicoContext) LOGICO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserLOGICO, 0)
}

func (s *Tipo_basicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_basicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_basicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterTipo_basico(s)
	}
}

func (s *Tipo_basicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitTipo_basico(s)
	}
}

func (p *CalcLexerParser) Tipo_basico() (localctx ITipo_basicoContext) {
	localctx = NewTipo_basicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CalcLexerParserRULE_tipo_basico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737628864512) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipo_basico_identContext is an interface to support dynamic dispatch.
type ITipo_basico_identContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo_basico() ITipo_basicoContext
	IDENT() antlr.TerminalNode

	// IsTipo_basico_identContext differentiates from other interfaces.
	IsTipo_basico_identContext()
}

type Tipo_basico_identContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipo_basico_identContext() *Tipo_basico_identContext {
	var p = new(Tipo_basico_identContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo_basico_ident
	return p
}

func InitEmptyTipo_basico_identContext(p *Tipo_basico_identContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo_basico_ident
}

func (*Tipo_basico_identContext) IsTipo_basico_identContext() {}

func NewTipo_basico_identContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_basico_identContext {
	var p = new(Tipo_basico_identContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_tipo_basico_ident

	return p
}

func (s *Tipo_basico_identContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_basico_identContext) Tipo_basico() ITipo_basicoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_basicoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_basicoContext)
}

func (s *Tipo_basico_identContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserIDENT, 0)
}

func (s *Tipo_basico_identContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_basico_identContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_basico_identContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterTipo_basico_ident(s)
	}
}

func (s *Tipo_basico_identContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitTipo_basico_ident(s)
	}
}

func (p *CalcLexerParser) Tipo_basico_ident() (localctx ITipo_basico_identContext) {
	localctx = NewTipo_basico_identContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CalcLexerParserRULE_tipo_basico_ident)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserLITERAL, CalcLexerParserINTEIRO, CalcLexerParserREAL, CalcLexerParserLOGICO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Tipo_basico()
		}

	case CalcLexerParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(CalcLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipo_estendidoContext is an interface to support dynamic dispatch.
type ITipo_estendidoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo_basico_ident() ITipo_basico_identContext

	// IsTipo_estendidoContext differentiates from other interfaces.
	IsTipo_estendidoContext()
}

type Tipo_estendidoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipo_estendidoContext() *Tipo_estendidoContext {
	var p = new(Tipo_estendidoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo_estendido
	return p
}

func InitEmptyTipo_estendidoContext(p *Tipo_estendidoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_tipo_estendido
}

func (*Tipo_estendidoContext) IsTipo_estendidoContext() {}

func NewTipo_estendidoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_estendidoContext {
	var p = new(Tipo_estendidoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_tipo_estendido

	return p
}

func (s *Tipo_estendidoContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_estendidoContext) Tipo_basico_ident() ITipo_basico_identContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_basico_identContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_basico_identContext)
}

func (s *Tipo_estendidoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_estendidoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_estendidoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterTipo_estendido(s)
	}
}

func (s *Tipo_estendidoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitTipo_estendido(s)
	}
}

func (p *CalcLexerParser) Tipo_estendido() (localctx ITipo_estendidoContext) {
	localctx = NewTipo_estendidoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CalcLexerParserRULE_tipo_estendido)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserT__1 {
		{
			p.SetState(170)
			p.Match(CalcLexerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(173)
		p.Tipo_basico_ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValor_constanteContext is an interface to support dynamic dispatch.
type IValor_constanteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CADEIA() antlr.TerminalNode
	NUM_INT() antlr.TerminalNode
	NUM_REAL() antlr.TerminalNode
	VERDADEIRO() antlr.TerminalNode
	FALSO() antlr.TerminalNode

	// IsValor_constanteContext differentiates from other interfaces.
	IsValor_constanteContext()
}

type Valor_constanteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValor_constanteContext() *Valor_constanteContext {
	var p = new(Valor_constanteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_valor_constante
	return p
}

func InitEmptyValor_constanteContext(p *Valor_constanteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_valor_constante
}

func (*Valor_constanteContext) IsValor_constanteContext() {}

func NewValor_constanteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Valor_constanteContext {
	var p = new(Valor_constanteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_valor_constante

	return p
}

func (s *Valor_constanteContext) GetParser() antlr.Parser { return s.parser }

func (s *Valor_constanteContext) CADEIA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserCADEIA, 0)
}

func (s *Valor_constanteContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserNUM_INT, 0)
}

func (s *Valor_constanteContext) NUM_REAL() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserNUM_REAL, 0)
}

func (s *Valor_constanteContext) VERDADEIRO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVERDADEIRO, 0)
}

func (s *Valor_constanteContext) FALSO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFALSO, 0)
}

func (s *Valor_constanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Valor_constanteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Valor_constanteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterValor_constante(s)
	}
}

func (s *Valor_constanteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitValor_constante(s)
	}
}

func (p *CalcLexerParser) Valor_constante() (localctx IValor_constanteContext) {
	localctx = NewValor_constanteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CalcLexerParserRULE_valor_constante)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-48)) & ^0x3f) == 0 && ((int64(1)<<(_la-48))&360451) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegistroContext is an interface to support dynamic dispatch.
type IRegistroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGISTRO() antlr.TerminalNode
	FIM_REGISTRO() antlr.TerminalNode
	AllVariavel() []IVariavelContext
	Variavel(i int) IVariavelContext

	// IsRegistroContext differentiates from other interfaces.
	IsRegistroContext()
}

type RegistroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegistroContext() *RegistroContext {
	var p = new(RegistroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_registro
	return p
}

func InitEmptyRegistroContext(p *RegistroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_registro
}

func (*RegistroContext) IsRegistroContext() {}

func NewRegistroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegistroContext {
	var p = new(RegistroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_registro

	return p
}

func (s *RegistroContext) GetParser() antlr.Parser { return s.parser }

func (s *RegistroContext) REGISTRO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserREGISTRO, 0)
}

func (s *RegistroContext) FIM_REGISTRO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_REGISTRO, 0)
}

func (s *RegistroContext) AllVariavel() []IVariavelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariavelContext); ok {
			len++
		}
	}

	tst := make([]IVariavelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariavelContext); ok {
			tst[i] = t.(IVariavelContext)
			i++
		}
	}

	return tst
}

func (s *RegistroContext) Variavel(i int) IVariavelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariavelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariavelContext)
}

func (s *RegistroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegistroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegistroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterRegistro(s)
	}
}

func (s *RegistroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitRegistro(s)
	}
}

func (p *CalcLexerParser) Registro() (localctx IRegistroContext) {
	localctx = NewRegistroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CalcLexerParserRULE_registro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(CalcLexerParserREGISTRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserIDENT {
		{
			p.SetState(178)
			p.Variavel()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(CalcLexerParserFIM_REGISTRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracao_globalContext is an interface to support dynamic dispatch.
type IDeclaracao_globalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROCEDIMENTO() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	FECHAPAR() antlr.TerminalNode
	FIM_PROCEDIMENTO() antlr.TerminalNode
	Parametros() IParametrosContext
	AllDeclaracao_local() []IDeclaracao_localContext
	Declaracao_local(i int) IDeclaracao_localContext
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext
	FUNCAO() antlr.TerminalNode
	DOIS_PONTOS() antlr.TerminalNode
	Tipo_estendido() ITipo_estendidoContext
	FIM_FUNCAO() antlr.TerminalNode

	// IsDeclaracao_globalContext differentiates from other interfaces.
	IsDeclaracao_globalContext()
}

type Declaracao_globalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracao_globalContext() *Declaracao_globalContext {
	var p = new(Declaracao_globalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_declaracao_global
	return p
}

func InitEmptyDeclaracao_globalContext(p *Declaracao_globalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_declaracao_global
}

func (*Declaracao_globalContext) IsDeclaracao_globalContext() {}

func NewDeclaracao_globalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracao_globalContext {
	var p = new(Declaracao_globalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_declaracao_global

	return p
}

func (s *Declaracao_globalContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracao_globalContext) PROCEDIMENTO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserPROCEDIMENTO, 0)
}

func (s *Declaracao_globalContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserIDENT, 0)
}

func (s *Declaracao_globalContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserABREPAR, 0)
}

func (s *Declaracao_globalContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFECHAPAR, 0)
}

func (s *Declaracao_globalContext) FIM_PROCEDIMENTO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_PROCEDIMENTO, 0)
}

func (s *Declaracao_globalContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *Declaracao_globalContext) AllDeclaracao_local() []IDeclaracao_localContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracao_localContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracao_localContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracao_localContext); ok {
			tst[i] = t.(IDeclaracao_localContext)
			i++
		}
	}

	return tst
}

func (s *Declaracao_globalContext) Declaracao_local(i int) IDeclaracao_localContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracao_localContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracao_localContext)
}

func (s *Declaracao_globalContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *Declaracao_globalContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *Declaracao_globalContext) FUNCAO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFUNCAO, 0)
}

func (s *Declaracao_globalContext) DOIS_PONTOS() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserDOIS_PONTOS, 0)
}

func (s *Declaracao_globalContext) Tipo_estendido() ITipo_estendidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_estendidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_estendidoContext)
}

func (s *Declaracao_globalContext) FIM_FUNCAO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_FUNCAO, 0)
}

func (s *Declaracao_globalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracao_globalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracao_globalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterDeclaracao_global(s)
	}
}

func (s *Declaracao_globalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitDeclaracao_global(s)
	}
}

func (p *CalcLexerParser) Declaracao_global() (localctx IDeclaracao_globalContext) {
	localctx = NewDeclaracao_globalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CalcLexerParserRULE_declaracao_global)
	var _la int

	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserPROCEDIMENTO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(CalcLexerParserPROCEDIMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(CalcLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(CalcLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CalcLexerParserVAR || _la == CalcLexerParserIDENT {
			{
				p.SetState(189)
				p.Parametros()
			}

		}
		{
			p.SetState(192)
			p.Match(CalcLexerParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70918500253696) != 0 {
			{
				p.SetState(193)
				p.Declaracao_local()
			}

			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
			{
				p.SetState(199)
				p.Cmd()
			}

			p.SetState(204)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(205)
			p.Match(CalcLexerParserFIM_PROCEDIMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalcLexerParserFUNCAO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(CalcLexerParserFUNCAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(CalcLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(CalcLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CalcLexerParserVAR || _la == CalcLexerParserIDENT {
			{
				p.SetState(209)
				p.Parametros()
			}

		}
		{
			p.SetState(212)
			p.Match(CalcLexerParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(CalcLexerParserDOIS_PONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Tipo_estendido()
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70918500253696) != 0 {
			{
				p.SetState(215)
				p.Declaracao_local()
			}

			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
			{
				p.SetState(221)
				p.Cmd()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(227)
			p.Match(CalcLexerParserFIM_FUNCAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParametro() []IParametroContext
	Parametro(i int) IParametroContext
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *ParametrosContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *ParametrosContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (p *CalcLexerParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CalcLexerParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Parametro()
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserVIRG {
		{
			p.SetState(232)
			p.Match(CalcLexerParserVIRG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Parametro()
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICorpoContext is an interface to support dynamic dispatch.
type ICorpoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaracao_local() []IDeclaracao_localContext
	Declaracao_local(i int) IDeclaracao_localContext
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCorpoContext differentiates from other interfaces.
	IsCorpoContext()
}

type CorpoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCorpoContext() *CorpoContext {
	var p = new(CorpoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_corpo
	return p
}

func InitEmptyCorpoContext(p *CorpoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_corpo
}

func (*CorpoContext) IsCorpoContext() {}

func NewCorpoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CorpoContext {
	var p = new(CorpoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_corpo

	return p
}

func (s *CorpoContext) GetParser() antlr.Parser { return s.parser }

func (s *CorpoContext) AllDeclaracao_local() []IDeclaracao_localContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracao_localContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracao_localContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracao_localContext); ok {
			tst[i] = t.(IDeclaracao_localContext)
			i++
		}
	}

	return tst
}

func (s *CorpoContext) Declaracao_local(i int) IDeclaracao_localContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracao_localContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracao_localContext)
}

func (s *CorpoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CorpoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CorpoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CorpoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CorpoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCorpo(s)
	}
}

func (s *CorpoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCorpo(s)
	}
}

func (p *CalcLexerParser) Corpo() (localctx ICorpoContext) {
	localctx = NewCorpoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CalcLexerParserRULE_corpo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70918500253696) != 0 {
		{
			p.SetState(239)
			p.Declaracao_local()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
		{
			p.SetState(245)
			p.Cmd()
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []IIdentificadorContext
	Identificador(i int) IIdentificadorContext
	DOIS_PONTOS() antlr.TerminalNode
	Tipo_estendido() ITipo_estendidoContext
	VAR() antlr.TerminalNode
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) AllIdentificador() []IIdentificadorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentificadorContext); ok {
			len++
		}
	}

	tst := make([]IIdentificadorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentificadorContext); ok {
			tst[i] = t.(IIdentificadorContext)
			i++
		}
	}

	return tst
}

func (s *ParametroContext) Identificador(i int) IIdentificadorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *ParametroContext) DOIS_PONTOS() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserDOIS_PONTOS, 0)
}

func (s *ParametroContext) Tipo_estendido() ITipo_estendidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_estendidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_estendidoContext)
}

func (s *ParametroContext) VAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVAR, 0)
}

func (s *ParametroContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *ParametroContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (p *CalcLexerParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CalcLexerParserRULE_parametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserVAR {
		{
			p.SetState(251)
			p.Match(CalcLexerParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(254)
		p.Identificador()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserVIRG {
		{
			p.SetState(255)
			p.Match(CalcLexerParserVIRG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Identificador()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(262)
		p.Match(CalcLexerParserDOIS_PONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Tipo_estendido()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdContext is an interface to support dynamic dispatch.
type ICmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CmdLeia() ICmdLeiaContext
	CmdEscreva() ICmdEscrevaContext
	CmdSe() ICmdSeContext
	CmdCaso() ICmdCasoContext
	CmdPara() ICmdParaContext
	CmdEnquanto() ICmdEnquantoContext
	CmdFaca() ICmdFacaContext
	CmdAtribuicao() ICmdAtribuicaoContext
	CmdChamada() ICmdChamadaContext
	CmdRetorne() ICmdRetorneContext

	// IsCmdContext differentiates from other interfaces.
	IsCmdContext()
}

type CmdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdContext() *CmdContext {
	var p = new(CmdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmd
	return p
}

func InitEmptyCmdContext(p *CmdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmd
}

func (*CmdContext) IsCmdContext() {}

func NewCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdContext {
	var p = new(CmdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmd

	return p
}

func (s *CmdContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdContext) CmdLeia() ICmdLeiaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdLeiaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdLeiaContext)
}

func (s *CmdContext) CmdEscreva() ICmdEscrevaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdEscrevaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdEscrevaContext)
}

func (s *CmdContext) CmdSe() ICmdSeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdSeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdSeContext)
}

func (s *CmdContext) CmdCaso() ICmdCasoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdCasoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdCasoContext)
}

func (s *CmdContext) CmdPara() ICmdParaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdParaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdParaContext)
}

func (s *CmdContext) CmdEnquanto() ICmdEnquantoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdEnquantoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdEnquantoContext)
}

func (s *CmdContext) CmdFaca() ICmdFacaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdFacaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdFacaContext)
}

func (s *CmdContext) CmdAtribuicao() ICmdAtribuicaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdAtribuicaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdAtribuicaoContext)
}

func (s *CmdContext) CmdChamada() ICmdChamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdChamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdChamadaContext)
}

func (s *CmdContext) CmdRetorne() ICmdRetorneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdRetorneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdRetorneContext)
}

func (s *CmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmd(s)
	}
}

func (s *CmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmd(s)
	}
}

func (p *CalcLexerParser) Cmd() (localctx ICmdContext) {
	localctx = NewCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CalcLexerParserRULE_cmd)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.CmdLeia()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.CmdEscreva()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(267)
			p.CmdSe()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(268)
			p.CmdCaso()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(269)
			p.CmdPara()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(270)
			p.CmdEnquanto()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(271)
			p.CmdFaca()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(272)
			p.CmdAtribuicao()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(273)
			p.CmdChamada()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(274)
			p.CmdRetorne()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdLeiaContext is an interface to support dynamic dispatch.
type ICmdLeiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEIA() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllIdentificador() []IIdentificadorContext
	Identificador(i int) IIdentificadorContext
	FECHAPAR() antlr.TerminalNode
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode

	// IsCmdLeiaContext differentiates from other interfaces.
	IsCmdLeiaContext()
}

type CmdLeiaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdLeiaContext() *CmdLeiaContext {
	var p = new(CmdLeiaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdLeia
	return p
}

func InitEmptyCmdLeiaContext(p *CmdLeiaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdLeia
}

func (*CmdLeiaContext) IsCmdLeiaContext() {}

func NewCmdLeiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdLeiaContext {
	var p = new(CmdLeiaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdLeia

	return p
}

func (s *CmdLeiaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdLeiaContext) LEIA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserLEIA, 0)
}

func (s *CmdLeiaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserABREPAR, 0)
}

func (s *CmdLeiaContext) AllIdentificador() []IIdentificadorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentificadorContext); ok {
			len++
		}
	}

	tst := make([]IIdentificadorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentificadorContext); ok {
			tst[i] = t.(IIdentificadorContext)
			i++
		}
	}

	return tst
}

func (s *CmdLeiaContext) Identificador(i int) IIdentificadorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *CmdLeiaContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFECHAPAR, 0)
}

func (s *CmdLeiaContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *CmdLeiaContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *CmdLeiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdLeiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdLeiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdLeia(s)
	}
}

func (s *CmdLeiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdLeia(s)
	}
}

func (p *CalcLexerParser) CmdLeia() (localctx ICmdLeiaContext) {
	localctx = NewCmdLeiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CalcLexerParserRULE_cmdLeia)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(CalcLexerParserLEIA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(CalcLexerParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserT__1 {
		{
			p.SetState(279)
			p.Match(CalcLexerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(282)
		p.Identificador()
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserVIRG {
		{
			p.SetState(283)
			p.Match(CalcLexerParserVIRG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CalcLexerParserT__1 {
			{
				p.SetState(284)
				p.Match(CalcLexerParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(287)
			p.Identificador()
		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(293)
		p.Match(CalcLexerParserFECHAPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdEscrevaContext is an interface to support dynamic dispatch.
type ICmdEscrevaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ESCREVA() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllExpressao() []IExpressaoContext
	Expressao(i int) IExpressaoContext
	FECHAPAR() antlr.TerminalNode
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode

	// IsCmdEscrevaContext differentiates from other interfaces.
	IsCmdEscrevaContext()
}

type CmdEscrevaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdEscrevaContext() *CmdEscrevaContext {
	var p = new(CmdEscrevaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdEscreva
	return p
}

func InitEmptyCmdEscrevaContext(p *CmdEscrevaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdEscreva
}

func (*CmdEscrevaContext) IsCmdEscrevaContext() {}

func NewCmdEscrevaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdEscrevaContext {
	var p = new(CmdEscrevaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdEscreva

	return p
}

func (s *CmdEscrevaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdEscrevaContext) ESCREVA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserESCREVA, 0)
}

func (s *CmdEscrevaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserABREPAR, 0)
}

func (s *CmdEscrevaContext) AllExpressao() []IExpressaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressaoContext); ok {
			len++
		}
	}

	tst := make([]IExpressaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressaoContext); ok {
			tst[i] = t.(IExpressaoContext)
			i++
		}
	}

	return tst
}

func (s *CmdEscrevaContext) Expressao(i int) IExpressaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdEscrevaContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFECHAPAR, 0)
}

func (s *CmdEscrevaContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *CmdEscrevaContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *CmdEscrevaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdEscrevaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdEscrevaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdEscreva(s)
	}
}

func (s *CmdEscrevaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdEscreva(s)
	}
}

func (p *CalcLexerParser) CmdEscreva() (localctx ICmdEscrevaContext) {
	localctx = NewCmdEscrevaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CalcLexerParserRULE_cmdEscreva)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(CalcLexerParserESCREVA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Match(CalcLexerParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Expressao()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserVIRG {
		{
			p.SetState(298)
			p.Match(CalcLexerParserVIRG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Expressao()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(305)
		p.Match(CalcLexerParserFECHAPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdSeContext is an interface to support dynamic dispatch.
type ICmdSeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SE() antlr.TerminalNode
	Expressao() IExpressaoContext
	ENTAO() antlr.TerminalNode
	FIM_SE() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext
	SENAO() antlr.TerminalNode

	// IsCmdSeContext differentiates from other interfaces.
	IsCmdSeContext()
}

type CmdSeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdSeContext() *CmdSeContext {
	var p = new(CmdSeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdSe
	return p
}

func InitEmptyCmdSeContext(p *CmdSeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdSe
}

func (*CmdSeContext) IsCmdSeContext() {}

func NewCmdSeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdSeContext {
	var p = new(CmdSeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdSe

	return p
}

func (s *CmdSeContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdSeContext) SE() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserSE, 0)
}

func (s *CmdSeContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdSeContext) ENTAO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserENTAO, 0)
}

func (s *CmdSeContext) FIM_SE() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_SE, 0)
}

func (s *CmdSeContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdSeContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdSeContext) SENAO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserSENAO, 0)
}

func (s *CmdSeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdSeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdSeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdSe(s)
	}
}

func (s *CmdSeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdSe(s)
	}
}

func (p *CalcLexerParser) CmdSe() (localctx ICmdSeContext) {
	localctx = NewCmdSeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CalcLexerParserRULE_cmdSe)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(CalcLexerParserSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Expressao()
	}
	{
		p.SetState(309)
		p.Match(CalcLexerParserENTAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
		{
			p.SetState(310)
			p.Cmd()
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserSENAO {
		{
			p.SetState(316)
			p.Match(CalcLexerParserSENAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
			{
				p.SetState(317)
				p.Cmd()
			}

			p.SetState(322)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(325)
		p.Match(CalcLexerParserFIM_SE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdCasoContext is an interface to support dynamic dispatch.
type ICmdCasoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASO() antlr.TerminalNode
	Exp_aritmetica() IExp_aritmeticaContext
	SEJA() antlr.TerminalNode
	Selecao() ISelecaoContext
	FIM_CASO() antlr.TerminalNode
	SENAO() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdCasoContext differentiates from other interfaces.
	IsCmdCasoContext()
}

type CmdCasoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdCasoContext() *CmdCasoContext {
	var p = new(CmdCasoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdCaso
	return p
}

func InitEmptyCmdCasoContext(p *CmdCasoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdCaso
}

func (*CmdCasoContext) IsCmdCasoContext() {}

func NewCmdCasoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdCasoContext {
	var p = new(CmdCasoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdCaso

	return p
}

func (s *CmdCasoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdCasoContext) CASO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserCASO, 0)
}

func (s *CmdCasoContext) Exp_aritmetica() IExp_aritmeticaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *CmdCasoContext) SEJA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserSEJA, 0)
}

func (s *CmdCasoContext) Selecao() ISelecaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelecaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelecaoContext)
}

func (s *CmdCasoContext) FIM_CASO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_CASO, 0)
}

func (s *CmdCasoContext) SENAO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserSENAO, 0)
}

func (s *CmdCasoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdCasoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdCasoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdCasoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdCasoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdCaso(s)
	}
}

func (s *CmdCasoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdCaso(s)
	}
}

func (p *CalcLexerParser) CmdCaso() (localctx ICmdCasoContext) {
	localctx = NewCmdCasoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CalcLexerParserRULE_cmdCaso)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(CalcLexerParserCASO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.Exp_aritmetica()
	}
	{
		p.SetState(329)
		p.Match(CalcLexerParserSEJA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Selecao()
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserSENAO {
		{
			p.SetState(331)
			p.Match(CalcLexerParserSENAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
			{
				p.SetState(332)
				p.Cmd()
			}

			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(340)
		p.Match(CalcLexerParserFIM_CASO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdParaContext is an interface to support dynamic dispatch.
type ICmdParaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARA() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ATRIB() antlr.TerminalNode
	AllExp_aritmetica() []IExp_aritmeticaContext
	Exp_aritmetica(i int) IExp_aritmeticaContext
	ATE() antlr.TerminalNode
	FACA() antlr.TerminalNode
	FIM_PARA() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdParaContext differentiates from other interfaces.
	IsCmdParaContext()
}

type CmdParaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdParaContext() *CmdParaContext {
	var p = new(CmdParaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdPara
	return p
}

func InitEmptyCmdParaContext(p *CmdParaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdPara
}

func (*CmdParaContext) IsCmdParaContext() {}

func NewCmdParaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdParaContext {
	var p = new(CmdParaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdPara

	return p
}

func (s *CmdParaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdParaContext) PARA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserPARA, 0)
}

func (s *CmdParaContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserIDENT, 0)
}

func (s *CmdParaContext) ATRIB() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserATRIB, 0)
}

func (s *CmdParaContext) AllExp_aritmetica() []IExp_aritmeticaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			len++
		}
	}

	tst := make([]IExp_aritmeticaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExp_aritmeticaContext); ok {
			tst[i] = t.(IExp_aritmeticaContext)
			i++
		}
	}

	return tst
}

func (s *CmdParaContext) Exp_aritmetica(i int) IExp_aritmeticaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *CmdParaContext) ATE() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserATE, 0)
}

func (s *CmdParaContext) FACA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFACA, 0)
}

func (s *CmdParaContext) FIM_PARA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_PARA, 0)
}

func (s *CmdParaContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdParaContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdParaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdParaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdParaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdPara(s)
	}
}

func (s *CmdParaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdPara(s)
	}
}

func (p *CalcLexerParser) CmdPara() (localctx ICmdParaContext) {
	localctx = NewCmdParaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CalcLexerParserRULE_cmdPara)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(CalcLexerParserPARA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(CalcLexerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Match(CalcLexerParserATRIB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Exp_aritmetica()
	}
	{
		p.SetState(346)
		p.Match(CalcLexerParserATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Exp_aritmetica()
	}
	{
		p.SetState(348)
		p.Match(CalcLexerParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
		{
			p.SetState(349)
			p.Cmd()
		}

		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(355)
		p.Match(CalcLexerParserFIM_PARA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdEnquantoContext is an interface to support dynamic dispatch.
type ICmdEnquantoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENQUANTO() antlr.TerminalNode
	Expressao() IExpressaoContext
	FACA() antlr.TerminalNode
	FIM_ENQUANTO() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdEnquantoContext differentiates from other interfaces.
	IsCmdEnquantoContext()
}

type CmdEnquantoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdEnquantoContext() *CmdEnquantoContext {
	var p = new(CmdEnquantoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdEnquanto
	return p
}

func InitEmptyCmdEnquantoContext(p *CmdEnquantoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdEnquanto
}

func (*CmdEnquantoContext) IsCmdEnquantoContext() {}

func NewCmdEnquantoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdEnquantoContext {
	var p = new(CmdEnquantoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdEnquanto

	return p
}

func (s *CmdEnquantoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdEnquantoContext) ENQUANTO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserENQUANTO, 0)
}

func (s *CmdEnquantoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdEnquantoContext) FACA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFACA, 0)
}

func (s *CmdEnquantoContext) FIM_ENQUANTO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFIM_ENQUANTO, 0)
}

func (s *CmdEnquantoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdEnquantoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdEnquantoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdEnquantoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdEnquantoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdEnquanto(s)
	}
}

func (s *CmdEnquantoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdEnquanto(s)
	}
}

func (p *CalcLexerParser) CmdEnquanto() (localctx ICmdEnquantoContext) {
	localctx = NewCmdEnquantoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CalcLexerParserRULE_cmdEnquanto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(CalcLexerParserENQUANTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Expressao()
	}
	{
		p.SetState(359)
		p.Match(CalcLexerParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
		{
			p.SetState(360)
			p.Cmd()
		}

		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(366)
		p.Match(CalcLexerParserFIM_ENQUANTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdFacaContext is an interface to support dynamic dispatch.
type ICmdFacaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FACA() antlr.TerminalNode
	ATE() antlr.TerminalNode
	Expressao() IExpressaoContext
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdFacaContext differentiates from other interfaces.
	IsCmdFacaContext()
}

type CmdFacaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdFacaContext() *CmdFacaContext {
	var p = new(CmdFacaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdFaca
	return p
}

func InitEmptyCmdFacaContext(p *CmdFacaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdFaca
}

func (*CmdFacaContext) IsCmdFacaContext() {}

func NewCmdFacaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdFacaContext {
	var p = new(CmdFacaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdFaca

	return p
}

func (s *CmdFacaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdFacaContext) FACA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFACA, 0)
}

func (s *CmdFacaContext) ATE() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserATE, 0)
}

func (s *CmdFacaContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdFacaContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdFacaContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdFacaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdFacaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdFacaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdFaca(s)
	}
}

func (s *CmdFacaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdFaca(s)
	}
}

func (p *CalcLexerParser) CmdFaca() (localctx ICmdFacaContext) {
	localctx = NewCmdFacaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CalcLexerParserRULE_cmdFaca)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(CalcLexerParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
		{
			p.SetState(369)
			p.Cmd()
		}

		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(375)
		p.Match(CalcLexerParserATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Expressao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdAtribuicaoContext is an interface to support dynamic dispatch.
type ICmdAtribuicaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() IIdentificadorContext
	ATRIB() antlr.TerminalNode
	Expressao() IExpressaoContext

	// IsCmdAtribuicaoContext differentiates from other interfaces.
	IsCmdAtribuicaoContext()
}

type CmdAtribuicaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdAtribuicaoContext() *CmdAtribuicaoContext {
	var p = new(CmdAtribuicaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdAtribuicao
	return p
}

func InitEmptyCmdAtribuicaoContext(p *CmdAtribuicaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdAtribuicao
}

func (*CmdAtribuicaoContext) IsCmdAtribuicaoContext() {}

func NewCmdAtribuicaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdAtribuicaoContext {
	var p = new(CmdAtribuicaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdAtribuicao

	return p
}

func (s *CmdAtribuicaoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdAtribuicaoContext) Identificador() IIdentificadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *CmdAtribuicaoContext) ATRIB() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserATRIB, 0)
}

func (s *CmdAtribuicaoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdAtribuicaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdAtribuicaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdAtribuicaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdAtribuicao(s)
	}
}

func (s *CmdAtribuicaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdAtribuicao(s)
	}
}

func (p *CalcLexerParser) CmdAtribuicao() (localctx ICmdAtribuicaoContext) {
	localctx = NewCmdAtribuicaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CalcLexerParserRULE_cmdAtribuicao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserT__1 {
		{
			p.SetState(378)
			p.Match(CalcLexerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(381)
		p.Identificador()
	}
	{
		p.SetState(382)
		p.Match(CalcLexerParserATRIB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Expressao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdChamadaContext is an interface to support dynamic dispatch.
type ICmdChamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllExpressao() []IExpressaoContext
	Expressao(i int) IExpressaoContext
	FECHAPAR() antlr.TerminalNode
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode

	// IsCmdChamadaContext differentiates from other interfaces.
	IsCmdChamadaContext()
}

type CmdChamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdChamadaContext() *CmdChamadaContext {
	var p = new(CmdChamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdChamada
	return p
}

func InitEmptyCmdChamadaContext(p *CmdChamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdChamada
}

func (*CmdChamadaContext) IsCmdChamadaContext() {}

func NewCmdChamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdChamadaContext {
	var p = new(CmdChamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdChamada

	return p
}

func (s *CmdChamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdChamadaContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserIDENT, 0)
}

func (s *CmdChamadaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserABREPAR, 0)
}

func (s *CmdChamadaContext) AllExpressao() []IExpressaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressaoContext); ok {
			len++
		}
	}

	tst := make([]IExpressaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressaoContext); ok {
			tst[i] = t.(IExpressaoContext)
			i++
		}
	}

	return tst
}

func (s *CmdChamadaContext) Expressao(i int) IExpressaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdChamadaContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFECHAPAR, 0)
}

func (s *CmdChamadaContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *CmdChamadaContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *CmdChamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdChamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdChamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdChamada(s)
	}
}

func (s *CmdChamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdChamada(s)
	}
}

func (p *CalcLexerParser) CmdChamada() (localctx ICmdChamadaContext) {
	localctx = NewCmdChamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CalcLexerParserRULE_cmdChamada)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(CalcLexerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(CalcLexerParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.Expressao()
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserVIRG {
		{
			p.SetState(388)
			p.Match(CalcLexerParserVIRG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.Expressao()
		}

		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(395)
		p.Match(CalcLexerParserFECHAPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdRetorneContext is an interface to support dynamic dispatch.
type ICmdRetorneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETORNE() antlr.TerminalNode
	Expressao() IExpressaoContext

	// IsCmdRetorneContext differentiates from other interfaces.
	IsCmdRetorneContext()
}

type CmdRetorneContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdRetorneContext() *CmdRetorneContext {
	var p = new(CmdRetorneContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdRetorne
	return p
}

func InitEmptyCmdRetorneContext(p *CmdRetorneContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_cmdRetorne
}

func (*CmdRetorneContext) IsCmdRetorneContext() {}

func NewCmdRetorneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdRetorneContext {
	var p = new(CmdRetorneContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_cmdRetorne

	return p
}

func (s *CmdRetorneContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdRetorneContext) RETORNE() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserRETORNE, 0)
}

func (s *CmdRetorneContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdRetorneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdRetorneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdRetorneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterCmdRetorne(s)
	}
}

func (s *CmdRetorneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitCmdRetorne(s)
	}
}

func (p *CalcLexerParser) CmdRetorne() (localctx ICmdRetorneContext) {
	localctx = NewCmdRetorneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CalcLexerParserRULE_cmdRetorne)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(CalcLexerParserRETORNE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Expressao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelecaoContext is an interface to support dynamic dispatch.
type ISelecaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllItem_selecao() []IItem_selecaoContext
	Item_selecao(i int) IItem_selecaoContext

	// IsSelecaoContext differentiates from other interfaces.
	IsSelecaoContext()
}

type SelecaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelecaoContext() *SelecaoContext {
	var p = new(SelecaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_selecao
	return p
}

func InitEmptySelecaoContext(p *SelecaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_selecao
}

func (*SelecaoContext) IsSelecaoContext() {}

func NewSelecaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelecaoContext {
	var p = new(SelecaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_selecao

	return p
}

func (s *SelecaoContext) GetParser() antlr.Parser { return s.parser }

func (s *SelecaoContext) AllItem_selecao() []IItem_selecaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItem_selecaoContext); ok {
			len++
		}
	}

	tst := make([]IItem_selecaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItem_selecaoContext); ok {
			tst[i] = t.(IItem_selecaoContext)
			i++
		}
	}

	return tst
}

func (s *SelecaoContext) Item_selecao(i int) IItem_selecaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_selecaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_selecaoContext)
}

func (s *SelecaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelecaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelecaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterSelecao(s)
	}
}

func (s *SelecaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitSelecao(s)
	}
}

func (p *CalcLexerParser) Selecao() (localctx ISelecaoContext) {
	localctx = NewSelecaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CalcLexerParserRULE_selecao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserT__2 || _la == CalcLexerParserNUM_INT {
		{
			p.SetState(400)
			p.Item_selecao()
		}

		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItem_selecaoContext is an interface to support dynamic dispatch.
type IItem_selecaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constantes() IConstantesContext
	DOIS_PONTOS() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsItem_selecaoContext differentiates from other interfaces.
	IsItem_selecaoContext()
}

type Item_selecaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_selecaoContext() *Item_selecaoContext {
	var p = new(Item_selecaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_item_selecao
	return p
}

func InitEmptyItem_selecaoContext(p *Item_selecaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_item_selecao
}

func (*Item_selecaoContext) IsItem_selecaoContext() {}

func NewItem_selecaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_selecaoContext {
	var p = new(Item_selecaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_item_selecao

	return p
}

func (s *Item_selecaoContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_selecaoContext) Constantes() IConstantesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantesContext)
}

func (s *Item_selecaoContext) DOIS_PONTOS() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserDOIS_PONTOS, 0)
}

func (s *Item_selecaoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *Item_selecaoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *Item_selecaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_selecaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Item_selecaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterItem_selecao(s)
	}
}

func (s *Item_selecaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitItem_selecao(s)
	}
}

func (p *CalcLexerParser) Item_selecao() (localctx IItem_selecaoContext) {
	localctx = NewItem_selecaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CalcLexerParserRULE_item_selecao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Constantes()
	}
	{
		p.SetState(407)
		p.Match(CalcLexerParserDOIS_PONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-9223367633101520895) != 0 {
		{
			p.SetState(408)
			p.Cmd()
		}

		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantesContext is an interface to support dynamic dispatch.
type IConstantesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumero_intervalo() []INumero_intervaloContext
	Numero_intervalo(i int) INumero_intervaloContext
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode

	// IsConstantesContext differentiates from other interfaces.
	IsConstantesContext()
}

type ConstantesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantesContext() *ConstantesContext {
	var p = new(ConstantesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_constantes
	return p
}

func InitEmptyConstantesContext(p *ConstantesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_constantes
}

func (*ConstantesContext) IsConstantesContext() {}

func NewConstantesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantesContext {
	var p = new(ConstantesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_constantes

	return p
}

func (s *ConstantesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantesContext) AllNumero_intervalo() []INumero_intervaloContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumero_intervaloContext); ok {
			len++
		}
	}

	tst := make([]INumero_intervaloContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumero_intervaloContext); ok {
			tst[i] = t.(INumero_intervaloContext)
			i++
		}
	}

	return tst
}

func (s *ConstantesContext) Numero_intervalo(i int) INumero_intervaloContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumero_intervaloContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumero_intervaloContext)
}

func (s *ConstantesContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *ConstantesContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *ConstantesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterConstantes(s)
	}
}

func (s *ConstantesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitConstantes(s)
	}
}

func (p *CalcLexerParser) Constantes() (localctx IConstantesContext) {
	localctx = NewConstantesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CalcLexerParserRULE_constantes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Numero_intervalo()
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserVIRG {
		{
			p.SetState(415)
			p.Match(CalcLexerParserVIRG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Numero_intervalo()
		}

		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumero_intervaloContext is an interface to support dynamic dispatch.
type INumero_intervaloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNUM_INT() []antlr.TerminalNode
	NUM_INT(i int) antlr.TerminalNode
	AllOp_unario() []IOp_unarioContext
	Op_unario(i int) IOp_unarioContext
	PONTOPONTO() antlr.TerminalNode

	// IsNumero_intervaloContext differentiates from other interfaces.
	IsNumero_intervaloContext()
}

type Numero_intervaloContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumero_intervaloContext() *Numero_intervaloContext {
	var p = new(Numero_intervaloContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_numero_intervalo
	return p
}

func InitEmptyNumero_intervaloContext(p *Numero_intervaloContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_numero_intervalo
}

func (*Numero_intervaloContext) IsNumero_intervaloContext() {}

func NewNumero_intervaloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numero_intervaloContext {
	var p = new(Numero_intervaloContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_numero_intervalo

	return p
}

func (s *Numero_intervaloContext) GetParser() antlr.Parser { return s.parser }

func (s *Numero_intervaloContext) AllNUM_INT() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserNUM_INT)
}

func (s *Numero_intervaloContext) NUM_INT(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserNUM_INT, i)
}

func (s *Numero_intervaloContext) AllOp_unario() []IOp_unarioContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp_unarioContext); ok {
			len++
		}
	}

	tst := make([]IOp_unarioContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp_unarioContext); ok {
			tst[i] = t.(IOp_unarioContext)
			i++
		}
	}

	return tst
}

func (s *Numero_intervaloContext) Op_unario(i int) IOp_unarioContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_unarioContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_unarioContext)
}

func (s *Numero_intervaloContext) PONTOPONTO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserPONTOPONTO, 0)
}

func (s *Numero_intervaloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numero_intervaloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numero_intervaloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterNumero_intervalo(s)
	}
}

func (s *Numero_intervaloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitNumero_intervalo(s)
	}
}

func (p *CalcLexerParser) Numero_intervalo() (localctx INumero_intervaloContext) {
	localctx = NewNumero_intervaloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CalcLexerParserRULE_numero_intervalo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserT__2 {
		{
			p.SetState(422)
			p.Op_unario()
		}

	}
	{
		p.SetState(425)
		p.Match(CalcLexerParserNUM_INT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserPONTOPONTO {
		{
			p.SetState(426)
			p.Match(CalcLexerParserPONTOPONTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CalcLexerParserT__2 {
			{
				p.SetState(427)
				p.Op_unario()
			}

		}
		{
			p.SetState(430)
			p.Match(CalcLexerParserNUM_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_unarioContext is an interface to support dynamic dispatch.
type IOp_unarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp_unarioContext differentiates from other interfaces.
	IsOp_unarioContext()
}

type Op_unarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_unarioContext() *Op_unarioContext {
	var p = new(Op_unarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_unario
	return p
}

func InitEmptyOp_unarioContext(p *Op_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_unario
}

func (*Op_unarioContext) IsOp_unarioContext() {}

func NewOp_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_unarioContext {
	var p = new(Op_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_op_unario

	return p
}

func (s *Op_unarioContext) GetParser() antlr.Parser { return s.parser }
func (s *Op_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterOp_unario(s)
	}
}

func (s *Op_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitOp_unario(s)
	}
}

func (p *CalcLexerParser) Op_unario() (localctx IOp_unarioContext) {
	localctx = NewOp_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CalcLexerParserRULE_op_unario)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)
		p.Match(CalcLexerParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExp_aritmeticaContext is an interface to support dynamic dispatch.
type IExp_aritmeticaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTermo() []ITermoContext
	Termo(i int) ITermoContext
	AllOp1() []IOp1Context
	Op1(i int) IOp1Context

	// IsExp_aritmeticaContext differentiates from other interfaces.
	IsExp_aritmeticaContext()
}

type Exp_aritmeticaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp_aritmeticaContext() *Exp_aritmeticaContext {
	var p = new(Exp_aritmeticaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_exp_aritmetica
	return p
}

func InitEmptyExp_aritmeticaContext(p *Exp_aritmeticaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_exp_aritmetica
}

func (*Exp_aritmeticaContext) IsExp_aritmeticaContext() {}

func NewExp_aritmeticaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_aritmeticaContext {
	var p = new(Exp_aritmeticaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_exp_aritmetica

	return p
}

func (s *Exp_aritmeticaContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_aritmeticaContext) AllTermo() []ITermoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermoContext); ok {
			len++
		}
	}

	tst := make([]ITermoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermoContext); ok {
			tst[i] = t.(ITermoContext)
			i++
		}
	}

	return tst
}

func (s *Exp_aritmeticaContext) Termo(i int) ITermoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoContext)
}

func (s *Exp_aritmeticaContext) AllOp1() []IOp1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp1Context); ok {
			len++
		}
	}

	tst := make([]IOp1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp1Context); ok {
			tst[i] = t.(IOp1Context)
			i++
		}
	}

	return tst
}

func (s *Exp_aritmeticaContext) Op1(i int) IOp1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp1Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp1Context)
}

func (s *Exp_aritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_aritmeticaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_aritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterExp_aritmetica(s)
	}
}

func (s *Exp_aritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitExp_aritmetica(s)
	}
}

func (p *CalcLexerParser) Exp_aritmetica() (localctx IExp_aritmeticaContext) {
	localctx = NewExp_aritmeticaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CalcLexerParserRULE_exp_aritmetica)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		p.Termo()
	}
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(436)
				p.Op1()
			}
			{
				p.SetState(437)
				p.Termo()
			}

		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermoContext is an interface to support dynamic dispatch.
type ITermoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFator() []IFatorContext
	Fator(i int) IFatorContext
	AllOp2() []IOp2Context
	Op2(i int) IOp2Context

	// IsTermoContext differentiates from other interfaces.
	IsTermoContext()
}

type TermoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermoContext() *TermoContext {
	var p = new(TermoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_termo
	return p
}

func InitEmptyTermoContext(p *TermoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_termo
}

func (*TermoContext) IsTermoContext() {}

func NewTermoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoContext {
	var p = new(TermoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_termo

	return p
}

func (s *TermoContext) GetParser() antlr.Parser { return s.parser }

func (s *TermoContext) AllFator() []IFatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFatorContext); ok {
			len++
		}
	}

	tst := make([]IFatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFatorContext); ok {
			tst[i] = t.(IFatorContext)
			i++
		}
	}

	return tst
}

func (s *TermoContext) Fator(i int) IFatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFatorContext)
}

func (s *TermoContext) AllOp2() []IOp2Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp2Context); ok {
			len++
		}
	}

	tst := make([]IOp2Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp2Context); ok {
			tst[i] = t.(IOp2Context)
			i++
		}
	}

	return tst
}

func (s *TermoContext) Op2(i int) IOp2Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp2Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp2Context)
}

func (s *TermoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterTermo(s)
	}
}

func (s *TermoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitTermo(s)
	}
}

func (p *CalcLexerParser) Termo() (localctx ITermoContext) {
	localctx = NewTermoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CalcLexerParserRULE_termo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Fator()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserT__4 || _la == CalcLexerParserT__5 {
		{
			p.SetState(445)
			p.Op2()
		}
		{
			p.SetState(446)
			p.Fator()
		}

		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFatorContext is an interface to support dynamic dispatch.
type IFatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParcela() []IParcelaContext
	Parcela(i int) IParcelaContext
	AllOp3() []IOp3Context
	Op3(i int) IOp3Context

	// IsFatorContext differentiates from other interfaces.
	IsFatorContext()
}

type FatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFatorContext() *FatorContext {
	var p = new(FatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_fator
	return p
}

func InitEmptyFatorContext(p *FatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_fator
}

func (*FatorContext) IsFatorContext() {}

func NewFatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FatorContext {
	var p = new(FatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_fator

	return p
}

func (s *FatorContext) GetParser() antlr.Parser { return s.parser }

func (s *FatorContext) AllParcela() []IParcelaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParcelaContext); ok {
			len++
		}
	}

	tst := make([]IParcelaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParcelaContext); ok {
			tst[i] = t.(IParcelaContext)
			i++
		}
	}

	return tst
}

func (s *FatorContext) Parcela(i int) IParcelaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcelaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcelaContext)
}

func (s *FatorContext) AllOp3() []IOp3Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp3Context); ok {
			len++
		}
	}

	tst := make([]IOp3Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp3Context); ok {
			tst[i] = t.(IOp3Context)
			i++
		}
	}

	return tst
}

func (s *FatorContext) Op3(i int) IOp3Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp3Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp3Context)
}

func (s *FatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterFator(s)
	}
}

func (s *FatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitFator(s)
	}
}

func (p *CalcLexerParser) Fator() (localctx IFatorContext) {
	localctx = NewFatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CalcLexerParserRULE_fator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Parcela()
	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserT__6 {
		{
			p.SetState(454)
			p.Op3()
		}
		{
			p.SetState(455)
			p.Parcela()
		}

		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp1Context is an interface to support dynamic dispatch.
type IOp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp1Context differentiates from other interfaces.
	IsOp1Context()
}

type Op1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp1Context() *Op1Context {
	var p = new(Op1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op1
	return p
}

func InitEmptyOp1Context(p *Op1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op1
}

func (*Op1Context) IsOp1Context() {}

func NewOp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op1Context {
	var p = new(Op1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_op1

	return p
}

func (s *Op1Context) GetParser() antlr.Parser { return s.parser }
func (s *Op1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterOp1(s)
	}
}

func (s *Op1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitOp1(s)
	}
}

func (p *CalcLexerParser) Op1() (localctx IOp1Context) {
	localctx = NewOp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CalcLexerParserRULE_op1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcLexerParserT__2 || _la == CalcLexerParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp2Context is an interface to support dynamic dispatch.
type IOp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp2Context differentiates from other interfaces.
	IsOp2Context()
}

type Op2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp2Context() *Op2Context {
	var p = new(Op2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op2
	return p
}

func InitEmptyOp2Context(p *Op2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op2
}

func (*Op2Context) IsOp2Context() {}

func NewOp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op2Context {
	var p = new(Op2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_op2

	return p
}

func (s *Op2Context) GetParser() antlr.Parser { return s.parser }
func (s *Op2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterOp2(s)
	}
}

func (s *Op2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitOp2(s)
	}
}

func (p *CalcLexerParser) Op2() (localctx IOp2Context) {
	localctx = NewOp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CalcLexerParserRULE_op2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcLexerParserT__4 || _la == CalcLexerParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp3Context is an interface to support dynamic dispatch.
type IOp3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp3Context differentiates from other interfaces.
	IsOp3Context()
}

type Op3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp3Context() *Op3Context {
	var p = new(Op3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op3
	return p
}

func InitEmptyOp3Context(p *Op3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op3
}

func (*Op3Context) IsOp3Context() {}

func NewOp3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op3Context {
	var p = new(Op3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_op3

	return p
}

func (s *Op3Context) GetParser() antlr.Parser { return s.parser }
func (s *Op3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterOp3(s)
	}
}

func (s *Op3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitOp3(s)
	}
}

func (p *CalcLexerParser) Op3() (localctx IOp3Context) {
	localctx = NewOp3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CalcLexerParserRULE_op3)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(CalcLexerParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcelaContext is an interface to support dynamic dispatch.
type IParcelaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Parcela_unario() IParcela_unarioContext
	Op_unario() IOp_unarioContext
	Parcela_nao_unario() IParcela_nao_unarioContext

	// IsParcelaContext differentiates from other interfaces.
	IsParcelaContext()
}

type ParcelaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcelaContext() *ParcelaContext {
	var p = new(ParcelaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela
	return p
}

func InitEmptyParcelaContext(p *ParcelaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela
}

func (*ParcelaContext) IsParcelaContext() {}

func NewParcelaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParcelaContext {
	var p = new(ParcelaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_parcela

	return p
}

func (s *ParcelaContext) GetParser() antlr.Parser { return s.parser }

func (s *ParcelaContext) Parcela_unario() IParcela_unarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcela_unarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcela_unarioContext)
}

func (s *ParcelaContext) Op_unario() IOp_unarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_unarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_unarioContext)
}

func (s *ParcelaContext) Parcela_nao_unario() IParcela_nao_unarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcela_nao_unarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcela_nao_unarioContext)
}

func (s *ParcelaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParcelaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParcelaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterParcela(s)
	}
}

func (s *ParcelaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitParcela(s)
	}
}

func (p *CalcLexerParser) Parcela() (localctx IParcelaContext) {
	localctx = NewParcelaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CalcLexerParserRULE_parcela)
	var _la int

	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserT__1, CalcLexerParserT__2, CalcLexerParserABREPAR, CalcLexerParserNUM_INT, CalcLexerParserNUM_REAL, CalcLexerParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CalcLexerParserT__2 {
			{
				p.SetState(468)
				p.Op_unario()
			}

		}
		{
			p.SetState(471)
			p.Parcela_unario()
		}

	case CalcLexerParserENDERECO, CalcLexerParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(472)
			p.Parcela_nao_unario()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcela_unarioContext is an interface to support dynamic dispatch.
type IParcela_unarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() IIdentificadorContext
	IDENT() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllExpressao() []IExpressaoContext
	Expressao(i int) IExpressaoContext
	FECHAPAR() antlr.TerminalNode
	AllVIRG() []antlr.TerminalNode
	VIRG(i int) antlr.TerminalNode
	NUM_INT() antlr.TerminalNode
	NUM_REAL() antlr.TerminalNode

	// IsParcela_unarioContext differentiates from other interfaces.
	IsParcela_unarioContext()
}

type Parcela_unarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcela_unarioContext() *Parcela_unarioContext {
	var p = new(Parcela_unarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela_unario
	return p
}

func InitEmptyParcela_unarioContext(p *Parcela_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela_unario
}

func (*Parcela_unarioContext) IsParcela_unarioContext() {}

func NewParcela_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_unarioContext {
	var p = new(Parcela_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_parcela_unario

	return p
}

func (s *Parcela_unarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_unarioContext) Identificador() IIdentificadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *Parcela_unarioContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserIDENT, 0)
}

func (s *Parcela_unarioContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserABREPAR, 0)
}

func (s *Parcela_unarioContext) AllExpressao() []IExpressaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressaoContext); ok {
			len++
		}
	}

	tst := make([]IExpressaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressaoContext); ok {
			tst[i] = t.(IExpressaoContext)
			i++
		}
	}

	return tst
}

func (s *Parcela_unarioContext) Expressao(i int) IExpressaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *Parcela_unarioContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFECHAPAR, 0)
}

func (s *Parcela_unarioContext) AllVIRG() []antlr.TerminalNode {
	return s.GetTokens(CalcLexerParserVIRG)
}

func (s *Parcela_unarioContext) VIRG(i int) antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVIRG, i)
}

func (s *Parcela_unarioContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserNUM_INT, 0)
}

func (s *Parcela_unarioContext) NUM_REAL() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserNUM_REAL, 0)
}

func (s *Parcela_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterParcela_unario(s)
	}
}

func (s *Parcela_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitParcela_unario(s)
	}
}

func (p *CalcLexerParser) Parcela_unario() (localctx IParcela_unarioContext) {
	localctx = NewParcela_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CalcLexerParserRULE_parcela_unario)
	var _la int

	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CalcLexerParserT__1 {
			{
				p.SetState(475)
				p.Match(CalcLexerParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(478)
			p.Identificador()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(479)
			p.Match(CalcLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.Match(CalcLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.Expressao()
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CalcLexerParserVIRG {
			{
				p.SetState(482)
				p.Match(CalcLexerParserVIRG)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(483)
				p.Expressao()
			}

			p.SetState(488)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(489)
			p.Match(CalcLexerParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(491)
			p.Match(CalcLexerParserNUM_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(492)
			p.Match(CalcLexerParserNUM_REAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(493)
			p.Match(CalcLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.Expressao()
		}
		{
			p.SetState(495)
			p.Match(CalcLexerParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcela_nao_unarioContext is an interface to support dynamic dispatch.
type IParcela_nao_unarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENDERECO() antlr.TerminalNode
	Identificador() IIdentificadorContext
	CADEIA() antlr.TerminalNode

	// IsParcela_nao_unarioContext differentiates from other interfaces.
	IsParcela_nao_unarioContext()
}

type Parcela_nao_unarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcela_nao_unarioContext() *Parcela_nao_unarioContext {
	var p = new(Parcela_nao_unarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela_nao_unario
	return p
}

func InitEmptyParcela_nao_unarioContext(p *Parcela_nao_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela_nao_unario
}

func (*Parcela_nao_unarioContext) IsParcela_nao_unarioContext() {}

func NewParcela_nao_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_nao_unarioContext {
	var p = new(Parcela_nao_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_parcela_nao_unario

	return p
}

func (s *Parcela_nao_unarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_nao_unarioContext) ENDERECO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserENDERECO, 0)
}

func (s *Parcela_nao_unarioContext) Identificador() IIdentificadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *Parcela_nao_unarioContext) CADEIA() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserCADEIA, 0)
}

func (s *Parcela_nao_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_nao_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_nao_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterParcela_nao_unario(s)
	}
}

func (s *Parcela_nao_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitParcela_nao_unario(s)
	}
}

func (p *CalcLexerParser) Parcela_nao_unario() (localctx IParcela_nao_unarioContext) {
	localctx = NewParcela_nao_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CalcLexerParserRULE_parcela_nao_unario)
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserENDERECO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Match(CalcLexerParserENDERECO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.Identificador()
		}

	case CalcLexerParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(501)
			p.Match(CalcLexerParserCADEIA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExp_relacionalContext is an interface to support dynamic dispatch.
type IExp_relacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp_aritmetica() []IExp_aritmeticaContext
	Exp_aritmetica(i int) IExp_aritmeticaContext
	Op_relacional() IOp_relacionalContext

	// IsExp_relacionalContext differentiates from other interfaces.
	IsExp_relacionalContext()
}

type Exp_relacionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp_relacionalContext() *Exp_relacionalContext {
	var p = new(Exp_relacionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_exp_relacional
	return p
}

func InitEmptyExp_relacionalContext(p *Exp_relacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_exp_relacional
}

func (*Exp_relacionalContext) IsExp_relacionalContext() {}

func NewExp_relacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_relacionalContext {
	var p = new(Exp_relacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_exp_relacional

	return p
}

func (s *Exp_relacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_relacionalContext) AllExp_aritmetica() []IExp_aritmeticaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			len++
		}
	}

	tst := make([]IExp_aritmeticaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExp_aritmeticaContext); ok {
			tst[i] = t.(IExp_aritmeticaContext)
			i++
		}
	}

	return tst
}

func (s *Exp_relacionalContext) Exp_aritmetica(i int) IExp_aritmeticaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *Exp_relacionalContext) Op_relacional() IOp_relacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_relacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_relacionalContext)
}

func (s *Exp_relacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_relacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_relacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterExp_relacional(s)
	}
}

func (s *Exp_relacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitExp_relacional(s)
	}
}

func (p *CalcLexerParser) Exp_relacional() (localctx IExp_relacionalContext) {
	localctx = NewExp_relacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CalcLexerParserRULE_exp_relacional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		p.Exp_aritmetica()
	}
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7938) != 0 {
		{
			p.SetState(505)
			p.Op_relacional()
		}
		{
			p.SetState(506)
			p.Exp_aritmetica()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_relacionalContext is an interface to support dynamic dispatch.
type IOp_relacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp_relacionalContext differentiates from other interfaces.
	IsOp_relacionalContext()
}

type Op_relacionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_relacionalContext() *Op_relacionalContext {
	var p = new(Op_relacionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_relacional
	return p
}

func InitEmptyOp_relacionalContext(p *Op_relacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_relacional
}

func (*Op_relacionalContext) IsOp_relacionalContext() {}

func NewOp_relacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_relacionalContext {
	var p = new(Op_relacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_op_relacional

	return p
}

func (s *Op_relacionalContext) GetParser() antlr.Parser { return s.parser }
func (s *Op_relacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_relacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_relacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterOp_relacional(s)
	}
}

func (s *Op_relacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitOp_relacional(s)
	}
}

func (p *CalcLexerParser) Op_relacional() (localctx IOp_relacionalContext) {
	localctx = NewOp_relacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, CalcLexerParserRULE_op_relacional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7938) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressaoContext is an interface to support dynamic dispatch.
type IExpressaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTermo_logico() []ITermo_logicoContext
	Termo_logico(i int) ITermo_logicoContext
	AllOp_logico_1() []IOp_logico_1Context
	Op_logico_1(i int) IOp_logico_1Context

	// IsExpressaoContext differentiates from other interfaces.
	IsExpressaoContext()
}

type ExpressaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressaoContext() *ExpressaoContext {
	var p = new(ExpressaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_expressao
	return p
}

func InitEmptyExpressaoContext(p *ExpressaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_expressao
}

func (*ExpressaoContext) IsExpressaoContext() {}

func NewExpressaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressaoContext {
	var p = new(ExpressaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_expressao

	return p
}

func (s *ExpressaoContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressaoContext) AllTermo_logico() []ITermo_logicoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermo_logicoContext); ok {
			len++
		}
	}

	tst := make([]ITermo_logicoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermo_logicoContext); ok {
			tst[i] = t.(ITermo_logicoContext)
			i++
		}
	}

	return tst
}

func (s *ExpressaoContext) Termo_logico(i int) ITermo_logicoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermo_logicoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermo_logicoContext)
}

func (s *ExpressaoContext) AllOp_logico_1() []IOp_logico_1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp_logico_1Context); ok {
			len++
		}
	}

	tst := make([]IOp_logico_1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp_logico_1Context); ok {
			tst[i] = t.(IOp_logico_1Context)
			i++
		}
	}

	return tst
}

func (s *ExpressaoContext) Op_logico_1(i int) IOp_logico_1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_logico_1Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_logico_1Context)
}

func (s *ExpressaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterExpressao(s)
	}
}

func (s *ExpressaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitExpressao(s)
	}
}

func (p *CalcLexerParser) Expressao() (localctx IExpressaoContext) {
	localctx = NewExpressaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, CalcLexerParserRULE_expressao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Termo_logico()
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserT__13 {
		{
			p.SetState(513)
			p.Op_logico_1()
		}
		{
			p.SetState(514)
			p.Termo_logico()
		}

		p.SetState(520)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermo_logicoContext is an interface to support dynamic dispatch.
type ITermo_logicoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFator_logico() []IFator_logicoContext
	Fator_logico(i int) IFator_logicoContext
	AllOp_logico_2() []IOp_logico_2Context
	Op_logico_2(i int) IOp_logico_2Context

	// IsTermo_logicoContext differentiates from other interfaces.
	IsTermo_logicoContext()
}

type Termo_logicoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermo_logicoContext() *Termo_logicoContext {
	var p = new(Termo_logicoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_termo_logico
	return p
}

func InitEmptyTermo_logicoContext(p *Termo_logicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_termo_logico
}

func (*Termo_logicoContext) IsTermo_logicoContext() {}

func NewTermo_logicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Termo_logicoContext {
	var p = new(Termo_logicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_termo_logico

	return p
}

func (s *Termo_logicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Termo_logicoContext) AllFator_logico() []IFator_logicoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFator_logicoContext); ok {
			len++
		}
	}

	tst := make([]IFator_logicoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFator_logicoContext); ok {
			tst[i] = t.(IFator_logicoContext)
			i++
		}
	}

	return tst
}

func (s *Termo_logicoContext) Fator_logico(i int) IFator_logicoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFator_logicoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFator_logicoContext)
}

func (s *Termo_logicoContext) AllOp_logico_2() []IOp_logico_2Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp_logico_2Context); ok {
			len++
		}
	}

	tst := make([]IOp_logico_2Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp_logico_2Context); ok {
			tst[i] = t.(IOp_logico_2Context)
			i++
		}
	}

	return tst
}

func (s *Termo_logicoContext) Op_logico_2(i int) IOp_logico_2Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_logico_2Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_logico_2Context)
}

func (s *Termo_logicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Termo_logicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Termo_logicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterTermo_logico(s)
	}
}

func (s *Termo_logicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitTermo_logico(s)
	}
}

func (p *CalcLexerParser) Termo_logico() (localctx ITermo_logicoContext) {
	localctx = NewTermo_logicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, CalcLexerParserRULE_termo_logico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(521)
		p.Fator_logico()
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcLexerParserT__14 {
		{
			p.SetState(522)
			p.Op_logico_2()
		}
		{
			p.SetState(523)
			p.Fator_logico()
		}

		p.SetState(529)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFator_logicoContext is an interface to support dynamic dispatch.
type IFator_logicoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Parcela_logica() IParcela_logicaContext

	// IsFator_logicoContext differentiates from other interfaces.
	IsFator_logicoContext()
}

type Fator_logicoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFator_logicoContext() *Fator_logicoContext {
	var p = new(Fator_logicoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_fator_logico
	return p
}

func InitEmptyFator_logicoContext(p *Fator_logicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_fator_logico
}

func (*Fator_logicoContext) IsFator_logicoContext() {}

func NewFator_logicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fator_logicoContext {
	var p = new(Fator_logicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_fator_logico

	return p
}

func (s *Fator_logicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Fator_logicoContext) Parcela_logica() IParcela_logicaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcela_logicaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcela_logicaContext)
}

func (s *Fator_logicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fator_logicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fator_logicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterFator_logico(s)
	}
}

func (s *Fator_logicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitFator_logico(s)
	}
}

func (p *CalcLexerParser) Fator_logico() (localctx IFator_logicoContext) {
	localctx = NewFator_logicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, CalcLexerParserRULE_fator_logico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcLexerParserT__12 {
		{
			p.SetState(530)
			p.Match(CalcLexerParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(533)
		p.Parcela_logica()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcela_logicaContext is an interface to support dynamic dispatch.
type IParcela_logicaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VERDADEIRO() antlr.TerminalNode
	FALSO() antlr.TerminalNode
	Exp_relacional() IExp_relacionalContext

	// IsParcela_logicaContext differentiates from other interfaces.
	IsParcela_logicaContext()
}

type Parcela_logicaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcela_logicaContext() *Parcela_logicaContext {
	var p = new(Parcela_logicaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela_logica
	return p
}

func InitEmptyParcela_logicaContext(p *Parcela_logicaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_parcela_logica
}

func (*Parcela_logicaContext) IsParcela_logicaContext() {}

func NewParcela_logicaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_logicaContext {
	var p = new(Parcela_logicaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_parcela_logica

	return p
}

func (s *Parcela_logicaContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_logicaContext) VERDADEIRO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserVERDADEIRO, 0)
}

func (s *Parcela_logicaContext) FALSO() antlr.TerminalNode {
	return s.GetToken(CalcLexerParserFALSO, 0)
}

func (s *Parcela_logicaContext) Exp_relacional() IExp_relacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_relacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_relacionalContext)
}

func (s *Parcela_logicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_logicaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_logicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterParcela_logica(s)
	}
}

func (s *Parcela_logicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitParcela_logica(s)
	}
}

func (p *CalcLexerParser) Parcela_logica() (localctx IParcela_logicaContext) {
	localctx = NewParcela_logicaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, CalcLexerParserRULE_parcela_logica)
	var _la int

	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcLexerParserFALSO, CalcLexerParserVERDADEIRO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(535)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CalcLexerParserFALSO || _la == CalcLexerParserVERDADEIRO) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case CalcLexerParserT__1, CalcLexerParserT__2, CalcLexerParserABREPAR, CalcLexerParserENDERECO, CalcLexerParserNUM_INT, CalcLexerParserNUM_REAL, CalcLexerParserIDENT, CalcLexerParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(536)
			p.Exp_relacional()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_logico_1Context is an interface to support dynamic dispatch.
type IOp_logico_1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp_logico_1Context differentiates from other interfaces.
	IsOp_logico_1Context()
}

type Op_logico_1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_logico_1Context() *Op_logico_1Context {
	var p = new(Op_logico_1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_logico_1
	return p
}

func InitEmptyOp_logico_1Context(p *Op_logico_1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_logico_1
}

func (*Op_logico_1Context) IsOp_logico_1Context() {}

func NewOp_logico_1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_logico_1Context {
	var p = new(Op_logico_1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_op_logico_1

	return p
}

func (s *Op_logico_1Context) GetParser() antlr.Parser { return s.parser }
func (s *Op_logico_1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_logico_1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_logico_1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterOp_logico_1(s)
	}
}

func (s *Op_logico_1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitOp_logico_1(s)
	}
}

func (p *CalcLexerParser) Op_logico_1() (localctx IOp_logico_1Context) {
	localctx = NewOp_logico_1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, CalcLexerParserRULE_op_logico_1)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.Match(CalcLexerParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_logico_2Context is an interface to support dynamic dispatch.
type IOp_logico_2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp_logico_2Context differentiates from other interfaces.
	IsOp_logico_2Context()
}

type Op_logico_2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_logico_2Context() *Op_logico_2Context {
	var p = new(Op_logico_2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_logico_2
	return p
}

func InitEmptyOp_logico_2Context(p *Op_logico_2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcLexerParserRULE_op_logico_2
}

func (*Op_logico_2Context) IsOp_logico_2Context() {}

func NewOp_logico_2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_logico_2Context {
	var p = new(Op_logico_2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcLexerParserRULE_op_logico_2

	return p
}

func (s *Op_logico_2Context) GetParser() antlr.Parser { return s.parser }
func (s *Op_logico_2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_logico_2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_logico_2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.EnterOp_logico_2(s)
	}
}

func (s *Op_logico_2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcLexerListener); ok {
		listenerT.ExitOp_logico_2(s)
	}
}

func (p *CalcLexerParser) Op_logico_2() (localctx IOp_logico_2Context) {
	localctx = NewOp_logico_2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, CalcLexerParserRULE_op_logico_2)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		p.Match(CalcLexerParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
