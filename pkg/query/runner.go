package query

import (
	"context"
	"encoding/json"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/lua-datasource/pkg/models"
	"github.com/grafana/lua-datasource/pkg/query/scenario"
	utl "github.com/grafana/lua-datasource/pkg/util"
)

func RunQuery(_ context.Context, settings models.PluginSettings, query backend.DataQuery) backend.DataResponse {
	response := backend.DataResponse{}

	// Unmarshal the JSON into our queryModel.
	var qm models.QueryModel

	err := json.Unmarshal(query.JSON, &qm)
	if err != nil {
		return backend.ErrDataResponse(backend.StatusBadRequest, "json unmarshal: "+err.Error())
	}

	utl.Log.Println("qm.RunnableQuery : " + qm.RunnableQuery)
	resStr, err := utl.RunLua(`
	local json = require("json")
	local http = require("http")
		
	function main()
		local response, err = http.request("GET", "https://reqres.in/api/users?page=2")
		if err then
			return nil, err
		end
		local res = response.body

		local jsonObj = json.decode(res)
		local jsonStr = json.encode(jsonObj["data"])

		print(jsonStr)

		return jsonStr
	end
	`)
	if err != nil {
		return backend.ErrDataResponse(backend.StatusInternal, " Query run failed "+err.Error())
	}
	utl.Log.Println("resultStr:", *resStr)

	var jsonData []map[string]string

	json.Unmarshal([]byte(*resStr), &jsonData)

	finalData := make(map[string][]string)

	for _, jsonD := range jsonData {
		for k, v := range jsonD {
			finalData[k] = append(finalData[k], v)
		}
	}

	frame := scenario.NewDataFrame(query, finalData)
	if frame == nil {
		return response
	}
	frame.RefID = query.RefID

	frame.Meta = &data.FrameMeta{
		ExecutedQueryString: qm.RunnableQuery,
	}

	// Add the frames to the response.
	response.Frames = append(response.Frames, frame)

	return response
}
