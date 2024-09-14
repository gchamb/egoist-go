package controllers

import (
	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func IsNewReport(w http.ResponseWriter, r *http.Request) {
	// get user id 
	uid := r.Context().Value("uid").(string)

	// get the request body
	var requestBody struct {
		Viewed bool `json:"viewed"`
		ReportId     string `json:"reportId"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestBody)

	db := database.ConnectDB()
	queries := queries.New(db)

	// viewed is true update the report id passed
	if requestBody.Viewed && requestBody.ReportId != "" {
		if err := queries.UpdateProgressReportViewed(requestBody.Viewed, requestBody.ReportId, uid); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	
	// if viewed is false or report id is the default empty string
	// we want to get the latest non viewed report
	

	report, err := queries.GetLatestProgressReport(uid)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return the report details
	res := map[string]interface{} {
		"report": report,
	}

	utils.ReturnJson(w, res, http.StatusOK)
}	