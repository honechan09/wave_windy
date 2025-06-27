package entity

type WeatherResponse struct {
	List []WeatherItem `json:"list"`
}

// 日時、風情報の取得
type WeatherItem struct {
	DtTxt string   `json:"dt_txt"`
	Wind  WindInfo `json:"wind"`
}

// 風情報の詳細を取得
type WindInfo struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}
