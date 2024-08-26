package controllers

import (
	"database/sql"
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
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.ConnectDB()
	queries := queries.New(db)

	// I want the user to be fully onboarded or not at all
	txn, err := queries.DB.BeginTx(r.Context(), &sql.TxOptions{})
	defer txn.Commit()

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updateUser := structs.UpdateUserRequest{GoalWeight: requestBody.GoalWeight, CurrentWeight: requestBody.CurrentWeight}
	if err := queries.UpdateUser(txn, r.Context(), updateUser , uid); err != nil {
		txn.Rollback()
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	entry := structs.ProgressEntry{AzureBlobKey: requestBody.Key, CurrentWeight: *requestBody.CurrentWeight, UserID: uid}
	if _, err := queries.CreateProgressEntry(txn, entry); err != nil {
		txn.Rollback()
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	
	w.WriteHeader(http.StatusOK)
}