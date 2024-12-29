package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { // Read the body of the request
		if err := json.Unmarshal([]byte(body), x); err != nil { // Unmarshal JSON into `x`
			return // If there is an error, do nothing (return)
		}
	}
}
