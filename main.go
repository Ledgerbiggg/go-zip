package main

import (
	"encoding/json"
	"goZipper/config"
	"goZipper/log"
	"goZipper/util"
	"sync"
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Println("配置文件加载失败")
	}
	log.InitLogStyle()
	log.Println("配置文件加载成功")
	marshal, err := json.Marshal(config.Configs)
	log.Println("配置文件内容: ", string(marshal))
	log.Println("zip enable: ", config.Configs.Zip.Enable)
	log.Println("unzip enable: ", config.Configs.Unzip.Enable)
	log.Println("tar.gz enable: ", config.Configs.TarGz.Enable)
	log.Println("untar.gz enable: ", config.Configs.UntarGz.Enable)
}

func main() {
	group := sync.WaitGroup{}
	group.Add(4)
	if config.Configs.Zip.Enable {
		go func() {
			defer group.Done()
			err := util.Zip(
				config.Configs.Zip.Name,
				config.Configs.Zip.Dir,
			)
			if err != nil {
				log.Println(err)
			}
		}()
	} else {
		group.Done()
	}
	if config.Configs.Unzip.Enable {
		go func() {
			defer group.Done()
			err := util.Unzip(
				config.Configs.Unzip.Name,
				config.Configs.Unzip.Dir)
			if err != nil {
				log.Println(err)
			}
		}()
	} else {
		group.Done()
	}
	if config.Configs.TarGz.Enable {
		go func() {
			defer group.Done()
			err := util.TarGz(
				config.Configs.TarGz.Name,
				config.Configs.TarGz.Dir,
			)
			if err != nil {
				log.Println(err)
			}
		}()
	} else {
		group.Done()
	}
	if config.Configs.UntarGz.Enable {
		go func() {
			defer group.Done()
			err := util.UnTarGz(
				config.Configs.UntarGz.Name,
				config.Configs.UntarGz.Dir)
			if err != nil {
				log.Println(err)
			}

		}()
	} else {
		group.Done()
	}

	group.Wait()
}
