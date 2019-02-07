package sum

import "testing"

var sum_tests_int8 = []struct {
	n1 	 int8
	n2	 int8
	expected int8
}
{
	{5, 1, 6},
	{12, -15, 3},
	{127, -3, 124},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val ! = v.expected
			t.Errorf("Sum(%d, %d) returned %d, expected %d", 
v.1, v.n2, val, v.expected)
		}
	}
}
