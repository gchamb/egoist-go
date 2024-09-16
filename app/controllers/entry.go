package controllers

import (
	"database/sql"
	"egoist/app"
	"egoist/internal/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func PutEntry(global *app.Globals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value("uid").(string)

		var requestBody structs.PutAssetRequest
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&requestBody)

		if err := requestBody.ValidPutAssetRequest(); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		txn, err := global.Queries.DB.BeginTx(r.Context(), &sql.TxOptions{})
		defer txn.Commit()

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		location, err := time.LoadLocation(requestBody.Timezone)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		todaysDateInTimezone := time.Now().In(location)
		date := fmt.Sprintf("%d-%d-%d", todaysDateInTimezone.Year(), todaysDateInTimezone.Month(), todaysDateInTimezone.Day())

		entry := structs.ProgressEntry{BlobKey: requestBody.Key, CurrentWeight: requestBody.CurrentWeight, UserID: uid, CreatedAt: date}

		if _, err := global.Queries.CreateProgressEntry(txn, entry); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := global.Queries.UpdateUser(txn, r.Context(), structs.UpdateUserRequest{CurrentWeight: &requestBody.CurrentWeight}, uid); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
