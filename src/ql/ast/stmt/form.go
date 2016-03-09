package stmt

import (
	"fmt"
	"ql/ast/vari"
)

type Form struct {
	Identifier vari.VarId
	Content    StmtList
}

func (f Form) String() string {
	return fmt.Sprintf("A form with identifier %s, statement list %v\n", f.Identifier, f.Content)
}