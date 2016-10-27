package tokenex

import "encoding/json"

type (
	TokenExResponse struct {
		Error           string
		ReferenceNumber string
		Success         bool
	}

	TokenResponse struct {
		TokenExResponse
		Token string
	}

	ValueResponse struct {
		TokenExResponse
		Value string
	}

	ValidateResponse struct {
		TokenExResponse
		Valid bool
	}

	DeleteResponse struct {
		TokenExResponse
	}
)

func Tokenize(data string, tokenScheme int) TokenResponse {
	tData := map[string]interface{}{
		"Data":        data,
		"TokenScheme": tokenScheme,
	}
	data = request("Tokenize", tData)
	response := TokenResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}

func Detokenize(token string) ValueResponse {
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("Detokenize", tData)
	response := ValueResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}

func Validate(token string) ValidateResponse {
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("ValidateToken", tData)
	response := ValidateResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}

func Delete(token string) DeleteResponse {
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("DeleteToken", tData)
	response := DeleteResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}
