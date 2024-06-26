package server

import (
	"github.com/timoruohomaki/open311togo/models"
	//	"github.com/timoruohomaki/open311togo/telemetry"
	"net/http"
	"strconv"
)

// open311/rest/v1/time

func HandleGetTime(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(models.GetServerTime()))

}

func HandleGetServicesXML(w http.ResponseWriter, r *http.Request) {

	u, err := r.URL.Query()

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	jid_str := q["jurisdiction_id"]

	jid, err := strconv.Atoi(jid_str)

	if err != nil {
		if jid > 0 {

		} else {
			// return 404 error with message "Jurisdiction with provided ID was not found."
		}

	} else {
		// return 400 error with message "Jurisdiction ID was not provided"

	}

}

func HandleGetServicesJSON(w http.ResponseWriter, r *http.Request) {

	u, err := r.URL.Query()

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	jid_str := q["jurisdiction_id"]

	jid, err := strconv.Atoi(jid_str)

	if err != nil {
		if jid > 0 {

		} else {
			// return 404 error with message "Jurisdiction with provided ID was not found."
		}

	} else {
		// return 400 error with message "Jurisdiction ID was not provided"

	}

}
