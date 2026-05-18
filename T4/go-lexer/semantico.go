package main

import (
	parser "go-lexer/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// JanderSemantico é um visitante que percorre a árvore de análise sintática gerada pelo ANTLR para realizar a análise semântica do programa, verificando tipos, escopos e outras regras semânticas definidas na linguagem.
type JanderSemantico struct {
	parser.BaseCalcLexerVisitor
	tabela *TabelaDeSimbolos
}

// NewJanderSemantico cria e inicializa uma nova instância do visitante semântico, configurando a tabela de símbolos para ser utilizada durante a análise semântica do programa.
func NewJanderSemantico() *JanderSemantico {
	return &JanderSemantico{
		tabela: nil,
	}
}

// VisitPrograma é o método de visita para o nó raiz do programa, onde a tabela de símbolos é inicializada e as declarações e corpo do programa são visitados para realizar a análise semântica.
func (j *JanderSemantico) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	j.tabela = NewTabelaDeSimbolos()

	if ctx.Declaracoes() != nil {
		ctx.Declaracoes().Accept(j)
	}

	if ctx.Corpo() != nil {
		ctx.Corpo().Accept(j)
	}

	return nil
}

// visitdeclaracoes é o método de visita para o nó de declarações, onde cada declaração local ou global é visitada para realizar a análise semântica das declarações do programa.
func (j *JanderSemantico) VisitDeclaracoes(ctx *parser.DeclaracoesContext) interface{} {
	for _, d := range ctx.AllDecl_local_global() {
		d.Accept(j)
	}

	return nil
}

// VisitDecl_local_global é o método de visita para o nó de declaração local ou global, onde as declarações locais e globais são processadas para verificar a validade das declarações e atualizar a tabela de símbolos conforme necessário.
func (j *JanderSemantico) VisitDecl_local_global(ctx *parser.Decl_local_globalContext) interface{} {

	if ctx.Declaracao_local() != nil {
		ctx.Declaracao_local().Accept(j)
	}

	if ctx.Declaracao_global() != nil {
		ctx.Declaracao_global().Accept(j)
	}

	return nil
}

// visitcorpo é o método de visita para o nó de corpo do programa, onde as declarações locais e os comandos dentro do corpo são visitados para realizar a análise semântica das operações e estruturas de controle presentes no programa.
func (j *JanderSemantico) VisitCorpo(ctx *parser.CorpoContext) interface{} {

	for _, d := range ctx.AllDeclaracao_local() {
		d.Accept(j)
	}

	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	return nil
}

// visitdecrlaracao_local é o método de visita para o nó de declaração local, onde as variáveis, constantes e tipos locais são processados para verificar a validade das declarações e atualizar a tabela de símbolos conforme necessário.
func (j *JanderSemantico) VisitDeclaracao_local(ctx *parser.Declaracao_localContext) interface{} {

	if ctx.Variavel() != nil {
		ctx.Variavel().Accept(j)
		return nil
	}

	if ctx.CONSTANTE() != nil {

		nome := ctx.IDENT().GetText()
		tipo := VerificarTipoBasico(ctx.Tipo_basico())

		if j.tabela.Existe(nome) {
			AdicionarErroSemantico(
				ctx.IDENT().GetSymbol().GetLine(),
				"identificador "+nome+" ja declarado anteriormente",
			)
			return nil
		}

		var valor interface{}

		if ctx.Valor_constante().NUM_INT() != nil {
			valor = ctx.Valor_constante().NUM_INT().GetText()
		}

		if ctx.Valor_constante().NUM_REAL() != nil {
			valor = ctx.Valor_constante().NUM_REAL().GetText()
		}

		if ctx.Valor_constante().CADEIA() != nil {
			valor = ctx.Valor_constante().CADEIA().GetText()
		}

		j.tabela.AdicionarConstante(nome, tipo, valor)

		return nil
	}

	if ctx.TIPO() != nil {

		nome := ctx.IDENT().GetText()

		if j.tabela.Existe(nome) {
			AdicionarErroSemantico(
				ctx.IDENT().GetSymbol().GetLine(),
				"identificador "+nome+" ja declarado anteriormente",
			)
			return nil
		}

		j.tabela.Adicionar(nome, REGISTRO_TIPO)

		if ctx.Tipo() != nil &&
			ctx.Tipo().Registro() != nil {

			for _, v := range ctx.Tipo().Registro().AllVariavel() {

				tipoCampo := VerificarTipo(j.tabela, v.Tipo())

				for _, id := range v.AllIdentificador() {

					nomeCampo := id.IDENT(0).GetText()

					j.tabela.AdicionarCampoRegistro(
						nome,
						nomeCampo,
						EntradaTabela{Tipo: tipoCampo},
					)
				}
			}
		}
	}

	return nil
}

// VisitVariavel é o método de visita para o nó de declaração de variável, onde as variáveis são processadas para verificar a validade das declarações, atualizar a tabela de símbolos e lidar com casos específicos como registros e tipos estendidos.
func (j *JanderSemantico) VisitVariavel(
	ctx *parser.VariavelContext,
) interface{} {

	tipo := VerificarTipo(j.tabela, ctx.Tipo())

	var nomeTipo string

	if ctx.Tipo().Tipo_estendido() != nil {

		tb := ctx.Tipo().Tipo_estendido().Tipo_basico_ident()

		if tb != nil && tb.IDENT() != nil {
			nomeTipo = tb.IDENT().GetText()
		}
	}

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.IDENT(0).GetText()

		if j.tabela.ExisteNoEscopoAtual(nome) {

			AdicionarErroSemantico(
				ident.GetStart().GetLine(),
				"identificador "+nome+" ja declarado anteriormente",
			)

			continue
		}

		j.tabela.Adicionar(nome, tipo)

		if nomeTipo != "" {

			entradaTipo, ok := j.tabela.ObterEntrada(nomeTipo)

			if ok {

				for campo, entradaCampo := range entradaTipo.CamposRegistro {

					j.tabela.AdicionarCampoRegistro(
						nome,
						campo,
						entradaCampo,
					)
				}
			}
		}

		if ctx.Tipo().Registro() != nil {

			for _, campoVar := range ctx.Tipo().Registro().AllVariavel() {

				tipoCampo := VerificarTipo(
					j.tabela,
					campoVar.Tipo(),
				)

				for _, idCampo := range campoVar.AllIdentificador() {

					nomeCampo := idCampo.IDENT(0).GetText()

					j.tabela.AdicionarCampoRegistro(
						nome,
						nomeCampo,
						EntradaTabela{
							Nome:           nomeCampo,
							Tipo:           tipoCampo,
							CamposRegistro: make(map[string]EntradaTabela),
						},
					)
				}
			}
		}
	}

	return nil
}

// VisitDeclaracao_global é o método de visita para o nó de declaração global, onde as funções e procedimentos são processados para verificar a validade das declarações, atualizar a tabela de símbolos, criar novos escopos para os corpos das funções e procedimentos e realizar a análise semântica dos parâmetros, corpo e comandos associados.
func (j *JanderSemantico) VisitDeclaracao_global(
	ctx *parser.Declaracao_globalContext,
) interface{} {

	nome := ctx.IDENT().GetText()

	escopoAtual := j.tabela.escopos[len(j.tabela.escopos)-1]

	if _, ok := escopoAtual[nome]; ok {

		AdicionarErroSemantico(
			ctx.IDENT().GetSymbol().GetLine(),
			"identificador "+nome+" ja declarado anteriormente",
		)

		return nil
	}

	var parametros []TipoJander

	if ctx.Parametros() != nil {

		for _, p := range ctx.Parametros().AllParametro() {

			tipoParam := VerificarTipoEstendido(
				j.tabela,
				p.Tipo_estendido(),
			)

			for range p.AllIdentificador() {
				parametros = append(parametros, tipoParam)
			}
		}
	}

	if ctx.PROCEDIMENTO() != nil {

		j.tabela.AdicionarFuncao(
			nome,
			INVALIDO,
			parametros,
		)

		j.tabela.NovoEscopo(false)

		if ctx.Parametros() != nil {

			for _, p := range ctx.Parametros().AllParametro() {

				j.VisitParametro(
					p.(*parser.ParametroContext),
				)
			}
		}

		for _, d := range ctx.AllDeclaracao_local() {
			d.Accept(j)
		}

		for _, c := range ctx.AllCmd() {
			c.Accept(j)
		}

		j.tabela.AbandonarEscopo()
	}

	if ctx.FUNCAO() != nil {

		tipoRetorno := VerificarTipoEstendido(
			j.tabela,
			ctx.Tipo_estendido(),
		)

		j.tabela.AdicionarFuncao(
			nome,
			tipoRetorno,
			parametros,
		)

		j.tabela.NovoEscopo(true)
		j.tabela.SetTipoRetornoFuncaoAtual(tipoRetorno)

		if ctx.Parametros() != nil {

			for _, p := range ctx.Parametros().AllParametro() {

				j.VisitParametro(
					p.(*parser.ParametroContext),
				)
			}
		}

		for _, d := range ctx.AllDeclaracao_local() {
			d.Accept(j)
		}

		for _, c := range ctx.AllCmd() {
			c.Accept(j)
		}

		j.tabela.AbandonarEscopo()
	}

	return nil

}

// VisitParametro é o método de visita para o nó de parâmetro, onde os parâmetros são processados para verificar a validade das declarações, atualizar a tabela de símbolos e lidar com casos específicos como registros e tipos estendidos.
func (j *JanderSemantico) VisitParametro(
	ctx *parser.ParametroContext,
) interface{} {

	tipo := VerificarTipoEstendido(
		j.tabela,
		ctx.Tipo_estendido(),
	)

	var entradaTipo EntradaTabela
	var temRegistro bool

	tb := ctx.Tipo_estendido().Tipo_basico_ident()

	if tb != nil && tb.IDENT() != nil {

		nomeTipo := tb.IDENT().GetText()

		entrada, ok := j.tabela.ObterEntrada(nomeTipo)
		if ok {

			if len(entrada.CamposRegistro) > 0 {

				entradaTipo = entrada
				temRegistro = true
			}
		}
	}

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.IDENT(0).GetText()

		if j.tabela.ExisteNoEscopoAtual(nome) {

			AdicionarErroSemantico(
				ident.GetStart().GetLine(),
				"identificador "+nome+" ja declarado anteriormente",
			)

			continue
		}

		j.tabela.Adicionar(nome, tipo)

		if temRegistro {

			for campo, entradaCampo := range entradaTipo.CamposRegistro {

				j.tabela.AdicionarCampoRegistro(
					nome,
					campo,
					entradaCampo,
				)
			}
		}
	}

	return nil
}

// VisitCmd é o método de visita para o nó de comando, onde os diferentes tipos de comandos são processados para verificar a validade das operações, atualizar a tabela de símbolos conforme necessário e realizar a análise semântica dos comandos presentes no programa.
func (j *JanderSemantico) VisitCmd(ctx *parser.CmdContext) interface{} {

	if ctx.CmdChamada() != nil {
		return ctx.CmdChamada().Accept(j)
	}

	if ctx.CmdEscreva() != nil {
		return ctx.CmdEscreva().Accept(j)
	}

	if ctx.CmdSe() != nil {
		return ctx.CmdSe().Accept(j)
	}

	if ctx.CmdAtribuicao() != nil {
		return ctx.CmdAtribuicao().Accept(j)
	}

	if ctx.CmdLeia() != nil {
		return ctx.CmdLeia().Accept(j)
	}

	if ctx.CmdFaca() != nil {
		return ctx.CmdFaca().Accept(j)
	}
	if ctx.CmdEnquanto() != nil {
		return ctx.CmdEnquanto().Accept(j)
	}
	if ctx.CmdRetorne() != nil {
		return ctx.CmdRetorne().Accept(j)
	}
	return j.VisitChildren(ctx)
}

// VisitCmdLeia é o método de visita para o nó de comando de leitura, onde os identificadores envolvidos na leitura são processados para verificar se estão declarados na tabela de símbolos e reportar erros semânticos caso contrário.
func (j *JanderSemantico) VisitCmdLeia(
	ctx *parser.CmdLeiaContext,
) interface{} {

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.GetText()

		existe := j.tabela.Existe(nome)

		if !existe {

			if !ErroJaReportado(nome) {

				AdicionarErroSemantico(
					ident.GetStart().GetLine(),
					"identificador "+nome+" nao declarado",
				)
			}
		}
	}
	return nil
}

// VisitCmdFaca é o método de visita para o nó de comando de faça, onde os comandos dentro do bloco de faça são processados para realizar a análise semântica dos comandos e a expressão de condição associada ao comando de faça.
func (j *JanderSemantico) VisitCmdFaca(
	ctx *parser.CmdFacaContext,
) interface{} {

	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	if ctx.Expressao() != nil {

		tipo := j.tipoExpressao(ctx.Expressao())

		if tipo != LOGICO &&
			tipo != INVALIDO {

			AdicionarErroSemantico(
				ctx.Expressao().GetStart().GetLine(),
				"condicao do faca deve ser do tipo logico",
			)
		}
	}

	return nil
}

// VisitCmdEscreva é o método de visita para o nó de comando de escrita, onde as expressões envolvidas na escrita são processadas para realizar a análise semântica das expressões e verificar a validade dos identificadores presentes nas expressões.
func (j *JanderSemantico) VisitCmdEscreva(
	ctx *parser.CmdEscrevaContext,
) interface{} {

	for _, expr := range ctx.AllExpressao() {

		j.tipoExpressao(expr)

		j.checkIdentificadores(expr.(antlr.Tree))
	}

	return nil
}

// visitcmdse é o método de visita para o nó de comando de seleção (se), onde a expressão de condição e os comandos associados aos blocos de então e senao são processados para realizar a análise semântica das estruturas de controle presentes no programa.
func (j *JanderSemantico) VisitCmdSe(
	ctx *parser.CmdSeContext,
) interface{} {

	if ctx.Expressao() != nil {

		tipo := j.tipoExpressao(ctx.Expressao())

		if tipo != LOGICO &&
			tipo != INVALIDO {

			AdicionarErroSemantico(
				ctx.Expressao().GetStart().GetLine(),
				"condicao do se deve ser do tipo logico",
			)
		}
	}

	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	return nil
}

// VisitCmdEnquanto é o método de visita para o nó de comando de repetição (enquanto), onde a expressão de condição e os comandos associados ao bloco de enquanto são processados para realizar a análise semântica das estruturas de controle presentes no programa.
func (j *JanderSemantico) VisitCmdEnquanto(
	ctx *parser.CmdEnquantoContext,
) interface{} {

	if ctx.Expressao() != nil {

		tipo := j.tipoExpressao(ctx.Expressao())

		if tipo != LOGICO &&
			tipo != INVALIDO {

			AdicionarErroSemantico(
				ctx.Expressao().GetStart().GetLine(),
				"condicao do enquanto deve ser do tipo logico",
			)
		}
	}

	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	return nil
}

// VisitCmdAtribuicao é o método de visita para o nó de comando de atribuição, onde os identificadores envolvidos na atribuição e a expressão associada são processados para verificar a validade da operação de atribuição, atualizar a tabela de símbolos conforme necessário e realizar a análise semântica das expressões presentes na atribuição.
func (j *JanderSemantico) VisitCmdAtribuicao(
	ctx *parser.CmdAtribuicaoContext,
) interface{} {

	nomeBruto := ctx.Identificador().GetText()

	nomeErro := nomeBruto

	if ctx.PONTEIRO() != nil {
		nomeErro = "^" + nomeBruto
	}

	if !j.tabela.Existe(nomeBruto) {

		if !ErroJaReportado(nomeBruto) {

			AdicionarErroSemantico(
				ctx.Identificador().GetStart().GetLine(),
				"identificador "+nomeBruto+" nao declarado",
			)
		}

		return nil
	}

	tipoVar := j.tabela.Verificar(nomeBruto)

	if ctx.PONTEIRO() != nil {

		switch tipoVar {
		case PONTEIRO_INTEIRO:
			tipoVar = INTEIRO

		case PONTEIRO_REAL:
			tipoVar = REAL

		case PONTEIRO_LITERAL:
			tipoVar = LITERAL

		case PONTEIRO_LOGICO:
			tipoVar = LOGICO

		default:
			tipoVar = INVALIDO
		}
	}

	tipoExpr := j.tipoExpressao(ctx.Expressao())

	if !Compatibilidade(tipoVar, tipoExpr) {

		AdicionarErroSemantico(
			ctx.Identificador().GetStart().GetLine(),
			"atribuicao nao compativel para "+nomeErro,
		)
	}

	return nil
}

// VisitCmdChamada é o método de visita para o nó de comando de chamada, onde os identificadores envolvidos na chamada de função ou procedimento e as expressões associadas aos parâmetros são processados para verificar a validade da operação de chamada, atualizar a tabela de símbolos conforme necessário e realizar a análise semântica das expressões presentes na chamada.
func (j *JanderSemantico) VisitCmdChamada(
	ctx *parser.CmdChamadaContext,
) interface{} {

	nome := ctx.IDENT().GetText()

	if !j.tabela.Existe(nome) {

		if !ErroJaReportado(nome) {
			AdicionarErroSemantico(
				ctx.IDENT().GetSymbol().GetLine(),
				"identificador "+nome+" nao declarado",
			)
		}

		return nil
	}

	parametros := j.tabela.ObterParametros(nome)

	exprs := ctx.AllExpressao()

	if len(parametros) != len(exprs) {

		AdicionarErroSemantico(
			ctx.GetStart().GetLine(),
			"incompatibilidade de parametros na chamada de "+nome,
		)

		return nil
	}

	for i, expr := range exprs {

		tipoExpr := j.tipoExpressao(expr)

		if !CompatibilidadeFuncao(parametros[i], tipoExpr) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"incompatibilidade de parametros na chamada de "+nome,
			)
			break
		}
	}

	return nil
}

// VisitCmdRetorne é o método de visita para o nó de comando de retorno, onde a expressão associada ao comando de retorno é processada para verificar a validade da operação de retorno, comparar o tipo da expressão com o tipo de retorno esperado da função atual e realizar a análise semântica das expressões presentes no comando de retorno.
func (j *JanderSemantico) VisitCmdRetorne(
	ctx *parser.CmdRetorneContext,
) interface{} {

	if !j.tabela.EstaEmFuncao() {

		AdicionarErroSemantico(
			ctx.GetStart().GetLine(),
			"comando retorne nao permitido nesse escopo",
		)

		return nil
	}

	tipoExpr := j.tipoExpressao(ctx.Expressao())
	tipoEsperado := j.tabela.ObterTipoRetornoFuncaoAtual()

	if !Compatibilidade(tipoEsperado, tipoExpr) {

		AdicionarErroSemantico(
			ctx.GetStart().GetLine(),
			"tipo de retorno incompativel",
		)
	}

	return nil
}

// tipoExpressao é um método auxiliar que recebe um contexto de expressão e retorna o tipo da expressão, realizando a análise semântica das expressões presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas expressões.
func (j *JanderSemantico) tipoExpressao(
	ctx parser.IExpressaoContext,
) TipoJander {
	return VerificarExpressao(j.tabela, ctx)
}

// tipoTermoLogico é um método auxiliar que recebe um contexto de termo lógico e retorna o tipo do termo lógico, realizando a análise semântica dos termos lógicos presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nos termos lógicos.
func (j *JanderSemantico) tipoTermoLogico(
	ctx parser.ITermo_logicoContext,
) TipoJander {

	for _, fator := range ctx.AllFator_logico() {

		t := j.tipoFatorLogico(fator)

		if t == INVALIDO {
			return INVALIDO
		}
	}

	return LOGICO
}

// tipoFatorLogico é um método auxiliar que recebe um contexto de fator lógico e retorna o tipo do fator lógico, realizando a análise semântica dos fatores lógicos presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nos fatores lógicos.
func (j *JanderSemantico) tipoFatorLogico(
	ctx parser.IFator_logicoContext,
) TipoJander {

	return j.tipoParcelaLogica(ctx.Parcela_logica())
}

// tipoParcelaLogica é um método auxiliar que recebe um contexto de parcela lógica e retorna o tipo da parcela lógica, realizando a análise semântica das parcelas lógicas presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas parcelas lógicas.
func (j *JanderSemantico) tipoParcelaLogica(
	ctx parser.IParcela_logicaContext,
) TipoJander {

	if ctx.VERDADEIRO() != nil ||
		ctx.FALSO() != nil {

		return LOGICO
	}

	return j.tipoExpRelacional(ctx.Exp_relacional())
}

// tipoExpRelacional é um método auxiliar que recebe um contexto de expressão relacional e retorna o tipo da expressão relacional, realizando a análise semântica das expressões relacionais presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas expressões relacionais.
func (j *JanderSemantico) tipoExpRelacional(
	ctx parser.IExp_relacionalContext,
) TipoJander {

	if ctx.Op_relacional() != nil {

		t1 := j.tipoExpAritmetica(
			ctx.Exp_aritmetica(0),
		)

		t2 := j.tipoExpAritmetica(
			ctx.Exp_aritmetica(1),
		)

		if Compatibilidade(t1, t2) {
			return LOGICO
		}

		return INVALIDO
	}

	return j.tipoExpAritmetica(
		ctx.Exp_aritmetica(0),
	)
}

// tipoExpAritmetica é um método auxiliar que recebe um contexto de expressão aritmética e retorna o tipo da expressão aritmética, realizando a análise semântica das expressões aritméticas presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas expressões aritméticas.
func (j *JanderSemantico) tipoExpAritmetica(
	ctx parser.IExp_aritmeticaContext,
) TipoJander {

	var tipo TipoJander = INVALIDO

	for i, termo := range ctx.AllTermo() {

		t := j.tipoTermo(termo)

		if i == 0 {
			tipo = t
		} else if !Compatibilidade(tipo, t) {
			return INVALIDO
		}
	}

	return tipo
}

// tipoTermo é um método auxiliar que recebe um contexto de termo e retorna o tipo do termo, realizando a análise semântica dos termos presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nos termos.
func (j *JanderSemantico) tipoTermo(
	ctx parser.ITermoContext,
) TipoJander {

	var tipo TipoJander = INVALIDO

	for i, fator := range ctx.AllFator() {

		t := j.tipoFator(fator)

		if i == 0 {
			tipo = t
		} else if !Compatibilidade(tipo, t) {
			return INVALIDO
		}
	}

	return tipo
}

// tipoFator é um método auxiliar que recebe um contexto de fator e retorna o tipo do fator, realizando a análise semântica dos fatores presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nos fatores.
func (j *JanderSemantico) tipoFator(
	ctx parser.IFatorContext,
) TipoJander {

	var tipo TipoJander = INVALIDO

	for i, parcela := range ctx.AllParcela() {

		t := j.tipoParcela(parcela)

		if i == 0 {
			tipo = t
		} else if !Compatibilidade(tipo, t) {
			return INVALIDO
		}
	}

	return tipo
}

// tipoParcela é um método auxiliar que recebe um contexto de parcela e retorna o tipo da parcela, realizando a análise semântica das parcelas presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas parcelas.
func (j *JanderSemantico) tipoParcela(
	ctx parser.IParcelaContext,
) TipoJander {

	if ctx.Parcela_unario() != nil {
		return j.tipoParcelaUnario(
			ctx.Parcela_unario(),
		)
	}

	return j.tipoParcelaNaoUnario(
		ctx.Parcela_nao_unario(),
	)
}

// tipoParcelaUnario é um método auxiliar que recebe um contexto de parcela unária e retorna o tipo da parcela unária, realizando a análise semântica das parcelas unárias presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas parcelas unárias.
func (j *JanderSemantico) tipoParcelaUnario(
	ctx parser.IParcela_unarioContext,
) TipoJander {

	if ctx.NUM_INT() != nil {

		return INTEIRO
	}

	if ctx.NUM_REAL() != nil {

		return REAL
	}

	if ctx.IDENT() != nil {

		nome := ctx.IDENT().GetText()

		exprs := ctx.AllExpressao()

		if ctx.ABREPAR() != nil {

			if !j.tabela.Existe(nome) {
				if !ErroJaReportado(nome) {
					AdicionarErroSemantico(
						ctx.GetStart().GetLine(),
						"identificador "+nome+" nao declarado",
					)
				}
				return INVALIDO
			}

			parametros := j.tabela.ObterParametros(nome)

			if len(parametros) != len(exprs) {
				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"incompatibilidade de parametros na chamada de "+nome,
				)
				return j.tabela.ObterTipoRetorno(nome)
			}

			for i, expr := range exprs {

				tipoExpr := j.tipoExpressao(expr)

				if !CompatibilidadeFuncao(parametros[i], tipoExpr) {

					AdicionarErroSemantico(
						ctx.GetStart().GetLine(),
						"incompatibilidade de parametros na chamada de "+nome,
					)
					break
				}
			}

			return j.tabela.ObterTipoRetorno(nome)
		}

		return j.tabela.Verificar(nome)
	}

	if len(ctx.AllExpressao()) > 0 {

		return j.tipoExpressao(
			ctx.Expressao(0),
		)
	}

	return INVALIDO

}

// tipoParcelaNaoUnario é um método auxiliar que recebe um contexto de parcela não unária e retorna o tipo da parcela não unária, realizando a análise semântica das parcelas não unárias presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas parcelas não unárias.
func (j *JanderSemantico) tipoParcelaNaoUnario(ctx parser.IParcela_nao_unarioContext) TipoJander {
	if ctx.CADEIA() != nil {
		return LITERAL
	}
	if ctx.Identificador() != nil {
		nomeBruto := ctx.Identificador().GetText()
		nomeBase := extrairNomeBase(nomeBruto)

		if !j.tabela.Existe(nomeBruto) {
			if !ErroJaReportado(nomeBase) {
				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"identificador "+nomeBruto+" nao declarado",
				)
			}
			return INVALIDO
		}
		return j.tabela.Verificar(nomeBruto)
	}

	if ctx.Identificador() != nil {
		nomeBruto := ctx.Identificador().GetText()

		if !j.tabela.Existe(nomeBruto) {
			if !ErroJaReportado(nomeBruto) {
				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"identificador "+nomeBruto+" nao declarado",
				)
			}
			return INVALIDO
		}
		return j.tabela.Verificar(nomeBruto)
	}

	return INVALIDO
}

// checkIdentificadores é um método auxiliar que percorre a árvore de análise sintática em busca de identificadores e verifica se eles estão declarados na tabela de símbolos, reportando erros semânticos caso contrário.
func (j *JanderSemantico) checkIdentificadores(t antlr.Tree) {
	if ident, ok := t.(parser.IIdentificadorContext); ok {
		nomeBruto := ident.GetText()

		if !j.tabela.Existe(nomeBruto) {
			if !ErroJaReportado(nomeBruto) {
				AdicionarErroSemantico(
					ident.GetStart().GetLine(),
					"identificador "+nomeBruto+" nao declarado",
				)
			}
		}
		return
	}

	for _, ch := range t.GetChildren() {
		j.checkIdentificadores(ch)
	}
}

// VisitIdentificador é o método de visita para o nó de identificador, onde o nome do identificador é processado para verificar se está declarado na tabela de símbolos e reportar erros semânticos caso contrário.
func (j *JanderSemantico) VisitIdentificador(
	ctx *parser.IdentificadorContext,
) interface{} {

	nomeBruto := ctx.GetText()

	if !j.tabela.Existe(nomeBruto) {

		if !ErroJaReportado(nomeBruto) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"identificador "+nomeBruto+" nao declarado",
			)
		}
	}

	return nil
}

// extrairNomeBase é uma função auxiliar que recebe um nome completo de identificador, que pode incluir operadores de acesso a campos (ponto) e índices de vetores (colchetes), e extrai o nome base do identificador para fins de verificação na tabela de símbolos.
func extrairNomeBase(nomeCompleto string) string {
	if idxPonto := strings.Index(nomeCompleto, "."); idxPonto != -1 {
		nomeCompleto = nomeCompleto[:idxPonto]
	}
	if idxColchete := strings.Index(nomeCompleto, "["); idxColchete != -1 {
		nomeCompleto = nomeCompleto[:idxColchete]
	}
	return strings.ReplaceAll(nomeCompleto, "^", "")
}
