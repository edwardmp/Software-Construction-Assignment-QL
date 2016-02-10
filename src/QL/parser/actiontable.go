
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
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Expr */
			reduce(6),		/* addop, reduce: Expr */
			reduce(6),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(6),		/* eqop, reduce: Expr */
			reduce(6),		/* neqop, reduce: Expr */
			reduce(6),		/* gtop, reduce: Expr */
			reduce(6),		/* ltop, reduce: Expr */
			reduce(6),		/* geqop, reduce: Expr */
			reduce(6),		/* leqop, reduce: Expr */
			reduce(6),		/* andop, reduce: Expr */
			reduce(6),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			shift(26),		/* mulop */
			shift(27),		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(16),		/* $, reduce: Expr */
			reduce(16),		/* addop, reduce: Expr */
			reduce(16),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(16),		/* eqop, reduce: Expr */
			reduce(16),		/* neqop, reduce: Expr */
			reduce(16),		/* gtop, reduce: Expr */
			reduce(16),		/* ltop, reduce: Expr */
			reduce(16),		/* geqop, reduce: Expr */
			reduce(16),		/* leqop, reduce: Expr */
			reduce(16),		/* andop, reduce: Expr */
			reduce(16),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(17),		/* $, reduce: Expr */
			reduce(17),		/* addop, reduce: Expr */
			reduce(17),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(17),		/* eqop, reduce: Expr */
			reduce(17),		/* neqop, reduce: Expr */
			reduce(17),		/* gtop, reduce: Expr */
			reduce(17),		/* ltop, reduce: Expr */
			reduce(17),		/* geqop, reduce: Expr */
			reduce(17),		/* leqop, reduce: Expr */
			reduce(17),		/* andop, reduce: Expr */
			reduce(17),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(20),		/* $, reduce: Term */
			reduce(20),		/* addop, reduce: Term */
			reduce(20),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(20),		/* eqop, reduce: Term */
			reduce(20),		/* neqop, reduce: Term */
			reduce(20),		/* gtop, reduce: Term */
			reduce(20),		/* ltop, reduce: Term */
			reduce(20),		/* geqop, reduce: Term */
			reduce(20),		/* leqop, reduce: Term */
			reduce(20),		/* andop, reduce: Term */
			reduce(20),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			reduce(20),		/* mulop, reduce: Term */
			reduce(20),		/* divop, reduce: Term */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(21),		/* $, reduce: Factor */
			reduce(21),		/* addop, reduce: Factor */
			reduce(21),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(21),		/* eqop, reduce: Factor */
			reduce(21),		/* neqop, reduce: Factor */
			reduce(21),		/* gtop, reduce: Factor */
			reduce(21),		/* ltop, reduce: Factor */
			reduce(21),		/* geqop, reduce: Factor */
			reduce(21),		/* leqop, reduce: Factor */
			reduce(21),		/* andop, reduce: Factor */
			reduce(21),		/* orop, reduce: Factor */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			reduce(21),		/* mulop, reduce: Factor */
			reduce(21),		/* divop, reduce: Factor */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(22),		/* $, reduce: Bool */
			reduce(22),		/* addop, reduce: Bool */
			reduce(22),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(22),		/* eqop, reduce: Bool */
			reduce(22),		/* neqop, reduce: Bool */
			reduce(22),		/* gtop, reduce: Bool */
			reduce(22),		/* ltop, reduce: Bool */
			reduce(22),		/* geqop, reduce: Bool */
			reduce(22),		/* leqop, reduce: Bool */
			reduce(22),		/* andop, reduce: Bool */
			reduce(22),		/* orop, reduce: Bool */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(23),		/* $, reduce: Bool */
			reduce(23),		/* addop, reduce: Bool */
			reduce(23),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(23),		/* eqop, reduce: Bool */
			reduce(23),		/* neqop, reduce: Bool */
			reduce(23),		/* gtop, reduce: Bool */
			reduce(23),		/* ltop, reduce: Bool */
			reduce(23),		/* geqop, reduce: Bool */
			reduce(23),		/* leqop, reduce: Bool */
			reduce(23),		/* andop, reduce: Bool */
			reduce(23),		/* orop, reduce: Bool */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
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
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(6),		/* lpar */
			nil,		/* rpar */
			shift(8),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			shift(11),		/* true */
			shift(12),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
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
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
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
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(10),		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			shift(62),		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(6),		/* addop, reduce: Expr */
			reduce(6),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(6),		/* eqop, reduce: Expr */
			reduce(6),		/* neqop, reduce: Expr */
			reduce(6),		/* gtop, reduce: Expr */
			reduce(6),		/* ltop, reduce: Expr */
			reduce(6),		/* geqop, reduce: Expr */
			reduce(6),		/* leqop, reduce: Expr */
			reduce(6),		/* andop, reduce: Expr */
			reduce(6),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(6),		/* rpar, reduce: Expr */
			nil,		/* ident */
			shift(66),		/* mulop */
			shift(67),		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(16),		/* addop, reduce: Expr */
			reduce(16),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(16),		/* eqop, reduce: Expr */
			reduce(16),		/* neqop, reduce: Expr */
			reduce(16),		/* gtop, reduce: Expr */
			reduce(16),		/* ltop, reduce: Expr */
			reduce(16),		/* geqop, reduce: Expr */
			reduce(16),		/* leqop, reduce: Expr */
			reduce(16),		/* andop, reduce: Expr */
			reduce(16),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(16),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(17),		/* addop, reduce: Expr */
			reduce(17),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(17),		/* eqop, reduce: Expr */
			reduce(17),		/* neqop, reduce: Expr */
			reduce(17),		/* gtop, reduce: Expr */
			reduce(17),		/* ltop, reduce: Expr */
			reduce(17),		/* geqop, reduce: Expr */
			reduce(17),		/* leqop, reduce: Expr */
			reduce(17),		/* andop, reduce: Expr */
			reduce(17),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(17),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(20),		/* addop, reduce: Term */
			reduce(20),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(20),		/* eqop, reduce: Term */
			reduce(20),		/* neqop, reduce: Term */
			reduce(20),		/* gtop, reduce: Term */
			reduce(20),		/* ltop, reduce: Term */
			reduce(20),		/* geqop, reduce: Term */
			reduce(20),		/* leqop, reduce: Term */
			reduce(20),		/* andop, reduce: Term */
			reduce(20),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(20),		/* rpar, reduce: Term */
			nil,		/* ident */
			reduce(20),		/* mulop, reduce: Term */
			reduce(20),		/* divop, reduce: Term */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(21),		/* addop, reduce: Factor */
			reduce(21),		/* subop, reduce: Factor */
			nil,		/* notop */
			reduce(21),		/* eqop, reduce: Factor */
			reduce(21),		/* neqop, reduce: Factor */
			reduce(21),		/* gtop, reduce: Factor */
			reduce(21),		/* ltop, reduce: Factor */
			reduce(21),		/* geqop, reduce: Factor */
			reduce(21),		/* leqop, reduce: Factor */
			reduce(21),		/* andop, reduce: Factor */
			reduce(21),		/* orop, reduce: Factor */
			nil,		/* lpar */
			reduce(21),		/* rpar, reduce: Factor */
			nil,		/* ident */
			reduce(21),		/* mulop, reduce: Factor */
			reduce(21),		/* divop, reduce: Factor */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(22),		/* addop, reduce: Bool */
			reduce(22),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(22),		/* eqop, reduce: Bool */
			reduce(22),		/* neqop, reduce: Bool */
			reduce(22),		/* gtop, reduce: Bool */
			reduce(22),		/* ltop, reduce: Bool */
			reduce(22),		/* geqop, reduce: Bool */
			reduce(22),		/* leqop, reduce: Bool */
			reduce(22),		/* andop, reduce: Bool */
			reduce(22),		/* orop, reduce: Bool */
			nil,		/* lpar */
			reduce(22),		/* rpar, reduce: Bool */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(23),		/* addop, reduce: Bool */
			reduce(23),		/* subop, reduce: Bool */
			nil,		/* notop */
			reduce(23),		/* eqop, reduce: Bool */
			reduce(23),		/* neqop, reduce: Bool */
			reduce(23),		/* gtop, reduce: Bool */
			reduce(23),		/* ltop, reduce: Bool */
			reduce(23),		/* geqop, reduce: Bool */
			reduce(23),		/* leqop, reduce: Bool */
			reduce(23),		/* andop, reduce: Bool */
			reduce(23),		/* orop, reduce: Bool */
			nil,		/* lpar */
			reduce(23),		/* rpar, reduce: Bool */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Expr */
			reduce(4),		/* addop, reduce: Expr */
			reduce(4),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(4),		/* eqop, reduce: Expr */
			reduce(4),		/* neqop, reduce: Expr */
			reduce(4),		/* gtop, reduce: Expr */
			reduce(4),		/* ltop, reduce: Expr */
			reduce(4),		/* geqop, reduce: Expr */
			reduce(4),		/* leqop, reduce: Expr */
			reduce(4),		/* andop, reduce: Expr */
			reduce(4),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			shift(26),		/* mulop */
			shift(27),		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(13),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(14),		/* $, reduce: Expr */
			shift(13),		/* addop */
			shift(14),		/* subop */
			nil,		/* notop */
			shift(15),		/* eqop */
			shift(16),		/* neqop */
			shift(17),		/* gtop */
			shift(18),		/* ltop */
			shift(19),		/* geqop */
			shift(20),		/* leqop */
			shift(21),		/* andop */
			shift(22),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(18),		/* $, reduce: Term */
			reduce(18),		/* addop, reduce: Term */
			reduce(18),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(18),		/* eqop, reduce: Term */
			reduce(18),		/* neqop, reduce: Term */
			reduce(18),		/* gtop, reduce: Term */
			reduce(18),		/* ltop, reduce: Term */
			reduce(18),		/* geqop, reduce: Term */
			reduce(18),		/* leqop, reduce: Term */
			reduce(18),		/* andop, reduce: Term */
			reduce(18),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			reduce(18),		/* mulop, reduce: Term */
			reduce(18),		/* divop, reduce: Term */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(19),		/* $, reduce: Term */
			reduce(19),		/* addop, reduce: Term */
			reduce(19),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(19),		/* eqop, reduce: Term */
			reduce(19),		/* neqop, reduce: Term */
			reduce(19),		/* gtop, reduce: Term */
			reduce(19),		/* ltop, reduce: Term */
			reduce(19),		/* geqop, reduce: Term */
			reduce(19),		/* leqop, reduce: Term */
			reduce(19),		/* andop, reduce: Term */
			reduce(19),		/* orop, reduce: Term */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			reduce(19),		/* mulop, reduce: Term */
			reduce(19),		/* divop, reduce: Term */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
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
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* addop */
			shift(30),		/* subop */
			shift(31),		/* notop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(33),		/* lpar */
			nil,		/* rpar */
			shift(35),		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			shift(38),		/* true */
			shift(39),		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(15),		/* $, reduce: Expr */
			reduce(15),		/* addop, reduce: Expr */
			reduce(15),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: Expr */
			reduce(15),		/* neqop, reduce: Expr */
			reduce(15),		/* gtop, reduce: Expr */
			reduce(15),		/* ltop, reduce: Expr */
			reduce(15),		/* geqop, reduce: Expr */
			reduce(15),		/* leqop, reduce: Expr */
			reduce(15),		/* andop, reduce: Expr */
			reduce(15),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(1),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S64
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(2),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S65
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(3),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S66
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
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
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S67
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
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
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			shift(37),		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S68
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			shift(81),		/* rpar */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S69
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(4),		/* addop, reduce: Expr */
			reduce(4),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(4),		/* eqop, reduce: Expr */
			reduce(4),		/* neqop, reduce: Expr */
			reduce(4),		/* gtop, reduce: Expr */
			reduce(4),		/* ltop, reduce: Expr */
			reduce(4),		/* geqop, reduce: Expr */
			reduce(4),		/* leqop, reduce: Expr */
			reduce(4),		/* andop, reduce: Expr */
			reduce(4),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(4),		/* rpar, reduce: Expr */
			nil,		/* ident */
			shift(66),		/* mulop */
			shift(67),		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S70
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(5),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S71
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(7),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S72
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(8),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S73
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(9),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S74
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(10),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S75
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(11),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S76
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(12),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S77
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(13),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S78
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(52),		/* addop */
			shift(53),		/* subop */
			nil,		/* notop */
			shift(54),		/* eqop */
			shift(55),		/* neqop */
			shift(56),		/* gtop */
			shift(57),		/* ltop */
			shift(58),		/* geqop */
			shift(59),		/* leqop */
			shift(60),		/* andop */
			shift(61),		/* orop */
			nil,		/* lpar */
			reduce(14),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S79
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(18),		/* addop, reduce: Term */
			reduce(18),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(18),		/* eqop, reduce: Term */
			reduce(18),		/* neqop, reduce: Term */
			reduce(18),		/* gtop, reduce: Term */
			reduce(18),		/* ltop, reduce: Term */
			reduce(18),		/* geqop, reduce: Term */
			reduce(18),		/* leqop, reduce: Term */
			reduce(18),		/* andop, reduce: Term */
			reduce(18),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(18),		/* rpar, reduce: Term */
			nil,		/* ident */
			reduce(18),		/* mulop, reduce: Term */
			reduce(18),		/* divop, reduce: Term */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S80
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(19),		/* addop, reduce: Term */
			reduce(19),		/* subop, reduce: Term */
			nil,		/* notop */
			reduce(19),		/* eqop, reduce: Term */
			reduce(19),		/* neqop, reduce: Term */
			reduce(19),		/* gtop, reduce: Term */
			reduce(19),		/* ltop, reduce: Term */
			reduce(19),		/* geqop, reduce: Term */
			reduce(19),		/* leqop, reduce: Term */
			reduce(19),		/* andop, reduce: Term */
			reduce(19),		/* orop, reduce: Term */
			nil,		/* lpar */
			reduce(19),		/* rpar, reduce: Term */
			nil,		/* ident */
			reduce(19),		/* mulop, reduce: Term */
			reduce(19),		/* divop, reduce: Term */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	actionRow{ // S81
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(15),		/* addop, reduce: Expr */
			reduce(15),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(15),		/* eqop, reduce: Expr */
			reduce(15),		/* neqop, reduce: Expr */
			reduce(15),		/* gtop, reduce: Expr */
			reduce(15),		/* ltop, reduce: Expr */
			reduce(15),		/* geqop, reduce: Expr */
			reduce(15),		/* leqop, reduce: Expr */
			reduce(15),		/* andop, reduce: Expr */
			reduce(15),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(15),		/* rpar, reduce: Expr */
			nil,		/* ident */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* lbrace */
			nil,		/* rbrace */
			nil,		/* type */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* col */
			nil,		/* form */
			nil,		/* literal */
			nil,		/* num_literal */
			nil,		/* semicol */
			
		},

	},
	
}

