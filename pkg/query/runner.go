package query

import (
	"context"
	"encoding/json"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
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

	utl.Log.Println("Query to run: ", qm.RawQuery)

	frame, backendResponse := scenario.NewDataFrame(query, qm)
	if backendResponse != nil {
		return *backendResponse
	}
	if frame == nil {
		return response
	}
	frame.RefID = query.RefID

	// Add the frames to the response.
	response.Frames = append(response.Frames, frame)

	return response
}
