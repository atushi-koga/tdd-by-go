package money

type currency string

type amount int

type money struct {
	currency
	amount
}

func (m money) times(amount amount) money {
	return money{
		currency: m.currency,
		amount: m.amount * amount,
	}
}

func (m money) equal(other money) bool {
	return m.currency == other.currency && m.amount == other.amount
}