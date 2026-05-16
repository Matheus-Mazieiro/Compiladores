// // LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
// package main

// import (
// 	"fmt"
// 	parser "go-lexer/parser"
// )

// var ErrosSemanticos []string
// var variaveisComErro = make(map[string]bool)
// var variaveisComTipoInvalido = make(map[string]bool)

// func AdicionarErroSemantico(linha int, msg string) {
// 	erro := fmt.Sprintf("Linha %d: %s", linha, msg)
// 	ErrosSemanticos = append(ErrosSemanticos, erro)
// }

// func ErroJaReportado(nome string) bool {
// 	if variaveisComErro[nome] {
// 		return true
// 	}
// 	variaveisComErro[nome] = true
// 	return false
// }

// // ================= TIPOS =================

// func VerificarTipoBasico(ctx parser.ITipo_basicoContext) TipoJander {
// 	if ctx == nil {
// 		return INVALIDO
// 	}

// 	texto := ctx.GetText()

// 	switch texto {
// 	case "literal":
// 		return LITERAL
// 	case "inteiro":
// 		return INTEIRO
// 	case "real":
// 		return REAL
// 	case "logico":
// 		return LOGICO
// 	default:
// 		return INVALIDO // Tipo inválido
// 	}
// }

// func VerificarTipo(tabela *TabelaDeSimbolos, ctx parser.ITipoContext) TipoJander {
// 	if ctx.Tipo_estendido() != nil {
// 		return VerificarTipoBasicoIdent(tabela, ctx.Tipo_estendido().Tipo_basico_ident())
// 	}
// 	return INVALIDO
// }

// func VerificarTipoBasicoIdent(tabela *TabelaDeSimbolos, ctx parser.ITipo_basico_identContext) TipoJander {
// 	if ctx.Tipo_basico() != nil {
// 		return VerificarTipoBasico(ctx.Tipo_basico())
// 	}

// 	if ctx.IDENT() != nil {
// 		nome := ctx.IDENT().GetText()

// 		if !tabela.Existe(nome) {
// 			if !ErroJaReportado(nome) {
// 				AdicionarErroSemantico(ctx.GetStart().GetLine(),
// 					"tipo "+nome+" nao declarado")
// 			}
// 			return INVALIDO
// 		}

// 		return tabela.Verificar(nome)
// 	}

//		return INVALIDO
//	}
//
// LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
package main

import (
	"fmt"
	parser "go-lexer/parser"
)

var ErrosSemanticos []string

var variaveisComErro = make(map[string]bool)
var tiposComErro = make(map[string]bool)

// ================= ERROS =================

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

// ================= COMPATIBILIDADE =================

func Compatibilidade(tipo1 TipoJander, tipo2 TipoJander) bool {

	if tipo1 == tipo2 {
		return true
	}

	// inteiro <-> real
	if (tipo1 == INTEIRO && tipo2 == REAL) ||
		(tipo1 == REAL && tipo2 == INTEIRO) {
		return true
	}

	// registros
	if (tipo1 == REGISTRO || tipo1 == REGISTRO_TIPO) &&
		(tipo2 == REGISTRO || tipo2 == REGISTRO_TIPO) {
		return true
	}

	return false
}

func CompatibilidadeFuncao(tipo1 TipoJander, tipo2 TipoJander) bool {
	return tipo1 == tipo2
}

// ================= TIPOS BASICOS =================

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

// ================= TIPO =================

func VerificarTipo(
	tabela *TabelaDeSimbolos,
	ctx parser.ITipoContext,
) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	// registro inline
	if ctx.Registro() != nil {
		return REGISTRO
	}

	// tipo estendido
	if ctx.Tipo_estendido() != nil {

		tipo := VerificarTipoEstendido(
			tabela,
			ctx.Tipo_estendido(),
		)

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

// ================= TIPO ESTENDIDO =================

func VerificarTipoEstendido(tabela *TabelaDeSimbolos, ctx parser.ITipo_estendidoContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	tipo := VerificarTipoBasicoIdent(tabela, ctx.Tipo_basico_ident())

	// verifica ponteiro pelo texto
	if len(ctx.GetText()) > 0 && ctx.GetText()[0] == '^' {
		return tipo
	}

	return tipo
}

// ================= TIPO BASICO IDENT =================

func VerificarTipoBasicoIdent(
	tabela *TabelaDeSimbolos,
	ctx parser.ITipo_basico_identContext,
) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	if ctx.Tipo_basico() != nil {
		return VerificarTipoBasico(
			ctx.Tipo_basico(),
		)
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

// ================= IDENTIFICADOR =================

func VerificarIdentificador(
	tabela *TabelaDeSimbolos,
	ctx parser.IIdentificadorContext,
) TipoJander {

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

// ================= PARCELA UNARIO =================

func VerificarParcelaUnario(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcela_unarioContext,
) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	// identificador
	if ctx.Identificador() != nil {
		return VerificarIdentificador(
			tabela,
			ctx.Identificador(),
		)
	}

	// inteiro
	if ctx.NUM_INT() != nil {
		return INTEIRO
	}

	// real
	if ctx.NUM_REAL() != nil {
		return REAL
	}

	// chamada de funcao
	if ctx.IDENT() != nil &&
		ctx.ABREPAR() != nil {

		nomeFunc := ctx.IDENT().GetText()

		if !tabela.Existe(nomeFunc) {

			if !ErroJaReportado(nomeFunc) {

				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"identificador "+nomeFunc+" nao declarado",
				)
			}

			return INVALIDO
		}

		tipoFunc := tabela.Verificar(nomeFunc)

		if tipoFunc != FUNCAO &&
			tipoFunc != PROCEDIMENTO {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"identificador "+nomeFunc+
					" nao e uma funcao ou procedimento",
			)

			return INVALIDO
		}

		parametrosEsperados :=
			tabela.ObterParametros(nomeFunc)

		argumentos :=
			ctx.AllExpressao()

		if len(parametrosEsperados) != len(argumentos) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"incompatibilidade de parametros na chamada de "+nomeFunc,
			)

		} else {

			for i := 0; i < len(parametrosEsperados); i++ {

				tipoEsperado :=
					parametrosEsperados[i]

				tipoRecebido :=
					VerificarExpressao(
						tabela,
						argumentos[i],
					)

				if !CompatibilidadeFuncao(
					tipoEsperado,
					tipoRecebido,
				) {

					AdicionarErroSemantico(
						ctx.GetStart().GetLine(),
						"incompatibilidade de parametros na chamada de "+nomeFunc,
					)

					break
				}
			}
		}

		return tabela.ObterTipoRetorno(nomeFunc)
	}

	// expressao entre parenteses
	if len(ctx.AllExpressao()) > 0 {

		tipo := INVALIDO

		for _, exp := range ctx.AllExpressao() {

			aux := VerificarExpressao(
				tabela,
				exp,
			)

			if tipo == INVALIDO {
				tipo = aux
			} else if !Compatibilidade(tipo, aux) {

				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"expressao incompativel",
				)

				return INVALIDO
			}
		}

		return tipo
	}

	return INVALIDO
}

// ================= PARCELA NAO UNARIO =================

func VerificarParcelaNaoUnario(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcela_nao_unarioContext,
) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	// endereco
	if ctx.ENDERECO() != nil &&
		ctx.Identificador() != nil {

		nome := ctx.Identificador().GetText()

		if !tabela.Existe(nome) {

			if !ErroJaReportado(nome) {

				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"identificador "+nome+" nao declarado",
				)
			}

			return INVALIDO
		}

		return ENDERECO
	}

	// cadeia
	if ctx.CADEIA() != nil {
		return LITERAL
	}

	return INVALIDO
}

// ================= PARCELA =================

func VerificarParcela(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcelaContext,
) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	if ctx.Parcela_unario() != nil {
		return VerificarParcelaUnario(
			tabela,
			ctx.Parcela_unario(),
		)
	}

	if ctx.Parcela_nao_unario() != nil {
		return VerificarParcelaNaoUnario(
			tabela,
			ctx.Parcela_nao_unario(),
		)
	}

	return INVALIDO
}

// ================= FATOR =================

func VerificarFator(
	tabela *TabelaDeSimbolos,
	ctx parser.IFatorContext,
) TipoJander {

	tipo := INVALIDO

	for _, parcela := range ctx.AllParcela() {

		aux := VerificarParcela(
			tabela,
			parcela,
		)

		if tipo == INVALIDO {
			tipo = aux
		} else if !Compatibilidade(tipo, aux) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"fator contem tipos incompativeis",
			)

			return INVALIDO
		}
	}

	return tipo
}

// ================= TERMO =================

func VerificarTermo(
	tabela *TabelaDeSimbolos,
	ctx parser.ITermoContext,
) TipoJander {

	tipo := INVALIDO

	for _, fator := range ctx.AllFator() {

		aux := VerificarFator(
			tabela,
			fator,
		)

		if tipo == INVALIDO {
			tipo = aux
		} else if !Compatibilidade(tipo, aux) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"termo contem tipos incompativeis",
			)

			return INVALIDO
		}
	}

	return tipo
}

// ================= EXPRESSAO ARITMETICA =================

func VerificarExpAritmetica(
	tabela *TabelaDeSimbolos,
	ctx parser.IExp_aritmeticaContext,
) TipoJander {

	tipo := INVALIDO

	for _, termo := range ctx.AllTermo() {

		aux := VerificarTermo(
			tabela,
			termo,
		)

		if tipo == INVALIDO {
			tipo = aux
		} else if !Compatibilidade(tipo, aux) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"expressao aritmetica contem tipos incompativeis",
			)

			return INVALIDO
		}
	}

	return tipo
}

// ================= EXPRESSAO RELACIONAL =================

func VerificarExpRelacional(
	tabela *TabelaDeSimbolos,
	ctx parser.IExp_relacionalContext,
) TipoJander {

	tipo1 := VerificarExpAritmetica(
		tabela,
		ctx.Exp_aritmetica(0),
	)

	if ctx.Op_relacional() != nil {

		tipo2 := VerificarExpAritmetica(
			tabela,
			ctx.Exp_aritmetica(1),
		)

		if !Compatibilidade(tipo1, tipo2) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"operacao relacional com tipos incompativeis",
			)
		}

		return LOGICO
	}

	return tipo1
}

// ================= EXPRESSAO =================

func VerificarExpressao(
	tabela *TabelaDeSimbolos,
	ctx parser.IExpressaoContext,
) TipoJander {

	if ctx == nil {
		return INVALIDO
	}

	tipo := INVALIDO

	for _, termo := range ctx.AllTermo_logico() {

		aux := VerificarTermoLogico(
			tabela,
			termo,
		)

		if tipo == INVALIDO {
			tipo = aux
		} else if !Compatibilidade(tipo, aux) {

			return INVALIDO
		}
	}

	return tipo
}

// ================= TERMO LOGICO =================

func VerificarTermoLogico(
	tabela *TabelaDeSimbolos,
	ctx parser.ITermo_logicoContext,
) TipoJander {

	tipo := INVALIDO

	for _, fator := range ctx.AllFator_logico() {

		aux := VerificarFatorLogico(
			tabela,
			fator,
		)

		if tipo == INVALIDO {
			tipo = aux
		} else if !Compatibilidade(tipo, aux) {
			return INVALIDO
		}
	}

	return tipo
}

// ================= FATOR LOGICO =================

func VerificarFatorLogico(
	tabela *TabelaDeSimbolos,
	ctx parser.IFator_logicoContext,
) TipoJander {

	if ctx.Parcela_logica() != nil {
		return VerificarParcelaLogica(
			tabela,
			ctx.Parcela_logica(),
		)
	}

	return INVALIDO
}

// ================= PARCELA LOGICA =================

func VerificarParcelaLogica(
	tabela *TabelaDeSimbolos,
	ctx parser.IParcela_logicaContext,
) TipoJander {

	if ctx.Exp_relacional() != nil {
		return VerificarExpRelacional(
			tabela,
			ctx.Exp_relacional(),
		)
	}

	if ctx.VERDADEIRO() != nil ||
		ctx.FALSO() != nil {

		return LOGICO
	}

	return INVALIDO
}
