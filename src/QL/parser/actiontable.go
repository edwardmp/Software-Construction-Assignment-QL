
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
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(3),		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(5),		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Form */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(6),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			shift(9),		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			shift(12),		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			reduce(12),		/* ident, reduce: StrLiteral */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(14),		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(46),		/* rbrace, reduce: Stmt */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(17),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(45),		/* str_lit, reduce: Stmt */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			reduce(45),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* lbrace */
			reduce(45),		/* rbrace, reduce: Stmt */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(6),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			shift(9),		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			shift(18),		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(42),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(43),		/* str_lit, reduce: StmtList */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			reduce(43),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* lbrace */
			reduce(43),		/* rbrace, reduce: StmtList */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(40),		/* rbrace, reduce: Question */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(20),		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(41),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(44),		/* rbrace, reduce: StmtList */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			shift(36),		/* integer */
			shift(37),		/* boolean */
			shift(38),		/* string */
			shift(39),		/* money */
			shift(40),		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			shift(41),		/* mulop */
			shift(42),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			shift(53),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(19),		/* rbrace, reduce: VarDecl */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(2),		/* rbrace, reduce: Type */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(3),		/* rbrace, reduce: Type */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(4),		/* rbrace, reduce: Type */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(5),		/* rbrace, reduce: Type */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(6),		/* rbrace, reduce: Type */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(21),		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			shift(22),		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			shift(27),		/* true */
			shift(28),		/* false */
			nil,		/* col */
			shift(31),		/* addop */
			shift(32),		/* subop */
			shift(33),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(34),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(71),		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(21),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(22),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(23),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			shift(72),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			shift(41),		/* mulop */
			shift(42),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			nil,		/* str_lit */
			shift(41),		/* mulop */
			shift(42),		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(26),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(27),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S64
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(28),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S65
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(29),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S66
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(30),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S67
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(31),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S68
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(32),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S69
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(43),		/* addop */
			shift(44),		/* subop */
			nil,		/* notop */
			shift(45),		/* eqop */
			shift(46),		/* neqop */
			shift(47),		/* gtop */
			shift(48),		/* ltop */
			shift(49),		/* geqop */
			shift(50),		/* leqop */
			shift(51),		/* andop */
			shift(52),		/* orop */
			nil,		/* lpar */
			reduce(33),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S70
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(73),		/* else */
			nil,		/* lbrace */
			reduce(38),		/* rbrace, reduce: IfStmt */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S71
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(6),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			shift(9),		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			shift(75),		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S72
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S73
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			shift(77),		/* lbrace */
			nil,		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S74
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(6),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			shift(9),		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			shift(78),		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S75
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(42),		/* rbrace, reduce: Block */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S76
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(39),		/* rbrace, reduce: IfStmt */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S77
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(6),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			shift(9),		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			shift(80),		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S78
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(41),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			reduce(41),		/* if, reduce: Block */
			reduce(41),		/* else, reduce: Block */
			nil,		/* lbrace */
			reduce(41),		/* rbrace, reduce: Block */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S79
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			shift(6),		/* str_lit */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			shift(9),		/* if */
			nil,		/* else */
			nil,		/* lbrace */
			shift(81),		/* rbrace */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S80
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
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
			nil,		/* lbrace */
			reduce(42),		/* rbrace, reduce: Block */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S81
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* money */
			nil,		/* date */
			nil,		/* integer_lit */
			nil,		/* money_lit */
			nil,		/* bool_lit */
			reduce(41),		/* str_lit, reduce: Block */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
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
			reduce(41),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* lbrace */
			reduce(41),		/* rbrace, reduce: Block */
			nil,		/* assign */
			nil,		/* semicol */
			
		},

	},
	
}

