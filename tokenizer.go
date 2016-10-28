package tokenex

import (
	"encoding/json"
	"errors"
)

func Tokenize(data string, tokenScheme int) (TokenResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Action":      "Tokenize",
		"Data":        data,
		"TokenScheme": tokenScheme,
	}
	response := TokenResponse{}
	data = response.get(tData)

	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func TokenizeFromEncryptedValue(data string, tokenScheme int) (TokenResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Action":       "TokenizeFromEncryptedValue",
		"EcryptedData": data,
		"TokenScheme":  tokenScheme,
	}
	response := TokenResponse{}
	data = response.get(tData)

	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func Detokenize(token string) (ValueResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Action": "Detokenize",
		"Token":  token,
	}
	response := ValueResponse{}
	data := response.get(tData)

	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func Validate(token string) (ValidateResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Action": "ValidateToken",
		"Token":  token,
	}
	response := ValidateResponse{}
	data := response.get(tData)

	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func Delete(token string) (DeleteResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Action": "DeleteToken",
		"Token":  token,
	}
	response := DeleteResponse{}
	data := response.get(tData)

	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}
