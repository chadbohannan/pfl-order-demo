package gopfl

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

// WriteJSONResponse expresses a struct as a propper HTTP JSON response or dies trying
func WriteJSONResponse(c context.Context, w http.ResponseWriter, data interface{}, statusCode int) {
	if dataJSON, jsonErr := json.Marshal(data); jsonErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(dataJSON)
	} else {
		log.Criticalf(c, "error serializing response: data=%#v, error=%s", data, jsonErr.Error())
		WriteJSONError(c, w, "internal server error serializing response")
	}
}

// WriteJSONError expresses an error string as an HTTP JSON  response
func WriteJSONError(c context.Context, w http.ResponseWriter, errMsg string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, errMsg)))
}

// GetURLContentBasicAuth executs a blocking GET of a url with a Basic Auth header
func GetURLContentBasicAuth(c context.Context, url, auth string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Basic "+auth)

	client := urlfetch.Client(c)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
