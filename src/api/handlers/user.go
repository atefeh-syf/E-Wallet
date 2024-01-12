package handlers

import (
	"encoding/json"
	"github.com/atefeh-syf/E-Wallet/api/dto"
	"github.com/atefeh-syf/E-Wallet/config"
	"github.com/atefeh-syf/E-Wallet/services"
	"net/http"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{service: service}
}

func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := new(dto.LoginByUsernameRequest)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.service.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Writ(helper.GenerateBaseResponse(token, true, helper.Success))
}
