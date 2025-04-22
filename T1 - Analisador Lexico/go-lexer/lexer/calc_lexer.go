// Code generated from CalcLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lexer

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'algoritmo'", "'fim_algoritmo'", "'declare'", "'leia'", "'escreva'",
		"'literal'", "'inteiro'", "'('", "')'", "'['", "']'", "':'", "','",
		"'.'",
	}
	staticData.SymbolicNames = []string{
		"", "ALG", "FIM_ALG", "DECLARE", "LEIA", "ESCREVA", "LITERAL", "INTEIRO",
		"ABREPAR", "FECHAPAR", "ABRECOL", "FECHACOL", "DOIS_PONTOS", "VIRG",
		"PONTO", "ARIT", "RELAC", "LOGIC", "INT", "FLOAT", "IDENT", "CADEIA",
		"ERRO_CADEIA", "WS", "COMENTARIO", "ERRO_COMENTARIO",
	}
	staticData.RuleNames = []string{
		"ALG", "FIM_ALG", "DECLARE", "LEIA", "ESCREVA", "LITERAL", "INTEIRO",
		"ABREPAR", "FECHAPAR", "ABRECOL", "FECHACOL", "DOIS_PONTOS", "VIRG",
		"PONTO", "ARIT", "RELAC", "LOGIC", "INT", "FLOAT", "IDENT", "CADEIA",
		"ESC_SEQ", "ERRO_CADEIA", "WS", "COMENTARIO", "ERRO_COMENTARIO",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 222, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 140, 8, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 148, 8, 16, 1, 17, 4, 17, 151, 8, 17, 11, 17,
		12, 17, 152, 1, 18, 4, 18, 156, 8, 18, 11, 18, 12, 18, 157, 1, 18, 1, 18,
		4, 18, 162, 8, 18, 11, 18, 12, 18, 163, 1, 19, 1, 19, 5, 19, 168, 8, 19,
		10, 19, 12, 19, 171, 9, 19, 1, 20, 1, 20, 1, 20, 5, 20, 176, 8, 20, 10,
		20, 12, 20, 179, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 5, 22, 189, 8, 22, 10, 22, 12, 22, 192, 9, 22, 1, 22, 1, 22, 1,
		23, 4, 23, 197, 8, 23, 11, 23, 12, 23, 198, 1, 23, 1, 23, 1, 24, 1, 24,
		5, 24, 205, 8, 24, 10, 24, 12, 24, 208, 9, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 5, 25, 216, 8, 25, 10, 25, 12, 25, 219, 9, 25, 1, 25,
		1, 25, 2, 206, 217, 0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 0, 45, 22, 47, 23, 49, 24, 51,
		25, 1, 0, 8, 4, 0, 37, 37, 42, 43, 45, 45, 47, 47, 1, 0, 48, 57, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 10,
		10, 34, 34, 39, 39, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 10, 10,
		1, 0, 125, 125, 238, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0,
		0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 53, 1, 0, 0,
		0, 3, 63, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 85, 1, 0, 0, 0, 9, 90, 1, 0,
		0, 0, 11, 98, 1, 0, 0, 0, 13, 106, 1, 0, 0, 0, 15, 114, 1, 0, 0, 0, 17,
		116, 1, 0, 0, 0, 19, 118, 1, 0, 0, 0, 21, 120, 1, 0, 0, 0, 23, 122, 1,
		0, 0, 0, 25, 124, 1, 0, 0, 0, 27, 126, 1, 0, 0, 0, 29, 128, 1, 0, 0, 0,
		31, 139, 1, 0, 0, 0, 33, 147, 1, 0, 0, 0, 35, 150, 1, 0, 0, 0, 37, 155,
		1, 0, 0, 0, 39, 165, 1, 0, 0, 0, 41, 172, 1, 0, 0, 0, 43, 182, 1, 0, 0,
		0, 45, 185, 1, 0, 0, 0, 47, 196, 1, 0, 0, 0, 49, 202, 1, 0, 0, 0, 51, 213,
		1, 0, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5, 108, 0, 0, 55, 56, 5, 103,
		0, 0, 56, 57, 5, 111, 0, 0, 57, 58, 5, 114, 0, 0, 58, 59, 5, 105, 0, 0,
		59, 60, 5, 116, 0, 0, 60, 61, 5, 109, 0, 0, 61, 62, 5, 111, 0, 0, 62, 2,
		1, 0, 0, 0, 63, 64, 5, 102, 0, 0, 64, 65, 5, 105, 0, 0, 65, 66, 5, 109,
		0, 0, 66, 67, 5, 95, 0, 0, 67, 68, 5, 97, 0, 0, 68, 69, 5, 108, 0, 0, 69,
		70, 5, 103, 0, 0, 70, 71, 5, 111, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73, 5,
		105, 0, 0, 73, 74, 5, 116, 0, 0, 74, 75, 5, 109, 0, 0, 75, 76, 5, 111,
		0, 0, 76, 4, 1, 0, 0, 0, 77, 78, 5, 100, 0, 0, 78, 79, 5, 101, 0, 0, 79,
		80, 5, 99, 0, 0, 80, 81, 5, 108, 0, 0, 81, 82, 5, 97, 0, 0, 82, 83, 5,
		114, 0, 0, 83, 84, 5, 101, 0, 0, 84, 6, 1, 0, 0, 0, 85, 86, 5, 108, 0,
		0, 86, 87, 5, 101, 0, 0, 87, 88, 5, 105, 0, 0, 88, 89, 5, 97, 0, 0, 89,
		8, 1, 0, 0, 0, 90, 91, 5, 101, 0, 0, 91, 92, 5, 115, 0, 0, 92, 93, 5, 99,
		0, 0, 93, 94, 5, 114, 0, 0, 94, 95, 5, 101, 0, 0, 95, 96, 5, 118, 0, 0,
		96, 97, 5, 97, 0, 0, 97, 10, 1, 0, 0, 0, 98, 99, 5, 108, 0, 0, 99, 100,
		5, 105, 0, 0, 100, 101, 5, 116, 0, 0, 101, 102, 5, 101, 0, 0, 102, 103,
		5, 114, 0, 0, 103, 104, 5, 97, 0, 0, 104, 105, 5, 108, 0, 0, 105, 12, 1,
		0, 0, 0, 106, 107, 5, 105, 0, 0, 107, 108, 5, 110, 0, 0, 108, 109, 5, 116,
		0, 0, 109, 110, 5, 101, 0, 0, 110, 111, 5, 105, 0, 0, 111, 112, 5, 114,
		0, 0, 112, 113, 5, 111, 0, 0, 113, 14, 1, 0, 0, 0, 114, 115, 5, 40, 0,
		0, 115, 16, 1, 0, 0, 0, 116, 117, 5, 41, 0, 0, 117, 18, 1, 0, 0, 0, 118,
		119, 5, 91, 0, 0, 119, 20, 1, 0, 0, 0, 120, 121, 5, 93, 0, 0, 121, 22,
		1, 0, 0, 0, 122, 123, 5, 58, 0, 0, 123, 24, 1, 0, 0, 0, 124, 125, 5, 44,
		0, 0, 125, 26, 1, 0, 0, 0, 126, 127, 5, 46, 0, 0, 127, 28, 1, 0, 0, 0,
		128, 129, 7, 0, 0, 0, 129, 30, 1, 0, 0, 0, 130, 140, 5, 62, 0, 0, 131,
		132, 5, 62, 0, 0, 132, 140, 5, 61, 0, 0, 133, 140, 5, 60, 0, 0, 134, 135,
		5, 60, 0, 0, 135, 140, 5, 61, 0, 0, 136, 137, 5, 60, 0, 0, 137, 140, 5,
		62, 0, 0, 138, 140, 5, 61, 0, 0, 139, 130, 1, 0, 0, 0, 139, 131, 1, 0,
		0, 0, 139, 133, 1, 0, 0, 0, 139, 134, 1, 0, 0, 0, 139, 136, 1, 0, 0, 0,
		139, 138, 1, 0, 0, 0, 140, 32, 1, 0, 0, 0, 141, 148, 5, 101, 0, 0, 142,
		143, 5, 111, 0, 0, 143, 148, 5, 117, 0, 0, 144, 145, 5, 110, 0, 0, 145,
		146, 5, 97, 0, 0, 146, 148, 5, 111, 0, 0, 147, 141, 1, 0, 0, 0, 147, 142,
		1, 0, 0, 0, 147, 144, 1, 0, 0, 0, 148, 34, 1, 0, 0, 0, 149, 151, 7, 1,
		0, 0, 150, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0,
		152, 153, 1, 0, 0, 0, 153, 36, 1, 0, 0, 0, 154, 156, 7, 1, 0, 0, 155, 154,
		1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0,
		0, 0, 158, 159, 1, 0, 0, 0, 159, 161, 5, 46, 0, 0, 160, 162, 7, 1, 0, 0,
		161, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 38, 1, 0, 0, 0, 165, 169, 7, 2, 0, 0, 166, 168, 7,
		3, 0, 0, 167, 166, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0,
		0, 169, 170, 1, 0, 0, 0, 170, 40, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172,
		177, 5, 34, 0, 0, 173, 176, 3, 43, 21, 0, 174, 176, 8, 4, 0, 0, 175, 173,
		1, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0,
		0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0,
		180, 181, 5, 34, 0, 0, 181, 42, 1, 0, 0, 0, 182, 183, 5, 92, 0, 0, 183,
		184, 5, 39, 0, 0, 184, 44, 1, 0, 0, 0, 185, 190, 5, 34, 0, 0, 186, 189,
		3, 43, 21, 0, 187, 189, 8, 4, 0, 0, 188, 186, 1, 0, 0, 0, 188, 187, 1,
		0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0,
		0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 5, 10, 0, 0, 194,
		46, 1, 0, 0, 0, 195, 197, 7, 5, 0, 0, 196, 195, 1, 0, 0, 0, 197, 198, 1,
		0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0,
		0, 200, 201, 6, 23, 0, 0, 201, 48, 1, 0, 0, 0, 202, 206, 5, 123, 0, 0,
		203, 205, 8, 6, 0, 0, 204, 203, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 209, 1, 0, 0, 0, 208, 206,
		1, 0, 0, 0, 209, 210, 5, 125, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 6,
		24, 0, 0, 212, 50, 1, 0, 0, 0, 213, 217, 5, 123, 0, 0, 214, 216, 8, 7,
		0, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0,
		217, 215, 1, 0, 0, 0, 218, 220, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220,
		221, 5, 10, 0, 0, 221, 52, 1, 0, 0, 0, 14, 0, 139, 147, 152, 157, 163,
		169, 175, 177, 188, 190, 198, 206, 217, 1, 6, 0, 0,
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

// CalcLexerInit initializes any static state used to implement CalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalcLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CalcLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerALG             = 1
	CalcLexerFIM_ALG         = 2
	CalcLexerDECLARE         = 3
	CalcLexerLEIA            = 4
	CalcLexerESCREVA         = 5
	CalcLexerLITERAL         = 6
	CalcLexerINTEIRO         = 7
	CalcLexerABREPAR         = 8
	CalcLexerFECHAPAR        = 9
	CalcLexerABRECOL         = 10
	CalcLexerFECHACOL        = 11
	CalcLexerDOIS_PONTOS     = 12
	CalcLexerVIRG            = 13
	CalcLexerPONTO           = 14
	CalcLexerARIT            = 15
	CalcLexerRELAC           = 16
	CalcLexerLOGIC           = 17
	CalcLexerINT             = 18
	CalcLexerFLOAT           = 19
	CalcLexerIDENT           = 20
	CalcLexerCADEIA          = 21
	CalcLexerERRO_CADEIA     = 22
	CalcLexerWS              = 23
	CalcLexerCOMENTARIO      = 24
	CalcLexerERRO_COMENTARIO = 25
)
