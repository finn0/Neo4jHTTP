package conf

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config sets up parameters for db connection
type Config struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Path     string `json:"path"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetConfig reads configuration from json file
func GetConfig() *Config {
	fp, _ := filepath.Abs("../conf/conf.json")
	f, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	dc := json.NewDecoder(f)
	cfg := &Config{}
	err = dc.Decode(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

// URL returns a request url
func (c *Config) URL() string {
	return "http://" + c.Hostname + ":" + c.Port + c.Path
}
