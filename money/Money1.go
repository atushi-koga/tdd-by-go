package money

// ・確認事項

type currency string

type amount int

type expression interface {
	reduce(to currency) money
}

type money struct {
	currency
	amount
}

func (m money) reduce(to currency) money {
	return m
}

type sum struct {
	augend money
	addend money
}

func (s sum) reduce(to currency) money {
	return money{
		currency: to,
		amount:   s.augend.amount + s.addend.amount,
	}
}

type bank struct {
}

func (b bank) reduce(e expression, c currency) money {
	return e.reduce(c)
}

func (m money) plus(m2 money) expression {
	return sum{augend: m, addend: m2}
}

func (m money) times(amount amount) money {
	return money{
		currency: m.currency,
		amount:   m.amount * amount,
	}
}

func (m money) equal(other money) bool {
	return m.currency == other.currency && m.amount == other.amount
}
