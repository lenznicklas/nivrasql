package sqlcommands

import (
	"encoding/json"
	"os"
)

func (db *Database) Save(path string) error {
	data, err := json.MarshalIndent(db, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
