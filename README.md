# tokenex-go [![CircleCI](https://circleci.com/gh/cliffom/tokenex-go.svg?style=svg)](https://circleci.com/gh/cliffom/tokenex-go)
A Go client for the [TokenEx](https://www.tokenex.com) API.

## TokenEx API Documentation

[Token Services](http://docs.tokenex.com/#tokenex-api-token-services)

## Installation

```bash
$ go get github.com/cliffom/tokenex-go
```

## Usage

```go
config := tokenex.NewConfig(
  "https://test-api.tokenex.com/TokenServices.svc/REST/",
  "YOUR_TOKENEX_ID",
  "YOUR_TOKENEX_API_KEY")
token, tokenErr := tokenex.Tokenize("4242424242424242", tokenex.SIXTOKENFOUR, config)
data, dataErr := tokenex.Detokenize(token.Token, config)
validate, validateErr  := tokenex.Validate(token.Token, config)
delete, deleteErr := tokenex.Delete(token.Token, config)
```

## License

tokenex-go is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
