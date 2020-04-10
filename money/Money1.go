package money

type currency string

type amount int

type money struct {
	currency
	amount
}

type expression interface {

}

type bank struct {

}

func (b bank) reduce(e expression, c currency) money {
	return money{
		currency: "USD",
		amount:   10,
	}
}

func (m money) plus(m2 money) expression {
	return money{
		currency: m.currency,
		amount:   m.amount + m2.amount,
	}
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