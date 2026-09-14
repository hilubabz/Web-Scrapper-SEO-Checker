package audit

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type CreateAuditRequest struct {
	URL string `json:"url"`
}

func (h *Handler) CreateAudit(w http.ResponseWriter, r *http.Request){
	if r.Method!=http.MethodPost{
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	req := CreateAuditRequest{}
	
	err := json.NewDecoder(r.Body).Decode(&req)
	
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.URL==""{
		writeError(w, http.StatusBadRequest, "URL is required")
		return
	}
	result, err := h.service.Run(req.URL)

	err = json.NewEncoder(w).Encode(result)
	if err!=nil{
		writeError(w, http.StatusBadRequest, err.Error())
		return 
	}

	writeJSON(w, http.StatusOK, result)
}