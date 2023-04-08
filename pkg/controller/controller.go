package controller

import (
	"scrooge/pkg/func_builder"
	"scrooge/pkg/struct_builder"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
)

type Controller struct {
	Struct      *struct_builder.Struct
	Constructor *func_builder.Func
	Methods     []*Method
}

func New() *Controller {
	return &Controller{}
}

func (r *Controller) SetStruct(str *struct_builder.Struct) *Controller {
	r.Struct = str
	return r
}

func (r *Controller) SetConstructor(constructor *func_builder.Func) *Controller {
	r.Constructor = constructor
	return r
}

func (r *Controller) AddMethod(method *Method) *Controller {
	r.Methods = append(r.Methods, method)
	return r
}

func (r *Controller) Generate() string {
	return utils.MustExecTemplate(templates.Repo, r)
}
