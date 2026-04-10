// Code generated from CalcLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcLexer

import "github.com/antlr4-go/antlr/v4"

// BaseCalcLexerListener is a complete listener for a parse tree produced by CalcLexerParser.
type BaseCalcLexerListener struct{}

var _ CalcLexerListener = &BaseCalcLexerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcLexerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcLexerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcLexerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcLexerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseCalcLexerListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseCalcLexerListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterDeclaracoes is called when production declaracoes is entered.
func (s *BaseCalcLexerListener) EnterDeclaracoes(ctx *DeclaracoesContext) {}

// ExitDeclaracoes is called when production declaracoes is exited.
func (s *BaseCalcLexerListener) ExitDeclaracoes(ctx *DeclaracoesContext) {}

// EnterDecl_local_global is called when production decl_local_global is entered.
func (s *BaseCalcLexerListener) EnterDecl_local_global(ctx *Decl_local_globalContext) {}

// ExitDecl_local_global is called when production decl_local_global is exited.
func (s *BaseCalcLexerListener) ExitDecl_local_global(ctx *Decl_local_globalContext) {}

// EnterDeclaracao_local is called when production declaracao_local is entered.
func (s *BaseCalcLexerListener) EnterDeclaracao_local(ctx *Declaracao_localContext) {}

// ExitDeclaracao_local is called when production declaracao_local is exited.
func (s *BaseCalcLexerListener) ExitDeclaracao_local(ctx *Declaracao_localContext) {}

// EnterVariavel is called when production variavel is entered.
func (s *BaseCalcLexerListener) EnterVariavel(ctx *VariavelContext) {}

// ExitVariavel is called when production variavel is exited.
func (s *BaseCalcLexerListener) ExitVariavel(ctx *VariavelContext) {}

// EnterIdentificador is called when production identificador is entered.
func (s *BaseCalcLexerListener) EnterIdentificador(ctx *IdentificadorContext) {}

// ExitIdentificador is called when production identificador is exited.
func (s *BaseCalcLexerListener) ExitIdentificador(ctx *IdentificadorContext) {}

// EnterDimensao is called when production dimensao is entered.
func (s *BaseCalcLexerListener) EnterDimensao(ctx *DimensaoContext) {}

// ExitDimensao is called when production dimensao is exited.
func (s *BaseCalcLexerListener) ExitDimensao(ctx *DimensaoContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseCalcLexerListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseCalcLexerListener) ExitTipo(ctx *TipoContext) {}

// EnterTipo_basico is called when production tipo_basico is entered.
func (s *BaseCalcLexerListener) EnterTipo_basico(ctx *Tipo_basicoContext) {}

// ExitTipo_basico is called when production tipo_basico is exited.
func (s *BaseCalcLexerListener) ExitTipo_basico(ctx *Tipo_basicoContext) {}

// EnterTipo_basico_ident is called when production tipo_basico_ident is entered.
func (s *BaseCalcLexerListener) EnterTipo_basico_ident(ctx *Tipo_basico_identContext) {}

// ExitTipo_basico_ident is called when production tipo_basico_ident is exited.
func (s *BaseCalcLexerListener) ExitTipo_basico_ident(ctx *Tipo_basico_identContext) {}

// EnterTipo_estendido is called when production tipo_estendido is entered.
func (s *BaseCalcLexerListener) EnterTipo_estendido(ctx *Tipo_estendidoContext) {}

// ExitTipo_estendido is called when production tipo_estendido is exited.
func (s *BaseCalcLexerListener) ExitTipo_estendido(ctx *Tipo_estendidoContext) {}

// EnterValor_constante is called when production valor_constante is entered.
func (s *BaseCalcLexerListener) EnterValor_constante(ctx *Valor_constanteContext) {}

// ExitValor_constante is called when production valor_constante is exited.
func (s *BaseCalcLexerListener) ExitValor_constante(ctx *Valor_constanteContext) {}

// EnterRegistro is called when production registro is entered.
func (s *BaseCalcLexerListener) EnterRegistro(ctx *RegistroContext) {}

// ExitRegistro is called when production registro is exited.
func (s *BaseCalcLexerListener) ExitRegistro(ctx *RegistroContext) {}

// EnterDeclaracao_global is called when production declaracao_global is entered.
func (s *BaseCalcLexerListener) EnterDeclaracao_global(ctx *Declaracao_globalContext) {}

// ExitDeclaracao_global is called when production declaracao_global is exited.
func (s *BaseCalcLexerListener) ExitDeclaracao_global(ctx *Declaracao_globalContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseCalcLexerListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseCalcLexerListener) ExitParametros(ctx *ParametrosContext) {}

// EnterCorpo is called when production corpo is entered.
func (s *BaseCalcLexerListener) EnterCorpo(ctx *CorpoContext) {}

// ExitCorpo is called when production corpo is exited.
func (s *BaseCalcLexerListener) ExitCorpo(ctx *CorpoContext) {}

// EnterParametro is called when production parametro is entered.
func (s *BaseCalcLexerListener) EnterParametro(ctx *ParametroContext) {}

// ExitParametro is called when production parametro is exited.
func (s *BaseCalcLexerListener) ExitParametro(ctx *ParametroContext) {}

// EnterCmd is called when production cmd is entered.
func (s *BaseCalcLexerListener) EnterCmd(ctx *CmdContext) {}

// ExitCmd is called when production cmd is exited.
func (s *BaseCalcLexerListener) ExitCmd(ctx *CmdContext) {}

// EnterCmdLeia is called when production cmdLeia is entered.
func (s *BaseCalcLexerListener) EnterCmdLeia(ctx *CmdLeiaContext) {}

// ExitCmdLeia is called when production cmdLeia is exited.
func (s *BaseCalcLexerListener) ExitCmdLeia(ctx *CmdLeiaContext) {}

// EnterCmdEscreva is called when production cmdEscreva is entered.
func (s *BaseCalcLexerListener) EnterCmdEscreva(ctx *CmdEscrevaContext) {}

// ExitCmdEscreva is called when production cmdEscreva is exited.
func (s *BaseCalcLexerListener) ExitCmdEscreva(ctx *CmdEscrevaContext) {}

// EnterCmdSe is called when production cmdSe is entered.
func (s *BaseCalcLexerListener) EnterCmdSe(ctx *CmdSeContext) {}

// ExitCmdSe is called when production cmdSe is exited.
func (s *BaseCalcLexerListener) ExitCmdSe(ctx *CmdSeContext) {}

// EnterCmdCaso is called when production cmdCaso is entered.
func (s *BaseCalcLexerListener) EnterCmdCaso(ctx *CmdCasoContext) {}

// ExitCmdCaso is called when production cmdCaso is exited.
func (s *BaseCalcLexerListener) ExitCmdCaso(ctx *CmdCasoContext) {}

// EnterCmdPara is called when production cmdPara is entered.
func (s *BaseCalcLexerListener) EnterCmdPara(ctx *CmdParaContext) {}

// ExitCmdPara is called when production cmdPara is exited.
func (s *BaseCalcLexerListener) ExitCmdPara(ctx *CmdParaContext) {}

// EnterCmdEnquanto is called when production cmdEnquanto is entered.
func (s *BaseCalcLexerListener) EnterCmdEnquanto(ctx *CmdEnquantoContext) {}

// ExitCmdEnquanto is called when production cmdEnquanto is exited.
func (s *BaseCalcLexerListener) ExitCmdEnquanto(ctx *CmdEnquantoContext) {}

// EnterCmdFaca is called when production cmdFaca is entered.
func (s *BaseCalcLexerListener) EnterCmdFaca(ctx *CmdFacaContext) {}

// ExitCmdFaca is called when production cmdFaca is exited.
func (s *BaseCalcLexerListener) ExitCmdFaca(ctx *CmdFacaContext) {}

// EnterCmdAtribuicao is called when production cmdAtribuicao is entered.
func (s *BaseCalcLexerListener) EnterCmdAtribuicao(ctx *CmdAtribuicaoContext) {}

// ExitCmdAtribuicao is called when production cmdAtribuicao is exited.
func (s *BaseCalcLexerListener) ExitCmdAtribuicao(ctx *CmdAtribuicaoContext) {}

// EnterCmdChamada is called when production cmdChamada is entered.
func (s *BaseCalcLexerListener) EnterCmdChamada(ctx *CmdChamadaContext) {}

// ExitCmdChamada is called when production cmdChamada is exited.
func (s *BaseCalcLexerListener) ExitCmdChamada(ctx *CmdChamadaContext) {}

// EnterCmdRetorne is called when production cmdRetorne is entered.
func (s *BaseCalcLexerListener) EnterCmdRetorne(ctx *CmdRetorneContext) {}

// ExitCmdRetorne is called when production cmdRetorne is exited.
func (s *BaseCalcLexerListener) ExitCmdRetorne(ctx *CmdRetorneContext) {}

// EnterSelecao is called when production selecao is entered.
func (s *BaseCalcLexerListener) EnterSelecao(ctx *SelecaoContext) {}

// ExitSelecao is called when production selecao is exited.
func (s *BaseCalcLexerListener) ExitSelecao(ctx *SelecaoContext) {}

// EnterItem_selecao is called when production item_selecao is entered.
func (s *BaseCalcLexerListener) EnterItem_selecao(ctx *Item_selecaoContext) {}

// ExitItem_selecao is called when production item_selecao is exited.
func (s *BaseCalcLexerListener) ExitItem_selecao(ctx *Item_selecaoContext) {}

// EnterConstantes is called when production constantes is entered.
func (s *BaseCalcLexerListener) EnterConstantes(ctx *ConstantesContext) {}

// ExitConstantes is called when production constantes is exited.
func (s *BaseCalcLexerListener) ExitConstantes(ctx *ConstantesContext) {}

// EnterNumero_intervalo is called when production numero_intervalo is entered.
func (s *BaseCalcLexerListener) EnterNumero_intervalo(ctx *Numero_intervaloContext) {}

// ExitNumero_intervalo is called when production numero_intervalo is exited.
func (s *BaseCalcLexerListener) ExitNumero_intervalo(ctx *Numero_intervaloContext) {}

// EnterOp_unario is called when production op_unario is entered.
func (s *BaseCalcLexerListener) EnterOp_unario(ctx *Op_unarioContext) {}

// ExitOp_unario is called when production op_unario is exited.
func (s *BaseCalcLexerListener) ExitOp_unario(ctx *Op_unarioContext) {}

// EnterExp_aritmetica is called when production exp_aritmetica is entered.
func (s *BaseCalcLexerListener) EnterExp_aritmetica(ctx *Exp_aritmeticaContext) {}

// ExitExp_aritmetica is called when production exp_aritmetica is exited.
func (s *BaseCalcLexerListener) ExitExp_aritmetica(ctx *Exp_aritmeticaContext) {}

// EnterTermo is called when production termo is entered.
func (s *BaseCalcLexerListener) EnterTermo(ctx *TermoContext) {}

// ExitTermo is called when production termo is exited.
func (s *BaseCalcLexerListener) ExitTermo(ctx *TermoContext) {}

// EnterFator is called when production fator is entered.
func (s *BaseCalcLexerListener) EnterFator(ctx *FatorContext) {}

// ExitFator is called when production fator is exited.
func (s *BaseCalcLexerListener) ExitFator(ctx *FatorContext) {}

// EnterOp1 is called when production op1 is entered.
func (s *BaseCalcLexerListener) EnterOp1(ctx *Op1Context) {}

// ExitOp1 is called when production op1 is exited.
func (s *BaseCalcLexerListener) ExitOp1(ctx *Op1Context) {}

// EnterOp2 is called when production op2 is entered.
func (s *BaseCalcLexerListener) EnterOp2(ctx *Op2Context) {}

// ExitOp2 is called when production op2 is exited.
func (s *BaseCalcLexerListener) ExitOp2(ctx *Op2Context) {}

// EnterOp3 is called when production op3 is entered.
func (s *BaseCalcLexerListener) EnterOp3(ctx *Op3Context) {}

// ExitOp3 is called when production op3 is exited.
func (s *BaseCalcLexerListener) ExitOp3(ctx *Op3Context) {}

// EnterParcela is called when production parcela is entered.
func (s *BaseCalcLexerListener) EnterParcela(ctx *ParcelaContext) {}

// ExitParcela is called when production parcela is exited.
func (s *BaseCalcLexerListener) ExitParcela(ctx *ParcelaContext) {}

// EnterParcela_unario is called when production parcela_unario is entered.
func (s *BaseCalcLexerListener) EnterParcela_unario(ctx *Parcela_unarioContext) {}

// ExitParcela_unario is called when production parcela_unario is exited.
func (s *BaseCalcLexerListener) ExitParcela_unario(ctx *Parcela_unarioContext) {}

// EnterParcela_nao_unario is called when production parcela_nao_unario is entered.
func (s *BaseCalcLexerListener) EnterParcela_nao_unario(ctx *Parcela_nao_unarioContext) {}

// ExitParcela_nao_unario is called when production parcela_nao_unario is exited.
func (s *BaseCalcLexerListener) ExitParcela_nao_unario(ctx *Parcela_nao_unarioContext) {}

// EnterExp_relacional is called when production exp_relacional is entered.
func (s *BaseCalcLexerListener) EnterExp_relacional(ctx *Exp_relacionalContext) {}

// ExitExp_relacional is called when production exp_relacional is exited.
func (s *BaseCalcLexerListener) ExitExp_relacional(ctx *Exp_relacionalContext) {}

// EnterOp_relacional is called when production op_relacional is entered.
func (s *BaseCalcLexerListener) EnterOp_relacional(ctx *Op_relacionalContext) {}

// ExitOp_relacional is called when production op_relacional is exited.
func (s *BaseCalcLexerListener) ExitOp_relacional(ctx *Op_relacionalContext) {}

// EnterExpressao is called when production expressao is entered.
func (s *BaseCalcLexerListener) EnterExpressao(ctx *ExpressaoContext) {}

// ExitExpressao is called when production expressao is exited.
func (s *BaseCalcLexerListener) ExitExpressao(ctx *ExpressaoContext) {}

// EnterTermo_logico is called when production termo_logico is entered.
func (s *BaseCalcLexerListener) EnterTermo_logico(ctx *Termo_logicoContext) {}

// ExitTermo_logico is called when production termo_logico is exited.
func (s *BaseCalcLexerListener) ExitTermo_logico(ctx *Termo_logicoContext) {}

// EnterFator_logico is called when production fator_logico is entered.
func (s *BaseCalcLexerListener) EnterFator_logico(ctx *Fator_logicoContext) {}

// ExitFator_logico is called when production fator_logico is exited.
func (s *BaseCalcLexerListener) ExitFator_logico(ctx *Fator_logicoContext) {}

// EnterParcela_logica is called when production parcela_logica is entered.
func (s *BaseCalcLexerListener) EnterParcela_logica(ctx *Parcela_logicaContext) {}

// ExitParcela_logica is called when production parcela_logica is exited.
func (s *BaseCalcLexerListener) ExitParcela_logica(ctx *Parcela_logicaContext) {}

// EnterOp_logico_1 is called when production op_logico_1 is entered.
func (s *BaseCalcLexerListener) EnterOp_logico_1(ctx *Op_logico_1Context) {}

// ExitOp_logico_1 is called when production op_logico_1 is exited.
func (s *BaseCalcLexerListener) ExitOp_logico_1(ctx *Op_logico_1Context) {}

// EnterOp_logico_2 is called when production op_logico_2 is entered.
func (s *BaseCalcLexerListener) EnterOp_logico_2(ctx *Op_logico_2Context) {}

// ExitOp_logico_2 is called when production op_logico_2 is exited.
func (s *BaseCalcLexerListener) ExitOp_logico_2(ctx *Op_logico_2Context) {}
