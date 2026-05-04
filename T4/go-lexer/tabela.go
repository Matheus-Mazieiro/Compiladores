// LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
package main

type TipoJander int

const (
	INTEIRO TipoJander = iota
	REAL
	LITERAL
	LOGICO
	INVALIDO
)

type EntradaTabela struct {
	Nome string
	Tipo TipoJander
}

type TabelaDeSimbolos struct {
	tabela map[string]EntradaTabela
}

func NewTabelaDeSimbolos() *TabelaDeSimbolos {
	return &TabelaDeSimbolos{
		tabela: make(map[string]EntradaTabela),
	}
}

func (t *TabelaDeSimbolos) Adicionar(nome string, tipo TipoJander) {
	t.tabela[nome] = EntradaTabela{nome, tipo}
}

func (t *TabelaDeSimbolos) Existe(nome string) bool {
	_, ok := t.tabela[nome]
	return ok
}

func (t *TabelaDeSimbolos) Verificar(nome string) TipoJander {
	if val, ok := t.tabela[nome]; ok {
		return val.Tipo
	}
	return INVALIDO
}