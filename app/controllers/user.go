package controllers

import (
	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func OnboardUser(w http.ResponseWriter, r *http.Request){
	// get user id 
	uid := r.Context().Value("uid").(string)

	// get the request body
	var requestBody structs.UpdateUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestBody)

	// validate inputs
	if err := requestBody.ValidateUpdateUserReq(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.ConnectDB()
	queries := queries.New(db)

	// Update the user
	if err := queries.UpdateUser(r.Context(), requestBody, uid); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
}