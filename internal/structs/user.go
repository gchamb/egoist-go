package structs

type OnboardUserRequest struct {
	GoalWeight    *float32 `json:"goal_weight"`
	CurrentWeight *float32 `json:"current_weight"`
	Key           string   `json:"key"`
	Tz            string   `json:"tz"`
}

type UpdateUserRequest struct {
	GoalWeight    *float32 `json:"goal_weight"`
	CurrentWeight *float32 `json:"current_weight"`
	ExpoToken     *string  `json:"expo_token"`
	Key           string   `json:"key"`
}