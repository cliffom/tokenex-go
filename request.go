package tokenex

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

func request(action string, data map[string]interface{}) string {
	baseUrl := config.baseUrl
	m := map[string]interface{}{
		"TokenExID": config.id,
		"APIKey":    config.apiKey,
	}

	for key, value := range data {
		m[key] = value
	}

	mJson, _ := json.Marshal(m)
	request := gorequest.New()
	_, body, _ := request.Post(baseUrl + "/" + action).
		Send(string(mJson)).
		End()
	return body
}
