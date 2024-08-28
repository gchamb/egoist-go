package controllers

import (
	"database/sql"
	"egoist/internal/azure"
	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/structs"
	"egoist/internal/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

// should be able to do the following based on query params:
// - get only images (progress-entry)
// - get only videos (progress-videos)
// 	 -- able to filter by weekly or monthly
// - get both (progress-entry and progress-videos)
// - able to determine how of each they want
func GetAssets(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)
	assetType := r.URL.Query().Get("type")

	// validate inputs
	take, page, frequency, err := structs.ValidateGetAssetsParams(w, r)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	
	skip := (page - 1 ) * take;

	db := database.ConnectDB()
	queries := queries.New(db)

	res := map[string]interface{} {}
	if strings.Contains(assetType, "progress-entry") && strings.Contains(assetType, "progress-video"){
		videos, err := queries.GetProgressVideos(uid, take, skip, frequency)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		entries, err := queries.GetProgressEntries(uid, take, skip)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	

		// map over videos and entries to return readable sas url links
		entries = utils.Map(entries, func (_ int, item structs.ProgressEntry) structs.ProgressEntry {

			blobClient, err := azure.GetBlobClient(azure.PROGRESS_ENTRY_CONTAINER, item.AzureBlobKey, true)
		
			permissions := sas.BlobPermissions{Read: true}
			sas, err := blobClient.GetSASURL(permissions, time.Now().Add(time.Hour), &blob.GetSASURLOptions{})

			item.AzureBlobKey = sas

			return item
		})

		res["videos"] = videos
		res["entries"] = entries
	}else if strings.Contains(assetType, "progress-entry"){
		entries, err := queries.GetProgressEntries(uid, take, skip)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// map over entries to return readable sas url links
		res["entries"] = entries
	}else if strings.Contains(assetType, "progress-video"){
		videos, err := queries.GetProgressVideos(uid, take, skip, frequency)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// map over entries to return readable sas url links
		res["videos"] = videos
	}else{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.ReturnJson(w, res, http.StatusOK)
}