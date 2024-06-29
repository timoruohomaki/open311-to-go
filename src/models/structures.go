package models

import (
	//	"encoding/json"
	//	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"time"
	// "golang.org/x/text/internal/language"
)

// NOTE: Open311 structures follow GeoReport v2 whenever possible https://wiki.open311.org/GeoReport_v2/

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

type ServerTime struct {
	SqlDateTime string `json:"SQLDateTime"`
	TimeZone    string `json:"TimeZone"`
	IsDST       bool   `json:"DST"`
	Info        string `json:"Info"`
}

// open311 data structures
// resource path: https://api.city.gov/dev/v2/services.xml?jurisdiction_id=city.gov

type Open311Service struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" xml:"id,attr"`
	Name string             `json:"name,omitempty" bson:"name,omitempty" xml:"name"`
}

type Open311ServiceRequest struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" xml:"id,attr"`
	Jurisdiction_id int                `json:"jurisdiction_id"`
	Service_code    int                `json:"service_code"`
}

type Open311ServiceRequestResponse struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" xml:"id,attr"`
	Service_request_id int                `json:"serviceRequestId"`
	Service_notice     int                `json:"serviceNotice"`
	Account_id         string             `json:"accountId"`
}

type CesSchema struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" xml:"id,attr"`
	Question_id   int                `json:"question_id"`
	Rating_id     int                `json:"rating_id"`
	CesQuestion   string             `json:"cesQuestion"`
	CesQuestionId int                `json:"cesQuestionId"`
	CesRating     string             `json:"cesRating"`
	CesComment    string             `json:"cesComment"`
	Language      string             `json:"language"` // ISO 639

}

// geospatial structures

type SpatialGeometry struct {
	// json as in GeoJSON (RFC 7946)

	Id           int       `json:"id"`
	GeometryType string    `json:"type"`
	Coordinates  []float64 `json:"coordinates"`
}

type SpatialFeature struct {
	Id                   int             `json:"id"`
	AuthorityName        string          `json:"authorityName"`
	AuthorityResourceURI string          `json:"authorityResourceURI"`
	FeatureType          string          `json:"featureType"`
	FeatureGeometry      SpatialGeometry `json:"featureGeometry"`
}

type CustomFeatureProperties struct {
	Id int `json:"id"`
}

type KeyName struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type ServiceDefinitionAttribute struct {
	Variable             string    `json:"variable"`
	Code                 string    `json:"code"`
	Datatype             string    `json:"datatype"`
	Required             bool      `json:"required"`
	Datatype_description string    `json:"datatypeDescription"`
	Order                int       `json:"order"`
	Description          string    `json:"description"`
	Values               []KeyName `json:"values"`
}

// as in https://docs.ogc.org/is/18-088/18-088.html#featureofinterest

type FeatureOfInterest struct {
	Id                  int                     `json:"id"`
	Name                string                  `json:"name"`
	Description         string                  `json:"description"`
	FeatureEncodingType string                  `json:"featureEncodinType"`
	Feature             SpatialFeature          `json:"feature"`
	Properties          CustomFeatureProperties `json:"properties"`
}

type ServiceDefinition struct {
	Service_code string                     `json:"serviceCode"`
	PTV_code_URI string                     `json:"ptvCodeURI"`
	Attribute    ServiceDefinitionAttribute `json:"attributeList"`
}

/*
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
} */
