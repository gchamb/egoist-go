package queries

import (
	"egoist/internal/structs"
	"errors"
)

func (q *Queries) GetLatestProgressReport(uid string) (structs.ProgressReport, error) {
	var query string
	var report []structs.ProgressReport

	if uid == ""{
		return report[0], errors.New("uid wasn't provided for query")
	}

	
	query = "SELECT * FROM progress_report where user_id = ? and viewed = ? LIMIT 1";
	
	if err := q.DB.Select(&report, query, uid, 0); err != nil {
		return report[0], err
	}

	return report[0], nil
}

func (q *Queries) UpdateProgressReportViewed(viewed bool, reportId string, uid string) error {
	var query string
	

	if reportId == ""{
		return errors.New("report id wasn't provided for query.")
	}

	
	query = "UPDATE progress_report SET viewed = ? WHERE id = ? and user_id = ?";
	
	_, err := q.DB.Exec(query, viewed, reportId, uid)
		
	return err
}