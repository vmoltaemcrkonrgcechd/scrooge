package converter

import "strings"

type Converter struct {
	goDict map[string]string
}

func New() *Converter {
	c := &Converter{}

	c.goDictInit()

	return c
}

// PgTypeToGo smallint = int, text = string.
func (c *Converter) PgTypeToGo(typ string) string {
	return c.goDict[typ]
}

// PgToPascalCase units_in_stock = UnitsInStock, name = Name.
func (c *Converter) PgToPascalCase(fieldName string) (r string) {
	for _, w := range strings.Split(fieldName, "_") {
		r += strings.ToUpper(w[:1]) + w[1:]
	}

	return r
}

// PgToCamelCase units_in_stock = unitsInStock, name = name.
func (c *Converter) PgToCamelCase(fieldName string) string {
	r := c.PgToPascalCase(fieldName)

	return strings.ToLower(r[:1]) + r[1:]
}
