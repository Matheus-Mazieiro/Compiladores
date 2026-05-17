// esse é o semantico.go
package main

import (
	parser "go-lexer/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type JanderSemantico struct {
	parser.BaseCalcLexerVisitor
	tabela *TabelaDeSimbolos
}

func NewJanderSemantico() *JanderSemantico {
	return &JanderSemantico{
		tabela: nil,
	}
}

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

// ================= DECLARAÇÕES =================

func (j *JanderSemantico) VisitDeclaracoes(ctx *parser.DeclaracoesContext) interface{} {
	for _, d := range ctx.AllDecl_local_global() {
		d.Accept(j)
	}

	return nil
}

func (j *JanderSemantico) VisitDecl_local_global(ctx *parser.Decl_local_globalContext) interface{} {

	if ctx.Declaracao_local() != nil {
		ctx.Declaracao_local().Accept(j)
	}

	if ctx.Declaracao_global() != nil {
		ctx.Declaracao_global().Accept(j)
	}

	return nil
}

func (j *JanderSemantico) VisitCorpo(ctx *parser.CorpoContext) interface{} {

	for _, d := range ctx.AllDeclaracao_local() {
		d.Accept(j)
	}

	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	return nil
}

func (j *JanderSemantico) VisitDeclaracao_local(ctx *parser.Declaracao_localContext) interface{} {

	// variável
	if ctx.Variavel() != nil {
		ctx.Variavel().Accept(j)
		return nil
	}

	// constante
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

	// tipo
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
						tipoCampo,
					)
				}
			}
		}
	}

	return nil
}

// ================= FUNÇÕES / PROCEDIMENTOS =================
func (j *JanderSemantico) VisitVariavel(
	ctx *parser.VariavelContext,
) interface{} {

	tipo := VerificarTipo(j.tabela, ctx.Tipo())

	var nomeTipo string

	// pega o nome do tipo declarado
	if ctx.Tipo().Tipo_estendido() != nil {

		tb := ctx.Tipo().Tipo_estendido().Tipo_basico_ident()

		if tb != nil && tb.IDENT() != nil {
			nomeTipo = tb.IDENT().GetText()
		}
	}

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.IDENT(0).GetText()

		// redeclaração no escopo atual
		if j.tabela.ExisteNoEscopoAtual(nome) {

			AdicionarErroSemantico(
				ident.GetStart().GetLine(),
				"identificador "+nome+" ja declarado anteriormente",
			)

			continue
		}

		j.tabela.Adicionar(nome, tipo)

		// ================================
		// COPIA CAMPOS DO TIPO REGISTRO
		// ================================
		if nomeTipo != "" {

			entradaTipo, ok := j.tabela.ObterEntrada(nomeTipo)

			if ok {

				for campo, tipoCampo := range entradaTipo.CamposRegistro {

					j.tabela.AdicionarCampoRegistro(
						nome,
						campo,
						tipoCampo,
					)
				}
			}
		}

		// ================================
		// REGISTRO INLINE
		// ================================
		if ctx.Tipo().Registro() != nil {

			for _, campo := range ctx.Tipo().Registro().AllVariavel() {

				tipoCampo := VerificarTipo(
					j.tabela,
					campo.Tipo(),
				)

				for _, idCampo := range campo.AllIdentificador() {

					nomeCampo := idCampo.IDENT(0).GetText()

					j.tabela.AdicionarCampoRegistro(
						nome,
						nomeCampo,
						tipoCampo,
					)
				}
			}
		}
	}

	return nil
}

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

	// procedimento
	// procedimento
	if ctx.PROCEDIMENTO() != nil {

		j.tabela.AdicionarFuncao(
			nome,
			INVALIDO,
			parametros,
		)

		j.tabela.NovoEscopo(false)

		// adiciona parâmetros
		if ctx.Parametros() != nil {

			for _, p := range ctx.Parametros().AllParametro() {

				j.VisitParametro(
					p.(*parser.ParametroContext),
				)
			}
		}

		// declarações locais
		for _, d := range ctx.AllDeclaracao_local() {
			d.Accept(j)
		}

		// comandos
		for _, c := range ctx.AllCmd() {
			c.Accept(j)
		}

		j.tabela.AbandonarEscopo()
	}

	// função
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

		// parâmetros
		if ctx.Parametros() != nil {

			for _, p := range ctx.Parametros().AllParametro() {

				j.VisitParametro(
					p.(*parser.ParametroContext),
				)
			}
		}

		// declarações locais
		for _, d := range ctx.AllDeclaracao_local() {
			d.Accept(j)
		}

		// comandos
		for _, c := range ctx.AllCmd() {
			c.Accept(j)
		}

		j.tabela.AbandonarEscopo()
	}

	return nil

}

func (j *JanderSemantico) VisitParametro(
	ctx *parser.ParametroContext,
) interface{} {

	tipo := VerificarTipoEstendido(
		j.tabela,
		ctx.Tipo_estendido(),
	)

	var entradaTipo EntradaTabela
	var temRegistro bool

	// pega typedef ANTES de adicionar variável
	tb := ctx.Tipo_estendido().Tipo_basico_ident()

	if tb != nil && tb.IDENT() != nil {

		nomeTipo := tb.IDENT().GetText()

		entrada, ok := j.tabela.ObterEntrada(nomeTipo)

		if ok {

			if entrada.CamposRegistro != nil {

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
		println("PARAMETRO ADICIONADO:", nome)

		entrada, ok := j.tabela.ObterEntrada(nome)

		if ok {

			println("Campos do parametro:")

			for campo := range entrada.CamposRegistro {
				println(" ->", campo)
			}
		} else {
			println("NAO ACHOU PARAMETRO")
		}
		// copia campos do registro
		if temRegistro {

			for campo, tipoCampo := range entradaTipo.CamposRegistro {

				j.tabela.AdicionarCampoRegistro(
					nome,
					campo,
					tipoCampo,
				)
			}
		}
	}

	return nil
}

// ================= COMANDOS =================

func (j *JanderSemantico) VisitCmd(ctx *parser.CmdContext) interface{} {

	println("ENTROU EM CMD")

	if ctx.CmdLeia() != nil {
		println("ACHOU CMDLEIA")
		return ctx.CmdLeia().Accept(j)
	}

	if ctx.CmdEscreva() != nil {
		println("ACHOU CMDESCREVA")
		return ctx.CmdEscreva().Accept(j)
	}
	if ctx.CmdFaca() != nil {
		return ctx.CmdFaca().Accept(j)
	}
	if ctx.CmdSe() != nil {
		println("ACHOU CMDSE")
		return ctx.CmdSe().Accept(j)
	}

	if ctx.CmdEnquanto() != nil {
		println("ACHOU CMDENQUANTO")
		return ctx.CmdEnquanto().Accept(j)
	}

	if ctx.CmdAtribuicao() != nil {
		println("ACHOU CMDATRIB")
		return ctx.CmdAtribuicao().Accept(j)
	}

	if ctx.CmdChamada() != nil {
		println("ACHOU CMDCHAMADA")
		return ctx.CmdChamada().Accept(j)
	}

	if ctx.CmdRetorne() != nil {
		println("ACHOU CMDRETORNE")
		return ctx.CmdRetorne().Accept(j)
	}

	return j.VisitChildren(ctx)
}

func (j *JanderSemantico) VisitCmdLeia(
	ctx *parser.CmdLeiaContext,
) interface{} {

	println("===== CMD LEIA =====")

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.GetText()

		println("LEIA IDENT:", nome)

		existe := j.tabela.Existe(nome)

		println("EXISTE?", existe)

		if !existe {

			println(">>> VAI REPORTAR ERRO")

			if !ErroJaReportado(nome) {

				AdicionarErroSemantico(
					ident.GetStart().GetLine(),
					"identificador "+nome+" nao declarado",
				)
			}
		}
	}
	println("QTD IDENT:", len(ctx.AllIdentificador()))
	return nil
}
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
func (j *JanderSemantico) VisitCmdEscreva(
	ctx *parser.CmdEscrevaContext,
) interface{} {

	for _, expr := range ctx.AllExpressao() {
		j.checkIdentificadores(expr.(antlr.Tree))
	}

	return nil
}

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

func (j *JanderSemantico) VisitCmdAtribuicao(
	ctx *parser.CmdAtribuicaoContext,
) interface{} {

	nomeBruto := ctx.Identificador().GetText()

	nomeErro := nomeBruto

	if ctx.PONTEIRO() != nil {
		nomeErro = "^" + nomeBruto
	}

	// verifica identificador COMPLETO
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

	// desreferenciamento
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

	if len(parametros) != len(ctx.AllExpressao()) {

		AdicionarErroSemantico(
			ctx.GetStart().GetLine(),
			"incompatibilidade de parametros na chamada de "+nome,
		)

		return nil
	}

	for i, expr := range ctx.AllExpressao() {

		tipoExpr := j.tipoExpressao(expr)

		if !CompatibilidadeFuncao(
			parametros[i],
			tipoExpr,
		) {

			AdicionarErroSemantico(
				ctx.GetStart().GetLine(),
				"incompatibilidade de parametros na chamada de "+nome,
			)

			break
		}
	}

	return nil
}

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

// ================= EXPRESSÕES =================

func (j *JanderSemantico) tipoExpressao(
	ctx parser.IExpressaoContext,
) TipoJander {
	return VerificarExpressao(j.tabela, ctx)
}

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

func (j *JanderSemantico) tipoFatorLogico(
	ctx parser.IFator_logicoContext,
) TipoJander {

	return j.tipoParcelaLogica(ctx.Parcela_logica())
}

func (j *JanderSemantico) tipoParcelaLogica(
	ctx parser.IParcela_logicaContext,
) TipoJander {

	if ctx.VERDADEIRO() != nil ||
		ctx.FALSO() != nil {

		return LOGICO
	}

	return j.tipoExpRelacional(ctx.Exp_relacional())
}

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

func (j *JanderSemantico) tipoParcelaUnario(ctx parser.IParcela_unarioContext) TipoJander {
	if ctx.NUM_INT() != nil {
		return INTEIRO
	}

	if ctx.NUM_REAL() != nil {
		return REAL
	}

	// Dentro de tipoParcelaUnario:
	if ctx.Identificador() != nil {
		nomeBruto := ctx.Identificador().GetText()
		nomeBase := extrairNomeBase(nomeBruto)

		if !j.tabela.Existe(nomeBruto) { // Valida usando a string completa modificada
			if !ErroJaReportado(nomeBase) {
				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"identificador "+nomeBruto+" nao declarado",
				)
			}
			return INVALIDO
		}
		return j.tabela.Verificar(nomeBruto) // Garante o retorno do tipo do campo (.preco)
	}

	// Dentro de tipoParcelaNaoUnario:
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
		nomeBase := extrairNomeBase(nomeBruto) // <-- Limpeza aplicada aqui

		if !j.tabela.Existe(nomeBase) {
			if !ErroJaReportado(nomeBase) {
				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"identificador "+nomeBruto+" nao declarado",
				)
			}
			return INVALIDO
		}
		return j.tabela.Verificar(nomeBruto) // Passa nomeBruto para a tabela resolver internamente os subcampos
	}

	if len(ctx.AllExpressao()) > 0 {
		return j.tipoExpressao(ctx.Expressao(0))
	}

	if ctx.IDENT() != nil {
		nome := ctx.IDENT().GetText()
		if !j.tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(
					ctx.GetStart().GetLine(),
					"identificador "+nome+" nao declarado",
				)
			}
			return INVALIDO
		}
		return j.tabela.ObterTipoRetorno(nome)
	}

	return INVALIDO
}

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
		nomeBase := extrairNomeBase(nomeBruto) // <-- Limpeza aplicada aqui

		if !j.tabela.Existe(nomeBase) {
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

	return INVALIDO
}

// ================= HELPERS =================
func (j *JanderSemantico) checkIdentificadores(t antlr.Tree) {
	if ident, ok := t.(parser.IIdentificadorContext); ok {
		nomeBruto := ident.GetText()
		nomeBase := extrairNomeBase(nomeBruto) // <-- Limpeza aplicada aqui

		if !j.tabela.Existe(nomeBruto) {
			if !ErroJaReportado(nomeBase) {
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
func extrairNomeBase(nomeCompleto string) string {
	if idxPonto := strings.Index(nomeCompleto, "."); idxPonto != -1 {
		nomeCompleto = nomeCompleto[:idxPonto]
	}
	if idxColchete := strings.Index(nomeCompleto, "["); idxColchete != -1 {
		nomeCompleto = nomeCompleto[:idxColchete]
	}
	return strings.ReplaceAll(nomeCompleto, "^", "")
}
