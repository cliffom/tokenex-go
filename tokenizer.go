// Package tokenex provides a set of methods for interacting with the TokenEx
// api (http://docs.tokenex.com/#tokenex-api-token-services).
package tokenex

// NewConfig returns a Config structure to be used for making reqs to the
// TokenEx API.
func NewConfig(baseUrl string, id string, apiKey string) Config {
	var config Config
	config.init(baseUrl, id, apiKey)
	return config
}

// Tokenize returns the tokenized representation of d using the given
// token scheme t
func Tokenize(d string, t int, c Config) (TokenResponse, error) {
	res := TokenResponse{}
	req := map[string]interface{}{
		"Action":      "Tokenize",
		"Data":        d,
		"TokenScheme": t,
	}
	err := res.get(&res, req, c)
	return res, err
}

// TokenizeFromEncryptedValue returns the tokenized representation of d using
// the given token scheme t where d was previously encrypted through browser
// based encryption.
func TokenizeFromEncryptedValue(d string, t int, c Config) (TokenResponse, error) {
	res := TokenResponse{}
	req := map[string]interface{}{
		"Action":       "TokenizeFromEncryptedValue",
		"EcryptedData": d,
		"TokenScheme":  t,
	}
	err := res.get(&res, req, c)
	return res, err
}

// Detokenize returns the sensitive data from the given token t
func Detokenize(t string, c Config) (ValueResponse, error) {
	res := ValueResponse{}
	req := map[string]interface{}{
		"Action": "Detokenize",
		"Token":  t,
	}
	err := res.get(&res, req, c)
	return res, err
}

// Validate returns whether the given token exists in the token vault
func Validate(t string, c Config) (ValidateResponse, error) {
	res := ValidateResponse{}
	req := map[string]interface{}{
		"Action": "ValidateToken",
		"Token":  t,
	}
	err := res.get(&res, req, c)
	return res, err
}

// Delete removes the given token from the token vault
func Delete(t string, c Config) (DeleteResponse, error) {
	res := DeleteResponse{}
	req := map[string]interface{}{
		"Action": "DeleteToken",
		"Token":  t,
	}
	err := res.get(&res, req, c)
	return res, err
}
