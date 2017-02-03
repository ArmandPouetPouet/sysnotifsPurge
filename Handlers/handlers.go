package Handlers

import (
	elastic "SysnotifsPurge/Elasticsearch"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

//Index just say hello
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to this AWS Elasticsearch sysnotifs logs cleaner service !")
}

//CheckRatio function : get user information
func CheckRatio(w http.ResponseWriter, r *http.Request) {
	ratio, err := elastic.CleanES()
	if err != nil {
		SetResponse(w, http.StatusInternalServerError, "Failed with error :  "+fmt.Sprint(err))
	} else {
		SetResponse(w, http.StatusOK, "Work done, ratio is now "+strconv.Itoa(ratio)+"%")
	}
}

//SetResponse function with json output
func SetResponse(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if i != nil {
		if err := json.NewEncoder(w).Encode(i); err != nil {
			errors.Wrap(err, "Encoding response failed")
		}
	}
}
