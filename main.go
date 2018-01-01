package main

import "github.com/pjgg/cryptocurrency-monitor/service"
import "fmt"

//Example of usage
func main() {

	cryptoCurrencyService := service.Instance()
	cryptoCurrencyList := cryptoCurrencyService.RetrieveCurrencyInfo()
	crytocurrencyNames := make([]string, len(cryptoCurrencyList))

	for i, element := range cryptoCurrencyList {
		fmt.Println(element)
		crytocurrencyNames[i] = element.Name
	}

	stockValues := cryptoCurrencyService.RetrieveCurrencyStockValue(crytocurrencyNames)
	for _, element := range stockValues {
		fmt.Println(element)
	}
}
