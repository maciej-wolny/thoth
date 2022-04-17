package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Form3ClientConfig struct {
	BaseURL   string
	Transport http.RoundTripper
	Timeout   time.Duration
}

var defaultConfig = Form3ClientConfig{
	BaseURL:   "http://localhost:8080/v1/",
	Transport: http.DefaultTransport,
	Timeout:   3 * time.Second,
}

func NewDefaultConfig() *Form3ClientConfig {
	return &defaultConfig
}

func GetConfig(path string) *Form3ClientConfig {
	config := defaultConfig
	file, err := os.Open(path)
	if err != nil {
		log.Printf("error reading configuration file: %s", err)
		log.Printf("using default config")
		return &config
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Printf("error parsing configuration: %s", err)
		log.Printf("using default config")
	}
	return &config
}
