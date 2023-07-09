# Lua as Query

### Frontend

1. Install dependencies

   ```bash
   yarn install
   ```

2. Build plugin in development mode or run in watch mode

   ```bash
   yarn dev
   ```

   or

   ```bash
   yarn watch
   ```

3. Build plugin in production mode

   ```bash
   yarn build
   ```

### Backend

   ```bash
   go get -u github.com/grafana/grafana-plugin-sdk-go
   go mod tidy
   ```

2. Build backend plugin binaries for Linux, Windows and Darwin:

   ```bash
   mage -v
   ```

3. List all available Mage targets for additional commands:

   ```bash
   mage -l
   ```


#### Example Lua query

```lua
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

   return jsonStr
end

```