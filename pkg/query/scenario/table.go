package scenario

import (
	"fmt"

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

func newLuaLogFrame(values []interface{}) *data.Frame {

	fields := []*data.Field{}

	strings := make([]string, len(values))
	for i, v := range values {
		strings[i] = fmt.Sprintf("%v", v)
	}

	fields = append(fields, data.NewField("log", data.Labels{}, strings))

	frame := data.NewFrame("data", fields...)

	return frame
}
