package userControllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/henriquemdimer/go-crud/auth"
	userModel "github.com/henriquemdimer/go-crud/models/user"
)

func returnError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user userModel.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		returnError(w, err)
		return
	}

	id, err := userModel.Insert(user.Name, user.Password)
	if err != nil {
		returnError(w, err)
		return
	}

	token, err := auth.CreateToken(id)
	if err != nil {
		returnError(w, err)
		return
	}

	res := map[string]any{
		"token": token,
		"id":    id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
