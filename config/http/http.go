package httplib

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//HandleAble serves as a place holder for the http request and response
type HandleAble struct {
	Req *http.Request
	Res http.ResponseWriter
}

//ArgMap defines a json type formate
type ArgMap map[string]interface{}

//BindJSON decodes http request body to a given object
func (c *HandleAble) BindJSON(data interface{}) {
	json.NewDecoder(c.Req.Body).Decode(data)
}

//ResponseJSON returns a http response encoded in application/json format to the response writer
func ResponseJSON(res http.ResponseWriter, status int, object interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(object)
}

//Params maps routes params to mux and returns the value of the key
func (c *HandleAble) Params(key string) string {
	return mux.Vars(c.Req)[key]
}

func MapToJsonBytes(values map[string]interface{}) *bytes.Buffer {
	authParamsJson, error := json.Marshal(values)
	if error != nil {
		panic(error)
	}

	authParamsBytes := bytes.NewBuffer(authParamsJson)
	return authParamsBytes
}
