package main

import (
	"nivrasql/engine"
)

func main() {
	cols := []engine.Column{}
	cols = append(cols, engine.Column{"c1", "INT"})
	cols = append(cols, engine.Column{"c2", "BOOL"})
	db := engine.NewDatabase()
	db.CreateTable("test", cols)
	db.Insert("test", engine.Row{9, true})
	db.Save("test.json")
}
