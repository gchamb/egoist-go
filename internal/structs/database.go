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
	ID            string  `db:"id" json:"id"`
	AzureBlobKey  string  `db:"azure_blob_key" json:"azureBlobKey"`
	CurrentWeight float32 `db:"current_weight" json:"currentWeight"`
	UserID        string  `db:"user_id" json:"userId"`
	CreatedAt     string  `db:"created_at" json:"createdAt"`
}

type ProgressReport struct {
	ID             string  `db:"id" json:"id"`
	GoalWeight     float32 `db:"goal_weight" json:"goalWeight"`
	CurrentWeight  float32 `db:"current_weight" json:"currentWeight"`
	LastWeekWeight float32 `db:"last_week_weight" json:"lastWeekWeight"`
	UserID         string  `db:"user_id" json:"userId"`
	CreatedAt      string  `db:"created_at" json:"createdAt"`
}

type ProgressVideo struct {
	ID           string `db:"id" json:"id"`
	AzureBlobKey string `db:"azure_blob_key" json:"azureBlobKey"`
	Frequency    string `db:"frequency" json:"frequency"`
	UserID       string `db:"user_id" json:"userId"`
	CreatedAt    string `db:"created_at" json:"createdAt"`
}

type Assets interface {
	ProgressEntry | ProgressVideo
}