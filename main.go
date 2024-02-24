package main

import (
	"goZipper/config"
	"goZipper/log"
	"goZipper/util"
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Println("配置文件加载失败")
	}
	log.InitLogStyle()
	log.Println("zip enable: ", config.Configs.Zip.Enable)
	log.Println("unzip enable: ", config.Configs.Unzip.Enable)
}

func main() {
	if config.Configs.Zip.Enable {
		err := util.Zip(
			config.Configs.Zip.Name,
			config.Configs.Zip.Dir,
		)

		if err != nil {
			log.Println(err)
		}
	}
	if config.Configs.Unzip.Enable {
		err := util.Unzip(
			config.Configs.Unzip.Name,
			config.Configs.Unzip.Dir)
		if err != nil {
			log.Println(err)
		}
	}
}
