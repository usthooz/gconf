package oozgconf

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/usthooz/oozgconf/json"
	"github.com/usthooz/oozgconf/xml"
	"github.com/usthooz/oozgconf/yaml"
)

const (
	JsonSub string = "json"
	YamlSub string = "yaml"
	TomlSub string = "toml"
	XmlSub  string = "xml"
)

const (
	// default config path
	DefaultConfPath = "./config/config.yaml"
)

type OozGconf struct {
	// ConfPath config file path->default: ./config/config.yaml
	ConfPath string
	// Subffix config file subffix
	Subffix string
}

// NewConf new conf object
func NewConf(confParam *OozGconf) *OozGconf {
	if len(confParam.ConfPath) == 0 {
		confParam.ConfPath = DefaultConfPath
	}
	return confParam
}

/*
	confpath: config file path->default: ./config/config.yaml
	subffix: config file subffie->option
*/
func (oozConf *OozGconf) GetConf(conf interface{}) error {
	// read config file
	bs, err := ioutil.ReadFile(oozConf.ConfPath)
	if err != nil {
		return err
	}
	if len(oozConf.Subffix) == 0 {
		// get file subffix
		oozConf.Subffix = strings.Trim(path.Ext(path.Base(oozConf.ConfPath)), ".")
	}
	// check analy
	switch oozConf.Subffix {
	case JsonSub:
		err = json.Unmarshal(bs, conf)
	case TomlSub:
		err = toml.Unmarshal(bs, conf)
	case YamlSub:
		err = yaml.Unmarshal(bs, conf)
	case XmlSub:
		err = xml.Unmarshal(bs, conf)
	default:
		err = fmt.Errorf("GetConf: non support this file type...")
	}
	return err
}
