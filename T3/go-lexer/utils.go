// LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
package main

import (
	"fmt"
	parser "go-lexer/parser"
)

var ErrosSemanticos []string
var variaveisComErro = make(map[string]bool)
var variaveisComTipoInvalido = make(map[string]bool)

func AdicionarErroSemantico(linha int, msg string) {
	erro := fmt.Sprintf("Linha %d: %s", linha, msg)
	fmt.Println("DEBUG ERRO:", erro)
	ErrosSemanticos = append(ErrosSemanticos, erro)
}

func ErroJaReportado(nome string) bool {
	if variaveisComErro[nome] {
		return true
	}
	variaveisComErro[nome] = true
	return false
}

// ================= TIPOS =================

func VerificarTipoBasico(ctx parser.ITipo_basicoContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	texto := ctx.GetText()

	switch texto {
	case "literal":
		return LITERAL
	case "inteiro":
		return INTEIRO
	case "real":
		return REAL
	case "logico":
		return LOGICO   
	default:
		return INVALIDO // Tipo inválido
	}
}

func VerificarTipo(tabela *TabelaDeSimbolos, ctx parser.ITipoContext) TipoJander {
	if ctx.Tipo_estendido() != nil {
		return VerificarTipoBasicoIdent(tabela, ctx.Tipo_estendido().Tipo_basico_ident())
	}
	return INVALIDO
}

func VerificarTipoBasicoIdent(tabela *TabelaDeSimbolos, ctx parser.ITipo_basico_identContext) TipoJander {
	if ctx.Tipo_basico() != nil {
		return VerificarTipoBasico(ctx.Tipo_basico())
	}

	if ctx.IDENT() != nil {
		nome := ctx.IDENT().GetText()

		if !tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(ctx.GetStart().GetLine(),
					"tipo "+nome+" nao declarado")
			}
			return INVALIDO
		}

		return tabela.Verificar(nome)
	}

	return INVALIDO
}