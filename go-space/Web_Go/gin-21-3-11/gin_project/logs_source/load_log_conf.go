package logs_source

import (
	"encoding/json"
	"os"
)

// {
//     "log_dir": "logs/gin_project.log",
//     "log_level": "info"
// }

type LogConf struct {
	LogDir   string `json:"log_dir"`
	LogLevel string `json:"log_level"`
}

func LoadLogConf() *LogConf {
	var log_conf LogConf
	// get bs
	bs, err := os.ReadFile("./conf/log_conf.json")
	if err != nil {
		panic(err)
	}

	// json
	if err := json.Unmarshal(bs, &log_conf); err != nil {
		panic(err)
	}

	return &log_conf
}
