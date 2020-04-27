package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type config struct {
	Database struct {
		Driver   string `yaml:"driver" envconfig:"DB_DRIVER"`
		Database string `yaml:"database" envconfig:"DB_DATABASE"`
	} `yaml:"database"`
}

var (
	cfgpath  string
	instance *config
	once     sync.Once
)

func Load(path string) *config {
	cfgpath = path
	instance := &config{}
	instance.readFromFile()
	instance.readFromEnv()

	return instance
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func (cfg *config) readFromFile() {
	f, err := os.Open(cfgpath)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func (cfg *config) readFromEnv() {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
