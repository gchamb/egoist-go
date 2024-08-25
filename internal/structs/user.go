package structs

type OnboardUserRequest struct {
	GoalWeight    *float32 `json:"goal_weight"`
	CurrentWeight *float32 `json:"current_weight"`
	Key           string   `json:"key"`
}

type UpdateUserRequest struct {
	GoalWeight    *float32 `json:"goal_weight"`
	CurrentWeight *float32 `json:"current_weight"`
	Key           string   `json:"key"`
}