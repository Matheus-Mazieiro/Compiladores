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
	Nome        string
	Tipo        TipoJander
	Parametros  []TipoJander
	TipoRetorno TipoJander
	//CamposRegistro map[string]TipoJander
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

	//nome = normalizar(nome)

	entrada := EntradaTabela{
		Nome: nome,
		Tipo: tipo,
	}

	//entrada.CamposRegistro = make(map[string]TipoJander)
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

// func (t *TabelaDeSimbolos) Existe(nomeCompleto string) bool {

// 	println("====== EXISTE ======")
// 	println("BUSCANDO:", nomeCompleto)

// 	nome := strings.ReplaceAll(nomeCompleto, "^", "")

// 	// remove dimensões
// 	if idx := strings.Index(nome, "["); idx != -1 {
// 		nome = nome[:idx]
// 	}

// 	println("NOME LIMPO:", nome)

// 	partes := strings.Split(nome, ".")

// 	nomeBase := partes[0]

// 	println("NOME BASE:", nomeBase)

// 	for i := len(t.escopos) - 1; i >= 0; i-- {

// 		println("ESCOPO:", i)

// 		for k, entrada := range t.escopos[i] {

// 			println("  TESTANDO:", k)

// 			if k == nomeBase {

// 				println("  ENCONTROU BASE:", k)

// 				// identificador simples
// 				if len(partes) == 1 {

// 					println("  EXISTE SIMPLES OK")
// 					return true
// 				}

// 				// acesso campo registro
// 				campo := partes[1]

// 				println("  BUSCANDO CAMPO:", campo)

// 				if entrada.CamposRegistro == nil {

// 					println("  CAMPOS NIL")
// 					return false
// 				}

// 				println("  CAMPOS DISPONIVEIS:")

// 				for c := range entrada.CamposRegistro {
// 					println("   ->", c)
// 				}

// 				_, ok := entrada.CamposRegistro[campo]

// 				if ok {
// 					println("  CAMPO ENCONTRADO")
// 				} else {
// 					println("  CAMPO NAO ENCONTRADO")
// 				}
// 				println("CAMPO BUSCADO:", campo)

// 				for c := range entrada.CamposRegistro {
// 					println("CAMPO REAL:", c)
// 				}
// 				return ok
// 			}
// 		}
// 	}

// 	println("NAO ENCONTROU:", nomeCompleto)
// 	return false
// }
func (t *TabelaDeSimbolos) Verificar(nomeCompleto string) TipoJander {

	nome := strings.ReplaceAll(nomeCompleto, "^", "")

	if idx := strings.Index(nome, "["); idx != -1 {
		nome = nome[:idx]
	}

	partes := strings.Split(nome, ".")

	if len(partes) == 0 {
		return INVALIDO
	}

	// procura variável base
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

	// variável simples
	if len(partes) == 1 {
		return entradaAtual.Tipo
	}

	// percorre campos aninhados
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

// func (t *TabelaDeSimbolos) Verificar(nomeCompleto string) TipoJander {

// 	println("====== VERIFICAR ======")
// 	println("VERIFICANDO:", nomeCompleto)

// 	nome := strings.ReplaceAll(nomeCompleto, "^", "")

// 	if idx := strings.Index(nome, "["); idx != -1 {
// 		nome = nome[:idx]
// 	}

// 	println("NOME LIMPO:", nome)

// 	partes := strings.Split(nome, ".")

// 	nomeBase := partes[0]

// 	println("NOME BASE:", nomeBase)

// 	for i := len(t.escopos) - 1; i >= 0; i-- {

// 		println("ESCOPO:", i)

// 		for k, entrada := range t.escopos[i] {

// 			println("  TESTANDO:", k)

// 			if k == nomeBase {

// 				println("  ENCONTROU BASE:", k)

// 				// variável simples
// 				if len(partes) == 1 {

// 					println("  RETORNANDO TIPO:", entrada.Tipo)

// 					return entrada.Tipo
// 				}

// 				// campo registro
// 				campo := partes[1]

// 				println("  BUSCANDO CAMPO:", campo)

// 				if entrada.CamposRegistro == nil {

// 					println("  CAMPOS NIL")
// 					return INVALIDO
// 				}

// 				entradaCampo, ok := entrada.CamposRegistro[campo]

// 				if !ok {

// 					println("  CAMPO NAO ENCONTRADO")
// 					return INVALIDO
// 				}

// 				println("  CAMPO ENCONTRADO TIPO:", entradaCampo.Tipo)

// 				return entradaCampo.Tipo
// 			}
// 		}
// 	}

// 	println("TIPO INVALIDO:", nomeCompleto)

// 	return INVALIDO
// }
func (t *TabelaDeSimbolos) DebugTabela() {
	println("======= TABELA =======")

	for i, escopo := range t.escopos {

		println("ESCOPO", i)

		for nome, entrada := range escopo {

			println("  Nome:", nome)
			println("  Tipo:", entrada.Tipo)

			if entrada.CamposRegistro != nil {
			}

			for c, campoEntrada := range entrada.CamposRegistro {
				println("   ->", c, "Tipo:", campoEntrada.Tipo)

			}
		}
	}

	println("======================")
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

	//nomeRegistro = normalizar(nomeRegistro)
	//campo = normalizar(campo)

	for i := len(t.escopos) - 1; i >= 0; i-- {

		if entrada, ok := t.escopos[i][nomeRegistro]; ok {

			if entrada.CamposRegistro == nil {
				entrada.CamposRegistro = make(map[string]EntradaTabela)
			}
			println(
				"ADICIONANDO CAMPO:",
				nomeRegistro,
				".",
				campo,
			)

			//entrada.CamposRegistro[campo] = tipo
			entrada.CamposRegistro[campo] = entradaCampo
			println("CAMPOS AGORA:")

			for c := range entrada.CamposRegistro {
				println(" ->", c)
			}
			entrada.CamposRegistro[campo] = entradaCampo
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
