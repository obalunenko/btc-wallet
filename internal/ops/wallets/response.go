package wallets

type response struct {
	Address string `json:"address"`
	Balance struct {
		USD string `json:"usd"`
		BTC string `json:"btc"`
	} `json:"balance"`
}
