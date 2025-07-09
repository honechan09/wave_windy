package entity

type HourlyData struct {
	Time               []string  `json:"time"`
	SwellWaveHeight    []float64 `json:"swell_wave_height"`
	SwellWaveDirection []float64 `json:"swell_wave_direction"`
}

type SurgeResponse struct {
	Hourly HourlyData `json:"hourly"`
}

type Agg struct {
	Count     int
	HeightSum float64
	DirSum    float64
}
