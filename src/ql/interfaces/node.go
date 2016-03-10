package interfaces

type Node interface {
	Accept(Visitor, interface{}) interface{}
}
