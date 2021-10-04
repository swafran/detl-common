package parsers

import (
	"encoding/json"
)

//JSONParser implements Parser; handles incoming json
type JSONParser struct{}

//Parse returns interface{} of data from json string
func (*JSONParser) Parse(input string) interface{} {
	var out interface{}
	_ = json.Unmarshal([]byte(input), &out)
	//_ = json.Unmarshal([]byte(`{"joey":"heya"}`), &out)
	return out
}
