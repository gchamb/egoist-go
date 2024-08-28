package queries

import (
	"database/sql"
	"egoist/internal/structs"

	"github.com/google/uuid"
)

func (q *Queries) InsertProgressEntry(txn *sql.Tx, user structs.ProgressEntry) (error){
	var query string 
	
	if txn == nil {
		query = `INSERT INTO progress_entry (id, azure_blob_key, current_weight, user_id)
        VALUES (:id, :azure_blob_key, :current_weight, :user_id)`
		_, err := q.DB.NamedExec(query, user)
		return err
	}	
	
	query = `INSERT INTO progress_entry (id, azure_blob_key, current_weight, user_id)
        VALUES (?, ?, ?, ?)`

	_, err := txn.Exec(query, user.ID, user.AzureBlobKey, user.CurrentWeight, user.UserID);
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
	query := "SELECT * from progress_entry where user_id = ? LIMIT ? OFFSET ?"

	entries := []structs.ProgressEntry{}
    err := q.DB.Select(&entries, query, uid, take, skip)

	return entries, err
}