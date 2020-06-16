package main

import (
	"github.com/swxctx/xlog"
	"github.com/usthooz/gconf"
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
		xlog.Errorf("GetConf Err: %v", err.Error())
	}
	xlog.Infof("Res: %v", conf)
}
