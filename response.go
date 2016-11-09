package tokenex

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
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

func (b *BaseResponse) get(v interface{}, request map[string]interface{}, config Config) error {
	err := json.Unmarshal([]byte(b.request(request, config)), &v)
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

func (b *BaseResponse) request(data map[string]interface{}, config Config) string {
	config.validate()
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
