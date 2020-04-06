package money

type franc struct {
	amount int
}

func (f franc) times(value int) franc {
	return franc{amount: f.amount * value}
}

func (f franc) equal(other franc) bool{
	return f.amount == other.amount
}