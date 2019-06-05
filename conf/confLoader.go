package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadConf(confPath string) (*Conf, error) {
	byteValue, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("ERROR: Cannot load config file %s", confPath)
		return nil, err
	}
	var conf Conf
	json.Unmarshal(byteValue, &conf)
	return &conf, nil
}
