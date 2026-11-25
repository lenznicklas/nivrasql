package sqlcommands

func (db *Database) CreateTable(name string, columns []Column) error {
	db.Tables[name] = &Table{
		Name:    name,
		Columns: columns,
		Rows:    []Row{},
	}

	return nil
}
