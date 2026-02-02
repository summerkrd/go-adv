package verify

import (
	"go-adv/3-validation-api/config"
	"net/http"
)

type Verifier struct {
	config config.Config
}

func NewVerifier(conf config.Config) *Verifier {
	return &Verifier{
		config: conf,
	}
}

func (v *Verifier) SendEmail(w http.ResponseWriter, r *http.Request) {

}

func (v *Verifier) VerifyHash(w http.ResponseWriter, r *http.Request) {

}
