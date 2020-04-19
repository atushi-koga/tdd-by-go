package money

// @todo:構造体埋め込みの宣言とそのユースケースを確認

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
	pair := currencyPair{from: m.currency, to: to}

	return money{
		currency: to,
		amount:   m.amount / b.ratio(pair),
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
	rates
}

func (b bank) addRate(pair currencyPair, ratio int) bank {
	rate := rate{currencyPair: pair, ratio: ratio}

	return bank{b.put(rate)}
}

func (b bank) reduce(e expression, c currency) money {
	return e.reduce(b, c)
}

func (b bank) ratio(pair currencyPair) int {
	return b.find(pair).ratio
}

type rates struct {
	rate []rate
}

func (rs rates) put(r rate) rates {
	return rates{rate: append(rs.rate, r)}
}

func (rs rates) find(pair currencyPair) rate {
	if pair.same() {
		return rate{pair, 1}
	}

	for i := 0; i < len(rs.rate); i++ {
		if rs.rate[i].currencyPair == pair {
			return rs.rate[i]
		}
	}
	// @todo: 実行時エラーとしてExceptionのようなものを吐かせたい(panic?)
	// errorを返してclient側でチェックはさせたくないから
	return rate{}
}

type rate struct {
	currencyPair
	ratio int
}

type currencyPair struct {
	from currency
	to   currency
}

func (c currencyPair) same() bool {
	return c.from == c.to
}
