package scenario

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func NewDataFrame(query backend.DataQuery, values map[string][]string) *data.Frame {
	switch query.QueryType {
	case Table:
		return newLuaTableFrame(query, values)
	}

	return nil
}
