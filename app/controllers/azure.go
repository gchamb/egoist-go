package controllers

import (
	"egoist/internal/azure"
	"egoist/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/google/uuid"
)

func GenerateUploadSaSUrl(w http.ResponseWriter, r *http.Request) {
	mimetype := r.URL.Query().Get("mimetype")
	
	if _, ok := utils.MIMETYPES[mimetype]; !ok {
		fmt.Println("invalid mimetype")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	blobKey := uuid.New().String() + utils.MIMETYPES[mimetype]
	blobClient, err := azure.GetBlobClient(azure.PROGRESS_ENTRY_CONTAINER, blobKey, true)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	permissions := sas.BlobPermissions{Write: true}
	sas, err := blobClient.GetSASURL(permissions, time.Now().Add(time.Minute), &blob.GetSASURLOptions{})
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := map[string]string {
		"url": sas,
		"key": blobKey,
	}

	utils.ReturnJson(w, res, http.StatusOK)
}