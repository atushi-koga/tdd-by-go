package money

import (
	"testing"
)

func TestSimpleAddition(t *testing.T) {
	fourDollar := money{currency: "USD", amount: 4}
	sum := fourDollar.plus(fourDollar)
	bank := bank{}
	reduced := bank.reduce(sum, "USD")
	if (reduced != money{currency: "USD", amount: 8}) {
		t.Error("fail")
	}
}

func TestMoneyMultiplication(t *testing.T) {
	// @todo: 構造体同士の比較はどのようにされる?
	if (money{currency: "USD", amount: 5}).times(2) != (money{currency: "USD", amount: 10}) {
		t.Error("fail")
	}

	if (money{currency: "CHF", amount: 5}).times(3) != (money{currency: "CHF", amount: 15}) {
		t.Error("fail")
	}
}

func TestEqual(t *testing.T) {
	if (money{currency: "USD", amount: 5}).equal(money{currency: "USD", amount: 5}) != true {
		t.Error("fail")
	}

	if ((money{currency: "USD", amount: 3})).equal((money{currency: "CHF", amount: 3})) != false {
		t.Error("fail")
	}
}

func TestPlusReturnsSum(t *testing.T) {
	fiveDollar := money{currency: "USD", amount: 5}
	threeDollar := money{currency: "USD", amount: 3}
	expected := sum{augend: fiveDollar, addend: threeDollar}
	if fiveDollar.plus(threeDollar) != expected {
		t.Error("fail")
	}
}

func TestReduceMoney(t *testing.T) {
	bank := bank{}
	threeDollar := money{currency: "USD", amount: 3}
	if (bank.reduce(threeDollar, "USD") != threeDollar) {
		t.Error("fail")
	}
}
