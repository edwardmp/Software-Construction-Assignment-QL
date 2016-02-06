
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(17),		/* $, reduce: Expr */
			reduce(17),		/* addop, reduce: Expr */
			reduce(17),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(17),		/* mulop, reduce: Expr */
			reduce(17),		/* divop, reduce: Expr */
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
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(18),		/* $, reduce: Expr */
			reduce(18),		/* addop, reduce: Expr */
			reduce(18),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(18),		/* mulop, reduce: Expr */
			reduce(18),		/* divop, reduce: Expr */
			reduce(18),		/* eqop, reduce: Expr */
			reduce(18),		/* neqop, reduce: Expr */
			reduce(18),		/* gtop, reduce: Expr */
			reduce(18),		/* ltop, reduce: Expr */
			reduce(18),		/* geqop, reduce: Expr */
			reduce(18),		/* leqop, reduce: Expr */
			reduce(18),		/* andop, reduce: Expr */
			reduce(18),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(19),		/* $, reduce: Val */
			reduce(19),		/* addop, reduce: Val */
			reduce(19),		/* subop, reduce: Val */
			nil,		/* notop */
			reduce(19),		/* mulop, reduce: Val */
			reduce(19),		/* divop, reduce: Val */
			reduce(19),		/* eqop, reduce: Val */
			reduce(19),		/* neqop, reduce: Val */
			reduce(19),		/* gtop, reduce: Val */
			reduce(19),		/* ltop, reduce: Val */
			reduce(19),		/* geqop, reduce: Val */
			reduce(19),		/* leqop, reduce: Val */
			reduce(19),		/* andop, reduce: Val */
			reduce(19),		/* orop, reduce: Val */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(20),		/* $, reduce: Val */
			reduce(20),		/* addop, reduce: Val */
			reduce(20),		/* subop, reduce: Val */
			nil,		/* notop */
			reduce(20),		/* mulop, reduce: Val */
			reduce(20),		/* divop, reduce: Val */
			reduce(20),		/* eqop, reduce: Val */
			reduce(20),		/* neqop, reduce: Val */
			reduce(20),		/* gtop, reduce: Val */
			reduce(20),		/* ltop, reduce: Val */
			reduce(20),		/* geqop, reduce: Val */
			reduce(20),		/* leqop, reduce: Val */
			reduce(20),		/* andop, reduce: Val */
			reduce(20),		/* orop, reduce: Val */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* addop */
			shift(3),		/* subop */
			shift(4),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
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
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(5),		/* lpar */
			nil,		/* rpar */
			shift(6),		/* int */
			shift(8),		/* true */
			shift(9),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			shift(58),		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(17),		/* addop, reduce: Expr */
			reduce(17),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(17),		/* mulop, reduce: Expr */
			reduce(17),		/* divop, reduce: Expr */
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
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(18),		/* addop, reduce: Expr */
			reduce(18),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(18),		/* mulop, reduce: Expr */
			reduce(18),		/* divop, reduce: Expr */
			reduce(18),		/* eqop, reduce: Expr */
			reduce(18),		/* neqop, reduce: Expr */
			reduce(18),		/* gtop, reduce: Expr */
			reduce(18),		/* ltop, reduce: Expr */
			reduce(18),		/* geqop, reduce: Expr */
			reduce(18),		/* leqop, reduce: Expr */
			reduce(18),		/* andop, reduce: Expr */
			reduce(18),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(18),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(19),		/* addop, reduce: Val */
			reduce(19),		/* subop, reduce: Val */
			nil,		/* notop */
			reduce(19),		/* mulop, reduce: Val */
			reduce(19),		/* divop, reduce: Val */
			reduce(19),		/* eqop, reduce: Val */
			reduce(19),		/* neqop, reduce: Val */
			reduce(19),		/* gtop, reduce: Val */
			reduce(19),		/* ltop, reduce: Val */
			reduce(19),		/* geqop, reduce: Val */
			reduce(19),		/* leqop, reduce: Val */
			reduce(19),		/* andop, reduce: Val */
			reduce(19),		/* orop, reduce: Val */
			nil,		/* lpar */
			reduce(19),		/* rpar, reduce: Val */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(20),		/* addop, reduce: Val */
			reduce(20),		/* subop, reduce: Val */
			nil,		/* notop */
			reduce(20),		/* mulop, reduce: Val */
			reduce(20),		/* divop, reduce: Val */
			reduce(20),		/* eqop, reduce: Val */
			reduce(20),		/* neqop, reduce: Val */
			reduce(20),		/* gtop, reduce: Val */
			reduce(20),		/* ltop, reduce: Val */
			reduce(20),		/* geqop, reduce: Val */
			reduce(20),		/* leqop, reduce: Val */
			reduce(20),		/* andop, reduce: Val */
			reduce(20),		/* orop, reduce: Val */
			nil,		/* lpar */
			reduce(20),		/* rpar, reduce: Val */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(13),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(14),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(15),		/* $, reduce: Expr */
			shift(10),		/* addop */
			shift(11),		/* subop */
			nil,		/* notop */
			shift(12),		/* mulop */
			shift(13),		/* divop */
			shift(14),		/* eqop */
			shift(15),		/* neqop */
			shift(16),		/* gtop */
			shift(17),		/* ltop */
			shift(18),		/* geqop */
			shift(19),		/* leqop */
			shift(20),		/* andop */
			shift(21),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(26),		/* addop */
			shift(27),		/* subop */
			shift(28),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(29),		/* lpar */
			nil,		/* rpar */
			shift(30),		/* int */
			shift(32),		/* true */
			shift(33),		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(16),		/* $, reduce: Expr */
			reduce(16),		/* addop, reduce: Expr */
			reduce(16),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(16),		/* mulop, reduce: Expr */
			reduce(16),		/* divop, reduce: Expr */
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
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(1),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(2),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(3),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			shift(75),		/* rpar */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(6),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S64
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(7),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S65
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(4),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S66
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(5),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S67
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(8),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S68
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(9),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S69
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(10),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S70
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(11),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S71
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(12),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S72
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(13),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S73
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(14),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S74
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(46),		/* addop */
			shift(47),		/* subop */
			nil,		/* notop */
			shift(48),		/* mulop */
			shift(49),		/* divop */
			shift(50),		/* eqop */
			shift(51),		/* neqop */
			shift(52),		/* gtop */
			shift(53),		/* ltop */
			shift(54),		/* geqop */
			shift(55),		/* leqop */
			shift(56),		/* andop */
			shift(57),		/* orop */
			nil,		/* lpar */
			reduce(15),		/* rpar, reduce: Expr */
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	actionRow{ // S75
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(16),		/* addop, reduce: Expr */
			reduce(16),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(16),		/* mulop, reduce: Expr */
			reduce(16),		/* divop, reduce: Expr */
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
			nil,		/* int */
			nil,		/* true */
			nil,		/* false */
			nil,		/* col */
			nil,		/* compop */
			nil,		/* else */
			nil,		/* equals */
			nil,		/* form */
			nil,		/* ident */
			nil,		/* if */
			nil,		/* lbrace */
			nil,		/* literal */
			nil,		/* logop */
			nil,		/* num_literal */
			nil,		/* rbrace */
			nil,		/* semicol */
			nil,		/* type */
			
		},

	},
	
}

