package func_builder

import (
	"scrooge/pkg/param"
	"scrooge/pkg/struct_builder"
	"scrooge/pkg/utils"
	"text/template"
)

type Body struct {
	Params    param.Params
	Returns   param.Params
	Tpl       *template.Template
	Struct    *struct_builder.Struct
	Recipient *param.Param
	Table     string
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

func (b *Body) SetStruct(str *struct_builder.Struct) *Body {
	b.Struct = str
	return b
}

func (b *Body) SetRecipient(recipient *param.Param) *Body {
	b.Recipient = recipient
	return b
}

func (b *Body) SetTable(table string) *Body {
	b.Table = table
	return b
}

func (b *Body) Generate() string {
	return utils.MustExecTemplate(b.Tpl, b)
}
