package gopfl

import (
	"fmt"
	"net/http"

	"github.com/chadbohannan/gae-session-store/gaess"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"google.golang.org/appengine"
)

const (
	apiKey         = "apikey"
	credentialsKey = "credentials"
	remoteHostKey  = "remoteHost"
	productsAPI    = "%s/products?apikey=%s"    // 2 params
	productAPI     = "%s/products/%s?apikey=%s" // 3 params
)

// TODO read product details
// TODO handle custom templates
// TODO place orders

// Init maps HTTP routes to handlers.
// router is  returned for chaining calls.
func Init(router *mux.Router) *mux.Router {
	gaess.SessionRoute(router, "GET", "/api/session", GetSessionUser, false)
	gaess.SessionRoute(router, "POST", "/api/session", PostSessionUser, false)
	gaess.SessionRoute(router, "GET", "/api/clear_session", ClearSessionUser, true)
	gaess.SessionRoute(router, "GET", "/api/products", GetProducts, false)
	gaess.SessionRoute(router, "GET", "/api/products/{ID}", GetProductDetail, false)
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
	c := appengine.NewContext(r)
	hostName, creds, apikey, err := getAccessParameters(c)
	if err != nil {
		WriteJSONError(c, w, err.Error())
		return
	}

	url := fmt.Sprintf(productsAPI, hostName, apikey)
	content, err := GetURLContentBasicAuth(c, url, creds)
	if err != nil {
		WriteJSONError(c, w, "GET products failed:"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(content)
}

// GetProductDetail fetches a more detailed product description than the GetProducts list
func GetProductDetail(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	c := appengine.NewContext(r)
	hostName, creds, apikey, err := getAccessParameters(c)
	if err != nil {
		WriteJSONError(c, w, err.Error())
		return
	}

	productID := mux.Vars(r)["ID"]
	url := fmt.Sprintf(productAPI, hostName, productID, apikey)
	content, err := GetURLContentBasicAuth(c, url, creds)
	if err != nil {
		WriteJSONError(c, w, "GET product "+productID+" failed:"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(content)
}
