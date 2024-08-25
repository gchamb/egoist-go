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
	var requestBody structs.OnboardUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestBody)

	// validate inputs
	if err := requestBody.ValidateOnboardUserReq(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.ConnectDB()
	queries := queries.New(db)

	// Update the user
	if err := queries.UpdateUser(r.Context(), structs.UpdateUserRequest{GoalWeight: requestBody.GoalWeight, CurrentWeight: requestBody.CurrentWeight}, uid); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// create the initial progress entry for today
	
	w.WriteHeader(http.StatusOK)
}