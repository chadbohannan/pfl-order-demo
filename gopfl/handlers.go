package gopfl

import (
	"net/http"

	"github.com/chadbohannan/gae-session-store/gaess"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"google.golang.org/appengine"
)

const (
	apiKey         = "apikey"
	credentialsKey = "credentials"
	productsAPI    = "https://testapi.pfl.com/products?apikey="
)

// TODO read product details
// TODO handle custom templates
// TODO place orders

// Init maps HTTP routes to handlers.
// router is  returned for chaining calls.
func Init(router *mux.Router) *mux.Router {
	gaess.SessionRoute(router, "GET", "/session", GetSessionUser, false)
	gaess.SessionRoute(router, "POST", "/session", PostSessionUser, false)
	gaess.SessionRoute(router, "GET", "/clear_session", ClearSessionUser, true)
	gaess.SessionRoute(router, "GET", "/products", GetProducts, false)
	gaess.SessionRoute(router, "GET", "/products/{ID}", GetProductDetail, true)
	return router
}

// GetSessionUser returns the User data associated with the given session
func GetSessionUser(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
}

// PostSessionUser creates or updates a User for the given session
func PostSessionUser(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
}

// ClearSessionUser dissassociates a User from the active sesion
func ClearSessionUser(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
}

// GetProducts is a gaess.EndpointHandler type function
func GetProducts(w http.ResponseWriter, r *http.Request, session *sessions.Session) {

	// TODO if session is new create a new user

	c := appengine.NewContext(r)
	creds, err := GetSetting(c, credentialsKey)
	if err != nil {
		WriteJSONError(c, w, "service configuration A missing, "+err.Error())
		return
	}

	apikey, err := GetSetting(c, apiKey)
	if err != nil {
		WriteJSONError(c, w, "service configuration B missing, "+err.Error())
		return
	}

	url := productsAPI + apikey
	content, err := GetURLContentBasicAuth(c, url, creds)
	if err != nil {
		WriteJSONError(c, w, "products GET failed:"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(content)
}

// GetProductDetail fetches a more detailed product description than the GetProducts list
func GetProductDetail(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	// productID := mux.Vars(r)["ID"]

}
