package queries

import "egoist/internal/structs"

func (q *Queries) GetProgressVideos(uid string, take int, skip int, frequency string) ([]structs.ProgressVideo, error){
	query := "SELECT * FROM progress_video WHERE user_id = ? AND frequency = ?  ORDER BY created_at DESC LIMIT ? OFFSET ?"

	videos := []structs.ProgressVideo{}
	err := q.DB.Select(&videos, query, uid, frequency, take, skip)

	return videos, err
}

func (q *Queries) GetAllProgressVideos(uid string) ([]structs.ProgressVideo, error){
	query := "SELECT * FROM progress_video WHERE user_id = ?"

	videos := []structs.ProgressVideo{}
	err := q.DB.Select(&videos, query, uid)

	return videos, err
}