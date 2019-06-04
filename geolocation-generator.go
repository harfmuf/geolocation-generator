package main

import (
	"encoding/json"
	"log"

	c "github.com/harfmuf/geolocation-generator/conf"
)

func main() {
	conf := c.LoadConf("conf/params.json")
	text, _ := json.Marshal(conf)
	log.Printf("config is %s", text)
}
