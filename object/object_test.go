package object

import (
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hahs keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have some hash keys")
	}
}

func TestHashableKey(t *testing.T) {
	var intObj Object = &Integer{Value: 1}

	a, ok := intObj.(Hashable)
	if !ok {
		t.Errorf("int must can be Hashable, %T", a)
	}

}
