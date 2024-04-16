package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/johanesalxd/lgt-project/http_server/model"
)

func newTable(input io.Reader) (model.League, error) {
	var table model.League

	err := json.NewDecoder(input).Decode(&table)
	if err != nil {
		err = fmt.Errorf("problem parsing table %v", err)
	}

	return table, err
}

func initDB(db *os.File) error {
	db.Seek(0, io.SeekStart)

	info, err := db.Stat()
	if err != nil {
		return fmt.Errorf("can't get info from file %s with error %v", db.Name(), err)
	}

	if info.Size() == 0 {
		db.Write([]byte("[]"))
		db.Seek(0, io.SeekStart)
	}

	return nil
}
