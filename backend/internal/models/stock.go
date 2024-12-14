package models

type Stock struct {
	Symbol            string  `json:"symbol"`
	CompanyName       string  `json:"companyName"`
	Price             float64 `json:"price"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Change            float64 `json:"change"`
	DayLow            float64 `json:"dayLow"`
	DayHigh           float64 `json:"dayHigh"`
	PreviousClose     float64 `json:"previousClose"`
	MarketCap         float64 `json:"marketCap"`
}

type StockApiResponse []*struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Change            float64 `json:"change"`
	DayLow            float64 `json:"dayLow"`
	DayHigh           float64 `json:"dayHigh"`
	PreviousClose     float64 `json:"previousClose"`
	MarketCap         float64 `json:"marketCap"`
	CompanyName       string  `json:"companyName"`
}
