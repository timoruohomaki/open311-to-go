package server

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/storage"
	"github.com/timoruohomaki/open311togo/storage/dbrepos"
	// "github.com/timoruohomaki/open311togo/telemetry"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	GetServices(w http.ResponseWriter, r *http.Request)
	CreateService(w http.ResponseWriter, r *http.Request)
	GetService(w http.ResponseWriter, r *http.Request)
	DeleteService(w http.ResponseWriter, r *http.Request)
	UpdateService(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	MG storage.DbMethod
}

func NewHandler( mg storage.DbInterface) Handlers {
	return &handler{
		// calling repo constructor
		MG: dbrepos.NewMongoDbRepo(mg, context.Background()),
	}
}

// Fetch all the available Open311 services and send it as response to user
func (h *handler) GetServices(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
	 limit = 10
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
	 page = 1
	}
	products, err := h.MG.GetServices(limit, page)
	if err != nil {
	 server.StatusInternalServerError(w, err.Error())
	 return
	}
	server.StatusOKAll(w, limit, page, products)
}

// Insert new Open311 Service into database

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request) {
	p := &models.Open311CreateUpdateService{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
	 server.StatusBadRequest(w, "error parsing json")
	 return
	}
	product, err := h.MG.CreateService(p)
	if err != nil {
	 server.StatusInternalServerError(w, err.Error())
	 return
	}
	server.StatusCreated(w, product)
   }

   // Get specific Service
   func (h *handler) GetService(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
	 server.StatusBadRequest(w, "invalid service id")
	 return
	}
	product, err := h.MG.GetService(id)
	if err != nil {
	 server.StatusNotFound(w, err.Error())
	 return
	}
	server.StatusOK(w, product)
   }

// delete a service using id
func (h *handler) DeleteService(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
	 server.StatusBadRequest(w, "invalid id")
	 return
	}
	if err := h.MG.DeleteService(id); err != nil {
	 server.StatusInternalServerError(w, err.Error())
	 return
	}
	server.StatusAcceptedMsg(w, "service deleted")
   }

// update a service using id
func (h *handler) UpdateService(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
	 server.StatusBadRequest(w, "invalid id")
	 return
	}
	update := &models.Open311CreateUpdateService{}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
	 server.StatusBadRequest(w, "error in parsing json")
	 return
	}
	if err := h.MG.UpdateService(id, update); err != nil {
	 server.StatusBadRequest(w, err.Error())
	 return
	}
	service, err := h.MG.GetService(id)
	if err != nil {
	 server.StatusNotFound(w, err.Error())
	 return
	}
	server.StatusAcceptedData(w, service)
   }