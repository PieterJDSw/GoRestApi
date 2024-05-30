package user

import (
	"net/http"

	"github.com/PieterJDSw/GoRestApi/types"
	"github.com/PieterJDSw/GoRestApi/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {

	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUserPayload

	if err := utils.ParseJson(r, payload); err != nil {

		utils.WriteError(w, http.StatusBadRequest, err)
	}

}
