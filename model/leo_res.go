package model

type LeoConfig struct {
	Rate           float32  `json:"rate"`
	ReleaseRate    float32  `json:"release_rate"`
	AvailableStage float32  `json:"available_stage"`
	DayPerStage    float32  `json:"day_per_stage"`
	Price          float32  `json:"price"`
	AllowTypes     string   `json:"allow_types"`
	Banners        []string `json:"banners"`
}
