package queries

import (
	"egoist/internal/structs"
	"errors"
)

func (q *Queries) GetReportById(id string, uid string) (structs.ProgressReport, error) {
	var query string
	var reports []structs.ProgressReport

	if uid == "" {
		return structs.ProgressReport{}, errors.New("uid wasn't provided for query")
	}

	if id == ""{
		return structs.ProgressReport{}, errors.New("report id wasn't provided for query")
	}

	
	query = "SELECT * FROM progress_report where user_id = ? and id = ?";
	
	if err := q.DB.Select(&reports, query, uid, id); err != nil {
		return structs.ProgressReport{}, err
	}

	if len(reports) == 0 {
		return structs.ProgressReport{}, errors.New("report doesn't exist")
	}


	return reports[0], nil
}
