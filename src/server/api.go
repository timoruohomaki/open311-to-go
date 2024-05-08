package server

import (
	"encoding/json"
	"net/http"
	"github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/telemetry"
)

func Init(address string) *http.Server {

	httpsrv := initServer()

	// ====== ROUTERS =======

	mux := http.NewServeMux()

	// POST and GET commit log

	mux.HandleFunc("POST /", httpsrv.handleProduce)
	mux.HandleFunc("GET /", httpsrv.handleConsume)

	// GET Service list

	mux.HandleFunc("GET /open311/rest/v1/services.xml", httpsrv.handleGetServicesXml)
	mux.HandleFunc("GET /open311/rest/v1/services.json", httpsrv.handleGetServicesJson)

	// GET Service definition
	// POST Service request
	// GET service_request_id from a token
	// GET Service requests
	// GET Service request

	telemetry.LogInfo("Starting http server...", "api")

	return &http.Server{
		Addr:	address,
		Handler:	mux,
	}
}

type httpServer struct {
	Log *Log
}

func initServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

// the following calls are for the commit log


func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {

	var req models.ProduceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	off, err := s.Log.Append(req.Record)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := models.ProduceResponse{Offset: off}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {

	var req models.ConsumeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := s.Log.Read(req.Offset)

	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := models.ConsumeResponse{Record: record}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// open311 request handlers

func (s *httpServer) handleGetServicesXml(w http.ResponseWriter, r *http.Request) {

	var req models.ConsumeServicesXmlRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := s.Log.Read(req.Offset)

	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := models.ConsumeServicesXmlResponse{models.RecordXml: record}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *httpServer) handleGetServicesJson(w http.ResponseWriter, r *http.Request) {

	var req models.ConsumeServicesJsonRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := s.Log.Read(req.Offset)

	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := models.ConsumeServicesJsonResponse{models.RecordJson: record}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}