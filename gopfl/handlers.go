package gopfl

import (
	"net/http"

	"github.com/chadbohannan/gae-session-store/gaess"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// TODO read product list
// TODO read product details
// TODO handle custom templates
// TODO place orders

// Init maps HTTP routes to handlers.
// router is  returned for chaining calls.
func Init(router *mux.Router) *mux.Router {
	gaess.SessionRoute(router, "GET", "/session", GetSessionUser, false)
	gaess.SessionRoute(router, "POST", "/session", PostSessionUser, false)
	gaess.SessionRoute(router, "GET", "/clear_session", ClearSessionUser, true)
	gaess.SessionRoute(router, "GET", "/products", GetProducts, true)
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
}

// GetProductDetail fetches a more detailed product description than the GetProducts list
func GetProductDetail(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	// productID := mux.Vars(r)["ID"]
}
