package wallets

type responseWallet struct {
	Address   string  `json:"address"`
	Balance   balance `json:"balance"`
	Available balance `json:"available_balance"`
}

type balance struct {
	USD string `json:"usd"`
	BTC string `json:"btc"`
}

type responseWallets struct {
	Wallets []string `json:"wallets"`
	Count   int      `json:"count"`
}
