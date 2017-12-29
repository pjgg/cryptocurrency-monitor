package main

import "github.com/pjgg/cryptocurrency-monitor/service"
import "fmt"

//Example of usage
func main() {

	cryptoCurrencyService := new(service.CryptoCurrencyService)
	cryptoCurrencyList := cryptoCurrencyService.RetrieveCurrencyInfo()

	for k := range cryptoCurrencyList {
		fmt.Println(k)
	}
}
