package handlers

import (
	"encoding/json"
	"github.com/atefeh-syf/E-Wallet/api/dto"
	"github.com/gorilla/mux"
	"strconv"

	//"github.com/atefeh-syf/E-Wallet/api/helper"
	"github.com/atefeh-syf/E-Wallet/services"
	"net/http"
)

type WalletHandler struct {
	service *services.WalletService
}

func (h *WalletHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	req := new(dto.TransactionAction)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s := services.NewWalletService()
	result, err := s.Deposit(req)
	//r, err = json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Convert the result to JSON
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	//w.Write(helper.GenerateBaseResponse(res, true, helper.Success))
}

func (h *WalletHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	req := new(dto.TransactionAction)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s := services.NewWalletService()
	result, err := s.Withdraw(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Convert the result to JSON
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *WalletHandler) ForceWithdraw(w http.ResponseWriter, r *http.Request) {
	req := new(dto.TransactionAction)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s := services.NewWalletService()
	result, err := s.Withdraw(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Convert the result to JSON
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *WalletHandler) GetWalletByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["user_id"]
	user_id, _ := strconv.Atoi(UserId)
	s := services.NewWalletService()
	result, err := s.FindWalletByUserId(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Convert the result to JSON
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	//w.Write(helper.GenerateBaseResponse(res, true, helper.Success))
}
