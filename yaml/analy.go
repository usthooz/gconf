package yaml

import (
	"gopkg.in/yaml.v2"
)

// Unmarshal 解析yaml
func Unmarshal(data []byte, res interface{}) error {
	return yaml.Unmarshal(data, res)
}
