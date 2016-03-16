
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			reduce(13),		/* ident, reduce: StrLiteral */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(44),		/* str_lit, reduce: Stmt */
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
			reduce(44),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(44),		/* rbrace, reduce: Stmt */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(43),		/* str_lit, reduce: Stmt */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
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
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(40),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(41),		/* str_lit, reduce: StmtList */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(37),		/* str_lit, reduce: Question */
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
			reduce(37),		/* if, reduce: Question */
			nil,		/* else */
			shift(22),		/* assign */
			nil,		/* lbrace */
			reduce(37),		/* rbrace, reduce: Question */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(39),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(42),		/* str_lit, reduce: StmtList */
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
			reduce(42),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(42),		/* rbrace, reduce: StmtList */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(40),		/* integer */
			shift(41),		/* boolean */
			shift(42),		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(11),		/* mulop, reduce: Factor */
			reduce(11),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: Factor */
			reduce(11),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(11),		/* eqop, reduce: Factor */
			reduce(11),		/* neqop, reduce: Factor */
			reduce(11),		/* gtop, reduce: Factor */
			reduce(11),		/* ltop, reduce: Factor */
			reduce(11),		/* geqop, reduce: Factor */
			reduce(11),		/* leqop, reduce: Factor */
			reduce(11),		/* andop, reduce: Factor */
			reduce(11),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(11),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			shift(59),		/* mulop */
			shift(60),		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			reduce(7),		/* mulop, reduce: Term */
			reduce(7),		/* divop, reduce: Term */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(7),		/* addop, reduce: Term */
			reduce(7),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(7),		/* eqop, reduce: Term */
			reduce(7),		/* neqop, reduce: Term */
			reduce(7),		/* gtop, reduce: Term */
			reduce(7),		/* ltop, reduce: Term */
			reduce(7),		/* geqop, reduce: Term */
			reduce(7),		/* leqop, reduce: Term */
			reduce(7),		/* andop, reduce: Term */
			reduce(7),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(7),		/* rpar, reduce: Term */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(8),		/* mulop, reduce: Factor */
			reduce(8),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(8),		/* addop, reduce: Factor */
			reduce(8),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(8),		/* eqop, reduce: Factor */
			reduce(8),		/* neqop, reduce: Factor */
			reduce(8),		/* gtop, reduce: Factor */
			reduce(8),		/* ltop, reduce: Factor */
			reduce(8),		/* geqop, reduce: Factor */
			reduce(8),		/* leqop, reduce: Factor */
			reduce(8),		/* andop, reduce: Factor */
			reduce(8),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(8),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(9),		/* mulop, reduce: Factor */
			reduce(9),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: Factor */
			reduce(9),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(9),		/* eqop, reduce: Factor */
			reduce(9),		/* neqop, reduce: Factor */
			reduce(9),		/* gtop, reduce: Factor */
			reduce(9),		/* ltop, reduce: Factor */
			reduce(9),		/* geqop, reduce: Factor */
			reduce(9),		/* leqop, reduce: Factor */
			reduce(9),		/* andop, reduce: Factor */
			reduce(9),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(9),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(10),		/* mulop, reduce: Factor */
			reduce(10),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(10),		/* addop, reduce: Factor */
			reduce(10),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(10),		/* eqop, reduce: Factor */
			reduce(10),		/* neqop, reduce: Factor */
			reduce(10),		/* gtop, reduce: Factor */
			reduce(10),		/* ltop, reduce: Factor */
			reduce(10),		/* geqop, reduce: Factor */
			reduce(10),		/* leqop, reduce: Factor */
			reduce(10),		/* andop, reduce: Factor */
			reduce(10),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(10),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(12),		/* mulop, reduce: NumLiteral */
			reduce(12),		/* divop, reduce: NumLiteral */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(12),		/* addop, reduce: NumLiteral */
			reduce(12),		/* subop, reduce: NumLiteral */
			nil,		/* notop */
			reduce(12),		/* eqop, reduce: NumLiteral */
			reduce(12),		/* neqop, reduce: NumLiteral */
			reduce(12),		/* gtop, reduce: NumLiteral */
			reduce(12),		/* ltop, reduce: NumLiteral */
			reduce(12),		/* geqop, reduce: NumLiteral */
			reduce(12),		/* leqop, reduce: NumLiteral */
			reduce(12),		/* andop, reduce: NumLiteral */
			reduce(12),		/* orop, reduce: NumLiteral */
			nil,		/* lpar */
			reduce(12),		/* rpar, reduce: NumLiteral */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(13),		/* mulop, reduce: StrLiteral */
			reduce(13),		/* divop, reduce: StrLiteral */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(13),		/* addop, reduce: StrLiteral */
			reduce(13),		/* subop, reduce: StrLiteral */
			nil,		/* notop */
			reduce(13),		/* eqop, reduce: StrLiteral */
			reduce(13),		/* neqop, reduce: StrLiteral */
			reduce(13),		/* gtop, reduce: StrLiteral */
			reduce(13),		/* ltop, reduce: StrLiteral */
			reduce(13),		/* geqop, reduce: StrLiteral */
			reduce(13),		/* leqop, reduce: StrLiteral */
			reduce(13),		/* andop, reduce: StrLiteral */
			reduce(13),		/* orop, reduce: StrLiteral */
			nil,		/* lpar */
			reduce(13),		/* rpar, reduce: StrLiteral */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(14),		/* mulop, reduce: BoolLiteral */
			reduce(14),		/* divop, reduce: BoolLiteral */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(14),		/* addop, reduce: BoolLiteral */
			reduce(14),		/* subop, reduce: BoolLiteral */
			nil,		/* notop */
			reduce(14),		/* eqop, reduce: BoolLiteral */
			reduce(14),		/* neqop, reduce: BoolLiteral */
			reduce(14),		/* gtop, reduce: BoolLiteral */
			reduce(14),		/* ltop, reduce: BoolLiteral */
			reduce(14),		/* geqop, reduce: BoolLiteral */
			reduce(14),		/* leqop, reduce: BoolLiteral */
			reduce(14),		/* andop, reduce: BoolLiteral */
			reduce(14),		/* orop, reduce: BoolLiteral */
			nil,		/* lpar */
			reduce(14),		/* rpar, reduce: BoolLiteral */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(15),		/* mulop, reduce: BoolLiteral */
			reduce(15),		/* divop, reduce: BoolLiteral */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(15),		/* addop, reduce: BoolLiteral */
			reduce(15),		/* subop, reduce: BoolLiteral */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: BoolLiteral */
			reduce(15),		/* neqop, reduce: BoolLiteral */
			reduce(15),		/* gtop, reduce: BoolLiteral */
			reduce(15),		/* ltop, reduce: BoolLiteral */
			reduce(15),		/* geqop, reduce: BoolLiteral */
			reduce(15),		/* leqop, reduce: BoolLiteral */
			reduce(15),		/* andop, reduce: BoolLiteral */
			reduce(15),		/* orop, reduce: BoolLiteral */
			nil,		/* lpar */
			reduce(15),		/* rpar, reduce: BoolLiteral */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(17),		/* mulop, reduce: VarId */
			reduce(17),		/* divop, reduce: VarId */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			shift(71),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(16),		/* str_lit, reduce: VarDecl */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(2),		/* str_lit, reduce: Type */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(3),		/* str_lit, reduce: Type */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(4),		/* str_lit, reduce: Type */
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
			reduce(11),		/* mulop, reduce: Factor */
			reduce(11),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			reduce(11),		/* str_lit, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: Factor */
			reduce(11),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(11),		/* eqop, reduce: Factor */
			reduce(11),		/* neqop, reduce: Factor */
			reduce(11),		/* gtop, reduce: Factor */
			reduce(11),		/* ltop, reduce: Factor */
			reduce(11),		/* geqop, reduce: Factor */
			reduce(11),		/* leqop, reduce: Factor */
			reduce(11),		/* andop, reduce: Factor */
			reduce(11),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(11),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(11),		/* rbrace, reduce: Factor */
			
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
			shift(76),		/* mulop */
			shift(77),		/* divop */
			nil,		/* integer_lit */
			reduce(32),		/* str_lit, reduce: Expr */
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
			reduce(7),		/* mulop, reduce: Term */
			reduce(7),		/* divop, reduce: Term */
			nil,		/* integer_lit */
			reduce(7),		/* str_lit, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(7),		/* addop, reduce: Term */
			reduce(7),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(7),		/* eqop, reduce: Term */
			reduce(7),		/* neqop, reduce: Term */
			reduce(7),		/* gtop, reduce: Term */
			reduce(7),		/* ltop, reduce: Term */
			reduce(7),		/* geqop, reduce: Term */
			reduce(7),		/* leqop, reduce: Term */
			reduce(7),		/* andop, reduce: Term */
			reduce(7),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(7),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(7),		/* rbrace, reduce: Term */
			
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
			reduce(8),		/* mulop, reduce: Factor */
			reduce(8),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			reduce(8),		/* str_lit, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(8),		/* addop, reduce: Factor */
			reduce(8),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(8),		/* eqop, reduce: Factor */
			reduce(8),		/* neqop, reduce: Factor */
			reduce(8),		/* gtop, reduce: Factor */
			reduce(8),		/* ltop, reduce: Factor */
			reduce(8),		/* geqop, reduce: Factor */
			reduce(8),		/* leqop, reduce: Factor */
			reduce(8),		/* andop, reduce: Factor */
			reduce(8),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(8),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(8),		/* rbrace, reduce: Factor */
			
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
			reduce(9),		/* mulop, reduce: Factor */
			reduce(9),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			reduce(9),		/* str_lit, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: Factor */
			reduce(9),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(9),		/* eqop, reduce: Factor */
			reduce(9),		/* neqop, reduce: Factor */
			reduce(9),		/* gtop, reduce: Factor */
			reduce(9),		/* ltop, reduce: Factor */
			reduce(9),		/* geqop, reduce: Factor */
			reduce(9),		/* leqop, reduce: Factor */
			reduce(9),		/* andop, reduce: Factor */
			reduce(9),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(9),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(9),		/* rbrace, reduce: Factor */
			
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
			reduce(10),		/* mulop, reduce: Factor */
			reduce(10),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			reduce(10),		/* str_lit, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(10),		/* addop, reduce: Factor */
			reduce(10),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(10),		/* eqop, reduce: Factor */
			reduce(10),		/* neqop, reduce: Factor */
			reduce(10),		/* gtop, reduce: Factor */
			reduce(10),		/* ltop, reduce: Factor */
			reduce(10),		/* geqop, reduce: Factor */
			reduce(10),		/* leqop, reduce: Factor */
			reduce(10),		/* andop, reduce: Factor */
			reduce(10),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(10),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(10),		/* rbrace, reduce: Factor */
			
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
			reduce(12),		/* mulop, reduce: NumLiteral */
			reduce(12),		/* divop, reduce: NumLiteral */
			nil,		/* integer_lit */
			reduce(12),		/* str_lit, reduce: NumLiteral */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(12),		/* addop, reduce: NumLiteral */
			reduce(12),		/* subop, reduce: NumLiteral */
			nil,		/* notop */
			reduce(12),		/* eqop, reduce: NumLiteral */
			reduce(12),		/* neqop, reduce: NumLiteral */
			reduce(12),		/* gtop, reduce: NumLiteral */
			reduce(12),		/* ltop, reduce: NumLiteral */
			reduce(12),		/* geqop, reduce: NumLiteral */
			reduce(12),		/* leqop, reduce: NumLiteral */
			reduce(12),		/* andop, reduce: NumLiteral */
			reduce(12),		/* orop, reduce: NumLiteral */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(12),		/* if, reduce: NumLiteral */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(12),		/* rbrace, reduce: NumLiteral */
			
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
			reduce(13),		/* mulop, reduce: StrLiteral */
			reduce(13),		/* divop, reduce: StrLiteral */
			nil,		/* integer_lit */
			reduce(13),		/* str_lit, reduce: StrLiteral */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(13),		/* addop, reduce: StrLiteral */
			reduce(13),		/* subop, reduce: StrLiteral */
			nil,		/* notop */
			reduce(13),		/* eqop, reduce: StrLiteral */
			reduce(13),		/* neqop, reduce: StrLiteral */
			reduce(13),		/* gtop, reduce: StrLiteral */
			reduce(13),		/* ltop, reduce: StrLiteral */
			reduce(13),		/* geqop, reduce: StrLiteral */
			reduce(13),		/* leqop, reduce: StrLiteral */
			reduce(13),		/* andop, reduce: StrLiteral */
			reduce(13),		/* orop, reduce: StrLiteral */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(13),		/* if, reduce: StrLiteral */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(13),		/* rbrace, reduce: StrLiteral */
			
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
			reduce(14),		/* mulop, reduce: BoolLiteral */
			reduce(14),		/* divop, reduce: BoolLiteral */
			nil,		/* integer_lit */
			reduce(14),		/* str_lit, reduce: BoolLiteral */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(14),		/* addop, reduce: BoolLiteral */
			reduce(14),		/* subop, reduce: BoolLiteral */
			nil,		/* notop */
			reduce(14),		/* eqop, reduce: BoolLiteral */
			reduce(14),		/* neqop, reduce: BoolLiteral */
			reduce(14),		/* gtop, reduce: BoolLiteral */
			reduce(14),		/* ltop, reduce: BoolLiteral */
			reduce(14),		/* geqop, reduce: BoolLiteral */
			reduce(14),		/* leqop, reduce: BoolLiteral */
			reduce(14),		/* andop, reduce: BoolLiteral */
			reduce(14),		/* orop, reduce: BoolLiteral */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(14),		/* if, reduce: BoolLiteral */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(14),		/* rbrace, reduce: BoolLiteral */
			
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
			reduce(15),		/* mulop, reduce: BoolLiteral */
			reduce(15),		/* divop, reduce: BoolLiteral */
			nil,		/* integer_lit */
			reduce(15),		/* str_lit, reduce: BoolLiteral */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(15),		/* addop, reduce: BoolLiteral */
			reduce(15),		/* subop, reduce: BoolLiteral */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: BoolLiteral */
			reduce(15),		/* neqop, reduce: BoolLiteral */
			reduce(15),		/* gtop, reduce: BoolLiteral */
			reduce(15),		/* ltop, reduce: BoolLiteral */
			reduce(15),		/* geqop, reduce: BoolLiteral */
			reduce(15),		/* leqop, reduce: BoolLiteral */
			reduce(15),		/* andop, reduce: BoolLiteral */
			reduce(15),		/* orop, reduce: BoolLiteral */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(15),		/* if, reduce: BoolLiteral */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(15),		/* rbrace, reduce: BoolLiteral */
			
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
			reduce(17),		/* mulop, reduce: VarId */
			reduce(17),		/* divop, reduce: VarId */
			nil,		/* integer_lit */
			reduce(17),		/* str_lit, reduce: VarId */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(38),		/* str_lit, reduce: Question */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(38),		/* if, reduce: Question */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(38),		/* rbrace, reduce: Question */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* integer_lit */
			shift(30),		/* str_lit */
			shift(31),		/* booltrue_lit */
			shift(32),		/* boolfalse_lit */
			nil,		/* col */
			shift(33),		/* ident */
			shift(35),		/* addop */
			shift(36),		/* subop */
			shift(37),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(38),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			shift(107),		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(18),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(19),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(20),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			shift(108),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* integer_lit */
			shift(50),		/* str_lit */
			shift(51),		/* booltrue_lit */
			shift(52),		/* boolfalse_lit */
			nil,		/* col */
			shift(53),		/* ident */
			shift(55),		/* addop */
			shift(56),		/* subop */
			shift(57),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(58),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(18),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(18),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(18),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(19),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(19),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(19),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(20),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(20),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(20),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			shift(123),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(5),		/* mulop, reduce: Term */
			reduce(5),		/* divop, reduce: Term */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(5),		/* addop, reduce: Term */
			reduce(5),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(5),		/* eqop, reduce: Term */
			reduce(5),		/* neqop, reduce: Term */
			reduce(5),		/* gtop, reduce: Term */
			reduce(5),		/* ltop, reduce: Term */
			reduce(5),		/* geqop, reduce: Term */
			reduce(5),		/* leqop, reduce: Term */
			reduce(5),		/* andop, reduce: Term */
			reduce(5),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(5),		/* rpar, reduce: Term */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(9),		/* mulop, reduce: Factor */
			reduce(9),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: Factor */
			reduce(9),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(9),		/* eqop, reduce: Factor */
			reduce(9),		/* neqop, reduce: Factor */
			reduce(9),		/* gtop, reduce: Factor */
			reduce(9),		/* ltop, reduce: Factor */
			reduce(9),		/* geqop, reduce: Factor */
			reduce(9),		/* leqop, reduce: Factor */
			reduce(9),		/* andop, reduce: Factor */
			reduce(9),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(9),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(10),		/* mulop, reduce: Factor */
			reduce(10),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(10),		/* addop, reduce: Factor */
			reduce(10),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(10),		/* eqop, reduce: Factor */
			reduce(10),		/* neqop, reduce: Factor */
			reduce(10),		/* gtop, reduce: Factor */
			reduce(10),		/* ltop, reduce: Factor */
			reduce(10),		/* geqop, reduce: Factor */
			reduce(10),		/* leqop, reduce: Factor */
			reduce(10),		/* andop, reduce: Factor */
			reduce(10),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(10),		/* rpar, reduce: Factor */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			reduce(6),		/* mulop, reduce: Term */
			reduce(6),		/* divop, reduce: Term */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(6),		/* addop, reduce: Term */
			reduce(6),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(6),		/* eqop, reduce: Term */
			reduce(6),		/* neqop, reduce: Term */
			reduce(6),		/* gtop, reduce: Term */
			reduce(6),		/* ltop, reduce: Term */
			reduce(6),		/* geqop, reduce: Term */
			reduce(6),		/* leqop, reduce: Term */
			reduce(6),		/* andop, reduce: Term */
			reduce(6),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(6),		/* rpar, reduce: Term */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			shift(59),		/* mulop */
			shift(60),		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			shift(59),		/* mulop */
			shift(60),		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(23),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(24),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(25),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(26),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(27),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(28),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(29),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(61),		/* addop */
			shift(62),		/* subop */
			nil,		/* notop */
			shift(63),		/* eqop */
			shift(64),		/* neqop */
			shift(65),		/* gtop */
			shift(66),		/* ltop */
			shift(67),		/* geqop */
			shift(68),		/* leqop */
			shift(69),		/* andop */
			shift(70),		/* orop */
			nil,		/* lpar */
			reduce(30),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(35),		/* str_lit, reduce: IfStmt */
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
			shift(124),		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(35),		/* rbrace, reduce: IfStmt */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			reduce(5),		/* mulop, reduce: Term */
			reduce(5),		/* divop, reduce: Term */
			nil,		/* integer_lit */
			reduce(5),		/* str_lit, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(5),		/* addop, reduce: Term */
			reduce(5),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(5),		/* eqop, reduce: Term */
			reduce(5),		/* neqop, reduce: Term */
			reduce(5),		/* gtop, reduce: Term */
			reduce(5),		/* ltop, reduce: Term */
			reduce(5),		/* geqop, reduce: Term */
			reduce(5),		/* leqop, reduce: Term */
			reduce(5),		/* andop, reduce: Term */
			reduce(5),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(5),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(5),		/* rbrace, reduce: Term */
			
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
			reduce(9),		/* mulop, reduce: Factor */
			reduce(9),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			reduce(9),		/* str_lit, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: Factor */
			reduce(9),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(9),		/* eqop, reduce: Factor */
			reduce(9),		/* neqop, reduce: Factor */
			reduce(9),		/* gtop, reduce: Factor */
			reduce(9),		/* ltop, reduce: Factor */
			reduce(9),		/* geqop, reduce: Factor */
			reduce(9),		/* leqop, reduce: Factor */
			reduce(9),		/* andop, reduce: Factor */
			reduce(9),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(9),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(9),		/* rbrace, reduce: Factor */
			
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
			reduce(10),		/* mulop, reduce: Factor */
			reduce(10),		/* divop, reduce: Factor */
			nil,		/* integer_lit */
			reduce(10),		/* str_lit, reduce: Factor */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(10),		/* addop, reduce: Factor */
			reduce(10),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(10),		/* eqop, reduce: Factor */
			reduce(10),		/* neqop, reduce: Factor */
			reduce(10),		/* gtop, reduce: Factor */
			reduce(10),		/* ltop, reduce: Factor */
			reduce(10),		/* geqop, reduce: Factor */
			reduce(10),		/* leqop, reduce: Factor */
			reduce(10),		/* andop, reduce: Factor */
			reduce(10),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(10),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(10),		/* rbrace, reduce: Factor */
			
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
			reduce(6),		/* mulop, reduce: Term */
			reduce(6),		/* divop, reduce: Term */
			nil,		/* integer_lit */
			reduce(6),		/* str_lit, reduce: Term */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(6),		/* addop, reduce: Term */
			reduce(6),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(6),		/* eqop, reduce: Term */
			reduce(6),		/* neqop, reduce: Term */
			reduce(6),		/* gtop, reduce: Term */
			reduce(6),		/* ltop, reduce: Term */
			reduce(6),		/* geqop, reduce: Term */
			reduce(6),		/* leqop, reduce: Term */
			reduce(6),		/* andop, reduce: Term */
			reduce(6),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(6),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(6),		/* rbrace, reduce: Term */
			
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
			shift(76),		/* mulop */
			shift(77),		/* divop */
			nil,		/* integer_lit */
			reduce(21),		/* str_lit, reduce: Expr */
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
			shift(76),		/* mulop */
			shift(77),		/* divop */
			nil,		/* integer_lit */
			reduce(22),		/* str_lit, reduce: Expr */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(23),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(23),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(23),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(24),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(24),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(24),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(25),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(25),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(25),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(26),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(26),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(26),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(27),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(27),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(27),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(28),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(28),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(28),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(29),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(29),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(29),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(30),		/* str_lit, reduce: Expr */
			nil,		/* booltrue_lit */
			nil,		/* boolfalse_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(78),		/* addop */
			shift(79),		/* subop */
			nil,		/* notop */
			shift(80),		/* eqop */
			shift(81),		/* neqop */
			shift(82),		/* gtop */
			shift(83),		/* ltop */
			shift(84),		/* geqop */
			shift(85),		/* leqop */
			shift(86),		/* andop */
			shift(87),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(30),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(30),		/* rbrace, reduce: Expr */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(31),		/* str_lit, reduce: Expr */
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			nil,		/* str_lit */
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
			shift(128),		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
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
			shift(129),		/* rbrace */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(40),		/* str_lit, reduce: Block */
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
			reduce(40),		/* if, reduce: Block */
			reduce(40),		/* else, reduce: Block */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(40),		/* rbrace, reduce: Block */
			
		},

	},
	actionRow{ // S127
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(36),		/* str_lit, reduce: IfStmt */
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
			reduce(36),		/* if, reduce: IfStmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(36),		/* rbrace, reduce: IfStmt */
			
		},

	},
	actionRow{ // S128
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
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
			shift(131),		/* rbrace */
			
		},

	},
	actionRow{ // S129
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(39),		/* str_lit, reduce: Block */
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
			
		},

	},
	actionRow{ // S130
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			shift(8),		/* str_lit */
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
			shift(132),		/* rbrace */
			
		},

	},
	actionRow{ // S131
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(40),		/* str_lit, reduce: Block */
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
			reduce(40),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(40),		/* rbrace, reduce: Block */
			
		},

	},
	actionRow{ // S132
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* integer_lit */
			reduce(39),		/* str_lit, reduce: Block */
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
			
		},

	},
	
}

