package verify

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"go-adv/3-validation-api/config"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Verifier struct {
	config config.Config
}

func NewVerifier(conf config.Config) *Verifier {
	return &Verifier{
		config: conf,
	}
}

var eres = &EmailResponse{}

func (v *Verifier) SendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ereq := EmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&ereq)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = IsValid(ereq)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	newHash := generateNewHash()
	eres.Hash = newHash
	e := email.NewEmail()
	e.From = fmt.Sprintf("Anatoliy <%s>", v.config.Email)
	e.To = []string{ereq.Email}
	e.HTML = []byte("http://localhost:8081/verify/" + newHash + ".html")
	err = e.Send(v.config.Address, smtp.PlainAuth("", v.config.Email, v.config.Password, "smtp.gmail.com"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (v *Verifier) VerifyHash(w http.ResponseWriter, r *http.Request) bool {
	hash := r.PathValue("hash")
	if hash != eres.Hash {
		eres.Hash = ""
		return false
	}
	return true
}

func generateNewHash() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
