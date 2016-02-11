package stmt

import "fmt"

type Form struct {
	Identifier string
	Content    StmtList
}

func (f Form) String() string {
	return fmt.Sprintf("A form with identifier %s, statement list %v", f.Identifier, f.Content)
}
