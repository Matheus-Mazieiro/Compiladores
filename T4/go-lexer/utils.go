// utils.go

package main

import (
	"fmt"
	parser "go-lexer/parser"
	"strings"
)

var ErrosSemanticos []string

var variaveisComErro = make(map[string]bool)
var tiposComErro = make(map[string]bool)

func AdicionarErroSemantico(linha int, msg string) {
	erro := fmt.Sprintf("Linha %d: %s", linha, msg)
	ErrosSemanticos = append(ErrosSemanticos, erro)
}

func ErroJaReportado(nome string) bool {
	if variaveisComErro[nome] {
		return true
	}

	variaveisComErro[nome] = true
	return false
}

func ResetarErros() {
	ErrosSemanticos = []string{}
	variaveisComErro = make(map[string]bool)
	tiposComErro = make(map[string]bool)
}

func Compatibilidade(tipo1 TipoJander, tipo2 TipoJander) bool {

	if tipo1 == tipo2 {
		return true
	}

	if (tipo1 == INTEIRO && tipo2 == REAL) ||
		(tipo1 == REAL && tipo2 == INTEIRO) {
		return true
	}

	if (tipo1 == REGISTRO || tipo1 == REGISTRO_TIPO) &&
		(tipo2 == REGISTRO || tipo2 == REGISTRO_TIPO) {
		return true
	}

	return false
}

func CompatibilidadeFuncao(tipo1 TipoJander, tipo2 TipoJander) bool {
	return Compatibilidade(tipo1, tipo2)
}

func VerificarTipoBasico(ctx parser.ITipo_basicoContext) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	switch ctx.GetText() {
	case "literal":
		return LITERAL
	case "inteiro":
		return INTEIRO
	case "real":
		return REAL
	case "logico":
		return LOGICO
	default:
		return INVALIDO
	}
}

func VerificarTipo(tabela *TabelaDeSimbolos, ctx parser.ITipoContext) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	if ctx.Registro() != nil {
		return REGISTRO
	}

	if ctx.Tipo_estendido() != nil {

		tipo := VerificarTipoEstendido(tabela, ctx.Tipo_estendido())

		tb := ctx.Tipo_estendido().Tipo_basico_ident()

		if tb != nil && tb.IDENT() != nil {

			nome := tb.IDENT().GetText()

			if tabela.Existe(nome) &&
				tabela.Verificar(nome) == REGISTRO_TIPO {

				return REGISTRO_TIPO
			}
		}

		return tipo
	}

	return INVALIDO
}

func VerificarTipoEstendido(tabela *TabelaDeSimbolos, ctx parser.ITipo_estendidoContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	base := VerificarTipoBasicoIdent(tabela, ctx.Tipo_basico_ident())

	texto := ctx.GetText()

	if len(texto) > 0 && texto[0] == '^' {
		switch base {
		case INTEIRO:
			return PONTEIRO_INTEIRO
		case REAL:
			return PONTEIRO_REAL
		case LITERAL:
			return PONTEIRO_LITERAL
		case LOGICO:
			return PONTEIRO_LOGICO
		}
	}

	return base
}

func VerificarTipoBasicoIdent(tabela *TabelaDeSimbolos, ctx parser.ITipo_basico_identContext) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	if ctx.Tipo_basico() != nil {
		return VerificarTipoBasico(ctx.Tipo_basico())
	}

	if ctx.IDENT() != nil {

		nome := ctx.IDENT().GetText()

		if !tabela.Existe(nome) {

			if !tiposComErro[nome] {

				tiposComErro[nome] = true

				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"tipo "+nome+" nao declarado",
				)
			}

			return INVALIDO
		}

		return tabela.Verificar(nome)
	}

	return INVALIDO
}

func VerificarIdentificador(tabela *TabelaDeSimbolos, ctx parser.IIdentificadorContext) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	nome := ctx.GetText()

	if !tabela.Existe(nome) {

		if !ErroJaReportado(nome) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"identificador "+nome+" nao declarado",
			)
		}

		return INVALIDO
	}

	return tabela.Verificar(nome)
}

func splitNome(nome string) []string {
	return strings.Split(nome, ".")
}
func VerificarExpressao(
	tabela *TabelaDeSimbolos,
	ctx parser.IExpressaoContext,
) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	termos := ctx.AllTermo_logico()

	if len(termos) == 0 {
		return INVALIDO
	}

	// se tiver operador logico → resultado logico
	if len(termos) > 1 {
		return LOGICO
	}

	return VerificarTermoLogico(
		tabela,
		termos[0],
	)
}
func VerificarTermoLogico(
	tabela *TabelaDeSimbolos,
	ctx parser.ITermo_logicoContext,
) TipoJander {

	fatores := ctx.AllFator_logico()

	if len(fatores) == 0 {
		return INVALIDO
	}

	if len(fatores) > 1 {
		return LOGICO
	}

	return VerificarFatorLogico(
		tabela,
		fatores[0],
	)
}
func VerificarFatorLogico(
	tabela *TabelaDeSimbolos,
	ctx parser.IFator_logicoContext,
) TipoJander {

	return VerificarParcelaLogica(
		tabela,
		ctx.Parcela_logica(),
	)
}

func VerificarParcelaLogica(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcela_logicaContext,
) TipoJander {

	if ctx.VERDADEIRO() != nil ||
		ctx.FALSO() != nil {

		return LOGICO
	}

	return VerificarExpRelacional(
		tabela,
		ctx.Exp_relacional(),
	)
}

func VerificarExpRelacional(
	tabela *TabelaDeSimbolos,
	ctx parser.IExp_relacionalContext,
) TipoJander {

	if ctx.Op_relacional() != nil {

		t1 := VerificarExpAritmetica(
			tabela,
			ctx.Exp_aritmetica(0),
		)

		t2 := VerificarExpAritmetica(
			tabela,
			ctx.Exp_aritmetica(1),
		)

		if Compatibilidade(t1, t2) {
			return LOGICO
		}

		return INVALIDO
	}

	return VerificarExpAritmetica(
		tabela,
		ctx.Exp_aritmetica(0),
	)
}

func VerificarExpAritmetica(
	tabela *TabelaDeSimbolos,
	ctx parser.IExp_aritmeticaContext,
) TipoJander {

	var tipo TipoJander = INVALIDO

	for i, termo := range ctx.AllTermo() {

		t := VerificarTermo(tabela, termo)

		if i == 0 {
			tipo = t
		} else if !Compatibilidade(tipo, t) {
			return INVALIDO
		}
	}

	return tipo
}

func VerificarTermo(
	tabela *TabelaDeSimbolos,
	ctx parser.ITermoContext,
) TipoJander {

	var tipo TipoJander = INVALIDO

	for i, fator := range ctx.AllFator() {

		t := VerificarFator(tabela, fator)

		if i == 0 {
			tipo = t
		} else if !Compatibilidade(tipo, t) {
			return INVALIDO
		}
	}

	return tipo
}

func VerificarFator(
	tabela *TabelaDeSimbolos,
	ctx parser.IFatorContext,
) TipoJander {

	var tipo TipoJander = INVALIDO

	for i, parcela := range ctx.AllParcela() {

		t := VerificarParcela(tabela, parcela)

		if i == 0 {
			tipo = t
		} else if !Compatibilidade(tipo, t) {
			return INVALIDO
		}
	}

	return tipo
}

func VerificarParcela(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcelaContext,
) TipoJander {

	if ctx.Parcela_unario() != nil {
		return VerificarParcelaUnario(
			tabela,
			ctx.Parcela_unario(),
		)
	}

	return VerificarParcelaNaoUnario(
		tabela,
		ctx.Parcela_nao_unario(),
	)
}

func VerificarParcelaUnario(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcela_unarioContext,
) TipoJander {

	if ctx.NUM_INT() != nil {
		return INTEIRO
	}

	if ctx.NUM_REAL() != nil {
		return REAL
	}

	if ctx.Identificador() != nil {
		return VerificarIdentificador(
			tabela,
			ctx.Identificador(),
		)
	}

	if ctx.IDENT() != nil {
		return tabela.ObterTipoRetorno(
			ctx.IDENT().GetText(),
		)
	}

	if len(ctx.AllExpressao()) > 0 {
		return VerificarExpressao(
			tabela,
			ctx.Expressao(0),
		)
	}

	return INVALIDO
}

func VerificarParcelaNaoUnario(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcela_nao_unarioContext,
) TipoJander {

	if ctx.CADEIA() != nil {
		return LITERAL
	}

	if ctx.Identificador() != nil {
		return VerificarIdentificador(
			tabela,
			ctx.Identificador(),
		)
	}

	return INVALIDO
}
