package tokenex

type (
	BaseResponse struct {
		Error           string
		ReferenceNumber string
		Success         bool
	}

	TokenResponse struct {
		BaseResponse
		Token string
	}

	ValueResponse struct {
		BaseResponse
		Value string
	}

	ValidateResponse struct {
		BaseResponse
		Valid bool
	}

	DeleteResponse struct {
		BaseResponse
	}
)
