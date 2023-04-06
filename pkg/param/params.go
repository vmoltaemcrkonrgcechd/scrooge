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

func (p Params) Generate() string {
	return utils.MustExecTemplate(templates.Params, p)
}
