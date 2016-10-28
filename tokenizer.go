package tokenex

import (
	"encoding/json"
	"errors"
)

func Tokenize(data string, tokenScheme int) (TokenResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Data":        data,
		"TokenScheme": tokenScheme,
	}
	data = request("Tokenize", tData)
	response := TokenResponse{}
	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func TokenizeFromEncryptedValue(data string, tokenScheme int) (TokenResponse, error) {
	var err error
	tData := map[string]interface{}{
		"EcryptedData": data,
		"TokenScheme":  tokenScheme,
	}
	data = request("TokenizeFromEncryptedValue", tData)
	response := TokenResponse{}
	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func Detokenize(token string) (ValueResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("Detokenize", tData)
	response := ValueResponse{}
	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func Validate(token string) (ValidateResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("ValidateToken", tData)
	response := ValidateResponse{}
	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}

func Delete(token string) (DeleteResponse, error) {
	var err error
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("DeleteToken", tData)
	response := DeleteResponse{}
	json.Unmarshal([]byte(data), &response)
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err
}
