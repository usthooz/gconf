package yaml

import (
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Unmarshal 解析yaml
func Unmarshal(data []byte, res interface{}) error {
	var (
		conf interface{}
	)
	yaml.Unmarshal(data, &conf)
	return mapstructure.Decode(conf, res)
}
