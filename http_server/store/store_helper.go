package store

import (
	"encoding/json"
	"fmt"
	"io"

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
