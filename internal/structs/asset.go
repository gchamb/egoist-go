package structs

type PutAssetRequest struct {
	Key           string  `json:"key"`
	CurrentWeight float32 `json:"current_weight"`
	Timezone      string  `json:"timezone"`
}
