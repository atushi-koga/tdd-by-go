package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	// @todo: 構造体同士の比較はどのようにされる?
	if (money{amount: 5}).times(2) != (money{amount: 10}) {
		t.Error("fail")
	}

	if (money{amount: 5}).times(3) != (money{amount: 15}) {
		t.Error("fail")
	}

	fiveDollar := dollar{m: money{amount: 5}}
	tenDollar := dollar{m: money{amount: 10}}
	if fiveDollar.times(2) != tenDollar {
		t.Error("fail")
	}

	fiveFranc := franc{m: money{amount: 5}}
	tenFranc := franc{m: money{amount: 10}}
	if fiveFranc.times(2) != tenFranc {
		t.Error("fail")
	}
}

func TestEqual(t *testing.T) {
	if (money{amount: 5}).equal(money{amount: 5}) != true {
		t.Error("fail")
	}

	if (money{amount: 5}).equal(money{amount: 6}) != false {
		t.Error("fail")
	}

	fiveDollar := dollar{m: money{amount: 5}}
	tenDollar := dollar{m: money{amount: 10}}
	if fiveDollar.equal(fiveDollar) != true {
		t.Error("fail")
	}
	if fiveDollar.equal(tenDollar) != false {
		t.Error("fail")
	}

	fiveFranc := franc{m: money{amount: 5}}
	tenFranc := franc{m: money{amount: 10}}
	if fiveFranc.equal(fiveFranc) != true {
		t.Error("fail")
	}
	if fiveFranc.equal(tenFranc) != false {
		t.Error("fail")
	}
}
