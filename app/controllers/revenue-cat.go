package controllers

import (
	"egoist/app"
	"egoist/internal/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func RevenueCatWebhook(global *app.Globals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var eventData structs.RevenueCatEvent
		eventDecoder := json.NewDecoder(r.Body)
		eventDecoder.Decode(&eventData)

		fmt.Println("Event Type", eventData.Event.Type)

		switch eventData.Event.Type {
		case "INITIAL_PURCHASE":
			uid := eventData.Event.AppUserID

			if uid == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			fmt.Println("User ID", eventData.Event.AppUserID)

			sub := structs.RevenueCatSubscriber{
				ID:             eventData.Event.ID,
				ProductID:      eventData.Event.ProductID,
				TransactionID:  eventData.Event.TransactionID,
				ExpirationAtMs: eventData.Event.ExpirationAtMs,
				PurchasedAtMs:  eventData.Event.PurchasedAtMs,
				UserID:         uid,
			}

			err := global.Queries.CreateSubscriber(sub)

			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
		case "RENEWAL":
			uid := eventData.Event.AppUserID

			if uid == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("User ID", eventData.Event.AppUserID)

			err := global.Queries.UpdateSubscriber(eventData.Event.ExpirationAtMs, uid)

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

}
