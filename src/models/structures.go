package models

import (
	"sync"
)

// commit log structures

type ProduceRequest struct {
	Record	Record	`json:"record"`
}

type ProduceResponse struct {
	Offset	uint64	`json:"offset"`
}

type ConsumeRequest struct {
	Offset	uint64	`json:"offset"`
}

type ConsumeResponse struct {
	Record	Record	`json"Record"`
}

// open311 structures

type SpatialGeometry struct {
	// json as in GeoJSON (RFC 7946)

	id           int       `json:"id"`
	geometryType string    `json:"type"`
	coordinates  []float64 `json:"coordinates"`
}

type SpatialFeature struct {
	id                   int
	authorityName        string
	authorityResourceURI string
	featureType          string
	featureGeometry      SpatialGeometry
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

type ServiceRequest struct {
	id int
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

	requests map[int]ServiceRequest
}
