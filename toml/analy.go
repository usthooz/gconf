package toml

import (
	"github.com/BurntSushi/toml"
)

// Unmarshal 解析toml
func Unmarshal(data []byte, res interface{}) error {
	_, err := toml.Decode(string(data), res)
	return err
}
