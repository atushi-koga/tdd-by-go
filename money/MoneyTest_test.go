package money

import (
	"testing"
)

func TestDollarMultiplication(t *testing.T) {
	if (dollar{amount: 5}).times(2) != (dollar{amount: 10}) {
		t.Error("fail")
	}
	if (dollar{amount: 5}).times(3) != (dollar{amount: 15}) {
		t.Error("fail")
	}
}

func TestFrancMultiplication(t *testing.T) {
	if (franc{amount: 5}).times(2) != (franc{amount: 10}) {
		t.Error("fail")
	}
	if (franc{amount: 5}).times(3) != (franc{amount: 15}) {
		t.Error("fail")
	}
}

func TestEqual(t *testing.T) {
	if (dollar{amount: 5}).equal(dollar{amount: 5}) != true {
		t.Error("fail")
	}
	if (dollar{amount: 5}).equal(dollar{amount: 6}) != false {
		t.Error("fail")
	}
}
