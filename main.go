package main

import (
	"flag"
	config2 "mini-douyin-rebuild/config"
	"mini-douyin-rebuild/db"
	"mini-douyin-rebuild/utils"
)

var FlagConfig = flag.String("f", "./etc/server_config.yaml", "choose config file (.yaml)")

func main() {
	var config config2.ServerConfig
	utils.YamlParse(FlagConfig, &config)
	db.MysqlInit(&config)
	GenInit()
}
