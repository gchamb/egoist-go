package structs

type UpdateUserRequest struct {
	GoalWeight    *float32 `json:"goal_weight"`
	CurrentWeight *float32 `json:"current_weight"`
}