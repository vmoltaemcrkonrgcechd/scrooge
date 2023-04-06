package struct_builder

import (
	"scrooge/pkg/param"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
)

type Struct struct {
	Name   string
	Fields param.Params
}

func New() *Struct {
	return &Struct{}
}

func (s *Struct) SetName(name string) *Struct {
	s.Name = name
	return s
}

func (s *Struct) AddField(field *param.Param) *Struct {
	s.Fields = append(s.Fields, field)
	return s
}

func (s *Struct) Generate() string {
	return utils.MustExecTemplate(templates.Struct, s)
}
