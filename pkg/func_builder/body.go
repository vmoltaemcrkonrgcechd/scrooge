package func_builder

import (
	"scrooge/pkg/param"
	"scrooge/pkg/utils"
	"text/template"
)

type Body struct {
	Params  param.Params
	Returns param.Params
	Tpl     *template.Template
}

func NewBody() *Body {
	return &Body{}
}

func (b *Body) SetParams(params param.Params) *Body {
	b.Params = params
	return b
}

func (b *Body) SetReturns(returns param.Params) *Body {
	b.Returns = returns
	return b
}

func (b *Body) SetTpl(tpl *template.Template) *Body {
	b.Tpl = tpl
	return b
}

func (b *Body) Generate() string {
	return utils.MustExecTemplate(b.Tpl, b)
}
