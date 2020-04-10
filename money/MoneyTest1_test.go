package money

import (
	"testing"
)

func TestSimpleAddition(t *testing.T) {
	fiveDollar := money{currency: "USD", amount: 5}
	expression := fiveDollar.plus(fiveDollar)
	bank := bank{}
	reduced := bank.reduce(expression, "USD")
	if(reduced != money{currency: "USD", amount: 10}){
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
