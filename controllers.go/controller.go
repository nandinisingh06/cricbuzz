package controllers

import (
	"cricbuzz/models"
	"encoding/json"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)
	u := NewUser.NewUser()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
