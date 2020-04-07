package money

type franc struct {
	m money
}

func (f franc) times(value int) franc {
	return franc{m: f.m.times(value)}
}

func (f franc) equal(o franc) bool{
	return f.m.equal(o.m)
}