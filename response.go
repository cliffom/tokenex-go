package tokenex

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"log"
)

type (
	BaseResponse struct {
		Data            map[string]interface{}
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

func (b *BaseResponse) result(v interface{}) error {
	err := json.Unmarshal([]byte(request(b.Data)), &v)
	errStr := ""
	if err == nil {
		switch v.(type) {
		case *TokenResponse:
			errStr = v.(*TokenResponse).Error
		case *ValueResponse:
			errStr = v.(*ValueResponse).Error
		case *ValidateResponse:
			errStr = v.(*ValidateResponse).Error
		case *DeleteResponse:
			errStr = v.(*DeleteResponse).Error
		}
		if errStr != "" {
			err = errors.New(errStr)
		}
	}

	return err
}

func request(data map[string]interface{}) string {
	validateConfig()
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

func validateConfig() {
	if config.baseUrl == "" {
		log.Fatalf("config.baseUrl not set")
	} else if config.id == "" {
		log.Fatalf("config.id not set")
	} else if config.apiKey == "" {
		log.Fatalf("config.apiKey not set")
	}
}
