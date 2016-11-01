package tokenex

import "log"

type Config struct {
	baseUrl string
	id      string
	apiKey  string
}

func (c *Config) init(baseUrl string, id string, apiKey string) {
	c.baseUrl = baseUrl
	c.id = id
	c.apiKey = apiKey
}

func (c *Config) validate() {
	if c.baseUrl == "" {
		log.Fatalf("config.baseUrl not set")
	} else if c.id == "" {
		log.Fatalf("config.id not set")
	} else if c.apiKey == "" {
		log.Fatalf("config.apiKey not set")
	}
}
