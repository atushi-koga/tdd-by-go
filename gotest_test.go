package gotest

import "testing"

func TestDivision1(t *testing.T) {
	if i, e := division(6, 2); i != 3 || e != nil {
		t.Error("divisionテストが失敗しました")
	}else{
		t.Log("divisionテストが成功しました")
	}
}

func TestDivision2(t *testing.T) {
	if _, e := division(6, 0); e == nil {
		t.Error("divisionテストが失敗しました")
	}else{
		t.Log("divisionテストが成功しました.", e)
	}
}