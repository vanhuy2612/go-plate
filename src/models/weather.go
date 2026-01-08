package models

type Weather struct {
	Current any `json:"current"`
	Daily []any `json:"daily"`
	Hourly []any `json:"hourly"`
	Lat float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Minutely []any   `json:"minutely"`
	Timezone string  `json:"timezone"`
	TimezoneOffset int64 `json:"timezone_offset"`
}
