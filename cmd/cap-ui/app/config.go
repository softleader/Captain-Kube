package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	DefaultValue struct {
		CaptainUrl     string   `yaml:"captainUrl"`
		UiPort         int      `yaml:"uiPort"`
		Inventory      string   `yaml:"inventory"`
		Tags           []string `yaml:"tags"`
		Namespace      string   `yaml:"namespace"`
		Version        string   `yaml:"version"`
		Verbose        bool     `yaml:"verbose"`
		SourceRegistry string   `yaml:"sourceRegistry"`
		Registry       string   `yaml:"registry"`
	} `yaml:"defaultValue"`
}

func GetConfig(configYamlPath string) (c *config, err error) {
	raw, err := ioutil.ReadFile(configYamlPath)
	if err != nil {
		log.Fatalln("load config '", configYamlPath, "' failed, abort to up serve, error: ", err)
	}

	yaml.Unmarshal(raw, &c)
	return
}
