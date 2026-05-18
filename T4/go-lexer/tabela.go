package main

import "strings"

type TipoJander int

// Tipos de dados suportados
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

// EntradaTabela representa uma entrada na tabela de símbolos, contendo informações sobre variáveis, funções, procedimentos, registros e constantes.
type EntradaTabela struct {
	Nome           string
	Tipo           TipoJander
	Parametros     []TipoJander
	TipoRetorno    TipoJander
	CamposRegistro map[string]EntradaTabela
	Dimensoes      []int
	ValorConstante interface{}
}

// TabelaDeSimbolos gerencia os escopos e as entradas da tabela de símbolos, permitindo a adição, verificação e obtenção de informações sobre variáveis, funções, procedimentos, registros e constantes.
type TabelaDeSimbolos struct {
	escopos          []map[string]EntradaTabela
	escoposFuncao    []bool
	tipoRetornoAtual TipoJander
}

// NewTabelaDeSimbolos cria e inicializa uma nova tabela de símbolos, começando com um escopo global vazio.
func NewTabelaDeSimbolos() *TabelaDeSimbolos {
	t := &TabelaDeSimbolos{}
	t.NovoEscopo(false)
	return t
}

// NovoEscopo adiciona um novo escopo à tabela de símbolos, indicando se o escopo é de uma função ou não.
func (t *TabelaDeSimbolos) NovoEscopo(isFuncao bool) {
	t.escopos = append(t.escopos, make(map[string]EntradaTabela))
	t.escoposFuncao = append(t.escoposFuncao, isFuncao)
}

// AbandonarEscopo remove o escopo mais recente da tabela de símbolos, retornando ao escopo anterior.
func (t *TabelaDeSimbolos) AbandonarEscopo() {
	if len(t.escopos) > 0 {
		t.escopos = t.escopos[:len(t.escopos)-1]
	}

	if len(t.escoposFuncao) > 0 {
		t.escoposFuncao = t.escoposFuncao[:len(t.escoposFuncao)-1]
	}
}

// EstaEmFuncao verifica se o escopo atual é de uma função, retornando true se for o caso e false caso contrário.
func (t *TabelaDeSimbolos) EstaEmFuncao() bool {
	if len(t.escoposFuncao) == 0 {
		return false
	}

	return t.escoposFuncao[len(t.escoposFuncao)-1]
}

// SetTipoRetornoFuncaoAtual define o tipo de retorno da função atual, permitindo que seja verificado posteriormente durante a análise semântica.
func (t *TabelaDeSimbolos) SetTipoRetornoFuncaoAtual(tipo TipoJander) {
	t.tipoRetornoAtual = tipo
}

// ObterTipoRetornoFuncaoAtual retorna o tipo de retorno da função atual, que pode ser utilizado para validar as instruções de retorno dentro da função.
func (t *TabelaDeSimbolos) ObterTipoRetornoFuncaoAtual() TipoJander {
	return t.tipoRetornoAtual
}

// Adicionar insere uma nova entrada na tabela de símbolos para uma variável, registrando seu nome e tipo no escopo atual.
func (t *TabelaDeSimbolos) Adicionar(nome string, tipo TipoJander) {

	escopoAtual := t.escopos[len(t.escopos)-1]

	entrada := EntradaTabela{
		Nome: nome,
		Tipo: tipo,
	}

	entrada.CamposRegistro = make(map[string]EntradaTabela)
	escopoAtual[nome] = entrada
}

// AdicionarFuncao insere uma nova entrada na tabela de símbolos para uma função ou procedimento, registrando seu nome, tipo de retorno e tipos dos parâmetros no escopo atual.
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

// Verificar recebe um nome completo (que pode incluir campos de registros e índices de arrays) e retorna o tipo associado a esse nome na tabela de símbolos, ou INVALIDO se o nome não for encontrado ou for inválido.
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

// Existe verifica se um nome completo (que pode incluir campos de registros e índices de arrays) existe na tabela de símbolos, retornando true se o nome for encontrado e false caso contrário.
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

// ObterParametros retorna os tipos dos parâmetros de uma função ou procedimento com o nome especificado, ou um slice vazio se a função ou procedimento não for encontrado ou não tiver parâmetros.
func (t *TabelaDeSimbolos) ObterParametros(nome string) []TipoJander {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return entrada.Parametros
		}
	}

	return []TipoJander{}
}

// ObterTipoRetorno retorna o tipo de retorno de uma função com o nome especificado, ou INVALIDO se a função não for encontrada ou não tiver um tipo de retorno válido.
func (t *TabelaDeSimbolos) ObterTipoRetorno(nome string) TipoJander {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return entrada.TipoRetorno
		}
	}

	return INVALIDO
}

// ObterEntrada retorna a entrada completa da tabela de símbolos para um nome especificado, incluindo informações como tipo, parâmetros, tipo de retorno e campos de registro, ou um valor zero e false se o nome não for encontrado.
func (t *TabelaDeSimbolos) ObterEntrada(nome string) (EntradaTabela, bool) {

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nome]; ok {
			return entrada, true
		}
	}

	return EntradaTabela{}, false
}

// AdicionarConstante insere uma nova entrada na tabela de símbolos para uma constante, registrando seu nome, tipo e valor no escopo atual.
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

// AdicionarArray insere uma nova entrada na tabela de símbolos para um array, registrando seu nome, tipo e dimensões no escopo atual.
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

// AdicionarCampoRegistro adiciona um campo a um registro existente na tabela de símbolos, associando o nome do campo ao tipo e outras informações contidas na entradaCampo.
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

// ExisteNoEscopoAtual verifica se um nome existe no escopo mais recente da tabela de símbolos, retornando true se o nome for encontrado e false caso contrário.
func (t *TabelaDeSimbolos) ExisteNoEscopoAtual(nome string) bool {
	if len(t.escopos) == 0 {
		return false
	}
	_, existe := t.escopos[len(t.escopos)-1][nome]
	return existe
}

// ObterQtdParametros retorna a quantidade de parâmetros de uma função ou procedimento com o nome especificado, ou 0 se a função ou procedimento não for encontrado ou não tiver parâmetros.
func (t *TabelaDeSimbolos) ObterQtdParametros(nome string) int {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return len(entrada.Parametros)
		}
	}
	return 0
}

// ObterTipoParametroIdx retorna o tipo do parâmetro em uma posição específica de uma função ou procedimento com o nome especificado, ou INVALIDO se a função ou procedimento não for encontrado, não tiver parâmetros ou o índice for inválido.
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
