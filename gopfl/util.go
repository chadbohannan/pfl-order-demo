package gopfl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

var getCache = map[string][]byte{}

// GetURLContentBasicAuth executs a blocking GET of a url with a Basic Auth header
func GetURLContentBasicAuth(c context.Context, url, auth string) ([]byte, int, error) {
	if content, ok := getCache[url]; ok {
		// TODO check content metadata to allow cache expiration
		return content, 0, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	request.Header.Add("Authorization", "Basic "+auth)

	client := urlfetch.Client(c)
	response, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	// cache result DISABLED until cache expiration logic is used
	// getCache[url] = result
	return result, response.StatusCode, nil
}

// PostURLContentBasicAuth executs a blocking GET of a url with a Basic Auth header
func PostURLContentBasicAuth(c context.Context, url, auth string, body []byte) ([]byte, int, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, err
	}
	request.Header.Add("Authorization", "Basic "+auth)
	request.Header.Set("Content-Type", "application/json")

	client := urlfetch.Client(c)
	response, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	return result, response.StatusCode, nil
}

func getAccessParameters(c context.Context) (string, string, string, error) {
	remoteHost, err := GetSetting(c, remoteHostKey)
	if err != nil {
		return "", "", "", errors.New("service configuration A missing, " + err.Error())
	}
	creds, err := GetSetting(c, credentialsKey)
	if err != nil {
		return "", "", "", errors.New("service configuration B missing, " + err.Error())
	}
	apikey, err := GetSetting(c, apiKey)
	if err != nil {
		return "", "", "", errors.New("service configuration C missing, " + err.Error())
	}
	return remoteHost, creds, apikey, nil
}

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

func Stringify(obj interface{}) string {
	body, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
