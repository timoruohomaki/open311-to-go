package server

import (
	// "github.com/timoruohomaki/open311togo/models"
	//	"github.com/timoruohomaki/open311togo/telemetry"
	"net/http"
	// "strconv"
)

type Handlers interface {
	
}

type handler struct {

}



func HandleGetServiceListXML(w http.ResponseWriter, r *http.Request) {

	// u := r.URL.Query()

	// jid_str := u["jurisdiction_id"]

	// jid, err := strconv.Atoi([]jid_str)

	// if err != nil {
	// 	if jid > 0 {

	// 	} else {
	// 		// return 404 error with message "Jurisdiction with provided ID was not found."
	// 	}

	// } else {
	// 	// return 400 error with message "Jurisdiction ID was not provided"

	// }

}

func HandleGetServiceListJSON(w http.ResponseWriter, r *http.Request) {

		// u := r.URL.Query()

	   	// jid_str := u["jurisdiction_id"]

	   	// jid, err := strconv.Atoi([]jid_str)

	   	// if err != nil {
	   	// 	if jid > 0 {

	   	// 	} else {
	   	// 		// return 404 error with message "Jurisdiction with provided ID was not found."
	   	// 	}

	   	// } else {
	   	// 	// return 400 error with message "Jurisdiction ID was not provided"

	   	// }

}

func HandleGetServiceDefinitionXML(w http.ResponseWriter, r *http.Request) {

}

func HandleGetServiceDefinitionJSON(w http.ResponseWriter, r *http.Request) {

}

func HandlePostServiceDefinitionXML(w http.ResponseWriter, r *http.Request) {

}

func HandlePostServiceDefinitionJSON(w http.ResponseWriter, r *http.Request) {

}



