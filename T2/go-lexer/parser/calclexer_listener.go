// Code generated from CalcLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcLexer

import "github.com/antlr4-go/antlr/v4"

// CalcLexerListener is a complete listener for a parse tree produced by CalcLexerParser.
type CalcLexerListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterDeclaracoes is called when entering the declaracoes production.
	EnterDeclaracoes(c *DeclaracoesContext)

	// EnterDecl_local_global is called when entering the decl_local_global production.
	EnterDecl_local_global(c *Decl_local_globalContext)

	// EnterDeclaracao_local is called when entering the declaracao_local production.
	EnterDeclaracao_local(c *Declaracao_localContext)

	// EnterVariavel is called when entering the variavel production.
	EnterVariavel(c *VariavelContext)

	// EnterIdentificador is called when entering the identificador production.
	EnterIdentificador(c *IdentificadorContext)

	// EnterDimensao is called when entering the dimensao production.
	EnterDimensao(c *DimensaoContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterTipo_basico is called when entering the tipo_basico production.
	EnterTipo_basico(c *Tipo_basicoContext)

	// EnterTipo_basico_ident is called when entering the tipo_basico_ident production.
	EnterTipo_basico_ident(c *Tipo_basico_identContext)

	// EnterTipo_estendido is called when entering the tipo_estendido production.
	EnterTipo_estendido(c *Tipo_estendidoContext)

	// EnterValor_constante is called when entering the valor_constante production.
	EnterValor_constante(c *Valor_constanteContext)

	// EnterRegistro is called when entering the registro production.
	EnterRegistro(c *RegistroContext)

	// EnterDeclaracao_global is called when entering the declaracao_global production.
	EnterDeclaracao_global(c *Declaracao_globalContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterCorpo is called when entering the corpo production.
	EnterCorpo(c *CorpoContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterCmd is called when entering the cmd production.
	EnterCmd(c *CmdContext)

	// EnterCmdLeia is called when entering the cmdLeia production.
	EnterCmdLeia(c *CmdLeiaContext)

	// EnterCmdEscreva is called when entering the cmdEscreva production.
	EnterCmdEscreva(c *CmdEscrevaContext)

	// EnterCmdSe is called when entering the cmdSe production.
	EnterCmdSe(c *CmdSeContext)

	// EnterCmdCaso is called when entering the cmdCaso production.
	EnterCmdCaso(c *CmdCasoContext)

	// EnterCmdPara is called when entering the cmdPara production.
	EnterCmdPara(c *CmdParaContext)

	// EnterCmdEnquanto is called when entering the cmdEnquanto production.
	EnterCmdEnquanto(c *CmdEnquantoContext)

	// EnterCmdFaca is called when entering the cmdFaca production.
	EnterCmdFaca(c *CmdFacaContext)

	// EnterCmdAtribuicao is called when entering the cmdAtribuicao production.
	EnterCmdAtribuicao(c *CmdAtribuicaoContext)

	// EnterCmdChamada is called when entering the cmdChamada production.
	EnterCmdChamada(c *CmdChamadaContext)

	// EnterCmdRetorne is called when entering the cmdRetorne production.
	EnterCmdRetorne(c *CmdRetorneContext)

	// EnterSelecao is called when entering the selecao production.
	EnterSelecao(c *SelecaoContext)

	// EnterItem_selecao is called when entering the item_selecao production.
	EnterItem_selecao(c *Item_selecaoContext)

	// EnterConstantes is called when entering the constantes production.
	EnterConstantes(c *ConstantesContext)

	// EnterNumero_intervalo is called when entering the numero_intervalo production.
	EnterNumero_intervalo(c *Numero_intervaloContext)

	// EnterOp_unario is called when entering the op_unario production.
	EnterOp_unario(c *Op_unarioContext)

	// EnterExp_aritmetica is called when entering the exp_aritmetica production.
	EnterExp_aritmetica(c *Exp_aritmeticaContext)

	// EnterTermo is called when entering the termo production.
	EnterTermo(c *TermoContext)

	// EnterFator is called when entering the fator production.
	EnterFator(c *FatorContext)

	// EnterOp1 is called when entering the op1 production.
	EnterOp1(c *Op1Context)

	// EnterOp2 is called when entering the op2 production.
	EnterOp2(c *Op2Context)

	// EnterOp3 is called when entering the op3 production.
	EnterOp3(c *Op3Context)

	// EnterParcela is called when entering the parcela production.
	EnterParcela(c *ParcelaContext)

	// EnterParcela_unario is called when entering the parcela_unario production.
	EnterParcela_unario(c *Parcela_unarioContext)

	// EnterParcela_nao_unario is called when entering the parcela_nao_unario production.
	EnterParcela_nao_unario(c *Parcela_nao_unarioContext)

	// EnterExp_relacional is called when entering the exp_relacional production.
	EnterExp_relacional(c *Exp_relacionalContext)

	// EnterOp_relacional is called when entering the op_relacional production.
	EnterOp_relacional(c *Op_relacionalContext)

	// EnterExpressao is called when entering the expressao production.
	EnterExpressao(c *ExpressaoContext)

	// EnterTermo_logico is called when entering the termo_logico production.
	EnterTermo_logico(c *Termo_logicoContext)

	// EnterFator_logico is called when entering the fator_logico production.
	EnterFator_logico(c *Fator_logicoContext)

	// EnterParcela_logica is called when entering the parcela_logica production.
	EnterParcela_logica(c *Parcela_logicaContext)

	// EnterOp_logico_1 is called when entering the op_logico_1 production.
	EnterOp_logico_1(c *Op_logico_1Context)

	// EnterOp_logico_2 is called when entering the op_logico_2 production.
	EnterOp_logico_2(c *Op_logico_2Context)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitDeclaracoes is called when exiting the declaracoes production.
	ExitDeclaracoes(c *DeclaracoesContext)

	// ExitDecl_local_global is called when exiting the decl_local_global production.
	ExitDecl_local_global(c *Decl_local_globalContext)

	// ExitDeclaracao_local is called when exiting the declaracao_local production.
	ExitDeclaracao_local(c *Declaracao_localContext)

	// ExitVariavel is called when exiting the variavel production.
	ExitVariavel(c *VariavelContext)

	// ExitIdentificador is called when exiting the identificador production.
	ExitIdentificador(c *IdentificadorContext)

	// ExitDimensao is called when exiting the dimensao production.
	ExitDimensao(c *DimensaoContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitTipo_basico is called when exiting the tipo_basico production.
	ExitTipo_basico(c *Tipo_basicoContext)

	// ExitTipo_basico_ident is called when exiting the tipo_basico_ident production.
	ExitTipo_basico_ident(c *Tipo_basico_identContext)

	// ExitTipo_estendido is called when exiting the tipo_estendido production.
	ExitTipo_estendido(c *Tipo_estendidoContext)

	// ExitValor_constante is called when exiting the valor_constante production.
	ExitValor_constante(c *Valor_constanteContext)

	// ExitRegistro is called when exiting the registro production.
	ExitRegistro(c *RegistroContext)

	// ExitDeclaracao_global is called when exiting the declaracao_global production.
	ExitDeclaracao_global(c *Declaracao_globalContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitCorpo is called when exiting the corpo production.
	ExitCorpo(c *CorpoContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitCmd is called when exiting the cmd production.
	ExitCmd(c *CmdContext)

	// ExitCmdLeia is called when exiting the cmdLeia production.
	ExitCmdLeia(c *CmdLeiaContext)

	// ExitCmdEscreva is called when exiting the cmdEscreva production.
	ExitCmdEscreva(c *CmdEscrevaContext)

	// ExitCmdSe is called when exiting the cmdSe production.
	ExitCmdSe(c *CmdSeContext)

	// ExitCmdCaso is called when exiting the cmdCaso production.
	ExitCmdCaso(c *CmdCasoContext)

	// ExitCmdPara is called when exiting the cmdPara production.
	ExitCmdPara(c *CmdParaContext)

	// ExitCmdEnquanto is called when exiting the cmdEnquanto production.
	ExitCmdEnquanto(c *CmdEnquantoContext)

	// ExitCmdFaca is called when exiting the cmdFaca production.
	ExitCmdFaca(c *CmdFacaContext)

	// ExitCmdAtribuicao is called when exiting the cmdAtribuicao production.
	ExitCmdAtribuicao(c *CmdAtribuicaoContext)

	// ExitCmdChamada is called when exiting the cmdChamada production.
	ExitCmdChamada(c *CmdChamadaContext)

	// ExitCmdRetorne is called when exiting the cmdRetorne production.
	ExitCmdRetorne(c *CmdRetorneContext)

	// ExitSelecao is called when exiting the selecao production.
	ExitSelecao(c *SelecaoContext)

	// ExitItem_selecao is called when exiting the item_selecao production.
	ExitItem_selecao(c *Item_selecaoContext)

	// ExitConstantes is called when exiting the constantes production.
	ExitConstantes(c *ConstantesContext)

	// ExitNumero_intervalo is called when exiting the numero_intervalo production.
	ExitNumero_intervalo(c *Numero_intervaloContext)

	// ExitOp_unario is called when exiting the op_unario production.
	ExitOp_unario(c *Op_unarioContext)

	// ExitExp_aritmetica is called when exiting the exp_aritmetica production.
	ExitExp_aritmetica(c *Exp_aritmeticaContext)

	// ExitTermo is called when exiting the termo production.
	ExitTermo(c *TermoContext)

	// ExitFator is called when exiting the fator production.
	ExitFator(c *FatorContext)

	// ExitOp1 is called when exiting the op1 production.
	ExitOp1(c *Op1Context)

	// ExitOp2 is called when exiting the op2 production.
	ExitOp2(c *Op2Context)

	// ExitOp3 is called when exiting the op3 production.
	ExitOp3(c *Op3Context)

	// ExitParcela is called when exiting the parcela production.
	ExitParcela(c *ParcelaContext)

	// ExitParcela_unario is called when exiting the parcela_unario production.
	ExitParcela_unario(c *Parcela_unarioContext)

	// ExitParcela_nao_unario is called when exiting the parcela_nao_unario production.
	ExitParcela_nao_unario(c *Parcela_nao_unarioContext)

	// ExitExp_relacional is called when exiting the exp_relacional production.
	ExitExp_relacional(c *Exp_relacionalContext)

	// ExitOp_relacional is called when exiting the op_relacional production.
	ExitOp_relacional(c *Op_relacionalContext)

	// ExitExpressao is called when exiting the expressao production.
	ExitExpressao(c *ExpressaoContext)

	// ExitTermo_logico is called when exiting the termo_logico production.
	ExitTermo_logico(c *Termo_logicoContext)

	// ExitFator_logico is called when exiting the fator_logico production.
	ExitFator_logico(c *Fator_logicoContext)

	// ExitParcela_logica is called when exiting the parcela_logica production.
	ExitParcela_logica(c *Parcela_logicaContext)

	// ExitOp_logico_1 is called when exiting the op_logico_1 production.
	ExitOp_logico_1(c *Op_logico_1Context)

	// ExitOp_logico_2 is called when exiting the op_logico_2 production.
	ExitOp_logico_2(c *Op_logico_2Context)
}
