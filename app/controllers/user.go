package controllers

import (
	"database/sql"
	"egoist/internal/database"
	"egoist/internal/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// get user id 
	uid := r.Context().Value("uid").(string)

	// get the request body
	var requestBody structs.UpdateUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestBody)

	fmt.Println(*requestBody.CurrentWeight)
	// update the user
	db := database.ConnectDB()

	tx, err := db.BeginTx(r.Context(), &sql.TxOptions{})

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if requestBody.CurrentWeight == nil && requestBody.GoalWeight == nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if requestBody.CurrentWeight != nil {
		_, err := tx.Exec("UPDATE user SET current_weight = ? where id = ?", *requestBody.CurrentWeight, uid)
		if err != nil {
			tx.Rollback()
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if requestBody.GoalWeight != nil {
		_, err := tx.Exec("UPDATE user SET goal_weight = ? where id = ?", *requestBody.GoalWeight, uid)
		if err != nil {
			tx.Rollback()
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}