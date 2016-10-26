package tokenex

import "os"
import "encoding/json"
import gorequest "github.com/parnurzeal/gorequest"

func request(action string, data map[string]interface{}) string {
  baseUrl := os.Getenv("TOKENEX_BASE_URL")
  m := map[string]interface{}{
    "TokenExID": os.Getenv("TOKENEX_ID"),
    "APIKey": os.Getenv("TOKENEX_API_KEY"),
  }

  for key, value := range data {
    m[key] = value
  }

  mJson, _ := json.Marshal(m)
  request := gorequest.New()
  _, body, _ := request.Post(baseUrl + "/" + action).
    Send(string(mJson)).
    End()
  return body
}
