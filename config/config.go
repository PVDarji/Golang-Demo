package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

//GetConfig  config
func GetConfig() (*Configuration, error) {
	filename := "./config/config.json"
	flag.Parse()
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Print(err)
	}

	var configuration Configuration
	err = json.Unmarshal(data, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
