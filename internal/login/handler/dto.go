package handler

type (
	LoginCommand struct {
		CustomerID string `json:"customer_id"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)

func toLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		Token: token,
	}
}
