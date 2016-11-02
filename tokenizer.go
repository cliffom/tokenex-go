package tokenex

func NewConfig(baseUrl string, id string, apiKey string) Config {
	var config Config
	config.init(baseUrl, id, apiKey)
	return config
}

func Tokenize(data string, tokenScheme int, config Config) (TokenResponse, error) {
	response := TokenResponse{}
	response.Data = map[string]interface{}{
		"Action":      "Tokenize",
		"Data":        data,
		"TokenScheme": tokenScheme,
	}
	err := response.result(&response, config)
	return response, err
}

func TokenizeFromEncryptedValue(data string, tokenScheme int, config Config) (TokenResponse, error) {
	response := TokenResponse{}
	response.Data = map[string]interface{}{
		"Action":       "TokenizeFromEncryptedValue",
		"EcryptedData": data,
		"TokenScheme":  tokenScheme,
	}
	err := response.result(&response, config)
	return response, err
}

func Detokenize(token string, config Config) (ValueResponse, error) {
	response := ValueResponse{}
	response.Data = map[string]interface{}{
		"Action": "Detokenize",
		"Token":  token,
	}
	err := response.result(&response, config)
	return response, err
}

func Validate(token string, config Config) (ValidateResponse, error) {
	response := ValidateResponse{}
	response.Data = map[string]interface{}{
		"Action": "ValidateToken",
		"Token":  token,
	}
	err := response.result(&response, config)
	return response, err
}

func Delete(token string, config Config) (DeleteResponse, error) {
	response := DeleteResponse{}
	response.Data = map[string]interface{}{
		"Action": "DeleteToken",
		"Token":  token,
	}
	err := response.result(&response, config)
	return response, err
}
