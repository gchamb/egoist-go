package controllers

import (
	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/utils"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// func GetReports(w http.ResponseWriter, r *http.Request) {

// 	// get user id
// 	uid := r.Context().Value("uid").(string)

// 	db := database.ConnectDB()
// 	queries := queries.New(db)

// }

func GetReport(w http.ResponseWriter, r *http.Request) {
	// get user id 
	uid := r.Context().Value("uid").(string)
	reportId := chi.URLParam(r, "reportId")

	db := database.ConnectDB()
	queries := queries.New(db)

	report, err := queries.GetReportById(reportId, uid)
	
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.ReturnJson(w, report, http.StatusOK)
}	
