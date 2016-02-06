package main

import "QL/ast/expr"
import "fmt"
import "QL/lexer"
import "QL/token"
import "io/ioutil"

func main() {
	b, err := ioutil.ReadFile("example.ql")
	if err != nil {
		panic(err)
	}

	s := lexer.NewLexer(b)

	for {

		d := s.Scan()
		if d.Type == 1 {
			break
		}

		fmt.Println(token.TokMap.Id(d.Type))
	}
}

func test() {
	s := expr.Add{Lhs: 1, Rhs: 2}
	//fmt.Print(s.Eval())

	hoi := expr.Expr(s)
	fmt.Print(hoi)
}
