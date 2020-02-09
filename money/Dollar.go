package money

type dollar struct {
	amount int
}

func (d dollar) times(value int) dollar {
	return dollar{amount: d.amount * value}
}

func (d dollar) equal(other dollar) bool {
	return d.amount == other.amount
}