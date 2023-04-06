package func_builder

import (
	"scrooge/pkg/param"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
)

type Func struct {
	Name      string
	Recipient *param.Param
	Params    param.Params
	Returns   param.Params
	Body      *Body
}

func New() *Func {
	return &Func{}
}

func (f *Func) SetName(name string) *Func {
	f.Name = name
	return f
}

func (f *Func) SetRecipient(p *param.Param) *Func {
	f.Recipient = p
	return f
}

func (f *Func) SetBody(body *Body) *Func {
	f.Body = body
	return f
}

func (f *Func) AddParam(p *param.Param) *Func {
	f.Params = append(f.Params, p)
	return f
}

func (f *Func) AddReturn(p *param.Param) *Func {
	f.Returns = append(f.Params, p)
	return f
}

func (f *Func) Generate() string {
	return utils.MustExecTemplate(templates.Func, f)
}
