// LEMBRAR DE TIRAR PRINTFS DE DEBUG ANTES DE ENTREGAR O TRABALHO
package main

import (
	parser "go-lexer/parser"
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
func (j *JanderSemantico) VisitDecl_local_global(ctx *parser.Decl_local_globalContext) interface{} {
	if ctx.Declaracao_local() != nil {
		ctx.Declaracao_local().Accept(j)
	}
	return nil
}
func (j *JanderSemantico) VisitDeclaracao_local(ctx *parser.Declaracao_localContext) interface{} {
	// VisitDeclaracao_local

	if ctx.Variavel() != nil {
		ctx.Variavel().Accept(j)
		return nil
	}

	if ctx.IDENT() != nil {
		nome := ctx.IDENT().GetText()
		tipo := VerificarTipoBasico(ctx.Tipo_basico())

		if tipo == INVALIDO {
			AdicionarErroSemantico(ctx.Tipo_basico().GetStart().GetLine(),
				"tipo "+ctx.Tipo_basico().GetText()+" nao declarado")
		}

		if j.tabela.Existe(nome) {
			AdicionarErroSemantico(ctx.IDENT().GetSymbol().GetLine(),
				"identificador "+nome+" ja declarado anteriormente")
		} else {
			if tipo != INVALIDO {
				j.tabela.Adicionar(nome, tipo)
			}
		}
	}

	return nil
}
func (j *JanderSemantico) VisitDeclaracoes(ctx *parser.DeclaracoesContext) interface{} {
	for _, d := range ctx.AllDecl_local_global() {
		d.Accept(j)
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
func (j *JanderSemantico) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	j.tabela = NewTabelaDeSimbolos()

	if ctx.Declaracoes() != nil {
		ctx.Declaracoes().Accept(j)
	}

	// 🔥 VISITA CORPO
	if ctx.Corpo() != nil {
		ctx.Corpo().Accept(j)
	}

	return nil
}
func (j *JanderSemantico) VisitCmd(ctx *parser.CmdContext) interface{} {
	if ctx.CmdLeia() != nil {
		return ctx.CmdLeia().Accept(j)
	}
	if ctx.CmdEscreva() != nil {
		return ctx.CmdEscreva().Accept(j)
	}
	if ctx.CmdSe() != nil {
		return ctx.CmdSe().Accept(j)
	}
	if ctx.CmdEnquanto() != nil {
		return ctx.CmdEnquanto().Accept(j)
	}
	if ctx.CmdAtribuicao() != nil {
		return ctx.CmdAtribuicao().Accept(j)
	}

	return j.VisitChildren(ctx)
}

func (j *JanderSemantico) VisitVariavel(ctx *parser.VariavelContext) interface{} {
	tipo := VerificarTipo(j.tabela, ctx.Tipo())

	for _, ident := range ctx.AllIdentificador() {
		nome := ident.GetText()

		if j.tabela.Existe(nome) {
			AdicionarErroSemantico(ident.GetStart().GetLine(),
				"identificador "+nome+" ja declarado anteriormente")
		} else {
			if tipo != INVALIDO {
				j.tabela.Adicionar(nome, tipo)
			} else {
				// marca variável cujo tipo é inválido para evitar erros em cascata
				variaveisComTipoInvalido[nome] = true
			}
		}
	}

	return nil
}
func (j *JanderSemantico) VisitCmdLeia(ctx *parser.CmdLeiaContext) interface{} {
	for _, ident := range ctx.AllIdentificador() {
		nome := ident.GetText()

		if !j.tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(ident.GetStart().GetLine(),
					"identificador "+nome+" nao declarado")
			}
		}
	}

	return nil
}

func (j *JanderSemantico) VisitCmdEscreva(ctx *parser.CmdEscrevaContext) interface{} {
	for _, expr := range ctx.AllExpressao() {
		j.checkIdentificadores(expr.(antlr.Tree))
	}

	return nil
}

func (j *JanderSemantico) VisitCmdSe(ctx *parser.CmdSeContext) interface{} {
	if ctx == nil {
		return nil
	}

	// verifica identificadores na expressão de condição
	if ctx.Expressao() != nil {
		j.checkIdentificadores(ctx.Expressao().(antlr.Tree))
	}

	// visita todos os comandos do bloco (inclui SENAO quando presente)
	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	return nil
}

func (j *JanderSemantico) VisitCmdEnquanto(ctx *parser.CmdEnquantoContext) interface{} {
	if ctx == nil {
		return nil
	}

	if ctx.Expressao() != nil {
		j.checkIdentificadores(ctx.Expressao().(antlr.Tree))
	}

	for _, c := range ctx.AllCmd() {
		c.Accept(j)
	}

	return nil
}

func (j *JanderSemantico) VisitCmdAtribuicao(ctx *parser.CmdAtribuicaoContext) interface{} {
	if ctx == nil {
		return nil
	}

	ident := ctx.Identificador()
	if ident == nil {
		return nil
	}

	nome := ident.GetText()

	// se a variável tem tipo inválido previamente detectado, evite erros em cascata
	if variaveisComTipoInvalido[nome] {
		return nil
	}

	// tipo do lado esquerdo
	var left TipoJander
	if !j.tabela.Existe(nome) {
		if !ErroJaReportado(nome) {
			AdicionarErroSemantico(ident.GetStart().GetLine(), "identificador "+nome+" nao declarado")
		}
		left = INVALIDO
	} else {
		left = j.tabela.Verificar(nome)
	}

	// tipo do lado direito
	right := j.tipoExpressao(ctx.Expressao())

	// se o lado esquerdo é válido, reporta incompatibilidade quando o lado direito
	// é inválido (por mistura de tipos na expressão) ou quando os tipos não são compatíveis
	if left != INVALIDO {
		if right == INVALIDO || !j.tiposCompativeis(left, right) {
			AdicionarErroSemantico(ident.GetStart().GetLine(), "atribuicao nao compativel para "+nome)
		}
	}

	return nil
}

func (j *JanderSemantico) tiposCompativeis(dest, src TipoJander) bool {
	if dest == src {
		return true
	}
	// permite atribuir inteiro a real
	if dest == REAL && src == INTEIRO {
		return true
	}
	return false
}

func (j *JanderSemantico) tipoExpressao(ctx parser.IExpressaoContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	termos := ctx.AllTermo_logico()
	if len(termos) == 0 {
		return INVALIDO
	}

	// se houver operador 'ou', toda expressão deve ser lógica
	if len(ctx.AllOp_logico_1()) > 0 {
		for _, t := range termos {
			if j.tipoTermoLogico(t) != LOGICO {
				return INVALIDO
			}
		}
		return LOGICO
	}

	return j.tipoTermoLogico(termos[0])
}

func (j *JanderSemantico) tipoTermoLogico(ctx parser.ITermo_logicoContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	fatores := ctx.AllFator_logico()
	if len(ctx.AllOp_logico_2()) > 0 {
		for _, f := range fatores {
			if j.tipoFatorLogico(f) != LOGICO {
				return INVALIDO
			}
		}
		return LOGICO
	}

	return j.tipoFatorLogico(fatores[0])
}

func (j *JanderSemantico) tipoFatorLogico(ctx parser.IFator_logicoContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}
	return j.tipoParcelaLogica(ctx.Parcela_logica())
}

func (j *JanderSemantico) tipoParcelaLogica(ctx parser.IParcela_logicaContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}
	if ctx.VERDADEIRO() != nil || ctx.FALSO() != nil {
		return LOGICO
	}
	return j.tipoExpRelacional(ctx.Exp_relacional())
}

func (j *JanderSemantico) tipoExpRelacional(ctx parser.IExp_relacionalContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	exps := ctx.AllExp_aritmetica()
	if len(exps) == 0 {
		return INVALIDO
	}

	if ctx.Op_relacional() != nil {
		t1 := j.tipoExpAritmetica(exps[0])
		t2 := j.tipoExpAritmetica(exps[1])
		if t1 == INVALIDO || t2 == INVALIDO {
			return INVALIDO
		}
		// comparações: números entre si ou literais entre si
		if (t1 == INTEIRO || t1 == REAL) && (t2 == INTEIRO || t2 == REAL) {
			return LOGICO
		}
		if t1 == LITERAL && t2 == LITERAL {
			return LOGICO
		}
		return INVALIDO
	}

	return j.tipoExpAritmetica(exps[0])
}

func (j *JanderSemantico) tipoExpAritmetica(ctx parser.IExp_aritmeticaContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	termos := ctx.AllTermo()
	if len(termos) == 0 {
		return INVALIDO
	}

	t := j.tipoTermo(termos[0])
	// se houver operadores '+' ou '-', combinar tipos
	for i := 1; i < len(termos); i++ {
		t2 := j.tipoTermo(termos[i])
		// se qualquer um inválido, propaga
		if t == INVALIDO || t2 == INVALIDO {
			return INVALIDO
		}
		// se um for literal, os dois devem ser literal (concatenação)
		if t == LITERAL || t2 == LITERAL {
			if t == LITERAL && t2 == LITERAL {
				t = LITERAL
			} else {
				return INVALIDO
			}
		} else {
			// aritmética numérica: real domina
			if t == REAL || t2 == REAL {
				t = REAL
			} else {
				t = INTEIRO
			}
		}
	}

	return t
}

func (j *JanderSemantico) tipoTermo(ctx parser.ITermoContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	fatores := ctx.AllFator()
	if len(fatores) == 0 {
		return INVALIDO
	}

	t := j.tipoFator(fatores[0])
	for i := 1; i < len(fatores); i++ {
		t2 := j.tipoFator(fatores[i])
		if t == INVALIDO || t2 == INVALIDO {
			return INVALIDO
		}
		// multiplicação/divisão -> se qualquer real, resultado real
		if t == REAL || t2 == REAL {
			t = REAL
		} else {
			t = INTEIRO
		}
	}

	return t
}

func (j *JanderSemantico) tipoFator(ctx parser.IFatorContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	parcelas := ctx.AllParcela()
	if len(parcelas) == 0 {
		return INVALIDO
	}

	t := j.tipoParcela(parcelas[0])
	for i := 1; i < len(parcelas); i++ {
		t2 := j.tipoParcela(parcelas[i])
		if t == INVALIDO || t2 == INVALIDO {
			return INVALIDO
		}
		// operador '%' exige inteiros
		if t == INTEIRO && t2 == INTEIRO {
			t = INTEIRO
		} else if t == LITERAL || t2 == LITERAL {
			return INVALIDO
		} else if t == REAL || t2 == REAL {
			t = REAL
		} else {
			t = INTEIRO
		}
	}

	return t
}

func (j *JanderSemantico) tipoParcela(ctx parser.IParcelaContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}
	if ctx.Parcela_unario() != nil {
		return j.tipoParcelaUnario(ctx.Parcela_unario())
	}
	return j.tipoParcelaNaoUnario(ctx.Parcela_nao_unario())
}

func (j *JanderSemantico) tipoParcelaUnario(ctx parser.IParcela_unarioContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}

	if ctx.Identificador() != nil {
		nome := ctx.Identificador().GetText()
		if variaveisComTipoInvalido[nome] {
			return INVALIDO
		}
		if !j.tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(ctx.GetStart().GetLine(), "identificador "+nome+" nao declarado")
			}
			return INVALIDO
		}
		return j.tabela.Verificar(nome)
	}

	if ctx.IDENT() != nil {
		// chamada de função: verificar se existe na tabela e retornar seu tipo
		nome := ctx.IDENT().GetText()
		if !j.tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(ctx.GetStart().GetLine(), "identificador "+nome+" nao declarado")
			}
			return INVALIDO
		}
		return j.tabela.Verificar(nome)
	}

	if ctx.NUM_INT() != nil {
		return INTEIRO
	}
	if ctx.NUM_REAL() != nil {
		return REAL
	}

	// '(' expressao ')'
	if ctx.Expressao(0) != nil {
		return j.tipoExpressao(ctx.Expressao(0))
	}

	return INVALIDO
}

func (j *JanderSemantico) tipoParcelaNaoUnario(ctx parser.IParcela_nao_unarioContext) TipoJander {
	if ctx == nil {
		return INVALIDO
	}
	if ctx.CADEIA() != nil {
		return LITERAL
	}
	if ctx.Identificador() != nil {
		nome := ctx.Identificador().GetText()
		if variaveisComTipoInvalido[nome] {
			return INVALIDO
		}
		if !j.tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(ctx.GetStart().GetLine(), "identificador "+nome+" nao declarado")
			}
			return INVALIDO
		}
		return j.tabela.Verificar(nome)
	}
	return INVALIDO
}

func (j *JanderSemantico) checkIdentificadores(t antlr.Tree) {
	// se for um contexto de identificador, verifica na tabela
	if ident, ok := t.(parser.IIdentificadorContext); ok {
		nome := ident.GetText()
		// checa identificador
		if variaveisComTipoInvalido[nome] {
			return
		}

		if !j.tabela.Existe(nome) {
			if !ErroJaReportado(nome) {
				AdicionarErroSemantico(ident.GetStart().GetLine(),
					"identificador "+nome+" nao declarado")
			}
		}

		return
	}

	// caso contrário, percorre os filhos recursivamente
	for _, ch := range t.GetChildren() {
		j.checkIdentificadores(ch)
	}
}

func (j *JanderSemantico) VisitIdentificador(ctx *parser.IdentificadorContext) interface{} {
	nome := ctx.GetText()

	if variaveisComTipoInvalido[nome] {
		return nil
	}

	if !j.tabela.Existe(nome) {
		if !ErroJaReportado(nome) {
			AdicionarErroSemantico(ctx.GetStart().GetLine(),
				"identificador "+nome+" nao declarado")
		}
	}

	return nil
}
