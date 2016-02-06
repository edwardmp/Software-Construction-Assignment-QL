
/*
*/
package parser

const numNTSymbols = 3
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Expr
		7, // Val
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S2
		
		-1, // S'
		22, // Expr
		7, // Val
		

	},
	gotoRow{ // S3
		
		-1, // S'
		23, // Expr
		7, // Val
		

	},
	gotoRow{ // S4
		
		-1, // S'
		24, // Expr
		7, // Val
		

	},
	gotoRow{ // S5
		
		-1, // S'
		25, // Expr
		31, // Val
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S10
		
		-1, // S'
		34, // Expr
		7, // Val
		

	},
	gotoRow{ // S11
		
		-1, // S'
		35, // Expr
		7, // Val
		

	},
	gotoRow{ // S12
		
		-1, // S'
		36, // Expr
		7, // Val
		

	},
	gotoRow{ // S13
		
		-1, // S'
		37, // Expr
		7, // Val
		

	},
	gotoRow{ // S14
		
		-1, // S'
		38, // Expr
		7, // Val
		

	},
	gotoRow{ // S15
		
		-1, // S'
		39, // Expr
		7, // Val
		

	},
	gotoRow{ // S16
		
		-1, // S'
		40, // Expr
		7, // Val
		

	},
	gotoRow{ // S17
		
		-1, // S'
		41, // Expr
		7, // Val
		

	},
	gotoRow{ // S18
		
		-1, // S'
		42, // Expr
		7, // Val
		

	},
	gotoRow{ // S19
		
		-1, // S'
		43, // Expr
		7, // Val
		

	},
	gotoRow{ // S20
		
		-1, // S'
		44, // Expr
		7, // Val
		

	},
	gotoRow{ // S21
		
		-1, // S'
		45, // Expr
		7, // Val
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S26
		
		-1, // S'
		59, // Expr
		31, // Val
		

	},
	gotoRow{ // S27
		
		-1, // S'
		60, // Expr
		31, // Val
		

	},
	gotoRow{ // S28
		
		-1, // S'
		61, // Expr
		31, // Val
		

	},
	gotoRow{ // S29
		
		-1, // S'
		62, // Expr
		31, // Val
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S42
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S43
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S44
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S45
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S46
		
		-1, // S'
		63, // Expr
		31, // Val
		

	},
	gotoRow{ // S47
		
		-1, // S'
		64, // Expr
		31, // Val
		

	},
	gotoRow{ // S48
		
		-1, // S'
		65, // Expr
		31, // Val
		

	},
	gotoRow{ // S49
		
		-1, // S'
		66, // Expr
		31, // Val
		

	},
	gotoRow{ // S50
		
		-1, // S'
		67, // Expr
		31, // Val
		

	},
	gotoRow{ // S51
		
		-1, // S'
		68, // Expr
		31, // Val
		

	},
	gotoRow{ // S52
		
		-1, // S'
		69, // Expr
		31, // Val
		

	},
	gotoRow{ // S53
		
		-1, // S'
		70, // Expr
		31, // Val
		

	},
	gotoRow{ // S54
		
		-1, // S'
		71, // Expr
		31, // Val
		

	},
	gotoRow{ // S55
		
		-1, // S'
		72, // Expr
		31, // Val
		

	},
	gotoRow{ // S56
		
		-1, // S'
		73, // Expr
		31, // Val
		

	},
	gotoRow{ // S57
		
		-1, // S'
		74, // Expr
		31, // Val
		

	},
	gotoRow{ // S58
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S59
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S60
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S61
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S62
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S63
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S64
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S65
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S66
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S67
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S68
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S69
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S70
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S71
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S72
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S73
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S74
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	gotoRow{ // S75
		
		-1, // S'
		-1, // Expr
		-1, // Val
		

	},
	
}
