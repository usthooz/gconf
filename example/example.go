package main

import (
	"github.com/usthooz/gconf"
	ozlog "github.com/usthooz/oozlog/go"
)

type Config struct {
	Author string `yaml:"author"`
	Mysql  struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`
}

func main() {
	var (
		conf Config
	)
	ozconf := gconf.NewConf(&gconf.Gconf{
		ConfPath: "./config.yaml",
		Subffix:  "",
	})
	err := ozconf.GetConf(&conf)
	if err != nil {
		ozlog.Errorf("GetConf Err: %v", err.Error())
	}
	ozlog.Infof("Res: %v", conf)
}
