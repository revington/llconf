package promise

import "testing"

func TestNamedPromiseDesc(t *testing.T) {
	promise := NamedPromise{ "test", DummyPromise{ "Hello", true } }
	equals(t, promise.Desc([]Constant{}), "(test (dummy [Hello]))")
}

func TestNamedPromiseEval(t *testing.T) {
	promise := NamedPromise{ "test", DummyPromise{ "Hello", true } }
	result,_,_ := promise.Eval([]Constant{})
	equals(t, result, true)
}
