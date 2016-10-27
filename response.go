package tokenex

type (
	baseResponse struct {
		Error           string
		ReferenceNumber string
		Success         bool
	}

	tokenResponse struct {
		baseResponse
		Token string
	}

	valueResponse struct {
		baseResponse
		Value string
	}

	validateResponse struct {
		baseResponse
		Valid bool
	}

	deleteResponse struct {
		baseResponse
	}
)
