package service

import (
	"encoding/json"

	"github.com/pjgg/cryptocurrency-monitor/dto"
	"github.com/pjgg/cryptocurrency-monitor/utils"
)

type CryptoCurrencyService struct {
}

type CryptoCurrency interface {
	RetrieveCurrencyInfo() []dto.CryptoCurrencyDto
}

func (cryptoCurrencyService CryptoCurrencyService) RetrieveCurrencyInfo() (crptoCurrency []dto.CryptoCurrencyDto) {
	req, _ := utils.NewHTTPRequest("GET", "https://www.cryptocompare.com/api/data/coinlist/", nil, nil)
	if resp, err := utils.MakeHTTPQuery(req); err == nil {
		raw := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&raw)

		data := raw["Data"].(map[string]string)
		keys := make([]string, 0, len(data))

		for k := range data {
			keys = append(keys, k)
			c := &dto.CryptoCurrencyDto{
				Name:      data["Name"],
				Algorithm: data["Algorithm"],
				CoinName:  data["CoinName"],
				FullName:  data["FullName"],
				Symbol:    data["Symbol"],
				Logo:      "https://www.cryptocompare.com/" + data["ImageUrl"],
			}

			crptoCurrency = append(crptoCurrency, *c)

		}
	}
	return
}
