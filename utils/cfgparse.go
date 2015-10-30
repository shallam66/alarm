package utils

import  (
	"encoding/json"
	"os"
	"fmt"
	"github.com/golang/glog"
)

type GlobalConfig struct {
	Mysql		string		`json:"mysql"`
	Log_dir		string		`json:"log_dir"` 
	Kafka_host	[]string	`json:"kafka_host"`
	Es_host		string		`json:"es_host"`
	Bind_port	string		`json:"bind_port"`
}

Config := GlobalConfig{}

func ParseConfig() {
	file, err := os.Open("../cfg.json")
	if err != nil {
		panic(err)
	}

