package visit

type Visitor interface {
	Visit(t interface{}) interface{}
}
