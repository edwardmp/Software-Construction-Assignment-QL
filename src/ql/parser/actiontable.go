
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(20),		/* lbrace, reduce: VarId */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(7),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			reduce(12),		/* ident, reduce: StrLiteral */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(47),		/* str_lit, reduce: Stmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(47),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(47),		/* rbrace, reduce: Stmt */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(46),		/* str_lit, reduce: Stmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(46),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(46),		/* rbrace, reduce: Stmt */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(7),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(43),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(44),		/* str_lit, reduce: StmtList */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(44),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(44),		/* rbrace, reduce: StmtList */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(40),		/* str_lit, reduce: Question */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(40),		/* if, reduce: Question */
			nil,		/* else */
			shift(22),		/* assign */
			nil,		/* lbrace */
			reduce(40),		/* rbrace, reduce: Question */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			reduce(20),		/* col, reduce: VarId */
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
			nil,		/* money */
			nil,		/* date */
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(42),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(45),		/* str_lit, reduce: StmtList */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(45),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(45),		/* rbrace, reduce: StmtList */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money */
			nil,		/* date */
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(37),		/* addop, reduce: Expr */
			reduce(37),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(37),		/* eqop, reduce: Expr */
			reduce(37),		/* neqop, reduce: Expr */
			reduce(37),		/* gtop, reduce: Expr */
			reduce(37),		/* ltop, reduce: Expr */
			reduce(37),		/* geqop, reduce: Expr */
			reduce(37),		/* leqop, reduce: Expr */
			reduce(37),		/* andop, reduce: Expr */
			reduce(37),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(37),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			reduce(11),		/* mulop, reduce: NumLiteral */
			reduce(11),		/* divop, reduce: NumLiteral */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: NumLiteral */
			reduce(11),		/* subop, reduce: NumLiteral */
			nil,		/* notop */
			reduce(11),		/* eqop, reduce: NumLiteral */
			reduce(11),		/* neqop, reduce: NumLiteral */
			reduce(11),		/* gtop, reduce: NumLiteral */
			reduce(11),		/* ltop, reduce: NumLiteral */
			reduce(11),		/* geqop, reduce: NumLiteral */
			reduce(11),		/* leqop, reduce: NumLiteral */
			reduce(11),		/* andop, reduce: NumLiteral */
			reduce(11),		/* orop, reduce: NumLiteral */
			nil,		/* lpar */
			reduce(11),		/* rpar, reduce: NumLiteral */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			reduce(16),		/* mulop, reduce: Factor */
			reduce(16),		/* divop, reduce: Factor */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(16),		/* addop, reduce: Factor */
			reduce(16),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(16),		/* eqop, reduce: Factor */
			reduce(16),		/* neqop, reduce: Factor */
			reduce(16),		/* gtop, reduce: Factor */
			reduce(16),		/* ltop, reduce: Factor */
			reduce(16),		/* geqop, reduce: Factor */
			reduce(16),		/* leqop, reduce: Factor */
			reduce(16),		/* andop, reduce: Factor */
			reduce(16),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(16),		/* rpar, reduce: Factor */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			shift(57),		/* mulop */
			shift(58),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(35),		/* addop, reduce: Expr */
			reduce(35),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(35),		/* eqop, reduce: Expr */
			reduce(35),		/* neqop, reduce: Expr */
			reduce(35),		/* gtop, reduce: Expr */
			reduce(35),		/* ltop, reduce: Expr */
			reduce(35),		/* geqop, reduce: Expr */
			reduce(35),		/* leqop, reduce: Expr */
			reduce(35),		/* andop, reduce: Expr */
			reduce(35),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(35),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			reduce(15),		/* mulop, reduce: Term */
			reduce(15),		/* divop, reduce: Term */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(15),		/* addop, reduce: Term */
			reduce(15),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: Term */
			reduce(15),		/* neqop, reduce: Term */
			reduce(15),		/* gtop, reduce: Term */
			reduce(15),		/* ltop, reduce: Term */
			reduce(15),		/* geqop, reduce: Term */
			reduce(15),		/* leqop, reduce: Term */
			reduce(15),		/* andop, reduce: Term */
			reduce(15),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(15),		/* rpar, reduce: Term */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(36),		/* addop, reduce: Expr */
			reduce(36),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(36),		/* eqop, reduce: Expr */
			reduce(36),		/* neqop, reduce: Expr */
			reduce(36),		/* gtop, reduce: Expr */
			reduce(36),		/* ltop, reduce: Expr */
			reduce(36),		/* geqop, reduce: Expr */
			reduce(36),		/* leqop, reduce: Expr */
			reduce(36),		/* andop, reduce: Expr */
			reduce(36),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(36),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(17),		/* addop, reduce: Bool */
			reduce(17),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(17),		/* eqop, reduce: Bool */
			reduce(17),		/* neqop, reduce: Bool */
			reduce(17),		/* gtop, reduce: Bool */
			reduce(17),		/* ltop, reduce: Bool */
			reduce(17),		/* geqop, reduce: Bool */
			reduce(17),		/* leqop, reduce: Bool */
			reduce(17),		/* andop, reduce: Bool */
			reduce(17),		/* orop, reduce: Bool */
			nil,		/* lpar */
			reduce(17),		/* rpar, reduce: Bool */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(18),		/* addop, reduce: Bool */
			reduce(18),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(18),		/* eqop, reduce: Bool */
			reduce(18),		/* neqop, reduce: Bool */
			reduce(18),		/* gtop, reduce: Bool */
			reduce(18),		/* ltop, reduce: Bool */
			reduce(18),		/* geqop, reduce: Bool */
			reduce(18),		/* leqop, reduce: Bool */
			reduce(18),		/* andop, reduce: Bool */
			reduce(18),		/* orop, reduce: Bool */
			nil,		/* lpar */
			reduce(18),		/* rpar, reduce: Bool */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(20),		/* addop, reduce: VarId */
			reduce(20),		/* subop, reduce: VarId */
			nil,		/* notop */
			reduce(20),		/* eqop, reduce: VarId */
			reduce(20),		/* neqop, reduce: VarId */
			reduce(20),		/* gtop, reduce: VarId */
			reduce(20),		/* ltop, reduce: VarId */
			reduce(20),		/* geqop, reduce: VarId */
			reduce(20),		/* leqop, reduce: VarId */
			reduce(20),		/* andop, reduce: VarId */
			reduce(20),		/* orop, reduce: VarId */
			nil,		/* lpar */
			reduce(20),		/* rpar, reduce: VarId */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(19),		/* str_lit, reduce: VarDecl */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(19),		/* if, reduce: VarDecl */
			nil,		/* else */
			reduce(19),		/* assign, reduce: VarDecl */
			nil,		/* lbrace */
			reduce(19),		/* rbrace, reduce: VarDecl */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(2),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(3),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(4),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(5),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(6),		/* str_lit, reduce: Type */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(37),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(37),		/* addop, reduce: Expr */
			reduce(37),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(37),		/* eqop, reduce: Expr */
			reduce(37),		/* neqop, reduce: Expr */
			reduce(37),		/* gtop, reduce: Expr */
			reduce(37),		/* ltop, reduce: Expr */
			reduce(37),		/* geqop, reduce: Expr */
			reduce(37),		/* leqop, reduce: Expr */
			reduce(37),		/* andop, reduce: Expr */
			reduce(37),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(37),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(37),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(11),		/* str_lit, reduce: NumLiteral */
			reduce(11),		/* mulop, reduce: NumLiteral */
			reduce(11),		/* divop, reduce: NumLiteral */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: NumLiteral */
			reduce(11),		/* subop, reduce: NumLiteral */
			nil,		/* notop */
			reduce(11),		/* eqop, reduce: NumLiteral */
			reduce(11),		/* neqop, reduce: NumLiteral */
			reduce(11),		/* gtop, reduce: NumLiteral */
			reduce(11),		/* ltop, reduce: NumLiteral */
			reduce(11),		/* geqop, reduce: NumLiteral */
			reduce(11),		/* leqop, reduce: NumLiteral */
			reduce(11),		/* andop, reduce: NumLiteral */
			reduce(11),		/* orop, reduce: NumLiteral */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(11),		/* if, reduce: NumLiteral */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(11),		/* rbrace, reduce: NumLiteral */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(16),		/* str_lit, reduce: Factor */
			reduce(16),		/* mulop, reduce: Factor */
			reduce(16),		/* divop, reduce: Factor */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(16),		/* addop, reduce: Factor */
			reduce(16),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(16),		/* eqop, reduce: Factor */
			reduce(16),		/* neqop, reduce: Factor */
			reduce(16),		/* gtop, reduce: Factor */
			reduce(16),		/* ltop, reduce: Factor */
			reduce(16),		/* geqop, reduce: Factor */
			reduce(16),		/* leqop, reduce: Factor */
			reduce(16),		/* andop, reduce: Factor */
			reduce(16),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(16),		/* if, reduce: Factor */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(16),		/* rbrace, reduce: Factor */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(35),		/* str_lit, reduce: Expr */
			shift(74),		/* mulop */
			shift(75),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(35),		/* addop, reduce: Expr */
			reduce(35),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(35),		/* eqop, reduce: Expr */
			reduce(35),		/* neqop, reduce: Expr */
			reduce(35),		/* gtop, reduce: Expr */
			reduce(35),		/* ltop, reduce: Expr */
			reduce(35),		/* geqop, reduce: Expr */
			reduce(35),		/* leqop, reduce: Expr */
			reduce(35),		/* andop, reduce: Expr */
			reduce(35),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(35),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(35),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(15),		/* str_lit, reduce: Term */
			reduce(15),		/* mulop, reduce: Term */
			reduce(15),		/* divop, reduce: Term */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(15),		/* addop, reduce: Term */
			reduce(15),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: Term */
			reduce(15),		/* neqop, reduce: Term */
			reduce(15),		/* gtop, reduce: Term */
			reduce(15),		/* ltop, reduce: Term */
			reduce(15),		/* geqop, reduce: Term */
			reduce(15),		/* leqop, reduce: Term */
			reduce(15),		/* andop, reduce: Term */
			reduce(15),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(15),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(15),		/* rbrace, reduce: Term */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(36),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(36),		/* addop, reduce: Expr */
			reduce(36),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(36),		/* eqop, reduce: Expr */
			reduce(36),		/* neqop, reduce: Expr */
			reduce(36),		/* gtop, reduce: Expr */
			reduce(36),		/* ltop, reduce: Expr */
			reduce(36),		/* geqop, reduce: Expr */
			reduce(36),		/* leqop, reduce: Expr */
			reduce(36),		/* andop, reduce: Expr */
			reduce(36),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(36),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(36),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(17),		/* str_lit, reduce: Bool */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(17),		/* addop, reduce: Bool */
			reduce(17),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(17),		/* eqop, reduce: Bool */
			reduce(17),		/* neqop, reduce: Bool */
			reduce(17),		/* gtop, reduce: Bool */
			reduce(17),		/* ltop, reduce: Bool */
			reduce(17),		/* geqop, reduce: Bool */
			reduce(17),		/* leqop, reduce: Bool */
			reduce(17),		/* andop, reduce: Bool */
			reduce(17),		/* orop, reduce: Bool */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(17),		/* if, reduce: Bool */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(17),		/* rbrace, reduce: Bool */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(18),		/* str_lit, reduce: Bool */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(18),		/* addop, reduce: Bool */
			reduce(18),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(18),		/* eqop, reduce: Bool */
			reduce(18),		/* neqop, reduce: Bool */
			reduce(18),		/* gtop, reduce: Bool */
			reduce(18),		/* ltop, reduce: Bool */
			reduce(18),		/* geqop, reduce: Bool */
			reduce(18),		/* leqop, reduce: Bool */
			reduce(18),		/* andop, reduce: Bool */
			reduce(18),		/* orop, reduce: Bool */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(18),		/* if, reduce: Bool */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(18),		/* rbrace, reduce: Bool */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(20),		/* str_lit, reduce: VarId */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(20),		/* addop, reduce: VarId */
			reduce(20),		/* subop, reduce: VarId */
			nil,		/* notop */
			reduce(20),		/* eqop, reduce: VarId */
			reduce(20),		/* neqop, reduce: VarId */
			reduce(20),		/* gtop, reduce: VarId */
			reduce(20),		/* ltop, reduce: VarId */
			reduce(20),		/* geqop, reduce: VarId */
			reduce(20),		/* leqop, reduce: VarId */
			reduce(20),		/* andop, reduce: VarId */
			reduce(20),		/* orop, reduce: VarId */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(20),		/* if, reduce: VarId */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(20),		/* rbrace, reduce: VarId */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(41),		/* str_lit, reduce: Question */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(41),		/* if, reduce: Question */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(41),		/* rbrace, reduce: Question */
			
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			shift(24),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(29),		/* true */
			shift(30),		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(21),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(22),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			shift(44),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(49),		/* true */
			shift(50),		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(21),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(21),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(21),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(22),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(22),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(22),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(23),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			reduce(13),		/* mulop, reduce: Term */
			reduce(13),		/* divop, reduce: Term */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(13),		/* addop, reduce: Term */
			reduce(13),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(13),		/* eqop, reduce: Term */
			reduce(13),		/* neqop, reduce: Term */
			reduce(13),		/* gtop, reduce: Term */
			reduce(13),		/* ltop, reduce: Term */
			reduce(13),		/* geqop, reduce: Term */
			reduce(13),		/* leqop, reduce: Term */
			reduce(13),		/* andop, reduce: Term */
			reduce(13),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(13),		/* rpar, reduce: Term */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			reduce(14),		/* mulop, reduce: Term */
			reduce(14),		/* divop, reduce: Term */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(14),		/* addop, reduce: Term */
			reduce(14),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(14),		/* eqop, reduce: Term */
			reduce(14),		/* neqop, reduce: Term */
			reduce(14),		/* gtop, reduce: Term */
			reduce(14),		/* ltop, reduce: Term */
			reduce(14),		/* geqop, reduce: Term */
			reduce(14),		/* leqop, reduce: Term */
			reduce(14),		/* andop, reduce: Term */
			reduce(14),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(14),		/* rpar, reduce: Term */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			shift(57),		/* mulop */
			shift(58),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(24),		/* addop, reduce: Expr */
			reduce(24),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(24),		/* eqop, reduce: Expr */
			reduce(24),		/* neqop, reduce: Expr */
			reduce(24),		/* gtop, reduce: Expr */
			reduce(24),		/* ltop, reduce: Expr */
			reduce(24),		/* geqop, reduce: Expr */
			reduce(24),		/* leqop, reduce: Expr */
			reduce(24),		/* andop, reduce: Expr */
			reduce(24),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(24),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			shift(57),		/* mulop */
			shift(58),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(25),		/* addop, reduce: Expr */
			reduce(25),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(25),		/* eqop, reduce: Expr */
			reduce(25),		/* neqop, reduce: Expr */
			reduce(25),		/* gtop, reduce: Expr */
			reduce(25),		/* ltop, reduce: Expr */
			reduce(25),		/* geqop, reduce: Expr */
			reduce(25),		/* leqop, reduce: Expr */
			reduce(25),		/* andop, reduce: Expr */
			reduce(25),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(25),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(31),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(32),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(33),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(38),		/* str_lit, reduce: IfStmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(38),		/* if, reduce: IfStmt */
			shift(118),		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(38),		/* rbrace, reduce: IfStmt */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(7),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(34),		/* addop, reduce: Expr */
			reduce(34),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(34),		/* eqop, reduce: Expr */
			reduce(34),		/* neqop, reduce: Expr */
			reduce(34),		/* gtop, reduce: Expr */
			reduce(34),		/* ltop, reduce: Expr */
			reduce(34),		/* geqop, reduce: Expr */
			reduce(34),		/* leqop, reduce: Expr */
			reduce(34),		/* andop, reduce: Expr */
			reduce(34),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(34),		/* rpar, reduce: Expr */
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
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(13),		/* str_lit, reduce: Term */
			reduce(13),		/* mulop, reduce: Term */
			reduce(13),		/* divop, reduce: Term */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(13),		/* addop, reduce: Term */
			reduce(13),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(13),		/* eqop, reduce: Term */
			reduce(13),		/* neqop, reduce: Term */
			reduce(13),		/* gtop, reduce: Term */
			reduce(13),		/* ltop, reduce: Term */
			reduce(13),		/* geqop, reduce: Term */
			reduce(13),		/* leqop, reduce: Term */
			reduce(13),		/* andop, reduce: Term */
			reduce(13),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(13),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(13),		/* rbrace, reduce: Term */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(14),		/* str_lit, reduce: Term */
			reduce(14),		/* mulop, reduce: Term */
			reduce(14),		/* divop, reduce: Term */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(14),		/* addop, reduce: Term */
			reduce(14),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(14),		/* eqop, reduce: Term */
			reduce(14),		/* neqop, reduce: Term */
			reduce(14),		/* gtop, reduce: Term */
			reduce(14),		/* ltop, reduce: Term */
			reduce(14),		/* geqop, reduce: Term */
			reduce(14),		/* leqop, reduce: Term */
			reduce(14),		/* andop, reduce: Term */
			reduce(14),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(14),		/* if, reduce: Term */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(14),		/* rbrace, reduce: Term */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(24),		/* str_lit, reduce: Expr */
			shift(74),		/* mulop */
			shift(75),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(24),		/* addop, reduce: Expr */
			reduce(24),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(24),		/* eqop, reduce: Expr */
			reduce(24),		/* neqop, reduce: Expr */
			reduce(24),		/* gtop, reduce: Expr */
			reduce(24),		/* ltop, reduce: Expr */
			reduce(24),		/* geqop, reduce: Expr */
			reduce(24),		/* leqop, reduce: Expr */
			reduce(24),		/* andop, reduce: Expr */
			reduce(24),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(24),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(24),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(25),		/* str_lit, reduce: Expr */
			shift(74),		/* mulop */
			shift(75),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(25),		/* addop, reduce: Expr */
			reduce(25),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(25),		/* eqop, reduce: Expr */
			reduce(25),		/* neqop, reduce: Expr */
			reduce(25),		/* gtop, reduce: Expr */
			reduce(25),		/* ltop, reduce: Expr */
			reduce(25),		/* geqop, reduce: Expr */
			reduce(25),		/* leqop, reduce: Expr */
			reduce(25),		/* andop, reduce: Expr */
			reduce(25),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(25),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(25),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(26),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(27),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(28),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(29),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(30),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(31),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(31),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(31),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(32),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(32),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(32),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(33),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(33),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(33),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(34),		/* str_lit, reduce: Expr */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* ident */
			reduce(34),		/* addop, reduce: Expr */
			reduce(34),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(34),		/* eqop, reduce: Expr */
			reduce(34),		/* neqop, reduce: Expr */
			reduce(34),		/* gtop, reduce: Expr */
			reduce(34),		/* ltop, reduce: Expr */
			reduce(34),		/* geqop, reduce: Expr */
			reduce(34),		/* leqop, reduce: Expr */
			reduce(34),		/* andop, reduce: Expr */
			reduce(34),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(34),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(34),		/* rbrace, reduce: Expr */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(7),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(43),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(43),		/* if, reduce: Block */
			reduce(43),		/* else, reduce: Block */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(43),		/* rbrace, reduce: Block */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(39),		/* str_lit, reduce: IfStmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(39),		/* if, reduce: IfStmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(39),		/* rbrace, reduce: IfStmt */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(7),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(42),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(42),		/* if, reduce: Block */
			reduce(42),		/* else, reduce: Block */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(42),		/* rbrace, reduce: Block */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(7),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(43),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(43),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(43),		/* rbrace, reduce: Block */
			
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
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(42),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
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
			reduce(42),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(42),		/* rbrace, reduce: Block */
			
		},

	},
	
}

