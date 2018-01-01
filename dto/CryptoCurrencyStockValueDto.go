package dto

import (
	"time"
)

type CryptoCurrencyStockValueDto struct {
	CurrencyName string
	BTC          float64
	USD          float64
	EUR          float64
	CreatedAt    time.Time
}

func NewCryptoCurrencyStockValueDto(currencyName string, BTC, USD, EUR interface{}) *CryptoCurrencyStockValueDto {
	cryptoCurrencyStockValueDto := &CryptoCurrencyStockValueDto{}

	cryptoCurrencyStockValueDto.CurrencyName = currencyName
	if nil != BTC {
		cryptoCurrencyStockValueDto.BTC = BTC.(float64)
	}

	if nil != USD {
		cryptoCurrencyStockValueDto.USD = USD.(float64)
	}

	if nil != EUR {
		cryptoCurrencyStockValueDto.EUR = EUR.(float64)
	}

	cryptoCurrencyStockValueDto.CreatedAt = time.Now().UTC()

	return cryptoCurrencyStockValueDto
}
