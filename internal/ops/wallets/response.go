package wallets

type responseWallet struct {
	Address string `json:"address"`
	Balance struct {
		USD string `json:"usd"`
		BTC string `json:"btc"`
	} `json:"balance"`
}

type responseWallets struct {
	Wallets []string `json:"wallets"`
	Count   int      `json:"count"`
}
