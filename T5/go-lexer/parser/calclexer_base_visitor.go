// Code generated from CalcLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcLexer

import "github.com/antlr4-go/antlr/v4"

type BaseCalcLexerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcLexerVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitDeclaracoes(ctx *DeclaracoesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitDecl_local_global(ctx *Decl_local_globalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitDeclaracao_local(ctx *Declaracao_localContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitVariavel(ctx *VariavelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitIdentificador(ctx *IdentificadorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitDimensao(ctx *DimensaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitTipo_basico(ctx *Tipo_basicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitTipo_basico_ident(ctx *Tipo_basico_identContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitTipo_estendido(ctx *Tipo_estendidoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitValor_constante(ctx *Valor_constanteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitRegistro(ctx *RegistroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitDeclaracao_global(ctx *Declaracao_globalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCorpo(ctx *CorpoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitParametro(ctx *ParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmd(ctx *CmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdLeia(ctx *CmdLeiaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdEscreva(ctx *CmdEscrevaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdSe(ctx *CmdSeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdCaso(ctx *CmdCasoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdPara(ctx *CmdParaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdEnquanto(ctx *CmdEnquantoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdFaca(ctx *CmdFacaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdAtribuicao(ctx *CmdAtribuicaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdChamada(ctx *CmdChamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitCmdRetorne(ctx *CmdRetorneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitSelecao(ctx *SelecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitItem_selecao(ctx *Item_selecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitConstantes(ctx *ConstantesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitNumero_intervalo(ctx *Numero_intervaloContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitOp_unario(ctx *Op_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitExp_aritmetica(ctx *Exp_aritmeticaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitTermo(ctx *TermoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitFator(ctx *FatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitOp1(ctx *Op1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitOp2(ctx *Op2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitOp3(ctx *Op3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitParcela(ctx *ParcelaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitParcela_unario(ctx *Parcela_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitParcela_nao_unario(ctx *Parcela_nao_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitExp_relacional(ctx *Exp_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitOp_relacional(ctx *Op_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitExpressao(ctx *ExpressaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitTermo_logico(ctx *Termo_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitFator_logico(ctx *Fator_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitParcela_logica(ctx *Parcela_logicaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitOp_logico_1(ctx *Op_logico_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcLexerVisitor) VisitOp_logico_2(ctx *Op_logico_2Context) interface{} {
	return v.VisitChildren(ctx)
}
