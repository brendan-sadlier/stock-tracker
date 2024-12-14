package services

import (
	"encoding/json"
	"fmt"

	"stock-sleuth-backend/internal/models"
	"stock-sleuth-backend/pkg/api"
)

type StockService struct {
	apiKey     string
	httpClient api.HTTPClient
}

func NewStockService(apiKey string, client api.HTTPClient) *StockService {
	return &StockService{
		apiKey:     apiKey,
		httpClient: client,
	}
}

func (s *StockService) FetchStockInfo(symbol string) (*models.Stock, error) {
	url := fmt.Sprintf("https://financialmodelingprep.com/api/v3/quote/%s?apikey=%s", symbol, s.apiKey)

	resp, err := s.httpClient.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock info: %v", err)
	}

	defer resp.Body.Close()

	var apiResponse models.StockApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(apiResponse) == 0 {
		return nil, fmt.Errorf("no stock info found for symbol %s", symbol)
	}

	stockData := apiResponse[0]
	return &models.Stock{
		Symbol:            stockData.Symbol,
		CompanyName:       stockData.CompanyName,
		Price:             stockData.Price,
		ChangesPercentage: stockData.ChangesPercentage,
		Change:            stockData.Change,
		DayLow:            stockData.DayLow,
		DayHigh:           stockData.DayHigh,
		PreviousClose:     stockData.PreviousClose,
		MarketCap:         stockData.MarketCap,
	}, nil
}
