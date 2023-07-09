package utl

import (
	"fmt"
	"net/http"

	lhttp "github.com/cjoudrey/gluahttp"
	lua "github.com/yuin/gopher-lua"
	ljson "layeh.com/gopher-json"
)

func RunLua(luaCode string) (*string, error) {
	L := lua.NewState()
	defer L.Close()

	// Preload modules
	L.PreloadModule("http", lhttp.NewHttpModule(&http.Client{}).Loader)
	ljson.Preload(L)

	err := L.DoString(luaCode)
	if err != nil {
		return nil, fmt.Errorf("Error loading Lua string: %v ", err)
	}

	err = L.CallByParam(lua.P{
		Fn:      L.GetGlobal("main"),
		NRet:    1,
		Protect: true,
	})
	if err != nil {
		return nil, fmt.Errorf("Error calling Lua function: %v ", err)
	}

	result := L.Get(-1)
	L.Pop(1)
	resultStr := result.String()
	return &resultStr, nil
}
