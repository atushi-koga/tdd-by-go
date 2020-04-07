package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	if (money{amount: 5}).times(2) != (money{amount: 10}) {
		t.Error("fail")
	}
	if (money{amount: 5}).times(3) != (money{amount: 15}) {
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
}
