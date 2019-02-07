package sum

import "testing"

var sum_tests_float64 = []struct {
	n1	float64
	n2	float64
	expectd	float64
}
{
	{5, 2, 7},
	{5, -9, -4},
	{116, 0, 116},
}

func TestSumfloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%f, %f) returned %f, expected %f", 
v.n1, v.n2, val, v.expected)
		}
	}
}
