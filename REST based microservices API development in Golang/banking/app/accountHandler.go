package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	json.NewDecoder(r.Body).Decode(&request)
}
