package money

type money2 interface {
	equal(m money2) bool
	times(v int) money2
	amount() int
}

type dollar2 struct {
	a int
}

type franc2 struct {
	amount int
}

func (d dollar2) equal(m money2) bool {
	return d.a == m.amount()
}

func (d dollar2) times(v int) money2 {
	return dollar2{a: d.a * v}
}

func (d dollar2) amount() int {
	return d.a
}