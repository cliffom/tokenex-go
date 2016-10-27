package tokenex

type tokenexConfig struct {
	baseUrl string
	id      string
	apiKey  string
}

var config tokenexConfig

func Initialize(baseUrl string, id string, apiKey string) {
	config.baseUrl = baseUrl
	config.id = id
	config.apiKey = apiKey
}
