package scenario

import (
	"encoding/json"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/luaquery-datasource/pkg/models"

	utl "github.com/grafana/luaquery-datasource/pkg/util"
)

func NewDataFrame(query backend.DataQuery, qm models.QueryModel) (*data.Frame, *backend.DataResponse) {

	utl.Log.Println("QueryType ", query.QueryType)
	switch query.QueryType {
	case Table:
		return processTableQuery(qm, query)
	case Log:
		return processLogQuery(qm, query)
	}

	return nil, nil
}

func processLogQuery(qm models.QueryModel, query backend.DataQuery) (*data.Frame, *backend.DataResponse) {
	lr := utl.LuaRunner{}
	stdOut, _ := lr.RunLuaScript(qm.RawQuery)

	frame := newLuaLogFrame(stdOut)

	frame.Meta = &data.FrameMeta{
		ExecutedQueryString:    qm.RawQuery,
		PreferredVisualization: "Log",
	}
	return frame, nil
}

func processTableQuery(qm models.QueryModel, query backend.DataQuery) (*data.Frame, *backend.DataResponse) {
	lr := utl.LuaRunner{}
	resStr, err := lr.RunLuaFunc(qm.RawQuery)
	if err != nil {
		backendRes := backend.ErrDataResponse(backend.StatusInternal, " Query run failed "+err.Error())
		return nil, &backendRes
	}
	utl.Log.Println("Response Data ", *resStr)

	var jsonData []map[string]string

	json.Unmarshal([]byte(*resStr), &jsonData)

	finalData := make(map[string][]string)

	for _, jsonD := range jsonData {
		for k, v := range jsonD {
			finalData[k] = append(finalData[k], v)
		}
	}
	frame := newLuaTableFrame(query, finalData)

	frame.Meta = &data.FrameMeta{
		ExecutedQueryString: qm.RawQuery,
	}
	return frame, nil
}
