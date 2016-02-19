
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(4),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			shift(6),		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			reduce(17),		/* lbrace, reduce: VarId */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Form */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(13),		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(17),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			reduce(8),		/* ident, reduce: StrLiteral */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(43),		/* str_lit, reduce: Stmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(43),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(43),		/* rbrace, reduce: Stmt */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(18),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(42),		/* str_lit, reduce: Stmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(42),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(42),		/* rbrace, reduce: Stmt */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(19),		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(39),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(40),		/* str_lit, reduce: StmtList */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(40),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(40),		/* rbrace, reduce: StmtList */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			shift(21),		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(36),		/* str_lit, reduce: Question */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(36),		/* if, reduce: Question */
			nil,		/* else */
			shift(22),		/* assign */
			nil,		/* lbrace */
			reduce(36),		/* rbrace, reduce: Question */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			reduce(17),		/* col, reduce: VarId */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(38),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(41),		/* str_lit, reduce: StmtList */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(41),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(41),		/* rbrace, reduce: StmtList */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(38),		/* integer */
			shift(39),		/* boolean */
			shift(40),		/* string */
			shift(41),		/* money */
			shift(42),		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			reduce(13),		/* mulop, reduce: Factor */
			reduce(13),		/* divop, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(13),		/* addop, reduce: Factor */
			reduce(13),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(13),		/* eqop, reduce: Factor */
			reduce(13),		/* neqop, reduce: Factor */
			reduce(13),		/* gtop, reduce: Factor */
			reduce(13),		/* ltop, reduce: Factor */
			reduce(13),		/* geqop, reduce: Factor */
			reduce(13),		/* leqop, reduce: Factor */
			reduce(13),		/* andop, reduce: Factor */
			reduce(13),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(13),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			reduce(12),		/* mulop, reduce: Factor */
			reduce(12),		/* divop, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(12),		/* addop, reduce: Factor */
			reduce(12),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(12),		/* eqop, reduce: Factor */
			reduce(12),		/* neqop, reduce: Factor */
			reduce(12),		/* gtop, reduce: Factor */
			reduce(12),		/* ltop, reduce: Factor */
			reduce(12),		/* geqop, reduce: Factor */
			reduce(12),		/* leqop, reduce: Factor */
			reduce(12),		/* andop, reduce: Factor */
			reduce(12),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(12),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			reduce(7),		/* mulop, reduce: NumLiteral */
			reduce(7),		/* divop, reduce: NumLiteral */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(7),		/* addop, reduce: NumLiteral */
			reduce(7),		/* subop, reduce: NumLiteral */
			nil,		/* notop */
			reduce(7),		/* eqop, reduce: NumLiteral */
			reduce(7),		/* neqop, reduce: NumLiteral */
			reduce(7),		/* gtop, reduce: NumLiteral */
			reduce(7),		/* ltop, reduce: NumLiteral */
			reduce(7),		/* geqop, reduce: NumLiteral */
			reduce(7),		/* leqop, reduce: NumLiteral */
			reduce(7),		/* andop, reduce: NumLiteral */
			reduce(7),		/* orop, reduce: NumLiteral */
			nil,		/* lpar */
			reduce(7),		/* rpar, reduce: NumLiteral */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			shift(57),		/* mulop */
			shift(58),		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(32),		/* addop, reduce: Expr */
			reduce(32),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(32),		/* eqop, reduce: Expr */
			reduce(32),		/* neqop, reduce: Expr */
			reduce(32),		/* gtop, reduce: Expr */
			reduce(32),		/* ltop, reduce: Expr */
			reduce(32),		/* geqop, reduce: Expr */
			reduce(32),		/* leqop, reduce: Expr */
			reduce(32),		/* andop, reduce: Expr */
			reduce(32),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(32),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			reduce(11),		/* mulop, reduce: Term */
			reduce(11),		/* divop, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: Term */
			reduce(11),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(11),		/* eqop, reduce: Term */
			reduce(11),		/* neqop, reduce: Term */
			reduce(11),		/* gtop, reduce: Term */
			reduce(11),		/* ltop, reduce: Term */
			reduce(11),		/* geqop, reduce: Term */
			reduce(11),		/* leqop, reduce: Term */
			reduce(11),		/* andop, reduce: Term */
			reduce(11),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(11),		/* rpar, reduce: Term */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(33),		/* addop, reduce: Expr */
			reduce(33),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(33),		/* eqop, reduce: Expr */
			reduce(33),		/* neqop, reduce: Expr */
			reduce(33),		/* gtop, reduce: Expr */
			reduce(33),		/* ltop, reduce: Expr */
			reduce(33),		/* geqop, reduce: Expr */
			reduce(33),		/* leqop, reduce: Expr */
			reduce(33),		/* andop, reduce: Expr */
			reduce(33),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(33),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(14),		/* addop, reduce: Bool */
			reduce(14),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(14),		/* eqop, reduce: Bool */
			reduce(14),		/* neqop, reduce: Bool */
			reduce(14),		/* gtop, reduce: Bool */
			reduce(14),		/* ltop, reduce: Bool */
			reduce(14),		/* geqop, reduce: Bool */
			reduce(14),		/* leqop, reduce: Bool */
			reduce(14),		/* andop, reduce: Bool */
			reduce(14),		/* orop, reduce: Bool */
			nil,		/* lpar */
			reduce(14),		/* rpar, reduce: Bool */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(15),		/* addop, reduce: Bool */
			reduce(15),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: Bool */
			reduce(15),		/* neqop, reduce: Bool */
			reduce(15),		/* gtop, reduce: Bool */
			reduce(15),		/* ltop, reduce: Bool */
			reduce(15),		/* geqop, reduce: Bool */
			reduce(15),		/* leqop, reduce: Bool */
			reduce(15),		/* andop, reduce: Bool */
			reduce(15),		/* orop, reduce: Bool */
			nil,		/* lpar */
			reduce(15),		/* rpar, reduce: Bool */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			reduce(17),		/* mulop, reduce: VarId */
			reduce(17),		/* divop, reduce: VarId */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(17),		/* addop, reduce: VarId */
			reduce(17),		/* subop, reduce: VarId */
			nil,		/* notop */
			reduce(17),		/* eqop, reduce: VarId */
			reduce(17),		/* neqop, reduce: VarId */
			reduce(17),		/* gtop, reduce: VarId */
			reduce(17),		/* ltop, reduce: VarId */
			reduce(17),		/* geqop, reduce: VarId */
			reduce(17),		/* leqop, reduce: VarId */
			reduce(17),		/* andop, reduce: VarId */
			reduce(17),		/* orop, reduce: VarId */
			nil,		/* lpar */
			reduce(17),		/* rpar, reduce: VarId */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			shift(69),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(16),		/* str_lit, reduce: VarDecl */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(16),		/* if, reduce: VarDecl */
			nil,		/* else */
			reduce(16),		/* assign, reduce: VarDecl */
			nil,		/* lbrace */
			reduce(16),		/* rbrace, reduce: VarDecl */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(2),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(2),		/* if, reduce: Type */
			nil,		/* else */
			reduce(2),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(2),		/* rbrace, reduce: Type */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(3),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(3),		/* if, reduce: Type */
			nil,		/* else */
			reduce(3),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(3),		/* rbrace, reduce: Type */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(4),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(4),		/* if, reduce: Type */
			nil,		/* else */
			reduce(4),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(4),		/* rbrace, reduce: Type */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(5),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(5),		/* if, reduce: Type */
			nil,		/* else */
			reduce(5),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(5),		/* rbrace, reduce: Type */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(6),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(6),		/* if, reduce: Type */
			nil,		/* else */
			reduce(6),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(6),		/* rbrace, reduce: Type */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(13),		/* str_lit, reduce: Factor */
			reduce(13),		/* mulop, reduce: Factor */
			reduce(13),		/* divop, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(13),		/* addop, reduce: Factor */
			reduce(13),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(13),		/* eqop, reduce: Factor */
			reduce(13),		/* neqop, reduce: Factor */
			reduce(13),		/* gtop, reduce: Factor */
			reduce(13),		/* ltop, reduce: Factor */
			reduce(13),		/* geqop, reduce: Factor */
			reduce(13),		/* leqop, reduce: Factor */
			reduce(13),		/* andop, reduce: Factor */
			reduce(13),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(13),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(13),		/* rbrace, reduce: Factor */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(12),		/* str_lit, reduce: Factor */
			reduce(12),		/* mulop, reduce: Factor */
			reduce(12),		/* divop, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(12),		/* addop, reduce: Factor */
			reduce(12),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(12),		/* eqop, reduce: Factor */
			reduce(12),		/* neqop, reduce: Factor */
			reduce(12),		/* gtop, reduce: Factor */
			reduce(12),		/* ltop, reduce: Factor */
			reduce(12),		/* geqop, reduce: Factor */
			reduce(12),		/* leqop, reduce: Factor */
			reduce(12),		/* andop, reduce: Factor */
			reduce(12),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(12),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(12),		/* rbrace, reduce: Factor */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(7),		/* str_lit, reduce: NumLiteral */
			reduce(7),		/* mulop, reduce: NumLiteral */
			reduce(7),		/* divop, reduce: NumLiteral */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(7),		/* addop, reduce: NumLiteral */
			reduce(7),		/* subop, reduce: NumLiteral */
			nil,		/* notop */
			reduce(7),		/* eqop, reduce: NumLiteral */
			reduce(7),		/* neqop, reduce: NumLiteral */
			reduce(7),		/* gtop, reduce: NumLiteral */
			reduce(7),		/* ltop, reduce: NumLiteral */
			reduce(7),		/* geqop, reduce: NumLiteral */
			reduce(7),		/* leqop, reduce: NumLiteral */
			reduce(7),		/* andop, reduce: NumLiteral */
			reduce(7),		/* orop, reduce: NumLiteral */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(7),		/* if, reduce: NumLiteral */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(7),		/* rbrace, reduce: NumLiteral */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(32),		/* str_lit, reduce: Expr */
			shift(74),		/* mulop */
			shift(75),		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(32),		/* addop, reduce: Expr */
			reduce(32),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(32),		/* eqop, reduce: Expr */
			reduce(32),		/* neqop, reduce: Expr */
			reduce(32),		/* gtop, reduce: Expr */
			reduce(32),		/* ltop, reduce: Expr */
			reduce(32),		/* geqop, reduce: Expr */
			reduce(32),		/* leqop, reduce: Expr */
			reduce(32),		/* andop, reduce: Expr */
			reduce(32),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(32),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(32),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(11),		/* str_lit, reduce: Term */
			reduce(11),		/* mulop, reduce: Term */
			reduce(11),		/* divop, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: Term */
			reduce(11),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(11),		/* eqop, reduce: Term */
			reduce(11),		/* neqop, reduce: Term */
			reduce(11),		/* gtop, reduce: Term */
			reduce(11),		/* ltop, reduce: Term */
			reduce(11),		/* geqop, reduce: Term */
			reduce(11),		/* leqop, reduce: Term */
			reduce(11),		/* andop, reduce: Term */
			reduce(11),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(11),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(11),		/* rbrace, reduce: Term */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(33),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(33),		/* addop, reduce: Expr */
			reduce(33),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(33),		/* eqop, reduce: Expr */
			reduce(33),		/* neqop, reduce: Expr */
			reduce(33),		/* gtop, reduce: Expr */
			reduce(33),		/* ltop, reduce: Expr */
			reduce(33),		/* geqop, reduce: Expr */
			reduce(33),		/* leqop, reduce: Expr */
			reduce(33),		/* andop, reduce: Expr */
			reduce(33),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(33),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(33),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(14),		/* str_lit, reduce: Bool */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(14),		/* addop, reduce: Bool */
			reduce(14),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(14),		/* eqop, reduce: Bool */
			reduce(14),		/* neqop, reduce: Bool */
			reduce(14),		/* gtop, reduce: Bool */
			reduce(14),		/* ltop, reduce: Bool */
			reduce(14),		/* geqop, reduce: Bool */
			reduce(14),		/* leqop, reduce: Bool */
			reduce(14),		/* andop, reduce: Bool */
			reduce(14),		/* orop, reduce: Bool */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(14),		/* if, reduce: Bool */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(14),		/* rbrace, reduce: Bool */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(15),		/* str_lit, reduce: Bool */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(15),		/* addop, reduce: Bool */
			reduce(15),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: Bool */
			reduce(15),		/* neqop, reduce: Bool */
			reduce(15),		/* gtop, reduce: Bool */
			reduce(15),		/* ltop, reduce: Bool */
			reduce(15),		/* geqop, reduce: Bool */
			reduce(15),		/* leqop, reduce: Bool */
			reduce(15),		/* andop, reduce: Bool */
			reduce(15),		/* orop, reduce: Bool */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(15),		/* if, reduce: Bool */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(15),		/* rbrace, reduce: Bool */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(17),		/* str_lit, reduce: VarId */
			reduce(17),		/* mulop, reduce: VarId */
			reduce(17),		/* divop, reduce: VarId */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(17),		/* addop, reduce: VarId */
			reduce(17),		/* subop, reduce: VarId */
			nil,		/* notop */
			reduce(17),		/* eqop, reduce: VarId */
			reduce(17),		/* neqop, reduce: VarId */
			reduce(17),		/* gtop, reduce: VarId */
			reduce(17),		/* ltop, reduce: VarId */
			reduce(17),		/* geqop, reduce: VarId */
			reduce(17),		/* leqop, reduce: VarId */
			reduce(17),		/* andop, reduce: VarId */
			reduce(17),		/* orop, reduce: VarId */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(17),		/* if, reduce: VarId */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(17),		/* rbrace, reduce: VarId */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(37),		/* str_lit, reduce: Question */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(37),		/* if, reduce: Question */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(37),		/* rbrace, reduce: Question */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S64
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S65
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S66
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S67
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S68
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(25),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* booltrue_lit */
			shift(30),		/* boolfalse_lit */
			nil,		/* col */
			shift(31),		/* ident */
			shift(33),		/* addop */
			shift(34),		/* subop */
			shift(35),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(36),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S69
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			shift(103),		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S70
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(18),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S71
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(19),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S72
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(20),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S73
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			shift(104),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S74
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S75
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S76
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S77
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S78
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S79
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S80
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S81
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S82
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S83
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S84
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S85
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(45),		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* booltrue_lit */
			shift(50),		/* boolfalse_lit */
			nil,		/* col */
			shift(51),		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			shift(55),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(56),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S86
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(18),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(18),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(18),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S87
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(19),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(19),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(19),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S88
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(20),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(20),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(20),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S89
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			shift(117),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S90
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			reduce(9),		/* mulop, reduce: Term */
			reduce(9),		/* divop, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: Term */
			reduce(9),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(9),		/* eqop, reduce: Term */
			reduce(9),		/* neqop, reduce: Term */
			reduce(9),		/* gtop, reduce: Term */
			reduce(9),		/* ltop, reduce: Term */
			reduce(9),		/* geqop, reduce: Term */
			reduce(9),		/* leqop, reduce: Term */
			reduce(9),		/* andop, reduce: Term */
			reduce(9),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(9),		/* rpar, reduce: Term */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S91
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			reduce(10),		/* mulop, reduce: Term */
			reduce(10),		/* divop, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(10),		/* addop, reduce: Term */
			reduce(10),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(10),		/* eqop, reduce: Term */
			reduce(10),		/* neqop, reduce: Term */
			reduce(10),		/* gtop, reduce: Term */
			reduce(10),		/* ltop, reduce: Term */
			reduce(10),		/* geqop, reduce: Term */
			reduce(10),		/* leqop, reduce: Term */
			reduce(10),		/* andop, reduce: Term */
			reduce(10),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(10),		/* rpar, reduce: Term */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S92
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			shift(57),		/* mulop */
			shift(58),		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(21),		/* addop, reduce: Expr */
			reduce(21),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(21),		/* eqop, reduce: Expr */
			reduce(21),		/* neqop, reduce: Expr */
			reduce(21),		/* gtop, reduce: Expr */
			reduce(21),		/* ltop, reduce: Expr */
			reduce(21),		/* geqop, reduce: Expr */
			reduce(21),		/* leqop, reduce: Expr */
			reduce(21),		/* andop, reduce: Expr */
			reduce(21),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(21),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S93
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			shift(57),		/* mulop */
			shift(58),		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(22),		/* addop, reduce: Expr */
			reduce(22),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(22),		/* eqop, reduce: Expr */
			reduce(22),		/* neqop, reduce: Expr */
			reduce(22),		/* gtop, reduce: Expr */
			reduce(22),		/* ltop, reduce: Expr */
			reduce(22),		/* geqop, reduce: Expr */
			reduce(22),		/* leqop, reduce: Expr */
			reduce(22),		/* andop, reduce: Expr */
			reduce(22),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(22),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S94
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(23),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S95
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(24),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S96
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(25),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S97
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(26),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S98
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(27),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S99
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(28),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S100
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(29),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S101
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(59),		/* addop */
			shift(60),		/* subop */
			nil,		/* notop */
			shift(61),		/* eqop */
			shift(62),		/* neqop */
			shift(63),		/* gtop */
			shift(64),		/* ltop */
			shift(65),		/* geqop */
			shift(66),		/* leqop */
			shift(67),		/* andop */
			shift(68),		/* orop */
			nil,		/* lpar */
			reduce(30),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S102
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(34),		/* str_lit, reduce: IfStmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(34),		/* if, reduce: IfStmt */
			shift(118),		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(34),		/* rbrace, reduce: IfStmt */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S103
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(120),		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S104
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(31),		/* addop, reduce: Expr */
			reduce(31),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(31),		/* eqop, reduce: Expr */
			reduce(31),		/* neqop, reduce: Expr */
			reduce(31),		/* gtop, reduce: Expr */
			reduce(31),		/* ltop, reduce: Expr */
			reduce(31),		/* geqop, reduce: Expr */
			reduce(31),		/* leqop, reduce: Expr */
			reduce(31),		/* andop, reduce: Expr */
			reduce(31),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(31),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S105
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(9),		/* str_lit, reduce: Term */
			reduce(9),		/* mulop, reduce: Term */
			reduce(9),		/* divop, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: Term */
			reduce(9),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(9),		/* eqop, reduce: Term */
			reduce(9),		/* neqop, reduce: Term */
			reduce(9),		/* gtop, reduce: Term */
			reduce(9),		/* ltop, reduce: Term */
			reduce(9),		/* geqop, reduce: Term */
			reduce(9),		/* leqop, reduce: Term */
			reduce(9),		/* andop, reduce: Term */
			reduce(9),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(9),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(9),		/* rbrace, reduce: Term */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S106
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(10),		/* str_lit, reduce: Term */
			reduce(10),		/* mulop, reduce: Term */
			reduce(10),		/* divop, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(10),		/* addop, reduce: Term */
			reduce(10),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(10),		/* eqop, reduce: Term */
			reduce(10),		/* neqop, reduce: Term */
			reduce(10),		/* gtop, reduce: Term */
			reduce(10),		/* ltop, reduce: Term */
			reduce(10),		/* geqop, reduce: Term */
			reduce(10),		/* leqop, reduce: Term */
			reduce(10),		/* andop, reduce: Term */
			reduce(10),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(10),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(10),		/* rbrace, reduce: Term */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S107
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(21),		/* str_lit, reduce: Expr */
			shift(74),		/* mulop */
			shift(75),		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(21),		/* addop, reduce: Expr */
			reduce(21),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(21),		/* eqop, reduce: Expr */
			reduce(21),		/* neqop, reduce: Expr */
			reduce(21),		/* gtop, reduce: Expr */
			reduce(21),		/* ltop, reduce: Expr */
			reduce(21),		/* geqop, reduce: Expr */
			reduce(21),		/* leqop, reduce: Expr */
			reduce(21),		/* andop, reduce: Expr */
			reduce(21),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(21),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(21),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S108
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(22),		/* str_lit, reduce: Expr */
			shift(74),		/* mulop */
			shift(75),		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(22),		/* addop, reduce: Expr */
			reduce(22),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(22),		/* eqop, reduce: Expr */
			reduce(22),		/* neqop, reduce: Expr */
			reduce(22),		/* gtop, reduce: Expr */
			reduce(22),		/* ltop, reduce: Expr */
			reduce(22),		/* geqop, reduce: Expr */
			reduce(22),		/* leqop, reduce: Expr */
			reduce(22),		/* andop, reduce: Expr */
			reduce(22),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(22),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(22),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S109
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(23),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(23),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(23),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S110
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(24),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(24),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(24),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S111
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(25),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(25),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(25),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S112
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(26),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(26),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(26),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S113
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(27),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(27),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(27),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S114
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(28),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(28),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(28),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S115
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(29),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(29),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(29),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S116
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(30),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(76),		/* addop */
			shift(77),		/* subop */
			nil,		/* notop */
			shift(78),		/* eqop */
			shift(79),		/* neqop */
			shift(80),		/* gtop */
			shift(81),		/* ltop */
			shift(82),		/* geqop */
			shift(83),		/* leqop */
			shift(84),		/* andop */
			shift(85),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(30),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(30),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S117
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(31),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(31),		/* addop, reduce: Expr */
			reduce(31),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(31),		/* eqop, reduce: Expr */
			reduce(31),		/* neqop, reduce: Expr */
			reduce(31),		/* gtop, reduce: Expr */
			reduce(31),		/* ltop, reduce: Expr */
			reduce(31),		/* geqop, reduce: Expr */
			reduce(31),		/* leqop, reduce: Expr */
			reduce(31),		/* andop, reduce: Expr */
			reduce(31),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(31),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(31),		/* rbrace, reduce: Expr */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S118
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			shift(122),		/* lbrace */
			nil,		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S119
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(123),		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S120
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(39),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(39),		/* if, reduce: Block */
			reduce(39),		/* else, reduce: Block */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(39),		/* rbrace, reduce: Block */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S121
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(35),		/* str_lit, reduce: IfStmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(35),		/* if, reduce: IfStmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(35),		/* rbrace, reduce: IfStmt */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S122
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(125),		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S123
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(38),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(38),		/* if, reduce: Block */
			reduce(38),		/* else, reduce: Block */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(38),		/* rbrace, reduce: Block */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S124
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(126),		/* rbrace */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S125
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(39),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(39),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(39),		/* rbrace, reduce: Block */
			nil,		/* money_lit */
			
		},

	},
	actionRow{ // S126
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			reduce(38),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(38),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(38),		/* rbrace, reduce: Block */
			nil,		/* money_lit */
			
		},

	},
	
}

