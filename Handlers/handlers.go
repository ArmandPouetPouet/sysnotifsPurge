package Handlers

import (
	elastic "SysnotifsPurge/Elasticsearch"
	helpers "SysnotifsPurge/Helpers"
	"fmt"
	"net/http"
)

//Index just say hello
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to this AWS Elasticsearch sysnotifs logs cleaner service !")
}

//CheckRatio function : get user information
func CheckRatio(w http.ResponseWriter, r *http.Request) {

	elastic.CleanES()
	helpers.SetResponse(w, http.StatusOK, "work done")

}
