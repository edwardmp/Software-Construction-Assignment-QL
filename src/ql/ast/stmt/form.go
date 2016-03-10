package stmt

import (
	"fmt"
	"ql/interfaces"
)

type Form struct {
	Identifier interfaces.VarId
	Content    StmtList
}

func (f Form) String() string {
	return fmt.Sprintf("A form with identifier %s, statement list %v\n", f.Identifier, f.Content)
}
