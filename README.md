# gconf
[![Build Status](https://travis-ci.org/usth/gconf.svg?branch=master)](https://travis-ci.org/usthooz/gconf)
[![Go Report Card](https://goreportcard.com/badge/github.com/usthooz/gconf)](https://goreportcard.com/report/github.com/usthooz/gconf)
[![GoDoc](http://godoc.org/github.com/usthooz/gconf?status.svg)](http://godoc.org/github.com/usthooz/gconf)

Golang 多格式配置文件加载.

### install
```
$ go get -u github.com/usthooz/gconf
```

### 功能
1. Golang配置文件加载

### 支持文件格式
- .json
- .toml
- .xml
- .yaml

### 例程
- [example](https://github.com/usthooz/gconf/example)
```
import (
	"github.com/usthooz/gconf"
	"github.com/usthooz/oozlog/go"
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
	// new conf object
	ozconf := gconf.NewConf(&gconf.Gconf{
		ConfPath: "./config.json", // 可选，默认为./config/config.yaml
		Subffix:  "", // 可选，如果不指定则自动解析文件名获取
	})
	// get config
	err := ozconf.GetConf(&conf)
	if err != nil {
		uoozg.Errorf("GetConf Err: %v", err.Error())
	}
	uoozg.Infof("Res: %v", conf)
}
```

- 输出
```
➜  example git:(master) ✗ go run example.go
2018/09/06 12:20:10 [INFO] Res: [{ooz {ooz 123}}] [logger.go:101]
```
