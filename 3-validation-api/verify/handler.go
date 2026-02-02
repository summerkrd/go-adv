package verify

import "net/http"

type Verifier struct {
	srv *http.ServeMux
}

func NewVerifier(s *http.ServeMux) *Verifier {
	return &Verifier{srv: s}
}

func (v *Verifier) Handler(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method == http.MethodGet:
		if r.URL.Path == "/verify" {

			return
		}
		http.NotFound(w, r)

	case r.Method == http.MethodPost:
		if r.URL.Path == "/validate" {

			return
		}
		http.NotFound(w, r)
	}

}
