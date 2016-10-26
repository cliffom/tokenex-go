package tokenex

func Tokenize(data string, tokenScheme int) string {
  tData := map[string]interface{}{
    "Data": data,
    "TokenScheme": tokenScheme,
  }
  data = request("Tokenize", tData)
  return data
}
