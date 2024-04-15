package store

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

func (f *FSStore) newTable(input io.Reader) ([]server.Player, error) {
	var table []server.Player

	err := json.NewDecoder(input).Decode(&table)
	if err != nil {
		err = fmt.Errorf("problem parsing table %v", err)
	}

	return table, err
}
