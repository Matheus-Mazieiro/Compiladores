// // LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
// package main

// type TipoJander int

// const (
// 	INTEIRO TipoJander = iota
// 	REAL
// 	LITERAL
// 	LOGICO
// 	INVALIDO
// )

// type EntradaTabela struct {
// 	Nome string
// 	Tipo TipoJander
// }

// type TabelaDeSimbolos struct {
// 	tabela map[string]EntradaTabela
// }

// func NewTabelaDeSimbolos() *TabelaDeSimbolos {
// 	return &TabelaDeSimbolos{
// 		tabela: make(map[string]EntradaTabela),
// 	}
// }

// func (t *TabelaDeSimbolos) Adicionar(nome string, tipo TipoJander) {
// 	t.tabela[nome] = EntradaTabela{nome, tipo}
// }

// func (t *TabelaDeSimbolos) Existe(nome string) bool {
// 	_, ok := t.tabela[nome]
// 	return ok
// }

// func (t *TabelaDeSimbolos) Verificar(nome string) TipoJander {
// 	if val, ok := t.tabela[nome]; ok {
// 		return val.Tipo
// 	}
// 	return INVALIDO
// }
// LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
package main

type TipoJander int

const (
	INTEIRO TipoJander = iota
	REAL
	LITERAL
	LOGICO
	PONTEIRO
	REGISTRO
	REGISTRO_TIPO
	ENDERECO
	FUNCAO
	PROCEDIMENTO
	INVALIDO
)

type EntradaTabela struct {
	Nome string
	Tipo TipoJander

	// Funções/procedimentos
	Parametros  []TipoJander
	TipoRetorno TipoJander

	// Registros
	CamposRegistro map[string]TipoJander

	// Arrays
	Dimensoes []int

	// Constantes
	ValorConstante interface{}
}

type TabelaDeSimbolos struct {
	escopos []map[string]EntradaTabela

	// Controle de função
	escoposFuncao    []bool
	tipoRetornoAtual TipoJander
}

func NewTabelaDeSimbolos() *TabelaDeSimbolos {
	t := &TabelaDeSimbolos{
		escopos:           []map[string]EntradaTabela{},
		escoposFuncao:     []bool{},
		tipoRetornoAtual:  INVALIDO,
	}

	t.NovoEscopo(false)

	return t
}

// ================= ESCOPOS =================

func (t *TabelaDeSimbolos) NovoEscopo(isFuncao bool) {
	t.escopos = append(t.escopos, make(map[string]EntradaTabela))
	t.escoposFuncao = append(t.escoposFuncao, isFuncao)

	if isFuncao {
		t.tipoRetornoAtual = INVALIDO
	}
}

func (t *TabelaDeSimbolos) AbandonarEscopo() {
	if len(t.escopos) > 0 {
		t.escopos = t.escopos[:len(t.escopos)-1]
	}

	if len(t.escoposFuncao) > 0 {
		t.escoposFuncao = t.escoposFuncao[:len(t.escoposFuncao)-1]
	}

	if len(t.escoposFuncao) == 0 ||
		!t.escoposFuncao[len(t.escoposFuncao)-1] {
		t.tipoRetornoAtual = INVALIDO
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

// ================= ADIÇÃO =================

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

func (t *TabelaDeSimbolos) AdicionarConstante(
	nome string,
	tipo TipoJander,
	valor interface{},
) {
	escopoAtual := t.escopos[len(t.escopos)-1]

	escopoAtual[nome] = EntradaTabela{
		Nome:            nome,
		Tipo:            tipo,
		ValorConstante:  valor,
	}
}

func (t *TabelaDeSimbolos) AdicionarArray(
	nome string,
	tipo TipoJander,
	dimensoes []int,
) {
	escopoAtual := t.escopos[len(t.escopos)-1]

	escopoAtual[nome] = EntradaTabela{
		Nome:       nome,
		Tipo:       tipo,
		Dimensoes:  dimensoes,
	}
}

func (t *TabelaDeSimbolos) AdicionarFuncao(
	nome string,
	tipoRetorno TipoJander,
	parametros []TipoJander,
) {
	escopoAtual := t.escopos[len(t.escopos)-1]

	tipo := FUNCAO
	if tipoRetorno == INVALIDO {
		tipo = PROCEDIMENTO
	}

	escopoAtual[nome] = EntradaTabela{
		Nome:         nome,
		Tipo:         tipo,
		Parametros:   parametros,
		TipoRetorno:  tipoRetorno,
	}
}

// ================= REGISTROS =================

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

// ================= CONSULTAS =================

func (t *TabelaDeSimbolos) Existe(nomeCompleto string) bool {
	nomeBase := nomeCompleto

	// registro.campo
	if contains(nomeCompleto, ".") {
		nomeBase = split(nomeCompleto, ".")[0]
	}

	// vetor[1]
	if contains(nomeCompleto, "[") {
		nomeBase = nomeCompleto[:index(nomeCompleto, "[")]
	}

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeBase]; ok {

			// Campo de registro
			if contains(nomeCompleto, ".") {
				partes := split(nomeCompleto, ".")
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

			// Array
			if contains(nomeCompleto, "[") {
				return entrada.Dimensoes != nil
			}

			return true
		}
	}

	return false
}

func (t *TabelaDeSimbolos) Verificar(nomeCompleto string) TipoJander {
	nomeBase := nomeCompleto

	if contains(nomeCompleto, ".") {
		nomeBase = split(nomeCompleto, ".")[0]
	}

	if contains(nomeCompleto, "[") {
		nomeBase = nomeCompleto[:index(nomeCompleto, "[")]
	}

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeBase]; ok {

			// Campo de registro
			if contains(nomeCompleto, ".") {
				partes := split(nomeCompleto, ".")
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

			// Vetor
			if contains(nomeCompleto, "[") {
				if entrada.Dimensoes != nil {
					return entrada.Tipo
				}

				return INVALIDO
			}

			return entrada.Tipo
		}
	}

	return INVALIDO
}

func (t *TabelaDeSimbolos) ObterEntrada(nome string) *EntradaTabela {
	for i := len(t.escopos) - 1; i >= 0; i-- {
		if entrada, ok := t.escopos[i][nome]; ok {
			return &entrada
		}
	}

	return nil
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

// ================= HELPERS =================

// Essas helpers evitam importar strings em vários lugares.
// Se preferir, pode substituir por strings.Contains/Split/Index.

func contains(s, sub string) bool {
	return index(s, sub) >= 0
}

func index(s, sub string) int {
	n := len(sub)

	for i := 0; i+n <= len(s); i++ {
		if s[i:i+n] == sub {
			return i
		}
	}

	return -1
}

func split(s, sep string) []string {
	idx := index(s, sep)

	if idx < 0 {
		return []string{s}
	}

	return []string{
		s[:idx],
		s[idx+len(sep):],
	}
}