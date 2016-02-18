package visit

type Visitor interface {
	Visit(t interface{}, e interface{}) interface{}
}
