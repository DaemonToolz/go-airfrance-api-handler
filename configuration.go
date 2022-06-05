package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Config Section liée au fichier de configuration
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Key  string `json:"key"`
}

var appConfig Config
var logFile os.File

func (cfg *Config) httpListenUri() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

func loadConf(path string, output *interface{}) {
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(output)
}

func initConfiguration() {
	configFile, err := os.Open("./config/appConfig.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&appConfig)
}

func prepareLogs() {
	os.MkdirAll("./logs/", 0755)

	filename := fmt.Sprintf("./logs/%s.log", filepath.Base(os.Args[0]))
	logFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	log.SetOutput(logFile)
}

// constructHeaders : A remplacer par un Reverse-Proxy (NGINX)
func constructHeaders(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}
