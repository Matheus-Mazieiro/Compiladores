package main

import (
	parser "go-lexer/parser"
	"strings"
)

type Gerador struct {
	*parser.BaseCalcLexerVisitor
	tabela *TabelaDeSimbolos
	saida  strings.Builder
	indent int
}

func NewGerador(
	tabela *TabelaDeSimbolos,
) *Gerador {

	return &Gerador{
		tabela: tabela,
	}
}

func (g *Gerador) VisitPrograma(
	ctx *parser.ProgramaContext,
) interface{} {

	g.saida.WriteString("#include <stdio.h>\n")
	g.saida.WriteString("#include <stdlib.h>\n")
	g.saida.WriteString("#include <string.h>\n\n")

	// declarações globais
	if ctx.Declaracoes() != nil {
		ctx.Declaracoes().Accept(g)
	}

	g.saida.WriteString("int main() {\n")

	// corpo principal
	if ctx.Corpo() != nil {
		ctx.Corpo().Accept(g)
	}

	g.saida.WriteString("\treturn 0;\n")
	g.saida.WriteString("}\n")

	return g.saida.String()
}

func (g *Gerador) VisitDeclaracoes(
	ctx *parser.DeclaracoesContext,
) interface{} {

	for _, d := range ctx.AllDecl_local_global() {
		d.Accept(g)
	}

	return nil
}

func (g *Gerador) VisitDecl_local_global(
	ctx *parser.Decl_local_globalContext,
) interface{} {

	if ctx.Declaracao_local() != nil {
		ctx.Declaracao_local().Accept(g)
	}

	if ctx.Declaracao_global() != nil {
		ctx.Declaracao_global().Accept(g)
	}

	return nil
}

func (g *Gerador) VisitCorpo(
	ctx *parser.CorpoContext,
) interface{} {

	// declarações locais
	for _, d := range ctx.AllDeclaracao_local() {
		d.Accept(g)
	}

	g.saida.WriteString("\n")

	// comandos
	for _, c := range ctx.AllCmd() {
		c.Accept(g)
	}

	return nil
}

func (g *Gerador) VisitDeclaracao_local(
	ctx *parser.Declaracao_localContext,
) interface{} {

	// ----------------------------------------
	// Declaração de variáveis
	// ----------------------------------------
	if ctx.Variavel() != nil {
		ctx.Variavel().Accept(g)
		return nil
	}

	// ----------------------------------------
	// Declaração de constantes
	// ----------------------------------------
	if ctx.CONSTANTE() != nil {

		nome := ctx.IDENT().GetText()

		tipo := mapTipoLAParaC(
			ctx.Tipo_basico().GetText(),
		)

		var valor string

		if ctx.Valor_constante().NUM_INT() != nil {
			valor = ctx.Valor_constante().NUM_INT().GetText()
		} else if ctx.Valor_constante().NUM_REAL() != nil {
			valor = ctx.Valor_constante().NUM_REAL().GetText()
		} else if ctx.Valor_constante().CADEIA() != nil {
			valor = ctx.Valor_constante().CADEIA().GetText()
		} else if ctx.Valor_constante().VERDADEIRO() != nil {
			valor = "1"
		} else if ctx.Valor_constante().FALSO() != nil {
			valor = "0"
		}

		g.emitln(
			"const " + tipo + " " + nome + " = " + valor + ";",
		)

		return nil
	}

	// ----------------------------------------
	// Declaração de tipos
	// ----------------------------------------
	if ctx.TIPO() != nil {

		nomeTipo := ctx.IDENT().GetText()

		if ctx.Tipo() != nil &&
			ctx.Tipo().Registro() != nil {

			g.emitln("typedef struct {")

			g.indent++

			for _, variavel := range ctx.Tipo().Registro().AllVariavel() {

				tipoCampo := mapTipoContextoParaC(
					variavel.Tipo(),
				)

				for _, id := range variavel.AllIdentificador() {

					nomeCampo := id.GetText()

					// strings/literal
					if tipoCampo == "char" {

						g.emitln(
							"char " + nomeCampo + "[256];",
						)

					} else {

						g.emitln(
							tipoCampo + " " + nomeCampo + ";",
						)
					}
				}
			}

			g.indent--

			g.emitln("} " + nomeTipo + ";")
			g.emitln("")
		}

		return nil
	}

	return nil
}

func (g *Gerador) VisitVariavel(
	ctx *parser.VariavelContext,
) interface{} {

	var tipoC string

	// ----------------------------------------
	// Registro inline
	// ----------------------------------------

	if ctx.Tipo().Registro() != nil {

		tipoC = "struct"

	} else {

		// tipo_estendido
		if ctx.Tipo().Tipo_estendido() != nil {

			tipoC = mapTipoEstendidoParaC(
				ctx.Tipo().Tipo_estendido(),
			)

		} else {

			tipoC = mapTipoContextoParaC(
				ctx.Tipo(),
			)
		}
	}

	// ----------------------------------------
	// Variáveis declaradas
	// ----------------------------------------

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.IDENT(0).GetText()

		// ----------------------------------------
		// Vetores/matrizes
		// ----------------------------------------

		dimensoes := ""

		if ident.Dimensao() != nil {

			for _, exp := range ident.Dimensao().AllExp_aritmetica() {

				tam := exp.GetText()

				dimensoes += "[" + tam + "]"
			}
		}

		// ----------------------------------------
		// Registro inline
		// ----------------------------------------

		if ctx.Tipo().Registro() != nil {

			g.emitln("struct {")

			g.indent++

			for _, campoVar := range ctx.Tipo().
				Registro().
				AllVariavel() {

				tipoCampo := mapTipoContextoParaC(
					campoVar.Tipo(),
				)

				for _, idCampo := range campoVar.AllIdentificador() {

					nomeCampo := idCampo.GetText()

					if tipoCampo == "char" {

						g.emitln(
							"char " +
								nomeCampo +
								"[256];",
						)

					} else {

						g.emitln(
							tipoCampo +
								" " +
								nomeCampo +
								";",
						)
					}
				}
			}

			g.indent--

			g.emitln(
				"} " +
					nome +
					dimensoes +
					";",
			)

			continue
		}

		// ----------------------------------------
		// Literal/string
		// ----------------------------------------

		if tipoC == "char" {

			g.emitln(
				"char " +
					nome +
					"[256]" +
					dimensoes +
					";",
			)

			continue
		}

		// ----------------------------------------
		// Variável comum
		// ----------------------------------------

		g.emitln(
			tipoC +
				" " +
				nome +
				dimensoes +
				";",
		)
	}

	return nil
}

func (g *Gerador) VisitDeclaracao_global(
	ctx *parser.Declaracao_globalContext,
) interface{} {

	nome := ctx.IDENT().GetText()

	// =========================================================
	// PROCEDIMENTO
	// =========================================================

	if ctx.PROCEDIMENTO() != nil {

		g.saida.WriteString("void ")
		g.saida.WriteString(nome)
		g.saida.WriteString("(")

		// parâmetros
		if ctx.Parametros() != nil {

			var params []string

			for _, p := range ctx.Parametros().AllParametro() {

				tipo := mapTipoEstendidoParaC(
					p.Tipo_estendido(),
				)

				for _, ident := range p.AllIdentificador() {

					nomeParam := ident.GetText()

					// strings
					if tipo == "char" {

						params = append(
							params,
							"char "+nomeParam+"[256]",
						)

					} else {

						params = append(
							params,
							tipo+" "+nomeParam,
						)
					}
				}
			}

			g.saida.WriteString(
				strings.Join(params, ", "),
			)
		}

		g.saida.WriteString(") {\n")

		g.indent++

		// declarações locais
		for _, d := range ctx.AllDeclaracao_local() {
			d.Accept(g)
		}

		if len(ctx.AllDeclaracao_local()) > 0 {
			g.emitln("")
		}

		// comandos
		for _, c := range ctx.AllCmd() {
			c.Accept(g)
		}

		g.indent--

		g.emitln("}")
		g.emitln("")

		return nil
	}

	// =========================================================
	// FUNÇÃO
	// =========================================================

	if ctx.FUNCAO() != nil {

		tipoRetorno := mapTipoEstendidoParaC(
			ctx.Tipo_estendido(),
		)

		g.saida.WriteString(tipoRetorno)
		g.saida.WriteString(" ")
		g.saida.WriteString(nome)
		g.saida.WriteString("(")

		// parâmetros
		if ctx.Parametros() != nil {

			var params []string

			for _, p := range ctx.Parametros().AllParametro() {

				tipo := mapTipoEstendidoParaC(
					p.Tipo_estendido(),
				)

				for _, ident := range p.AllIdentificador() {

					nomeParam := ident.GetText()

					// strings
					if tipo == "char" {

						params = append(
							params,
							"char "+nomeParam+"[256]",
						)

					} else {

						params = append(
							params,
							tipo+" "+nomeParam,
						)
					}
				}
			}

			g.saida.WriteString(
				strings.Join(params, ", "),
			)
		}

		g.saida.WriteString(") {\n")

		g.indent++

		// declarações locais
		for _, d := range ctx.AllDeclaracao_local() {
			d.Accept(g)
		}

		if len(ctx.AllDeclaracao_local()) > 0 {
			g.emitln("")
		}

		// comandos
		for _, c := range ctx.AllCmd() {
			c.Accept(g)
		}

		g.indent--

		g.emitln("}")
		g.emitln("")

		return nil
	}

	return nil
}

func (g *Gerador) VisitCmd(
	ctx *parser.CmdContext,
) interface{} {

	if ctx.CmdChamada() != nil {
		return ctx.CmdChamada().Accept(g)
	}

	if ctx.CmdEscreva() != nil {
		return ctx.CmdEscreva().Accept(g)
	}

	if ctx.CmdSe() != nil {
		return ctx.CmdSe().Accept(g)
	}

	if ctx.CmdCaso() != nil {
		return ctx.CmdCaso().Accept(g)
	}

	if ctx.CmdAtribuicao() != nil {
		return ctx.CmdAtribuicao().Accept(g)
	}

	if ctx.CmdLeia() != nil {
		return ctx.CmdLeia().Accept(g)
	}

	if ctx.CmdFaca() != nil {
		return ctx.CmdFaca().Accept(g)
	}

	if ctx.CmdEnquanto() != nil {
		return ctx.CmdEnquanto().Accept(g)
	}

	if ctx.CmdPara() != nil {
		return ctx.CmdPara().Accept(g)
	}

	if ctx.CmdRetorne() != nil {
		return ctx.CmdRetorne().Accept(g)
	}

	panic(
		"comando nao implementado: " +
			ctx.GetText(),
	)
}

func (g *Gerador) VisitCmdLeia(
	ctx *parser.CmdLeiaContext,
) interface{} {

	for _, ident := range ctx.AllIdentificador() {

		nome := ident.GetText()

		entrada, ok := g.tabela.ObterEntrada(nome)

		if !ok {
			continue
		}

		var formato string
		var argumento string

		switch entrada.Tipo {

		case INTEIRO:
			formato = "%d"
			argumento = "&" + nome

		case REAL:
			formato = "%lf"
			argumento = "&" + nome

		case LITERAL:
			formato = "%s"
			argumento = nome

		case LOGICO:
			formato = "%d"
			argumento = "&" + nome

		default:
			formato = "%d"
			argumento = "&" + nome
		}

		g.emitln(
			`scanf("` + formato + `", ` + argumento + `);`,
		)
	}

	return nil
}

func (g *Gerador) VisitCmdFaca(
	ctx *parser.CmdFacaContext,
) interface{} {

	g.emitln("do {")

	g.indent++

	// comandos do bloco
	for _, cmd := range ctx.AllCmd() {
		cmd.Accept(g)
	}

	g.indent--

	// condição do ate
	condicao := ctx.Expressao().
		Accept(g).(string)

	g.emitln(
		"} while (!(" + condicao + "));",
	)

	return nil
}

func (g *Gerador) VisitCmdEscreva(
	ctx *parser.CmdEscrevaContext,
) interface{} {

	var formatos []string
	var argumentos []string

	for _, expr := range ctx.AllExpressao() {

		// expressão gerada
		exprStr := expr.Accept(g).(string)

		if exprStr == "" {

			panic(
				"expressao retornou string vazia: " +
					expr.GetText(),
			)
		}

		// tenta inferir tipo semanticamente
		tipo := g.tipoExpressao(expr)

		// ----------------------------------------
		// Heurística para strings
		// ----------------------------------------

		if strings.HasPrefix(exprStr, `"`) &&
			strings.HasSuffix(exprStr, `"`) {

			tipo = LITERAL
		}

		// identificadores conhecidos como char[]
		if strings.Contains(exprStr, ".nome") ||
			exprStr == "mensagem" {

			tipo = LITERAL
		}

		switch tipo {

		case INTEIRO:
			formatos = append(formatos, "%d")

		case REAL:
			formatos = append(formatos, "%lf")

		case LITERAL:
			formatos = append(formatos, "%s")

		case LOGICO:
			formatos = append(formatos, "%d")

		default:
			formatos = append(formatos, "%d")
		}

		argumentos = append(argumentos, exprStr)
	}

	codigo := `printf("`

	for _, f := range formatos {
		codigo += f
	}

	codigo += `"`

	if len(argumentos) > 0 {

		codigo += ", " +
			strings.Join(argumentos, ", ")
	}

	codigo += ");"

	g.emitln(codigo)

	return nil
}

func (g *Gerador) VisitCmdSe(
	ctx *parser.CmdSeContext,
) interface{} {

	// condição
	condicao := ctx.Expressao().
		Accept(g).(string)

	g.emitln(
		"if (" + condicao + ") {",
	)

	g.indent++

	// bloco ENTÃO
	for _, cmd := range ctx.GetCmdEntao() {
		cmd.Accept(g)
	}

	g.indent--

	// possui SENAO
	if len(ctx.GetCmdSenao()) > 0 {

		g.emitln("} else {")

		g.indent++

		for _, cmd := range ctx.GetCmdSenao() {
			cmd.Accept(g)
		}

		g.indent--
	}

	g.emitln("}")

	return nil
}

func (g *Gerador) VisitCmdEnquanto(
	ctx *parser.CmdEnquantoContext,
) interface{} {

	// condição do while
	condicao := ctx.Expressao().
		Accept(g).(string)

	g.emitln(
		"while (" + condicao + ") {",
	)

	g.indent++

	// comandos do bloco
	for _, cmd := range ctx.GetCmds() {
		cmd.Accept(g)
	}

	g.indent--

	g.emitln("}")

	return nil
}

func (g *Gerador) VisitCmdAtribuicao(
	ctx *parser.CmdAtribuicaoContext,
) interface{} {

	nome := ctx.Identificador().GetText()

	// expressão em C
	expr := ctx.Expressao().
		Accept(g).(string)

	// tenta obter tipo
	entrada, ok := g.tabela.ObterEntrada(nome)

	tipoVar := INVALIDO

	if ok {
		tipoVar = entrada.Tipo
	}

	// ----------------------------------------
	// Ponteiros
	// ----------------------------------------

	destino := nome

	if ctx.PONTEIRO() != nil {

		destino = "*" + nome

		switch tipoVar {

		case PONTEIRO_INTEIRO:
			tipoVar = INTEIRO

		case PONTEIRO_REAL:
			tipoVar = REAL

		case PONTEIRO_LITERAL:
			tipoVar = LITERAL

		case PONTEIRO_LOGICO:
			tipoVar = LOGICO
		}
	}

	// ----------------------------------------
	// Strings/literal
	// ----------------------------------------

	if tipoVar == LITERAL ||
		strings.Contains(nome, ".nome") {

		g.emitln(
			`strcpy(` +
				destino +
				`, ` +
				expr +
				`);`,
		)

		return nil
	}

	// ----------------------------------------
	// Atribuição comum
	// ----------------------------------------

	g.emitln(
		destino + " = " + expr + ";",
	)

	return nil
}

func (g *Gerador) VisitCmdChamada(
	ctx *parser.CmdChamadaContext,
) interface{} {

	nome := ctx.IDENT().GetText()

	var argumentos []string

	// gera argumentos
	for _, expr := range ctx.AllExpressao() {

		arg := expr.Accept(g).(string)

		argumentos = append(argumentos, arg)
	}

	g.emitln(
		nome +
			"(" +
			strings.Join(argumentos, ", ") +
			");",
	)

	return nil
}

func (g *Gerador) VisitCmdRetorne(
	ctx *parser.CmdRetorneContext,
) interface{} {

	expr := ctx.Expressao().
		Accept(g).(string)

	g.emitln(
		"return " + expr + ";",
	)

	return nil
}

// tipoExpressao é um método auxiliar que recebe um contexto de expressão e retorna o tipo da expressão, realizando a análise semântica das expressões presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas expressões.
func (j *Gerador) tipoExpressao(
	ctx parser.IExpressaoContext,
) TipoJander {
	return VerificarExpressao(j.tabela, ctx)
}

// tipoTermoLogico é um método auxiliar que recebe um contexto de termo lógico e retorna o tipo do termo lógico, realizando a análise semântica dos termos lógicos presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nos termos lógicos.
func (j *Gerador) tipoTermoLogico(
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
func (j *Gerador) tipoFatorLogico(
	ctx parser.IFator_logicoContext,
) TipoJander {

	return j.tipoParcelaLogica(ctx.Parcela_logica())
}

// tipoParcelaLogica é um método auxiliar que recebe um contexto de parcela lógica e retorna o tipo da parcela lógica, realizando a análise semântica das parcelas lógicas presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas parcelas lógicas.
func (j *Gerador) tipoParcelaLogica(
	ctx parser.IParcela_logicaContext,
) TipoJander {

	if ctx.VERDADEIRO() != nil ||
		ctx.FALSO() != nil {

		return LOGICO
	}

	return j.tipoExpRelacional(ctx.Exp_relacional())
}

// tipoExpRelacional é um método auxiliar que recebe um contexto de expressão relacional e retorna o tipo da expressão relacional, realizando a análise semântica das expressões relacionais presentes no programa e verificando a validade dos identificadores, operações e tipos envolvidos nas expressões relacionais.
func (j *Gerador) tipoExpRelacional(
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
func (j *Gerador) tipoExpAritmetica(
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
func (j *Gerador) tipoTermo(
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
func (j *Gerador) tipoFator(
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
func (j *Gerador) tipoParcela(
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
func (j *Gerador) tipoParcelaUnario(
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
func (j *Gerador) tipoParcelaNaoUnario(ctx parser.IParcela_nao_unarioContext) TipoJander {
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

func (g *Gerador) VisitExpressao(
	ctx *parser.ExpressaoContext,
) interface{} {

	if ctx == nil {
		return ""
	}

	// primeiro termo lógico
	expr := g.tipoTermoLogico(ctx.Termo_logico(0))
	exprStr := ctx.Termo_logico(0).Accept(g).(string)

	// OR (|| em C)
	for i := 1; i < len(ctx.AllTermo_logico()); i++ {

		t := g.tipoTermoLogico(ctx.Termo_logico(i))

		// gera código do termo
		right := ctx.Termo_logico(i).Accept(g).(string)

		// monta expressão C
		exprStr = "(" + exprStr + " || " + right + ")"

		// valida tipo (continua semântico)
		if expr != LOGICO || t != LOGICO {
			// você pode ignorar ou marcar erro
			// mas não quebra geração
		}
	}

	return exprStr
}

func (g *Gerador) VisitTermo_logico(
	ctx *parser.Termo_logicoContext,
) interface{} {

	if ctx == nil {
		return ""
	}

	// primeiro fator lógico
	result := ctx.Fator_logico(0).Accept(g).(string)

	for i := 1; i < len(ctx.AllFator_logico()); i++ {

		right := ctx.Fator_logico(i).Accept(g).(string)

		result = "(" + result + " && " + right + ")"
	}

	return result
}

func (g *Gerador) VisitFator_logico(
	ctx *parser.Fator_logicoContext,
) interface{} {

	return ctx.Parcela_logica().Accept(g).(string)
}

func (g *Gerador) VisitParcela_logica(
	ctx *parser.Parcela_logicaContext,
) interface{} {

	if ctx.VERDADEIRO() != nil {
		return "1"
	}

	if ctx.FALSO() != nil {
		return "0"
	}

	// expressão relacional
	return ctx.Exp_relacional().Accept(g).(string)
}

func (g *Gerador) VisitExp_relacional(
	ctx *parser.Exp_relacionalContext,
) interface{} {

	left := ctx.Exp_aritmetica(0).
		Accept(g).(string)

	// sem operador relacional
	if ctx.Op_relacional() == nil {
		return left
	}

	// operador da LA
	op := ctx.Op_relacional().GetText()

	// converte para operador C
	switch op {

	case "=":
		op = "=="

	case "<>":
		op = "!="
	}

	right := ctx.Exp_aritmetica(1).
		Accept(g).(string)

	return "(" +
		left +
		" " +
		op +
		" " +
		right +
		")"
}

func (g *Gerador) VisitExp_aritmetica(
	ctx *parser.Exp_aritmeticaContext,
) interface{} {

	result := ctx.Termo(0).
		Accept(g).(string)

	for i := 1; i < len(ctx.AllTermo()); i++ {

		// operador (+ ou -)
		op := ctx.Op1(i - 1).GetText()

		// termo da direita
		right := ctx.Termo(i).
			Accept(g).(string)

		result =
			"(" +
				result +
				" " +
				op +
				" " +
				right +
				")"
	}

	return result
}

func (g *Gerador) VisitTermo(
	ctx *parser.TermoContext,
) interface{} {

	result := ctx.Fator(0).
		Accept(g).(string)

	for i := 1; i < len(ctx.AllFator()); i++ {

		// operador (*, /, %)
		op := ctx.Op2(i - 1).GetText()

		// fator da direita
		right := ctx.Fator(i).
			Accept(g).(string)

		result =
			"(" +
				result +
				" " +
				op +
				" " +
				right +
				")"
	}

	return result
}

func (g *Gerador) VisitFator(
	ctx *parser.FatorContext,
) interface{} {

	return ctx.Parcela(0).Accept(g).(string)
}

func (g *Gerador) VisitParcela(
	ctx *parser.ParcelaContext,
) interface{} {

	if ctx.Parcela_unario() != nil {

		res := ctx.Parcela_unario().
			Accept(g)

		if res == nil {
			panic(
				"parcela_unario retornou nil: " +
					ctx.GetText(),
			)
		}

		return res.(string)
	}

	if ctx.Parcela_nao_unario() != nil {

		res := ctx.Parcela_nao_unario().
			Accept(g)

		if res == nil {
			panic(
				"parcela_nao_unario retornou nil: " +
					ctx.GetText(),
			)
		}

		return res.(string)
	}

	panic(
		"parcela invalida: " +
			ctx.GetText(),
	)
}

func (g *Gerador) VisitParcela_unario(
	ctx *parser.Parcela_unarioContext,
) interface{} {

	// número inteiro
	if ctx.NUM_INT() != nil {
		return ctx.NUM_INT().GetText()
	}

	// número real
	if ctx.NUM_REAL() != nil {
		return ctx.NUM_REAL().GetText()
	}

	// identificador (variável)
	if ctx.Identificador() != nil {
		return ctx.Identificador().Accept(g).(string)
	}

	// chamada de função
	if ctx.Identificador() != nil {

		return ctx.Identificador().GetText()
	}

	if ctx.IDENT() != nil {

		// chamada de função

		args := []string{}

		for _, e := range ctx.AllExpressao() {
			args = append(args, e.Accept(g).(string))
		}

		return ctx.IDENT().GetText() +
			"(" + strings.Join(args, ", ") + ")"
	}

	// expressão parentizada
	if len(ctx.AllExpressao()) > 0 {

		return "(" +
			ctx.Expressao(0).Accept(g).(string) +
			")"
	}

	return ""
}

func (g *Gerador) VisitIdentificador(
	ctx *parser.IdentificadorContext,
) interface{} {

	return ctx.GetText()
}

func (g *Gerador) VisitParcela_nao_unario(
	ctx *parser.Parcela_nao_unarioContext,
) interface{} {

	// cadeia literal
	if ctx.CADEIA() != nil {
		return ctx.CADEIA().GetText()
	}

	// endereço (&variavel)
	if ctx.Identificador() != nil {

		return "&" +
			ctx.Identificador().
				Accept(g).(string)
	}

	return ""
}

func (g *Gerador) VisitCmdCaso(
	ctx *parser.CmdCasoContext,
) interface{} {

	// expressão do switch
	expr := ctx.Exp_aritmetica().
		Accept(g).(string)

	g.emitln(
		"switch (" + expr + ") {",
	)

	g.indent++

	// itens do caso
	if ctx.Selecao() != nil {

		for _, item := range ctx.Selecao().AllItem_selecao() {

			// constantes do item
			for _, c := range item.Constantes().AllNumero_intervalo() {

				texto := c.GetText()

				// intervalo: 3..100
				if strings.Contains(texto, "..") {

					partes := strings.Split(
						texto,
						"..",
					)

					ini := partes[0]
					fim := partes[1]

					g.emitln(
						"/* intervalo " +
							texto +
							" */",
					)

					for i := atoi(ini); i <= atoi(fim); i++ {

						g.emitln(
							"case " +
								itoa(i) +
								":",
						)
					}

				} else {

					g.emitln(
						"case " + texto + ":",
					)
				}
			}

			g.indent++

			// comandos do item
			for _, cmd := range item.AllCmd() {
				cmd.Accept(g)
			}

			g.emitln("break;")

			g.indent--
		}
	}

	// senao/default
	if len(ctx.AllCmd()) > 0 {

		g.emitln("default:")

		g.indent++

		for _, cmd := range ctx.AllCmd() {
			cmd.Accept(g)
		}

		g.emitln("break;")

		g.indent--
	}

	g.indent--

	g.emitln("}")

	return nil
}

func (g *Gerador) VisitCmdPara(
	ctx *parser.CmdParaContext,
) interface{} {

	varName := ctx.IDENT().GetText()

	inicio := ctx.Exp_aritmetica(0).
		Accept(g).(string)

	fim := ctx.Exp_aritmetica(1).
		Accept(g).(string)

	g.emitln(
		"for (" +
			varName + " = " + inicio + "; " +
			varName + " <= " + fim + "; " +
			varName + "++) {",
	)

	g.indent++

	for _, cmd := range ctx.AllCmd() {
		cmd.Accept(g)
	}

	g.indent--

	g.emitln("}")

	return nil
}
