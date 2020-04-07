package money

type dollar struct {
	m money
}

func (d dollar) times(value int) dollar {
	return dollar{m: d.m.times(value)}
}

func (d dollar) equal(o dollar) bool {
	return d.m.equal(o.m)
}