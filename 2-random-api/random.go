package random_api

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type RandomApi struct{}

func NewRandomApi() *RandomApi {
	return &RandomApi{}
}

func (api *RandomApi) StartApi(router *http.ServeMux) {
	router.HandleFunc("/random", randomHandler)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	minValue := 1
	maxValue := 6

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := rnd.Intn(maxValue-minValue+1) + minValue
	_, err := w.Write([]byte(strconv.Itoa(i)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
