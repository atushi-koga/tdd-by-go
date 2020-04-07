package money

type money struct {
	amount int
}

func (m money) times(value int) money {
	return money{amount: m.amount * value}
}

func (m money) equal(other money) bool {
	return m.amount == other.amount
}