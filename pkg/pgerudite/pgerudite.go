package pgerudite

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PgErudite struct {
	tables map[string]PgTable
}

func New(pgUrl string) (*PgErudite, error) {
	db, err := sql.Open("postgres", pgUrl)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if rows, err = db.Query(`SELECT c.table_name, c.column_name, c.data_type, c.is_nullable, c.column_default,
tc.constraint_type, kcu2.table_name as ref_table, kcu2.column_name as ref_column
FROM information_schema.columns c
LEFT JOIN information_schema.key_column_usage kcu
ON c.table_name = kcu.table_name AND c.column_name = kcu.column_name
LEFT JOIN information_schema.table_constraints tc USING(constraint_name)
LEFT JOIN information_schema.referential_constraints USING(constraint_name)
LEFT JOIN information_schema.key_column_usage kcu2
ON unique_constraint_name = kcu2.constraint_name
WHERE c.table_schema = 'public'`); err != nil {
		return nil, err
	}

	var (
		tableName, isNullable string
		col                   PgColumn
		pgErudite             = PgErudite{make(map[string]PgTable)}
	)

	for rows.Next() {
		if err = rows.Scan(&tableName, &col.Name, &col.Type, &isNullable,
			&col.Default, &col.ConstraintType, &col.RefTable,
			&col.RefColumn); err != nil {
			return nil, err
		}

		if isNullable == "YES" {
			col.IsNullable = true
		}

		t := pgErudite.tables[tableName]

		t.Columns = append(t.Columns, col)

		t.Name = tableName

		pgErudite.tables[tableName] = t
	}

	return &pgErudite, nil
}

func (pg *PgErudite) TableExists(tableName string) bool {
	_, ok := pg.tables[tableName]

	return ok
}

func (pg *PgErudite) GetTable(tableName string) PgTable {
	return pg.tables[tableName]
}

func (pg *PgErudite) GetColumn(tableName, columnName string) PgColumn {
	return pg.tables[tableName].GetColumn(columnName)
}
