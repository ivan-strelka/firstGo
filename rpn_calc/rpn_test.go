package rpn_calc

import "testing"

func Test_rpn(t *testing.T) {
	t.Run("3 2 -", func(t *testing.T) {
		//rpn("3 2 -") // string
		//rpn("3", "2", "-") // []string
		result := rpn(3, 2, "-") // []interface{}
		if result != 1 {
			t.Error("Got ", result)
		}
	})
	t.Run("5 3 2 * +", func(t *testing.T) {
		// 5 [5]
		// 3 [5,3]
		// 2 [5,3,2]
		// * [5,6]
		// + [11]
		result := rpn(5, 3, 2, "*", "+") // []interface{}
		if result != 11 {
			t.Error("Got ", result)
		}

	})
	t.Run("12 3 - 3 /", func(t *testing.T) {

	})

}
