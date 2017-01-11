package helpers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

//SetResponse function with json output
func SetResponse(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if i != nil {
		if err := json.NewEncoder(w).Encode(i); err != nil {
			panic(err)
		}
	}
}

//GetID function
func GetID(r *http.Request, idName string) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[idName])

	if err != nil {
		panic(err)
	}

	return id
}

//GetString function
func GetString(r *http.Request, stringName string) string {
	vars := mux.Vars(r)
	stringValue := vars[stringName]

	return stringValue
}

//GetIsActive function
func GetIsActive(r *http.Request) bool {
	vars := mux.Vars(r)
	isactive, err := strconv.ParseBool(vars["isactive"])

	if err != nil {
		panic(err)
	}

	return isactive
}

//GetBody function
func GetBody(w http.ResponseWriter, r *http.Request, i interface{}) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &i); err != nil {
		SetResponse(w, 422, nil)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
}
