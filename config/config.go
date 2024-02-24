/*
@author: ledger
@since: 2024/1/29
*/

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Configs 使用全局的配置变量
var Configs *GConfig

// GConfig 结构体表示配置文件的结构
type GConfig struct {
	Zip     Zip     `yaml:"zip"`
	Unzip   Unzip   `yaml:"unzip"`
	TarGz   TarGz   `yaml:"tar-gz"`
	UntarGz UntarGz `yaml:"untar-gz"`
}

type Zip struct {
	Enable bool   `yaml:"enable"`
	Name   string `yaml:"name"`
	Dir    string `yaml:"dir"`
}

type Unzip struct {
	Enable bool   `yaml:"enable"`
	Name   string `yaml:"name"`
	Dir    string `yaml:"dir"`
}

type TarGz struct {
	Enable bool   `yaml:"enable"`
	Name   string `yaml:"name"`
	Dir    string `yaml:"dir"`
}

type UntarGz struct {
	Enable bool   `yaml:"enable"`
	Name   string `yaml:"name"`
	Dir    string `yaml:"dir"`
}

// LoadConfig viper读取yaml
func LoadConfig() error {
	// yaml
	vconfig := viper.New()
	//表示 先预加载匹配的环境变量
	vconfig.AutomaticEnv()
	//设置环境变量分割符，将点号和横杠替换为下划线
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	// 设置读取的配置文件
	vconfig.SetConfigName("config")
	// 添加读取的配置文件路径
	vconfig.AddConfigPath(".")
	// 读取文件类型
	vconfig.SetConfigType("yaml")
	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	if err = vconfig.Unmarshal(&Configs); err != nil {
		log.Panicln("unmarshal cng file fail " + err.Error())
	}
	// 赋值全局变量
	return err
}
