package money

// ・確認事項

type currency string

type amount int

type money struct {
	currency
	amount
}

type expression interface {
	augendValue() money
	addendValue() money
}

type sum struct {
	augend money
	addend money
}

func (s sum) augendValue() money {
	return s.augend
}

func (s sum) addendValue() money {
	return s.addend
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
