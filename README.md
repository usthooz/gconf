# oozgconf
[![Build Status](https://travis-ci.org/usth/oozgconf.svg?branch=master)](https://travis-ci.org/usthooz/oozgconf)
[![Go Report Card](https://goreportcard.com/badge/github.com/usthooz/oozgconf)](https://goreportcard.com/report/github.com/usthooz/oozgconf)
[![GoDoc](http://godoc.org/github.com/usthooz/oozgconf?status.svg)](http://godoc.org/github.com/usthooz/oozgconf)

Golang 多格式配置文件加载.

### install
```
$ go get -u github.com/usthooz/oozgconf
```

### 功能
1. Golang配置文件加载

### 支持文件格式
- .json
- .toml
- .xml
- .yaml

### 例程
- [example](https://github.com/usthooz/oozgconf/example)
```
import (
	"github.com/usthooz/oozgconf"
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
	ozconf := oozgconf.NewConf(&oozgconf.OozGconf{
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
