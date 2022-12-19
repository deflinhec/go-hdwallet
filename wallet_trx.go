package hdwallet

import (
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[TRX] = newETH
}

const (
	TronBytePrefix = byte(0x41)
)

type trx struct {
	name   string
	symbol string
	key    *Key

	// trx token
	contract string
}

func newTRX(key *Key) Wallet {
	return &trx{
		name:   "Tron",
		symbol: "TRX",
		key:    key,
	}
}

func (c *trx) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *trx) GetName() string {
	return c.name
}

func (c *trx) GetSymbol() string {
	return c.symbol
}

func (c *trx) GetKey() *Key {
	return c.key
}

func (c *trx) GetAddress() (string, error) {
	address := crypto.PubkeyToAddress(*c.key.PublicECDSA)
	addressTron := make([]byte, 0)
	addressTron = append(addressTron, TronBytePrefix)
	addressTron = append(addressTron, address.Bytes()...)
	checksum := sha256.Sum256(addressTron)
	checksum = sha256.Sum256(checksum[:])
	return base58.Encode(append(addressTron, checksum[0:4]...)), nil
}
