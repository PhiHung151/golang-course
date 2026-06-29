package service

type GenPass interface{}

type genPass struct{}

func NewGenPass() GenPass {
	return &genPass{}
}
