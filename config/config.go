package config

import (
	beegoConfig "github.com/astaxie/beego/config"
	"log"
	"os"
	"path"
)

var (
	iniConf beegoConfig.Configer
)

func init() {
	confPath := path.Join(path.Dir(os.Args[0]), "app.conf")
	var err error
	iniConf, err = beegoConfig.NewConfig("ini", confPath)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func GetConfig() beegoConfig.Configer {
	return iniConf
}
