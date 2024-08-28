package structs

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (req *AuthRequest) ValidateAuthRequest() error {
	if req.Email == "" {
		return errors.New("invalid email")
	}

	if req.Password == ""{
		return errors.New("password cannot be empty")
	}

	if len(req.Password) < 7 {
		return errors.New("password must be at least 7 characters long")
	}

	return nil
}

func (req *OnboardUserRequest) ValidateOnboardUserReq() error {
	if req.Key == "" || req.CurrentWeight == nil || req.GoalWeight == nil {
		return errors.New("invalid request")
	}

	if req.CurrentWeight != nil && *req.CurrentWeight  < 70 {
		return errors.New("invalid current weight value")
	}

	if req.GoalWeight != nil && *req.GoalWeight  < 70 {
		return errors.New("invalid goal weight value")
	}

	return nil
}

func ValidateGetAssetsParams(w http.ResponseWriter, r *http.Request) (int, int, string, error){
	var take int
	var page int
	var frequency string
	
	assetType := r.URL.Query().Get("type")

	if convertedTake, err := strconv.Atoi(r.URL.Query().Get("take")); err != nil {
		take = 5
	}else {
		take = convertedTake
	}
	if convertedPage, err := strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		page = 1
	}else {
		page = convertedPage
	}

	if strings.Contains(assetType, "progress-video"){
		frequency = r.URL.Query().Get("frequency")
	
		if frequency != "weekly" && frequency != "monthly"{
			return 0, 0, "", errors.New("invalid video type") 
		}
	}

	return take, page, frequency, nil
}