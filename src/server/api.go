package server

import (
	"encoding/json"
	"net/http"

	"github.com/timoruohomaki/open311togo/telemetry"
)

func Init(address string) *http.Server {

	httpsrv := initServer()

	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /", httpsrv.handleProduce)
	mux.HandleFunc("GET /", httpsrv.handleConsume)

	// ====== ROUTERS =======
	
	// TODO: POST/GET commit log

	// GET Service list
	
	// mux.HandleFunc("GET /open311/rest/v1/services.xml", func(w http.ResponseWriter, r *http.Request) {
	// 	// jurisdiction_id only required when endpoint serves multiple jurisdictions
	// 	params := r.FormValue("jurisdiction_id")
	// 	ua := r.UserAgent()
	// 	// todo: remove after testing
	// 	fmt.Fprint(w, "Services requested as XML by ",ua, " filtered by ",params)
	// 	telemetry.LogInfo("Services requested as XML by "+ua, "api")
	// })

	// mux.HandleFunc("GET /open311/rest/v1/services.json", func(w http.ResponseWriter, r *http.Request) {
	// 	// jurisdiction_id only required when endpoint serves multiple jurisdictions
	// 	params := r.FormValue("jurisdiction_id")
	// 	ua := r.UserAgent()
	// 	// todo: remove after testing
	// 	fmt.Fprint(w, "Services requested as JSON by ",ua, " filtered by ",params)
	// })

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

	// http.ListenAndServe("localhost:8080", mux)
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

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {

	var req ProduceRequest
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

	res := ProduceResponse{Offset: off}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {

	var req ConsumeRequest
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

	res := ConsumeResponse{Record: record}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}