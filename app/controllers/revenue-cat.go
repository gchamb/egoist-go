package controllers

import (
	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func RevenueCatWebhook(w http.ResponseWriter, r *http.Request) {
	
	var eventData structs.RevenueCatEvent
	eventDecoder := json.NewDecoder(r.Body)
	eventDecoder.Decode(&eventData)

	fmt.Println("Event Type", eventData.Event.Type)

	switch eventData.Event.Type {
	case "INITIAL_PURCHASE":
		uid := eventData.Event.AppUserID

		if uid == ""{
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("User ID", eventData.Event.AppUserID)
		// add to database
		db := database.ConnectDB()
		queries := queries.New(db)
		
		sub := structs.RevenueCatSubscriber {
			ID: eventData.Event.ID,
			TransactionID: eventData.Event.TransactionID,
			ExpirationAtMs: eventData.Event.ExpirationAtMs,
			PurchasedAtMs: eventData.Event.PurchasedAtMs,
			UserID: uid,
		}

		err := queries.CreateSubscriber(sub)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "RENEWAL":
		uid := eventData.Event.AppUserID

		if uid == ""{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("User ID", eventData.Event.AppUserID)

		db := database.ConnectDB()
		queries := queries.New(db)

		err := queries.UpdateSubscriber(eventData.Event.ExpirationAtMs, uid)


		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusOK)
	}

}