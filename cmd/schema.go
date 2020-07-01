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

// Total ...
type Total struct {
	Date             int `json:"date"`
	Pcr              int `json:"pcr"`
	Positive         int `json:"positive"`
	Symptom          int `json:"symptom"`
	Symptomless      int `json:"symptomless"`
	SymtomConfirming int `json:"symtomConfirming"`
	Hospitalize      int `json:"hospitalize"`
	Mild             int `json:"mild"`
	Severe           int `json:"severe"`
	Confirming       int `json:"confirming"`
	Waiting          int `json:"waiting"`
	Discharge        int `json:"discharge"`
	Death            int `json:"death"`
}
