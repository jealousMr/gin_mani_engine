package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configs struct {
	Server struct{
		Port string `yaml:"port"`
		Ip string `yaml:"ip"`
	}
	Database struct{
		Addr string `yaml:"addr"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
	Router struct{
		DefaultAllActionFile string
		OpenAllActionFile string
		DefaultImageActionFile string
	}
	Client struct{
		User string `yaml:"user"`
		Center string `yaml:"center"`
		Engine string `yaml:"engine"`
	}
}

func GetConfig() *Configs{
	gopath := os.Getenv("GOPATH")
	config := &Configs{}
	configUrl := fmt.Sprintf("%s/src/gin_mani_engine/conf/meta.yaml",gopath)
	content, err := ioutil.ReadFile(configUrl)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	config.Router.DefaultAllActionFile = fmt.Sprintf("%s/src/gin_mani_crm_api/static/default_all",gopath)
	config.Router.OpenAllActionFile = fmt.Sprintf("%s/src/gin_mani_crm_api/static/open_all",gopath)
	config.Router.DefaultImageActionFile = fmt.Sprintf("%s/src/gin_mani_crm_api/static/default_image",gopath)
	return config
}