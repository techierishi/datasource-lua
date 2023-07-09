package scenario

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func newLuaTableFrame(query backend.DataQuery, values map[string][]string) *data.Frame {

	fields := []*data.Field{}

	for k, v := range values {
		fields = append(fields, data.NewField(k, data.Labels{}, v))
	}

	return data.NewFrame("data", fields...)
}
