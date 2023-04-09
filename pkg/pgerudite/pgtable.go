package pgerudite

type PgTable struct {
	Name    string
	Columns []PgColumn
}

func (t PgTable) GetColumn(columnName string) (col PgColumn) {
	for _, c := range t.Columns {
		if c.Name == columnName {
			return c
		}
	}

	return col
}

func (t PgTable) Names() (names []string) {
	for _, c := range t.Columns {
		names = append(names, c.Name)
	}

	return names
}

func (t PgTable) GetColumns(fieldNames ...string) (columns []PgColumn) {
	if fieldNames == nil {
		return t.Columns
	}

	for _, name := range fieldNames {
		for _, col := range t.Columns {
			if name == col.Name {
				columns = append(columns, col)
			}
		}
	}

	return columns
}

func (t PgTable) GetPk() *PgColumn {
	for _, col := range t.Columns {
		if col.ConstraintType != nil && *col.ConstraintType == "PRIMARY KEY" {
			return &col
		}
	}

	return nil
}
