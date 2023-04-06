package param

import (
	"bytes"
	"fmt"
	"html/template"
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

func (p *Param) ToStructField() string {
	text := "{{if not .Embedded}}{{.Name}} {{end}}{{if .Slice}}[]{{end}}{{if .Pointer}}*{{end}}" +
		"{{$length := len .Pkg}}{{if ne $length 0}}{{.Pkg}}.{{end}}" +
		"{{.Typ}} {{$length := len .JSON}}{{if ne $length 0}}`json:\"{{.JSON}}\"`{{end}}"

	buf := new(bytes.Buffer)

	if err := template.Must(template.New("StructField").Parse(text)).Execute(buf, p); err != nil {
		fmt.Println(err)

		return ""
	}

	return buf.String()
}
