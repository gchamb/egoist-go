package controllers

import (
	"database/sql"
	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

	timeInTz, err := time.LoadLocation(requestBody.Tz)
	if err != nil {
		txn.Rollback()
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	
	todaysDateInTimezone := time.Now().In(timeInTz)
	createdAt := fmt.Sprintf("%d-%d-%d",todaysDateInTimezone.Year(), todaysDateInTimezone.Month(), todaysDateInTimezone.Day())
	entry := structs.ProgressEntry{BlobKey: requestBody.Key, CurrentWeight: *requestBody.CurrentWeight, UserID: uid , CreatedAt: createdAt}
	if _, err := queries.CreateProgressEntry(txn, entry); err != nil {
		txn.Rollback()
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	uid := r.Context().Value("uid").(string)

	// get the request body
	var requestBody struct {
		GoalWeight    *float32 `json:"goal_weight"`
		CurrentWeight *float32 `json:"current_weight"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestBody)
	fmt.Println(requestBody)

	if requestBody.GoalWeight == nil && requestBody.CurrentWeight == nil{
		fmt.Println("goal weight and current weight can't be empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if requestBody.GoalWeight != nil && (*requestBody.GoalWeight < 70 || *requestBody.GoalWeight > 500) {
		fmt.Println("invalid goal weight")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if requestBody.CurrentWeight != nil && (*requestBody.CurrentWeight < 70 || *requestBody.CurrentWeight > 500) {
		fmt.Println("invalid current weight")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.ConnectDB()
	queries := queries.New(db)

	updateUser := structs.UpdateUserRequest{GoalWeight: requestBody.GoalWeight, CurrentWeight: requestBody.CurrentWeight}
	if err := queries.UpdateUser(nil, r.Context(), updateUser , uid); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}