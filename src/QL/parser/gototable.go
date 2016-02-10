
/*
*/
package parser

const numNTSymbols = 10
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S2
		
		-1, // S'
		23, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S3
		
		-1, // S'
		24, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S4
		
		-1, // S'
		25, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S6
		
		-1, // S'
		28, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Expr
		40, // Term
		9, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S14
		
		-1, // S'
		41, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S15
		
		-1, // S'
		42, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S16
		
		-1, // S'
		43, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S17
		
		-1, // S'
		44, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S18
		
		-1, // S'
		45, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S19
		
		-1, // S'
		46, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S20
		
		-1, // S'
		47, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S21
		
		-1, // S'
		48, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S22
		
		-1, // S'
		49, // Expr
		5, // Term
		9, // Factor
		7, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // Expr
		-1, // Term
		50, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // Expr
		-1, // Term
		51, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S29
		
		-1, // S'
		63, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S30
		
		-1, // S'
		64, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S31
		
		-1, // S'
		65, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S33
		
		-1, // S'
		68, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S42
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S43
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S44
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S45
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S46
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S47
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S48
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S49
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S50
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S51
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S52
		
		-1, // S'
		-1, // Expr
		69, // Term
		36, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S53
		
		-1, // S'
		70, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S54
		
		-1, // S'
		71, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S55
		
		-1, // S'
		72, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S56
		
		-1, // S'
		73, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S57
		
		-1, // S'
		74, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S58
		
		-1, // S'
		75, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S59
		
		-1, // S'
		76, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S60
		
		-1, // S'
		77, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S61
		
		-1, // S'
		78, // Expr
		32, // Term
		36, // Factor
		34, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S62
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S63
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S64
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S65
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S66
		
		-1, // S'
		-1, // Expr
		-1, // Term
		79, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S67
		
		-1, // S'
		-1, // Expr
		-1, // Term
		80, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S68
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S69
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S70
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S71
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S72
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S73
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S74
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S75
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S76
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S77
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S78
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S79
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S80
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	gotoRow{ // S81
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Bool
		-1, // Form
		-1, // Question
		-1, // Ifstmt
		-1, // Elsestmt
		-1, // Body
		

	},
	
}
