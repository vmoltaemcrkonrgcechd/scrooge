package repo

import (
	"scrooge/pkg/func_builder"
	"scrooge/pkg/struct_builder"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
)

type Repo struct {
	Struct      *struct_builder.Struct
	Constructor *func_builder.Func
	Methods     []*func_builder.Func
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) SetStruct(str *struct_builder.Struct) *Repo {
	r.Struct = str
	return r
}

func (r *Repo) SetConstructor(constructor *func_builder.Func) *Repo {
	r.Constructor = constructor
	return r
}

func (r *Repo) AddMethod(method *func_builder.Func) *Repo {
	r.Methods = append(r.Methods, method)
	return r
}

func (r *Repo) Generate() string {
	return utils.MustExecTemplate(templates.Repo, r)
}
