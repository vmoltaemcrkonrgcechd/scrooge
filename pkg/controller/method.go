package controller

import (
	"scrooge/pkg/func_builder"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
	"strings"
)

type Method struct {
	*func_builder.Func
	Body       *MethodBody
	RepoMethod *func_builder.Func
	Typ        string
	ModName    string
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

func NewMethod(modName string) *Method {
	return &Method{
		Func:    func_builder.New(),
		ModName: modName,
	}
}

func (m *Method) SetRepoMethod(repoMethod *func_builder.Func) *Method {
	m.RepoMethod = repoMethod
	return m
}

func (m *Method) SetTyp(typ string) *Method {
	m.Typ = typ
	return m
}

func (m *Method) SetBody(body *MethodBody) *Method {
	m.Body = body
	return m
}

func (m *Method) Generate() string {
	return utils.MustExecTemplate(templates.Method, m)
}

func (b *MethodBody) Generate() string {
	return utils.MustExecTemplate(b.Tpl, b)
}

func (m *Method) LowerTyp() string {
	return strings.ToLower(m.Typ)
}
