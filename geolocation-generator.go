package main

import (
	"encoding/json"
	"log"

	c "github.com/harfmuf/geolocation-generator/conf"
)

func main() {
	conf, err := c.LoadConf("conf/params.json")
	if err != nil {
		panic(err)
	}
	text, _ := json.Marshal(conf)
	log.Printf("config is %s", text)
}
