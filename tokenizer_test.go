package tokenex

import (
	"os"
	"testing"
)

func TestTokenize(t *testing.T) {
	teConfig := NewConfig(
		os.Getenv("TOKENEX_BASE_URL"),
		os.Getenv("TOKENEX_ID"),
		os.Getenv("TOKENEX_API_KEY"))
	ccNum := "4242424242424242"
	token, tokenErr := Tokenize(ccNum, SIXTOKENFOUR, teConfig)
	data, dataErr := Detokenize(token.Token, teConfig)
	validate, validateErr := Validate(token.Token, teConfig)
	delete, deleteErr := Delete(token.Token, teConfig)

	if tokenErr != nil {
		t.Errorf(string(tokenErr.Error()))
	}

	if dataErr != nil {
		t.Errorf(string(dataErr.Error()))
	}

	if validateErr != nil {
		t.Errorf(string(validateErr.Error()))
	}

	if deleteErr != nil {
		t.Errorf(string(deleteErr.Error()))
	}

	if data.Value != ccNum {
		t.Errorf("Detokenized value is not equal to original value")
	}

	if !validate.Success {
		t.Errorf("Generated token is not valid")
	}

	if !delete.Success {
		t.Errorf("Could not delete generated token")
	}
}
