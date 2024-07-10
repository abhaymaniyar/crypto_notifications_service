package enums

type Coin string

const (
	BTC Coin = "BTC"
	ETH Coin = "ETH"
	XRP Coin = "XRP"
)

func (c Coin) IsValid() bool {
	switch c {
	case BTC, ETH, XRP:
		return true
	}

	return false
}
