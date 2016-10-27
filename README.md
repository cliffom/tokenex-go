# tokenex-go
A Go client for the [TokenEx](https://www.tokenex.com) API.

## TokenEx API Documentation

[Token Services](http://docs.tokenex.com/#tokenex-api-token-services)

## Installation

```bash
$ go get github.com/cliffom/tokenex-go
```

## Usage

```go
tokenex.Initialize(
  "https://test-api.tokenex.com/TokenServices.svc/REST/",
  "YOUR_TOKENEX_ID",
  "YOUR_TOKENEX_API_KEY")
token := tokenex.Tokenize("4242424242424242", tokenex.SIXTOKENFOUR)
data := tokenex.Detokenize(token.Token)
isValid := tokenex.Validate(token.Token)
deleted := tokenex.Delete(token.Token)
```

## License

tokenex-go is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
