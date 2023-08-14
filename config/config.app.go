package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var ConfigPath string

// Config struct for webapp config
type Config struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		AppID   string `yaml:"app_id"`
		Timeout struct {
			Write int `yaml:"write"`
			Read  int `yaml:"read"`
			Idle  int `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(ConfigPath string) *Config {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(ConfigPath)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("\033[31mlinha 39 in config.go!!!!!!: %v \033[0m\n", err)
		// u.ErrorLogger.Println("linha 39 in config.go!!!!!!: ", err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		log.Fatal(err)
		fmt.Printf("\033[31mlinha 50 in config.go!!!!!!: %v \033[0m\n", err)
		// u.ErrorLogger.Println("linha 50 in config.go!!!!!!: ", err)
	}

	return config
}

// Function will be run before main()
func init() {
	// Looking for config flag
	flag.StringVar(&ConfigPath, "config", "config.yml", "../config.yml")
}
