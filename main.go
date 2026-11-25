package main

import sqlcommands "nivrasql/sql_commands"

func main() {
	cols := []sqlcommands.Column{}
	cols = append(cols, sqlcommands.Column{"c1", "INT"})
	cols = append(cols, sqlcommands.Column{"c2", "BOOL"})
	db := sqlcommands.NewDatabase()
	db.CreateTable("test", cols)
	db.Insert("test", sqlcommands.Row{9, true})
	db.Save("test.json")
}
