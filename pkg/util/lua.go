package utl

import (
	"fmt"
	"net/http"

	lhttp "github.com/cjoudrey/gluahttp"
	lua "github.com/yuin/gopher-lua"
	ljson "layeh.com/gopher-json"
)

func printToGo(L *lua.LState) int {
	top := L.GetTop()
	args := make([]interface{}, top)

	for i := 1; i <= top; i++ {
		args[i-1] = L.Get(i).String()
	}
	fmt.Println(args...)
	return 0
}

func RunLuaScript(luaCode string) (*string, error) {

	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("print", L.NewFunction(printToGo))

	err := L.DoString(luaCode)
	if err != nil {
		fmt.Println("Error executing Lua script:", err)
		return nil, err
	}

	return nil, nil
}

func RunLuaFunc(luaCode string) (*string, error) {
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
