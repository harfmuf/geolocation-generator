package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadConf(confPath string) *Conf {
	byteValue, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalf("cannot load config file %s", confPath)
	}
	var conf Conf
	json.Unmarshal(byteValue, &conf)
	return &conf
}
