package controllers

import (
	egoistAws "egoist/internal/aws"
	"egoist/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GenerateUploadPresignedUrl(w http.ResponseWriter, r *http.Request) {
	mimetype := r.URL.Query().Get("mimetype")
	
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