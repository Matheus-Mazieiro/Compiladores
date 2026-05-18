package main

import "strings"

type TipoJander int

func normalizar(nome string) string {
	return nome
}

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
	CamposRegistro map[string]EntradaTabela
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

	entrada.CamposRegistro = make(map[string]EntradaTabela)
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

func (t *TabelaDeSimbolos) Verificar(nomeCompleto string) TipoJander {

	nome := strings.ReplaceAll(nomeCompleto, "^", "")

	if idx := strings.Index(nome, "["); idx != -1 {
		nome = nome[:idx]
	}

	partes := strings.Split(nome, ".")

	if len(partes) == 0 {
		return INVALIDO
	}

	var entradaAtual EntradaTabela
	encontrou := false

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][partes[0]]; ok {
			entradaAtual = entrada
			encontrou = true
			break
		}
	}

	if !encontrou {
		return INVALIDO
	}

	if len(partes) == 1 {
		return entradaAtual.Tipo
	}

	for i := 1; i < len(partes); i++ {

		if entradaAtual.CamposRegistro == nil {
			return INVALIDO
		}

		prox, ok := entradaAtual.CamposRegistro[partes[i]]

		if !ok {
			return INVALIDO
		}

		entradaAtual = prox
	}

	return entradaAtual.Tipo
}
func (t *TabelaDeSimbolos) Existe(nomeCompleto string) bool {

	nome := strings.ReplaceAll(nomeCompleto, "^", "")

	if idx := strings.Index(nome, "["); idx != -1 {
		nome = nome[:idx]
	}

	partes := strings.Split(nome, ".")

	if len(partes) == 0 {
		return false
	}

	var entradaAtual EntradaTabela
	encontrou := false

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][partes[0]]; ok {
			entradaAtual = entrada
			encontrou = true
			break
		}
	}

	if !encontrou {
		return false
	}

	if len(partes) == 1 {
		return true
	}

	for i := 1; i < len(partes); i++ {

		if entradaAtual.CamposRegistro == nil {
			return false
		}

		prox, ok := entradaAtual.CamposRegistro[partes[i]]

		if !ok {
			return false
		}

		entradaAtual = prox
	}

	return true
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
func (t *TabelaDeSimbolos) ObterEntrada(nome string) (EntradaTabela, bool) {

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nome]; ok {
			return entrada, true
		}
	}

	return EntradaTabela{}, false
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
	campo string,
	entradaCampo EntradaTabela,
) {

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeRegistro]; ok {

			if entrada.CamposRegistro == nil {
				entrada.CamposRegistro = make(map[string]EntradaTabela)
			}

			entrada.CamposRegistro[campo] = entradaCampo

			entrada.CamposRegistro[campo] = entradaCampo
			t.escopos[i][nomeRegistro] = entrada
			return
		}
	}
}

func (t *TabelaDeSimbolos) ExisteNoEscopoAtual(nome string) bool {
	if len(t.escopos) == 0 {
		return false
	}
	_, existe := t.escopos[len(t.escopos)-1][nome]
	return existe
}

func (t *TabelaDeSimbolos) ObterQtdParametros(nome string) int {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return len(entrada.Parametros)
		}
	}
	return 0
}

func (t *TabelaDeSimbolos) ObterTipoParametroIdx(nome string, idx int) TipoJander {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			if idx >= 0 && idx < len(entrada.Parametros) {
				return entrada.Parametros[idx]
			}
		}
	}
	return INVALIDO
}
