package auth

type SendCodeRequest struct {
	Phone string `json:"phone" validate:"required,min=10,max=15"`
}

type SendCodeResponse struct {
	SessionID string `json:"sessionId"`
}

type VerifyCodeRequest struct {
	SessionID string `json:"sessionId" validate:"required"`
	Code      int    `json:"code" validate:"required,min=1000,max=9999"`
}

type VerifyCodeResponse struct {
	Token string `json:"token"`
}
