package structs

type User struct {
	ID            string   `db:"id"`
	Email         string   `db:"email"`
	Password      *string  `db:"password"`
	GoalWeight    *float32 `db:"goal_weight"`
	CurrentWeight *float32 `db:"current_weight"`
	CreatedAt     string   `db:"created_at"`
}

type ProgressEntry struct {
	ID            string  `db:"id"`
	AzureBlobKey  string  `db:"azure_blob_key"`
	CurrentWeight float32 `db:"current_weight"`
	UserID        string  `db:"user_id"`
	CreatedAt     string  `db:"created_at"`
}

type ProgressReport struct {
	ID             string  `db:"id"`
	GoalWeight     float32 `db:"goal_weight"`
	CurrentWeight  float32 `db:"current_weight"`
	LastWeekWeight float32 `db:"last_week_weight"`
	UserID         string  `db:"user_id"`
	CreatedAt      string  `db:"created_at"`
}

type ProgressVideo struct {
	ID           string `db:"id"`
	AzureBlobKey string `db:"azure_blob_key"`
	Frequency    string `db:"frequency"`
	UserID       string `db:"user_id"`
	CreatedAt    string `db:"created_at"`
}