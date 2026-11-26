package main

import (
	"fmt"
	"nivrasql/engine"
)

func main() {
	cols := []engine.Column{}
	cols = append(cols, engine.Column{"c1", "INT"})
	cols = append(cols, engine.Column{"c2", "BOOL"})
	db := engine.NewDatabase()
	err := db.CreateTable("test", cols)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Insert("test", engine.Row{9, true})
	if err != nil {
		fmt.Println(err)
	}
	err = db.Save("test.json")
	if err != nil {
		fmt.Println(err)
	}
}
