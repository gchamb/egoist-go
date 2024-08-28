package queries

import "egoist/internal/structs"

func (q *Queries) GetProgressVideos(uid string, take int, skip int, frequency string) ([]structs.ProgressVideo, error){
	query := "SELECT * FROM progress_video WHERE user_id = ? AND frequency = ? LIMIT ? OFFSET ?"

	videos := []structs.ProgressVideo{}
	err := q.DB.Select(&videos, query, uid, frequency, take, skip)

	return videos, err
}