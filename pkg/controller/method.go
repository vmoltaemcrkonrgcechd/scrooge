package controller

import (
	"scrooge/pkg/func_builder"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
)

type Method struct {
	*func_builder.Func
	Body       *MethodBody
	RepoMethod *func_builder.Func
}

type MethodBody struct {
	*func_builder.Body
	RepoMethod *func_builder.Func
}

func NewMethodBody() *MethodBody {
	return &MethodBody{
		Body: func_builder.NewBody(),
	}
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

func (m *Method) SetBody(body *MethodBody) *Method {
	m.Body = body
	return m
}

func (m *Method) Generate() string {
	return utils.MustExecTemplate(templates.Func, m)
}

func (b *MethodBody) Generate() string {
	return utils.MustExecTemplate(templates.BodyControllerAdd, b)
}
