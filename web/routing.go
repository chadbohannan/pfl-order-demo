package quantitrack

import (
	"net/http"

	gopfl "github.com/chadbohannan/pfl-order-demo/gopfl"
	mx "github.com/gorilla/mux"
)

func init() {
	router := gopfl.Init(mx.NewRouter())
	http.Handle("/", router)
}
