package verify

type EmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type EmailResponse struct {
	Hash string `json:"hash"`
}
