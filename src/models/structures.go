package models

import (
	//	"encoding/json"
	//	"net/http"
	"encoding/json"
	"sync"
	"time"
)

// helper functions

func GetServerTime() (result string) {

	currentTime := time.Now()
	formattedTime := currentTime.Format(time.RFC3339)
	tzinfo := currentTime.Location().String()

	t := &ServerTime{
		SqlDateTime: formattedTime,
		TimeZone:    tzinfo,
		IsDST:       true,
		Info:        "",
	}

	s, _ := json.Marshal(t)

	result = string(s)
	return result

}

// log and telemetry

type TeleLog struct {
	mu      sync.Mutex
	records []Record
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// open311 data structures

type SpatialGeometry struct {
	// json as in GeoJSON (RFC 7946)

	id           int       `json:"id"`
	geometryType string    `json:"type"`
	coordinates  []float64 `json:"coordinates"`
}

type SpatialFeature struct {
	id                   int             `json:"id"`
	authorityName        string          `json:"authorityName"`
	authorityResourceURI string          `json:"authorityResourceURI"`
	featureType          string          `json:"featureType"`
	featureGeometry      SpatialGeometry `json:"featureGeometry"`
}

type CustomFeatureProperties struct {
	id int
}

type KeyName struct {
	key  string
	name string
}

type ServiceDefinitionAttribute struct {
	variable             string
	code                 string
	datatype             string
	required             bool
	datatype_description string
	order                int
	description          string
	values               []KeyName
}

// as in https://docs.ogc.org/is/18-088/18-088.html#featureofinterest

type FeatureOfInterest struct {
	id                  int                     `json:"id"`
	name                string                  `json:"name"`
	description         string                  `json:"description"`
	featureEncodingType string                  `json:"featureEncodinType"`
	feature             SpatialFeature          `json:"feature"`
	properties          CustomFeatureProperties `json:"properties"`
}

type ServiceDefinition struct {
	service_code string
	ptv_code_URI string
	attribute    ServiceDefinitionAttribute
}

type ServerTime struct {
	SqlDateTime string `json: "SQLDateTime"`
	TimeZone    string `json: "TimeZone"`
	IsDST       bool   `json: "DST"`
	Info        string `json: "Info"`
}

type Open311ServiceRequest struct {
	jurisdiction_id int `json:"jurisdiction_id"`
	service_code    int `json:"service_code"`
}

type Open311ServiceRequestResponse struct {
	service_request_id int
	service_notice     int
	account_id         string
}

type Defaults struct {
	mongoServiceCollection string
	mongoRequestCollection string
}

type ServiceStore struct {
	sync.Mutex

	services map[int]ServiceDefinition
	nextId   int
}

type RequestStore struct {
	sync.Mutex

	requests map[int]Open311ServiceRequest
}
