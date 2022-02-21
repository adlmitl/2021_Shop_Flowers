package http

import "net/http"

// RegisterRoutes - Router handlers.
func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/all", h.FindAll)
	mux.HandleFunc("/create", h.Create)
	mux.HandleFunc("/update", h.Update)
	mux.HandleFunc("/delete", h.Delete)
	mux.HandleFunc("/byId", h.FindById)
	mux.HandleFunc("/byLogin", h.FindByLogin)
}
