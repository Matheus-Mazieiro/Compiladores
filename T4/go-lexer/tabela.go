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

	//nome = normalizar(nome)

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
	nomeLimpo := nomeCompleto
	if strings.Contains(nomeCompleto, "[") {
		idxAbre := strings.Index(nomeCompleto, "[")
		idxFecha := strings.LastIndex(nomeCompleto, "]")
		if idxFecha > idxAbre {
			nomeLimpo = nomeCompleto[:idxAbre] + nomeCompleto[idxFecha+1:]
		}
	}

	nomeBase := nomeLimpo
	if strings.Contains(nomeLimpo, ".") {
		nomeBase = strings.Split(nomeLimpo, ".")[0]
	}

	// Força a busca da variável base em minúsculo para evitar o erro da linha 28
	nomeBaseBusca := strings.ToLower(nomeBase)

	for i := len(t.escopos) - 1; i >= 0; i-- {
		// Varre o escopo testando as chaves em LowerCase
		for k, entrada := range t.escopos[i] {
			if strings.ToLower(k) == nomeBaseBusca {

				if strings.Contains(nomeLimpo, ".") {
					partes := strings.Split(nomeLimpo, ".")
					if len(partes) < 2 {
						return false
					}
					campo := partes[1]

					if entrada.CamposRegistro != nil {
						_, existeCampo := entrada.CamposRegistro[campo]
						return existeCampo
					}
					return false
				}
				return true
			}
		}
	}
	return false
}

func (t *TabelaDeSimbolos) Verificar(nomeCompleto string) TipoJander {
	nomeLimpo := nomeCompleto
	if strings.Contains(nomeCompleto, "[") {
		idxAbre := strings.Index(nomeCompleto, "[")
		idxFecha := strings.LastIndex(nomeCompleto, "]")
		if idxFecha > idxAbre {
			nomeLimpo = nomeCompleto[:idxAbre] + nomeCompleto[idxFecha+1:]
		}
	}

	nomeBase := nomeLimpo
	if strings.Contains(nomeLimpo, ".") {
		nomeBase = strings.Split(nomeLimpo, ".")[0]
	}

	nomeBaseBusca := strings.ToLower(nomeBase)

	for i := len(t.escopos) - 1; i >= 0; i-- {
		for k, entrada := range t.escopos[i] {
			if strings.ToLower(k) == nomeBaseBusca {

				if strings.Contains(nomeLimpo, ".") {
					partes := strings.Split(nomeLimpo, ".")
					if len(partes) < 2 {
						return INVALIDO
					}
					campo := partes[1]

					if entrada.CamposRegistro != nil {
						if tipoCampo, existeCampo := entrada.CamposRegistro[campo]; existeCampo {
							return tipoCampo
						}
					}
					return INVALIDO
				}
				return entrada.Tipo
			}
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
	campo string,
	tipo TipoJander,
) {

	//nomeRegistro = normalizar(nomeRegistro)
	//campo = normalizar(campo)

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeRegistro]; ok {

			if entrada.CamposRegistro == nil {
				entrada.CamposRegistro = make(map[string]TipoJander)
			}

			entrada.CamposRegistro[campo] = tipo
			t.escopos[i][nomeRegistro] = entrada
			return
		}
	}
}

// Verifica se um identificador já existe estritamente no escopo atual (topo da pilha)
func (t *TabelaDeSimbolos) ExisteNoEscopoAtual(nome string) bool {
	if len(t.escopos) == 0 {
		return false
	}
	_, existe := t.escopos[len(t.escopos)-1][nome]
	return existe
}

// Retorna a quantidade de parâmetros formais que uma função possui
func (t *TabelaDeSimbolos) ObterQtdParametros(nome string) int {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return len(entrada.Parametros)
		}
	}
	return 0
}

// Retorna o tipo do parâmetro em uma posição específica (índice)
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
