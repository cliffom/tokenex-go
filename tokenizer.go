// Package tokenex provides a set of methods for interacting with the TokenEx
// api (http://docs.tokenex.com/#tokenex-api-token-services).
package tokenex

// NewConfig returns a Config structure to be used for making requests to the
// TokenEx API.
func NewConfig(baseUrl string, id string, apiKey string) Config {
	var config Config
	config.init(baseUrl, id, apiKey)
	return config
}

// Tokenize returns the tokenized representation of data using the given
// token scheme tokenScheme
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

// TokenizeFromEncryptedValue returns the tokenized representation of data using
// the given token scheme tokenScheme where data was previously encrypted
// through browser based encryption.
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

// Detokenize returns the sensitive data from the given token
func Detokenize(token string, config Config) (ValueResponse, error) {
	response := ValueResponse{}
	request := map[string]interface{}{
		"Action": "Detokenize",
		"Token":  token,
	}
	err := response.get(&response, request, config)
	return response, err
}

// Validate returns whether the given token exists in the token vault
func Validate(token string, config Config) (ValidateResponse, error) {
	response := ValidateResponse{}
	request := map[string]interface{}{
		"Action": "ValidateToken",
		"Token":  token,
	}
	err := response.get(&response, request, config)
	return response, err
}

// Delete removes the given token from the token vault
func Delete(token string, config Config) (DeleteResponse, error) {
	response := DeleteResponse{}
	request := map[string]interface{}{
		"Action": "DeleteToken",
		"Token":  token,
	}
	err := response.get(&response, request, config)
	return response, err
}
