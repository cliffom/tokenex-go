package tokenex

import "encoding/json"

func Tokenize(data string, tokenScheme int) tokenResponse {
	tData := map[string]interface{}{
		"Data":        data,
		"TokenScheme": tokenScheme,
	}
	data = request("Tokenize", tData)
	response := tokenResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}

func Detokenize(token string) valueResponse {
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("Detokenize", tData)
	response := valueResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}

func Validate(token string) validateResponse {
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("ValidateToken", tData)
	response := validateResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}

func Delete(token string) deleteResponse {
	tData := map[string]interface{}{
		"Token": token,
	}
	data := request("DeleteToken", tData)
	response := deleteResponse{}
	json.Unmarshal([]byte(data), &response)
	return response
}
