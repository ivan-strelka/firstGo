package rpn_calc

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, want interface{}, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Error("Got ", got, "but want ", want)
	}

}

func Test_rpn(t *testing.T) {
	t.Run("3 2 -", func(t *testing.T) {
		//rpn("3 2 -") // string
		//rpn("3", "2", "-") // []string
		result := rpn(3, 2, "-") // []interface{}
		assert(t, 1, result)

	})
	t.Run("5 3 2 * +", func(t *testing.T) {
		// 5 [5]
		// 3 [5,3]
		// 2 [5,3,2]
		// * [5,6]
		// + [11]
		result := rpn(5, 3, 2, "*", "+") // []interface{}
		assert(t, 11, result)
	})
	t.Run("12 3 - 3 /", func(t *testing.T) {
		result := rpn(12, 3, "-", 3, "/") // []interface{}
		assert(t, 3, result)
	})

	t.Run("4 DUP *", func(t *testing.T) {
		result := rpn(4, "DUP", "*") // []interface{}
		assert(t, 16, result)
	})
	t.Run("5 DUP *", func(t *testing.T) {
		result := rpn(5, "DUP", "*") // []interface{}
		assert(t, 25, result)
	})

}
