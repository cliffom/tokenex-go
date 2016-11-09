package tokenex

func NewConfig(baseUrl string, id string, apiKey string) Config {
	var config Config
	config.init(baseUrl, id, apiKey)
	return config
}

func Tokenize(data string, tokenScheme int, config Config) (TokenResponse, error) {
	response := TokenResponse{}
	request := map[string]interface{}{
		"Action":      "Tokenize",
		"Data":        data,
		"TokenScheme": tokenScheme,
	}
	err := response.get(&response, request, config)
	return response, err
}

func TokenizeFromEncryptedValue(data string, tokenScheme int, config Config) (TokenResponse, error) {
	response := TokenResponse{}
	request := map[string]interface{}{
		"Action":       "TokenizeFromEncryptedValue",
		"EcryptedData": data,
		"TokenScheme":  tokenScheme,
	}

	err := response.get(&response, request, config)
	return response, err
}

func Detokenize(token string, config Config) (ValueResponse, error) {
	response := ValueResponse{}
	request := map[string]interface{}{
		"Action": "Detokenize",
		"Token":  token,
	}
	err := response.get(&response, request, config)
	return response, err
}

func Validate(token string, config Config) (ValidateResponse, error) {
	response := ValidateResponse{}
	request := map[string]interface{}{
		"Action": "ValidateToken",
		"Token":  token,
	}
	err := response.get(&response, request, config)
	return response, err
}

func Delete(token string, config Config) (DeleteResponse, error) {
	response := DeleteResponse{}
	request := map[string]interface{}{
		"Action": "DeleteToken",
		"Token":  token,
	}
	err := response.get(&response, request, config)
	return response, err
}
