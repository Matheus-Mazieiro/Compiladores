package main

import "strings"

type TipoJander int

const (
	INTEIRO TipoJander = iota
	REAL
	LITERAL
	LOGICO
	PONTEIRO_INTEIRO
	PONTEIRO_REAL
	PONTEIRO_LITERAL
	PONTEIRO_LOGICO
	REGISTRO
	REGISTRO_TIPO
	ENDERECO
	FUNCAO
	PROCEDIMENTO
	INVALIDO
)

type EntradaTabela struct {
	Nome           string
	Tipo           TipoJander
	Parametros     []TipoJander
	TipoRetorno    TipoJander
	CamposRegistro map[string]TipoJander
	Dimensoes      []int
	ValorConstante interface{}
}

type TabelaDeSimbolos struct {
	escopos          []map[string]EntradaTabela
	escoposFuncao    []bool
	tipoRetornoAtual TipoJander
}

func NewTabelaDeSimbolos() *TabelaDeSimbolos {
	t := &TabelaDeSimbolos{}
	t.NovoEscopo(false)
	return t
}

func (t *TabelaDeSimbolos) NovoEscopo(isFuncao bool) {
	t.escopos = append(t.escopos, make(map[string]EntradaTabela))
	t.escoposFuncao = append(t.escoposFuncao, isFuncao)
}

func (t *TabelaDeSimbolos) AbandonarEscopo() {
	if len(t.escopos) > 0 {
		t.escopos = t.escopos[:len(t.escopos)-1]
	}

	if len(t.escoposFuncao) > 0 {
		t.escoposFuncao = t.escoposFuncao[:len(t.escoposFuncao)-1]
	}
}

func (t *TabelaDeSimbolos) EstaEmFuncao() bool {
	if len(t.escoposFuncao) == 0 {
		return false
	}

	return t.escoposFuncao[len(t.escoposFuncao)-1]
}

func (t *TabelaDeSimbolos) SetTipoRetornoFuncaoAtual(tipo TipoJander) {
	t.tipoRetornoAtual = tipo
}

func (t *TabelaDeSimbolos) ObterTipoRetornoFuncaoAtual() TipoJander {
	return t.tipoRetornoAtual
}

func (t *TabelaDeSimbolos) Adicionar(nome string, tipo TipoJander) {

	escopoAtual := t.escopos[len(t.escopos)-1]

	entrada := EntradaTabela{
		Nome: nome,
		Tipo: tipo,
	}

	if tipo == REGISTRO || tipo == REGISTRO_TIPO {
		entrada.CamposRegistro = make(map[string]TipoJander)
	}

	escopoAtual[nome] = entrada
}

func (t *TabelaDeSimbolos) AdicionarFuncao(nome string, tipoRetorno TipoJander, parametros []TipoJander) {

	escopoAtual := t.escopos[len(t.escopos)-1]

	tipo := FUNCAO

	if tipoRetorno == INVALIDO {
		tipo = PROCEDIMENTO
	}

	escopoAtual[nome] = EntradaTabela{
		Nome:        nome,
		Tipo:        tipo,
		Parametros:  parametros,
		TipoRetorno: tipoRetorno,
	}
}

func (t *TabelaDeSimbolos) Existe(nomeCompleto string) bool {

	nomeBase := nomeCompleto

	if strings.Contains(nomeCompleto, ".") {
		partes := strings.Split(nomeCompleto, ".")
		nomeBase = partes[0]
	}

	if strings.Contains(nomeCompleto, "[") {
		nomeBase = nomeCompleto[:strings.Index(nomeCompleto, "[")]
	}

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeBase]; ok {

			if strings.Contains(nomeCompleto, ".") {

				partes := strings.Split(nomeCompleto, ".")

				if len(partes) < 2 {
					return false
				}

				campo := partes[1]

				if entrada.CamposRegistro != nil {
					_, ok := entrada.CamposRegistro[campo]
					return ok
				}

				return false
			}

			return true
		}
	}

	return false
}

func (t *TabelaDeSimbolos) Verificar(nomeCompleto string) TipoJander {

	nomeBase := nomeCompleto

	if strings.Contains(nomeCompleto, ".") {
		partes := strings.Split(nomeCompleto, ".")
		nomeBase = partes[0]
	}

	if strings.Contains(nomeCompleto, "[") {
		nomeBase = nomeCompleto[:strings.Index(nomeCompleto, "[")]
	}

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeBase]; ok {

			if strings.Contains(nomeCompleto, ".") {

				partes := strings.Split(nomeCompleto, ".")

				if len(partes) < 2 {
					return INVALIDO
				}

				campo := partes[1]

				if entrada.CamposRegistro != nil {
					if tipoCampo, ok := entrada.CamposRegistro[campo]; ok {
						return tipoCampo
					}
				}

				return INVALIDO
			}

			return entrada.Tipo
		}
	}

	return INVALIDO
}

func (t *TabelaDeSimbolos) ObterParametros(nome string) []TipoJander {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return entrada.Parametros
		}
	}

	return []TipoJander{}
}

func (t *TabelaDeSimbolos) ObterTipoRetorno(nome string) TipoJander {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return entrada.TipoRetorno
		}
	}

	return INVALIDO
}
func (t *TabelaDeSimbolos) AdicionarConstante(
	nome string,
	tipo TipoJander,
	valor interface{},
) {

	escopoAtual := t.escopos[len(t.escopos)-1]

	escopoAtual[nome] = EntradaTabela{
		Nome:           nome,
		Tipo:           tipo,
		ValorConstante: valor,
	}
}

func (t *TabelaDeSimbolos) AdicionarArray(
	nome string,
	tipo TipoJander,
	dimensoes []int,
) {

	escopoAtual := t.escopos[len(t.escopos)-1]

	escopoAtual[nome] = EntradaTabela{
		Nome:      nome,
		Tipo:      tipo,
		Dimensoes: dimensoes,
	}
}

func (t *TabelaDeSimbolos) AdicionarCampoRegistro(
	nomeRegistro string,
	nomeCampo string,
	tipoCampo TipoJander,
) {

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeRegistro]; ok {

			if entrada.CamposRegistro == nil {
				entrada.CamposRegistro = make(map[string]TipoJander)
			}

			entrada.CamposRegistro[nomeCampo] = tipoCampo

			t.escopos[i][nomeRegistro] = entrada

			return
		}
	}
}
