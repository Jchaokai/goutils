package str

import (
	"testing"
)

func Test_ChildOfStr(t *testing.T) {
	if ChildOfStr("int", "int64") ==false{
		t.Error("用例不通过")
	}
	if ChildOfStr("Int", "int64") ==true {
		t.Error("用力不通过")
	}
}

func Test_SameNthFront(t *testing.T) {
	if SameNthFront("float", "float64", 5) == false{
		t.Error("不通过")
	}
	if SameNthFront("float","fffff",2) == true{
		t.Error("不通过")
	}
}
