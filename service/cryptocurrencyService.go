package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"

	"github.com/pjgg/cryptocurrency-monitor/dto"
	"github.com/pjgg/cryptocurrency-monitor/utils"
)

var cryptoCurrencyServiceSingleton sync.Once
var cryptoCurrencyService CryptoCurrencyService

type CryptoCurrencyService struct {
	parallelRequestCryptoCurrenciesAmount int
}

type CryptoCurrency interface {
	RetrieveCurrencyInfo() []dto.CryptoCurrencyDto
	RetrieveCurrencyStockValue(crytoCurrencyNames []string) []dto.CryptoCurrencyStockValueDto
}

func Instance() *CryptoCurrencyService {
	cryptoCurrencyServiceSingleton.Do(func() {
		cryptoCurrencyService.parallelRequestCryptoCurrenciesAmount = 20
	})

	return &cryptoCurrencyService
}

func (c *CryptoCurrencyService) RetrieveCurrencyInfo() (cryptoCurrency []dto.CryptoCurrencyDto) {
	req, _ := utils.NewHTTPRequest("GET", "https://www.cryptocompare.com/api/data/coinlist/", nil, nil)
	if resp, err := utils.MakeHTTPQuery(req); err == nil {

		//utils.PrintResponseBody(resp)

		raw := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&raw)

		data := raw["Data"].(map[string]interface{})
		keys := make([]string, 0, len(data))
		fmt.Println("Currency amount: " + strconv.Itoa(len(data)))
		for k := range data {
			keys = append(keys, k)
			currency := data[k].(map[string]interface{})
			fmt.Println("processing " + currency["Name"].(string) + " ...")
			c := dto.NewCryptoCurrencyDto(currency["Name"], currency["FullName"], currency["CoinName"], currency["ImageUrl"], currency["Symbol"], currency["Algorithm"])
			cryptoCurrency = append(cryptoCurrency, *c)
		}
	}
	return
}

func (c *CryptoCurrencyService) RetrieveCurrencyStockValue(crytoCurrencyNames []string) (stockValue []dto.CryptoCurrencyStockValueDto) {
	total := len(crytoCurrencyNames)
	fmt.Println("Retriving data. This could take a while...")
	for from := 0; from < total; from = from + c.parallelRequestCryptoCurrenciesAmount {

		slice := crytoCurrencyNames[from:c.getUntil(from, total)]
		watcher := observer.Observer{

			// Register a handler function for every next available item.
			NextHandler: func(resp interface{}) {
				raw := make(map[string]interface{})
				json.NewDecoder(resp.(*http.Response).Body).Decode(&raw)

				for k := range raw {
					currency := raw[k].(map[string]interface{})
					currencyValue := dto.NewCryptoCurrencyStockValueDto(k, currency["BTC"], currency["USD"], currency["EUR"])
					stockValue = append(stockValue, *currencyValue)
				}
			},

			// Register a handler for any emitted error.
			ErrHandler: func(err error) {
				fmt.Printf("Encountered error: %v\n", err)
			},

			// Register a handler when a stream is completed.
			DoneHandler: func() {
				fmt.Print("=")
			},
		}

		it, _ := iterable.New([]interface{}{prepareAndRunHttpquery(strings.Join(slice, ","))})
		source := observable.From(it)
		sub := source.Subscribe(watcher)
		<-sub
	}
	fmt.Print(" Done!")
	return
}

func prepareAndRunHttpquery(fsymsValues string) (result interface{}) {
	endpoint := "https://min-api.cryptocompare.com/data/pricemulti?fsyms=" + fsymsValues + "&tsyms=BTC,USD,EUR"
	req, _ := utils.NewHTTPRequest("GET", endpoint, nil, nil)
	if resp, err := utils.MakeHTTPQuery(req); err != nil {
		result = err
	} else {
		result = resp
	}
	return result
}

func (c *CryptoCurrencyService) getUntil(from, total int) (until int) {
	until = from + c.parallelRequestCryptoCurrenciesAmount
	if total < until {
		until = total
	}

	return
}
