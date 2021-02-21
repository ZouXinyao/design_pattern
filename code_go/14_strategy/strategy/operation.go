package strategy

type Strategy interface {
	DoOperation(a, b int) int
}

type OperationAdd struct {

}

func NewAdd() *OperationAdd {
	return &OperationAdd{}
}

func (o *OperationAdd) DoOperation(a, b int) int {
	return a + b
}

type OperationMultiply struct {

}

func NewMultiply() *OperationMultiply {
	return &OperationMultiply{}
}

func (o *OperationMultiply) DoOperation(a, b int) int {
	return a * b
}

type OperationSubtract struct {

}

func NewSubtract() *OperationSubtract {
	return &OperationSubtract{}
}

func (o *OperationSubtract) DoOperation(a, b int) int {
	return a - b
}
