package repo

import (
	"scrooge/pkg/struct_builder"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
)

type Repo struct {
	Struct *struct_builder.Struct
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) SetStruct(str *struct_builder.Struct) *Repo {
	r.Struct = str
	return r
}

func (r *Repo) Generate() string {
	return utils.MustExecTemplate(templates.Repo, r)
}
