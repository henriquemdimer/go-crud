package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/henriquemdimer/go-crud/auth"
	"github.com/henriquemdimer/go-crud/models"
)

func returnError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		returnError(w, err)
		return
	}

	if len(user.Name) > 100 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if len(user.Password) > 100 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	checkUser, _ := models.GetUserByName(user.Name)
	if checkUser.Name == user.Name {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	id, err := models.InsertUser(user.Name, user.Password)
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

func Login(w http.ResponseWriter, r *http.Request) {
	var receivedUser models.User

	err := json.NewDecoder(r.Body).Decode(&receivedUser)
	if err != nil {
		returnError(w, err)
		return
	}

	user, err := models.GetUserByName(receivedUser.Name)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if user.Password != receivedUser.Password {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateToken(user.Id)
	if err != nil {
		returnError(w, err)
		return
	}

	res := map[string]any{
		"token": token,
		"id":    user.Id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	user, err := models.GetUserById(userID)
	if err != nil {
		returnError(w, err)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
