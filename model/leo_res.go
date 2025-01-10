package model

type LeoConfig struct {
	Rate           float32  `json:"rate"`
	ReleaseRate    float32  `json:"release_rate"`
	AvailableStage int      `json:"available_stage"`
	DayPerStage    int      `json:"day_per_stage"`
	Price          float32  `json:"price"`
	AllowTypes     string   `json:"allow_types"`
	Banners        []string `json:"banners"`
	MinAmount      int64    `json:"min_amount"`
	MaxAmount      int64    `json:"max_amount"`
}
