
/*
*/
package parser

const numNTSymbols = 15
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		4, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		7, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		8, // IfStmt
		10, // Question
		-1, // Block
		11, // StmtList
		12, // Stmt
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		7, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		8, // IfStmt
		10, // Question
		-1, // Block
		-1, // StmtList
		16, // Stmt
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		26, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Form
		31, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		50, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		51, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		52, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		53, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		54, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		55, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		56, // Term
		22, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		57, // Term
		22, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		58, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S42
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		59, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S43
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		60, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S44
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		61, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S45
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		62, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S46
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		63, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S47
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		64, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S48
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		20, // NumLiteral
		-1, // StrLiteral
		21, // Term
		22, // Factor
		23, // Bool
		65, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S49
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		66, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S50
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S51
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S52
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S53
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S54
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S55
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S56
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S57
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S58
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S59
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S60
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S61
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S62
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S63
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S64
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S65
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S66
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S67
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		7, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		8, // IfStmt
		10, // Question
		-1, // Block
		70, // StmtList
		12, // Stmt
		

	},
	gotoRow{ // S68
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S69
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		71, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S70
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		7, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		8, // IfStmt
		10, // Question
		-1, // Block
		-1, // StmtList
		16, // Stmt
		

	},
	gotoRow{ // S71
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S72
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		7, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		8, // IfStmt
		10, // Question
		-1, // Block
		74, // StmtList
		12, // Stmt
		

	},
	gotoRow{ // S73
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	gotoRow{ // S74
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		7, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		8, // IfStmt
		10, // Question
		-1, // Block
		-1, // StmtList
		16, // Stmt
		

	},
	gotoRow{ // S75
		
		-1, // S'
		-1, // Form
		-1, // Type
		-1, // Literal
		-1, // NumLiteral
		-1, // StrLiteral
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Expr
		-1, // IfStmt
		-1, // Question
		-1, // Block
		-1, // StmtList
		-1, // Stmt
		

	},
	
}
