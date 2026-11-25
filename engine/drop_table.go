package engine

import "fmt"

func (db *Database) DropTable(name string) error {
	if _, exists := db.Tables[name]; !exists {
		return fmt.Errorf("table %s does not exist", name)
	}
	delete(db.Tables, name)
	return nil
}
