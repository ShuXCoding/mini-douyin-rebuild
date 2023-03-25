package utils

import (
	"gopkg.in/yaml.v3"
	"io"
	"log"
	config2 "mini-douyin-rebuild/config"
	"os"
)

// YamlParse 解析配置文件 etc/server_config.yaml
func YamlParse(yamlFile *string, ServerConfig *config2.ServerConfig) {

	content, err := os.ReadFile(*yamlFile)

	if err != nil && err != io.EOF {
		log.Fatal("yaml file read fail", err)
	}

	if err := yaml.Unmarshal(content, ServerConfig); err != nil {
		log.Fatal("yaml file parse fail\n", err)
	}

}
