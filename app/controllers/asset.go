package controllers

import (
	"net/http"
	"strings"
)

// should be able to do the following based on query params:
// - get only images (progress-entry)
// - get only videos (progress-videos)
// - get both (progress-entry and progress-videos)
// - able to determine how of each they want
func GetAssets(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)
	assetType := r.URL.Query().Get("type")
	take := r.URL.Query().Get("type")

	if strings.Contains(assetType, "progress-entry") && strings.Contains(assetType, "progress-video"){
		// both
	}else if strings.Contains(assetType, "progress-entry"){
		// only entries
	}else if strings.Contains(assetType, "progress-video"){
		// only videos
	}else{
		// none
	}



}