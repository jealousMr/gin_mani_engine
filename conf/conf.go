package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Configs struct {
	Server struct{
		Port string `yaml:"port"`
		Ip string `yaml:"ip"`
	}
	Database struct{
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
	Router struct{
		DefaultAllActionFile string `yaml:"default_all_action_file"`
		OpenAllActionFile string `yaml:"open_all_action_file"`
		DefaultImageActionFile string `yaml:"default_image_action_file"`
	}
}

func GetConfig() *Configs{
	config := &Configs{}
	content, err := ioutil.ReadFile("/home/xyl/src/gin_mani_engine/conf/meta.yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return config
}