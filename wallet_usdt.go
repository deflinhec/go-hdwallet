package hdwallet

func init() {
	coins[USDT] = newUSDT
}

type usdt struct {
	*trx
}

func newUSDT(key *Key) Wallet {
	token := newTRX(key).(*trx)
	token.name = "Tether"
	token.symbol = "USDT"
	token.contract = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"

	return &usdt{trx: token}
}
