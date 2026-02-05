package verify

type EmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type EmailResponse struct {
	Email string `json:"email" validate:"required,email"`
	Hash  string `json:"hash"`
}
