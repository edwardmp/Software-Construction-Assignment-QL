package stmt

import "fmt"

type Stmt interface {
	Eval()
	fmt.Stringer
}
