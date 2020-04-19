package money

type currency string

type expression interface {
	reduce(b bank, to currency) money
}

// @todo: amountフィールドは、intを拡張した独自型にしたい
// 独自型にした時、m.reduce()内でamountのフィールドに値をセットできなくて断念している
type money struct {
	currency
	amount int
}

func (m money) reduce(b bank, to currency) money {
	return money{
		currency: to,
		amount:   m.amount / b.rate(m.currency, to),
	}
}

func (m money) plus(m2 money) expression {
	return sum{augend: m, addend: m2}
}

func (m money) times(amount int) money {
	return money{
		currency: m.currency,
		amount:   m.amount * amount,
	}
}

func (m money) equal(other money) bool {
	return m.currency == other.currency && m.amount == other.amount
}

type sum struct {
	augend money
	addend money
}

func (s sum) reduce(b bank, to currency) money {
	return money{
		currency: to,
		amount:   s.augend.amount + s.addend.amount,
	}
}

type bank struct {
}

func (b bank) addRate(from currency, to currency, rate int) {
}

func (b bank) reduce(e expression, c currency) money {
	return e.reduce(b, c)
}

func (b bank) rate(from currency, to currency) int {
	if from == "CHF" && to == "USD" {
		return 2
	} else {
		return 1
	}
}
