package tokenex

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"log"
)

type (
	BaseResponse struct {
		Error           string
		ReferenceNumber string
		Success         bool
	}

	TokenResponse struct {
		BaseResponse
		Token string
	}

	ValueResponse struct {
		BaseResponse
		Value string
	}

	ValidateResponse struct {
		BaseResponse
		Valid bool
	}

	DeleteResponse struct {
		BaseResponse
	}
)

func (b *BaseResponse) get(data map[string]interface{}) string {
	if config.baseUrl == "" {
		log.Fatalf("config.baseUrl not set")
	} else if config.id == "" {
		log.Fatalf("config.id not set")
	} else if config.apiKey == "" {
		log.Fatalf("config.apiKey not set")
	}
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
	_, body, _ := request.Post(baseUrl + "/" + data["Action"].(string)).
		Send(string(mJson)).
		End()
	return body
}
