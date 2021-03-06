package gopfl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chadbohannan/gae-session-store/gaess"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	apiKey         = "apikey"
	credentialsKey = "credentials"
	remoteHostKey  = "remoteHost"
	productsAPI    = "%s/products?apikey=%s"    // 2 params
	productAPI     = "%s/products/%s?apikey=%s" // 3 params
	priceAPI       = "%s/price?apikey=%s"       // 2 params
	orderAPI       = "%s/orders?apikey=%s"      // 2 params
)

// Init maps HTTP routes to handlers.
// router is  returned for chaining calls.
func Init(router *mux.Router) *mux.Router {
	gaess.SessionRoute(router, "GET", "/api/session", GetSessionUser, false)
	gaess.SessionRoute(router, "POST", "/api/session", PostSessionUser, false)
	gaess.SessionRoute(router, "GET", "/api/clear_session", ClearSessionUser, true)
	gaess.SessionRoute(router, "GET", "/api/products", GetProducts, false)
	gaess.SessionRoute(router, "GET", "/api/products/{ID}", GetProductDetail, false)
	gaess.SessionRoute(router, "POST", "/api/price", PostPriceHandler, false)
	gaess.SessionRoute(router, "POST", "/api/order", PostOrderHandler, false)
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
	c := appengine.NewContext(r)
	if err := gaess.LogoutSession(r, w, session); err == nil {
		WriteJSONResponse(c, w, map[string]interface{}{
			"status": "logged out",
		}, 200)
	} else {
		WriteJSONError(c, w, err.Error())
	}
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
	content, statusCode, err := GetURLContentBasicAuth(c, url, creds)
	if err != nil {
		WriteJSONError(c, w, "GET products failed:"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
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
	content, statusCode, err := GetURLContentBasicAuth(c, url, creds)
	if err != nil {
		WriteJSONError(c, w, "GET product "+productID+" failed:"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
}

// PostPriceHandler passes an order entity to the PFL API and relays the response
func PostPriceHandler(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	c := appengine.NewContext(r)

	// parse upload to validate that it's legit JSON
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		order := &map[string]interface{}{}
		if err := json.Unmarshal([]byte(body), order); err != nil {
			WriteJSONError(c, w, err.Error())
			return
		}
	} else {
		WriteJSONError(c, w, err.Error())
		return
	}

	hostName, creds, apikey, err := getAccessParameters(c)
	if err != nil {
		WriteJSONError(c, w, err.Error())
		return
	}

	url := fmt.Sprintf(priceAPI, hostName, apikey)
	content, statusCode, err := PostURLContentBasicAuth(c, url, creds, body)
	if err != nil {
		WriteJSONError(c, w, "POST price failed:"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
}

// PostOrderHandler passes an order entity to the PFL API and relays the response
func PostOrderHandler(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	c := appengine.NewContext(r)

	// parse upload to validate that it's legit JSON
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		order := &map[string]interface{}{}
		if err := json.Unmarshal([]byte(body), order); err != nil {
			WriteJSONError(c, w, err.Error())
			return
		}
	} else {
		WriteJSONError(c, w, err.Error())
		return
	}

	hostName, creds, apikey, err := getAccessParameters(c)
	if err != nil {
		WriteJSONError(c, w, err.Error())
		return
	}

	url := fmt.Sprintf(orderAPI, hostName, apikey)
	log.Infof(c, "POST to %s", orderAPI)
	content, statusCode, err := PostURLContentBasicAuth(c, url, creds, body)
	if err != nil {
		WriteJSONError(c, w, "POST price failed:"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
}
