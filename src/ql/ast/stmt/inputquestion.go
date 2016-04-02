package stmt

import "ql/interfaces"

type InputQuestion struct {
	Question
}

func NewInputQuestion(label interfaces.StrLit, varDecl interfaces.VarDecl) InputQuestion {
	return InputQuestion{Question: NewQuestion(label, varDecl)}
}
