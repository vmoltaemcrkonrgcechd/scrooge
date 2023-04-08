package pgerudite

type PgColumn struct {
	Name, Type                                   string
	IsNullable                                   bool
	Default, ConstraintType, RefTable, RefColumn *string
}
