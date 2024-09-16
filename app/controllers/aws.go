package controllers

import (
	egoistAws "egoist/internal/aws"
	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GenerateUploadPresignedUrl(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)

	mimetype := r.URL.Query().Get("mimetype")
	tz := r.URL.Query().Get("tz")

	// make sure they haven't uploaded already using their tz
	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println("invalid timezone")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todaysDateInTimezone := time.Now().In(loc)
	date := fmt.Sprintf("%d-%d-%d",todaysDateInTimezone.Year(), todaysDateInTimezone.Month(), todaysDateInTimezone.Day())

	db := database.ConnectDB()
	queries := queries.New(db)

	entries, err := queries.GetProgressEntryByDate(date, uid)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(entries) > 0 {
		fmt.Println("Entry already exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	if _, ok := utils.MIMETYPES[mimetype]; !ok {
		fmt.Println("invalid mimetype")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	blobKey := uuid.New().String() + utils.MIMETYPES[mimetype]
	blobKey = fmt.Sprintf("%s/%s",egoistAws.PROGRESS_ENTRY_CONTAINER, blobKey)
	
	presignedReq, err := egoistAws.CreatePresignedUrl(blobKey, "WRITE", time.Now().Add(time.Minute))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := map[string]interface{} {
		"url": presignedReq.URL,
		"key": blobKey,
		"headers": map[string]string {
			"host": presignedReq.SignedHeader.Get("Host"),
			"expires": presignedReq.SignedHeader.Get("Expires"),
		},
	}

	utils.ReturnJson(w, res, http.StatusOK)
}