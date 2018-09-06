package json

import (
	"bytes"
	"encoding/json"
)

// Unmarshal json格式解析
func Unmarshal(data []byte, res interface{}) error {
	decod := json.NewDecoder(bytes.NewReader(data))
	decod.UseNumber()
	return decod.Decode(res)
}
