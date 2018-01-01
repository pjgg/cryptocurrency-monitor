package dto

type CryptoCurrencyDto struct {
	Name      string
	FullName  string
	CoinName  string
	Logo      string
	Symbol    string
	Algorithm string
}

func NewCryptoCurrencyDto(name, fullName, coinName, logo, symbol, algorithm interface{}) *CryptoCurrencyDto {
	cryptoCurrencyDto := &CryptoCurrencyDto{}
	if nil != name {
		cryptoCurrencyDto.Name = name.(string)
	}

	if nil != fullName {
		cryptoCurrencyDto.FullName = fullName.(string)
	}

	if nil != coinName {
		cryptoCurrencyDto.CoinName = coinName.(string)
	}

	if nil != logo {
		cryptoCurrencyDto.Logo = "https://www.cryptocompare.com/" + logo.(string)
	}

	if nil != symbol {
		cryptoCurrencyDto.Symbol = symbol.(string)
	}

	if nil != algorithm {
		cryptoCurrencyDto.Algorithm = algorithm.(string)
	}

	return cryptoCurrencyDto
}
