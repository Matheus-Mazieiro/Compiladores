// Code generated from CalcLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcLexer

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalcLexerParser.
type CalcLexerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcLexerParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#declaracoes.
	VisitDeclaracoes(ctx *DeclaracoesContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#decl_local_global.
	VisitDecl_local_global(ctx *Decl_local_globalContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#declaracao_local.
	VisitDeclaracao_local(ctx *Declaracao_localContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#variavel.
	VisitVariavel(ctx *VariavelContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#identificador.
	VisitIdentificador(ctx *IdentificadorContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#dimensao.
	VisitDimensao(ctx *DimensaoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#tipo_basico.
	VisitTipo_basico(ctx *Tipo_basicoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#tipo_basico_ident.
	VisitTipo_basico_ident(ctx *Tipo_basico_identContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#tipo_estendido.
	VisitTipo_estendido(ctx *Tipo_estendidoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#valor_constante.
	VisitValor_constante(ctx *Valor_constanteContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#registro.
	VisitRegistro(ctx *RegistroContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#declaracao_global.
	VisitDeclaracao_global(ctx *Declaracao_globalContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#corpo.
	VisitCorpo(ctx *CorpoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#parametro.
	VisitParametro(ctx *ParametroContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmd.
	VisitCmd(ctx *CmdContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdLeia.
	VisitCmdLeia(ctx *CmdLeiaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdEscreva.
	VisitCmdEscreva(ctx *CmdEscrevaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdSe.
	VisitCmdSe(ctx *CmdSeContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdCaso.
	VisitCmdCaso(ctx *CmdCasoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdPara.
	VisitCmdPara(ctx *CmdParaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdEnquanto.
	VisitCmdEnquanto(ctx *CmdEnquantoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdFaca.
	VisitCmdFaca(ctx *CmdFacaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdAtribuicao.
	VisitCmdAtribuicao(ctx *CmdAtribuicaoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdChamada.
	VisitCmdChamada(ctx *CmdChamadaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#cmdRetorne.
	VisitCmdRetorne(ctx *CmdRetorneContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#selecao.
	VisitSelecao(ctx *SelecaoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#item_selecao.
	VisitItem_selecao(ctx *Item_selecaoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#constantes.
	VisitConstantes(ctx *ConstantesContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#numero_intervalo.
	VisitNumero_intervalo(ctx *Numero_intervaloContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#op_unario.
	VisitOp_unario(ctx *Op_unarioContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#exp_aritmetica.
	VisitExp_aritmetica(ctx *Exp_aritmeticaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#termo.
	VisitTermo(ctx *TermoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#fator.
	VisitFator(ctx *FatorContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#op1.
	VisitOp1(ctx *Op1Context) interface{}

	// Visit a parse tree produced by CalcLexerParser#op2.
	VisitOp2(ctx *Op2Context) interface{}

	// Visit a parse tree produced by CalcLexerParser#op3.
	VisitOp3(ctx *Op3Context) interface{}

	// Visit a parse tree produced by CalcLexerParser#parcela.
	VisitParcela(ctx *ParcelaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#parcela_unario.
	VisitParcela_unario(ctx *Parcela_unarioContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#parcela_nao_unario.
	VisitParcela_nao_unario(ctx *Parcela_nao_unarioContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#exp_relacional.
	VisitExp_relacional(ctx *Exp_relacionalContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#op_relacional.
	VisitOp_relacional(ctx *Op_relacionalContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#expressao.
	VisitExpressao(ctx *ExpressaoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#termo_logico.
	VisitTermo_logico(ctx *Termo_logicoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#fator_logico.
	VisitFator_logico(ctx *Fator_logicoContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#parcela_logica.
	VisitParcela_logica(ctx *Parcela_logicaContext) interface{}

	// Visit a parse tree produced by CalcLexerParser#op_logico_1.
	VisitOp_logico_1(ctx *Op_logico_1Context) interface{}

	// Visit a parse tree produced by CalcLexerParser#op_logico_2.
	VisitOp_logico_2(ctx *Op_logico_2Context) interface{}
}
