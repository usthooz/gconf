package main

import (
	"github.com/usthooz/oozlog/go"

	"github.com/usthooz/oozgconf"
)

type Config struct {
	Author string
	Mysql  struct {
		User     string
		Password string
	}
}

func main() {
	var (
		conf Config
	)
	ozconf := oozgconf.NewConf(&oozgconf.OozGconf{
		ConfPath: "./config.json",
		Subffix:  "",
	})
	err := ozconf.GetConf(&conf)
	if err != nil {
		uoozg.Errorf("GetConf Err: %v", err.Error())
	}
	uoozg.Infof("Res: %v", conf)
}
