package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/AlastorTh/avTask/internal/app/apiserver"
	"gopkg.in/yaml.v2"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path,", "configs/apiserver.yml", "path to config file")
}

func main() {
	flag.Parse()

	// loggerMgr := initZapLog()
	// zap.ReplaceGlobals(loggerMgr)
	// defer loggerMgr.Sync()
	// logger := loggerMgr.Sugar()
	// logger.Debug("START!")

	config := apiserver.NewConfig()
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		log.Fatalln(err)
	}
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}

}
