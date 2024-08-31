package queries

import (
	"database/sql"
	"egoist/internal/structs"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (q *Queries) InsertProgressEntry(txn *sql.Tx, entry structs.ProgressEntry) (error){
	var query string


	if entry.CreatedAt == ""{
		date := time.Now()
		entry.CreatedAt = fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
	}
	
	if txn == nil {
		query = `INSERT INTO progress_entry (id, azure_blob_key, current_weight, user_id, created_at)
        VALUES (:id, :azure_blob_key, :current_weight, :user_id, :created_at)`

		_, err := q.DB.NamedExec(query, entry)
		return err
	}	
	
	query = `INSERT INTO progress_entry (id, azure_blob_key, current_weight, user_id, created_at)
        VALUES (?, ?, ?, ?, ?)`
		

	_, err := txn.Exec(query, entry.ID, entry.AzureBlobKey, entry.CurrentWeight, entry.UserID, entry.CreatedAt);
	return err
}

func (q *Queries) CreateProgressEntry(txn *sql.Tx, entry structs.ProgressEntry) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	entry.ID = id.String()

	err = q.InsertProgressEntry(txn, entry)
	if err != nil {
		return "", err
	}

	return id.String(), err
}

func (q *Queries) GetProgressEntries(uid string, take int, skip int) ([] structs.ProgressEntry, error){
	query := "SELECT * from progress_entry where user_id = ? ORDER BY created_at ASC LIMIT ? OFFSET ?"

	entries := []structs.ProgressEntry{}
    err := q.DB.Select(&entries, query, uid, take, skip)

	return entries, err
}