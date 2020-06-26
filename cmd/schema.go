package cmd

// Prefecture ...
type Prefecture struct {
	ID     int     `json:"id"`
	NameJa string  `json:"name_ja"`
	NameEn string  `json:"name_en"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Cases  int     `json:"cases"`
	Deaths int     `json:"deaths"`
}
