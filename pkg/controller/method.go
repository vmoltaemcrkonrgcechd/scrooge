package controller

import "scrooge/pkg/func_builder"

type Method struct {
	*func_builder.Func
	RepoMethod *func_builder.Func
}

func NewMethod() *Method {
	return &Method{
		Func: func_builder.New(),
	}
}

func (m *Method) SetRepoMethod(repoMethod *func_builder.Func) *Method {
	m.RepoMethod = repoMethod
	return m
}
