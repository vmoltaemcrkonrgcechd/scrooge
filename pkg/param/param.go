package param

import (
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
	"strings"
)

const (
	Body uint8 = iota
	Path
	Query
)

type Param struct {
	Name     string
	Typ      string
	Pkg      string
	JSON     string
	Pointer  bool
	Slice    bool
	Embedded bool
	In       uint8
	SQLName  string
}

func New() *Param {
	return &Param{}
}

func (p *Param) SetName(name string) *Param {
	p.Name = name
	return p
}

func (p *Param) SetTyp(typ string) *Param {
	p.Typ = typ
	return p
}

func (p *Param) SetPkg(pkg string) *Param {
	p.Pkg = pkg
	return p
}

func (p *Param) SetJSON(json string) *Param {
	p.JSON = json
	return p
}

func (p *Param) SetPointer(b bool) *Param {
	p.Pointer = b
	return p
}

func (p *Param) SetSlice(b bool) *Param {
	p.Slice = b
	return p
}

func (p *Param) SetEmbedded(b bool) *Param {
	p.Embedded = b
	return p
}

func (p *Param) SetIn(in uint8) *Param {
	p.In = in
	return p
}

func (p *Param) SetSQLName(sqlName string) *Param {
	p.SQLName = sqlName
	return p
}

func (p *Param) LowerCaseName() string {
	if len(p.Name) == 0 {
		return ""
	}

	return strings.ToLower(p.Name[:1]) + p.Name[1:]
}

func (p *Param) ToStructField() string {
	return utils.MustExecTemplate(templates.StructField, p)
}

func (p *Param) ToFuncParam() string {
	return utils.MustExecTemplate(templates.FuncParam, p)
}

func (p *Param) ToRecipient() string {
	return utils.MustExecTemplate(templates.Recipient, p)
}
