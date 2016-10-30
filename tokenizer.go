package tokenex

var config tokenexConfig

func Initialize(baseUrl string, id string, apiKey string) {
	config.init(baseUrl, id, apiKey)
}

func Tokenize(data string, tokenScheme int) (TokenResponse, error) {
	response := TokenResponse{}
	response.Data = map[string]interface{}{
		"Action":      "Tokenize",
		"Data":        data,
		"TokenScheme": tokenScheme,
	}
	err := response.result(&response)
	return response, err
}

func TokenizeFromEncryptedValue(data string, tokenScheme int) (TokenResponse, error) {
	response := TokenResponse{}
	response.Data = map[string]interface{}{
		"Action":       "TokenizeFromEncryptedValue",
		"EcryptedData": data,
		"TokenScheme":  tokenScheme,
	}
	err := response.result(&response)
	return response, err
}

func Detokenize(token string) (ValueResponse, error) {
	response := ValueResponse{}
	response.Data = map[string]interface{}{
		"Action": "Detokenize",
		"Token":  token,
	}
	err := response.result(&response)
	return response, err
}

func Validate(token string) (ValidateResponse, error) {
	response := ValidateResponse{}
	response.Data = map[string]interface{}{
		"Action": "ValidateToken",
		"Token":  token,
	}
	err := response.result(&response)
	return response, err
}

func Delete(token string) (DeleteResponse, error) {
	response := DeleteResponse{}
	response.Data = map[string]interface{}{
		"Action": "DeleteToken",
		"Token":  token,
	}
	err := response.result(&response)
	return response, err
}
