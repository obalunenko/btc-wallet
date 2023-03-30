package currency

import (
	"math/big"
	"strconv"

	"github.com/shopspring/decimal"
)

const (
	satoshiperbtc = 100_000_000
)

type Satoshi struct {
	i uint64
}

func (s Satoshi) ToBTC() (decimal.Decimal, error) {
	sat := strconv.FormatUint(s.i, 10)

	_, err := decimal.NewFromString(sat)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return decimal.Decimal{}, nil
}

func (s Satoshi) ParseInt() {
	i := big.Int{}
	i.SetUint64(s.i)
}
