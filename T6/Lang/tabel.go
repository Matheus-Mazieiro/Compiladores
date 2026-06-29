package main

type TipoSimbolo int

const (
	SET TipoSimbolo = iota
	PARAM
	VAR_BINARY
	VAR_INTEGER
	VAR_CONTINUOUS
	ITERATOR
	INVALIDO
)

type EntradaTabela struct {
	// Nome do símbolo
	Nome string

	// Tipo (SET, PARAM, VAR...)
	Tipo TipoSimbolo

	// Quantidade de índices
	//
	// Exemplo:
	//   x         -> 0
	//   x[i]      -> 1
	//   x[i,j]    -> 2
	Indices int

	// Domínios declarados
	//
	// Exemplo:
	//   x[i in I, j in J]
	//
	// Dominios = ["I", "J"]
	//
	// Para conjuntos simples permanece vazio.
	Dominios []string
}

type TabelaDeSimbolos struct {
	escopos []map[string]EntradaTabela
}

func NewTabela() *TabelaDeSimbolos {
	t := &TabelaDeSimbolos{}
	t.NovoEscopo()
	return t
}

func (t *TabelaDeSimbolos) NovoEscopo() {
	t.escopos = append(t.escopos, make(map[string]EntradaTabela))
}

func (t *TabelaDeSimbolos) AbandonarEscopo() {
	if len(t.escopos) > 0 {
		t.escopos = t.escopos[:len(t.escopos)-1]
	}
}

func (t *TabelaDeSimbolos) escopoAtual() map[string]EntradaTabela {
	return t.escopos[len(t.escopos)-1]
}

func (t *TabelaDeSimbolos) Adicionar(
	nome string,
	tipo TipoSimbolo,
	indices int,
	dominios []string,
) bool {

	escopo := t.escopoAtual()

	if _, existe := escopo[nome]; existe {
		return false
	}

	escopo[nome] = EntradaTabela{
		Nome:     nome,
		Tipo:     tipo,
		Indices:  indices,
		Dominios: dominios,
	}

	return true
}

func (t *TabelaDeSimbolos) ExisteNoEscopoAtual(nome string) bool {
	_, ok := t.escopoAtual()[nome]
	return ok
}

func (t *TabelaDeSimbolos) Existe(nome string) bool {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if _, ok := t.escopos[i][nome]; ok {
			return true
		}
	}

	return false
}

func (t *TabelaDeSimbolos) ObterEntrada(nome string) (EntradaTabela, bool) {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return entrada, true
		}
	}

	return EntradaTabela{
		Tipo: INVALIDO,
	}, false
}

func (t *TabelaDeSimbolos) Verificar(nome string) TipoSimbolo {
	if entrada, ok := t.ObterEntrada(nome); ok {
		return entrada.Tipo
	}

	return INVALIDO
}

func (t *TabelaDeSimbolos) RemoverEscopoAtual() {
	t.AbandonarEscopo()
}
