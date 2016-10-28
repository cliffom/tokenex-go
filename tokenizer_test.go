package tokenex

import (
  "testing"
  "os"
)

func TestTokenize(t *testing.T) {
  Initialize(
		os.Getenv("TOKENEX_BASE_URL"),
		os.Getenv("TOKENEX_ID"),
		os.Getenv("TOKENEX_API_KEY"))
  ccNum := "4242424242424242"
	token := Tokenize(ccNum, SIXTOKENFOUR)
	data := Detokenize(token.Token)
  validate := Validate(token.Token)
  delete := Delete(token.Token)

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
