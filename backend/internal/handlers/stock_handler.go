package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock-sleuth-backend/internal/services"
)

type StockHandler struct {
	stockService *services.StockService
}

func NewStockHandler(stockService *services.StockService) *StockHandler {
	return &StockHandler{stockService: stockService}
}

func (handler *StockHandler) GetStockInfo(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Stock symbol is required"})
		return
	}

	stock, err := handler.stockService.FetchStockInfo(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stock)
}
