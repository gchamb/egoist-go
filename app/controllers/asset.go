package controllers

import (
	"database/sql"
	"egoist/app"
	"egoist/internal/aws"
	"egoist/internal/structs"
	"egoist/internal/utils"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// should be able to do the following based on query params:
//   - get only images (progress-entry)
//   - get only videos (progress-videos)
//     -- able to filter by weekly or monthly
//   - get both (progress-entry and progress-videos)
//   - able to determine how of each they want
func GetAssets(global *app.Globals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value("uid").(string)
		assetType := r.URL.Query().Get("type")

		// validate inputs
		take, page, frequency, err := structs.ValidateGetAssetsParams(w, r)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		skip := (page - 1) * take

		res := map[string]interface{}{}
		if strings.Contains(assetType, "progress-entry") && strings.Contains(assetType, "progress-video") {
			videos, err := global.Queries.GetProgressVideos(uid, take, skip, frequency)
			if err != nil && err != sql.ErrNoRows {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			entries, err := global.Queries.GetProgressEntries(uid, take, skip)
			if err != nil && err != sql.ErrNoRows {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// map over videos and entries to return readable sas url links
			entries = utils.Map(entries, func(_ int, item structs.ProgressEntry) structs.ProgressEntry {
				presignedReq, err := aws.CreatePresignedUrl(item.BlobKey, "READ", time.Now().Add(time.Duration(time.Hour * 24 * 7)))

				if err != nil {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					panic(err.Error())
				}

				item.BlobKey = presignedReq.URL

				return item
			})

			videos = utils.Map(videos, func(_ int, item structs.ProgressVideo) structs.ProgressVideo {

				presignedReq, err := aws.CreatePresignedUrl(item.BlobKey, "READ", time.Now().Add(time.Duration(time.Hour * 24 * 7)))

				if err != nil {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					panic(err.Error())
				}

				item.BlobKey = presignedReq.URL

				return item
			})

			res["videos"] = videos
			res["entries"] = entries
		} else if strings.Contains(assetType, "progress-entry") {
			entries, err := global.Queries.GetProgressEntries(uid, take, skip)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			entries = utils.Map(entries, func(_ int, item structs.ProgressEntry) structs.ProgressEntry {
				presignedReq, err := aws.CreatePresignedUrl(item.BlobKey, "READ", time.Now().Add(time.Duration(time.Hour * 24 * 7)))

				if err != nil {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					panic(err.Error())
				}

				item.BlobKey = presignedReq.URL

				return item
			})

			// map over entries to return readable sas url links
			res["entries"] = entries
		} else if strings.Contains(assetType, "progress-video") {
			videos, err := global.Queries.GetProgressVideos(uid, take, skip, frequency)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			videos = utils.Map(videos, func(_ int, item structs.ProgressVideo) structs.ProgressVideo {
				presignedReq, err := aws.CreatePresignedUrl(item.BlobKey, "READ", time.Now().Add(time.Duration(time.Hour * 24 * 7)))

				if err != nil {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					panic(err.Error())
				}

				item.BlobKey = presignedReq.URL

				return item
			})

			// map over entries to return readable sas url links
			res["videos"] = videos
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		utils.ReturnJson(w, res, http.StatusOK)
	}

}
