package param

import (
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
)

type Params []*Param

func (p Params) Body() *Param {
	for _, val := range p {
		if val.In == Body {
			return val
		}
	}

	return nil
}

func (p Params) Path() *Param {
	for _, val := range p {
		if val.In == Path {
			return val
		}
	}

	return nil
}

func (p Params) Generate() string {
	return utils.MustExecTemplate(templates.Params, p)
}

func (p Params) ParamNames() string {
	return utils.MustExecTemplate(templates.ParamNames, p)
}

func (p Params) ParamSQLNames() string {
	return utils.MustExecTemplate(templates.ParamSQLNames, p)
}

func (p Params) Types() string {
	return utils.MustExecTemplate(templates.Repo, p)
}
