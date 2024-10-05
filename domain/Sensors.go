package domain

type Sensors struct {
	CurrentMoistureLevel   int     `json:"currentMoistureLevel"`
	CurrentMoisturePercent int     `json:"currentMoisturePercent"`
	CurrentWaterVolume     float64 `json:"currentWaterVolume"`
	MaximumWaterVolume     float64 `json:"maximumWaterVolume"`
}
