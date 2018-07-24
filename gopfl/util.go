package gopfl

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/log"
)

// WriteJSONResponse expresses a struct as a propper HTTP JSON response or dies trying
func WriteJSONResponse(c context.Context, w http.ResponseWriter, data interface{}, statusCode int) {
	if dataJSON, jsonErr := json.Marshal(data); jsonErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(dataJSON)
	} else {
		log.Criticalf(c, "error serializing response: data=%#v, error=%s", data, jsonErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "error": "internal server error serializing response"}`))
	}
}
