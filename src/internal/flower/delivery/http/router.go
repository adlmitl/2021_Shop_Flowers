package http

import "net/http"

// RegisterRoutes - Router handlers.
func (h *FlowerHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/flower/all", h.FindAll)
	mux.HandleFunc("/flower/create", h.Create)
	mux.HandleFunc("/flower/update", h.Update)
	mux.HandleFunc("/flower/delete", h.Delete)
	mux.HandleFunc("/flower/byId", h.FindById)
	mux.HandleFunc("/flower/byLogin", h.FindByName)
}
