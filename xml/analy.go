package xml

import (
	"encoding/xml"
)

// Unmarshal 解析xml
func Unmarshal(data []byte, res interface{}) error {
	return xml.Unmarshal(data, res)
}
