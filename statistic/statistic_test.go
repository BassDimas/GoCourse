package statistic

import "testing"

type testpair struct {
	values  []float64
	average float64
	sum     float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 6},
	{[]float64{-1, 1}, 0, 0},
}

func TestAverageSumSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		s := Sum(pair.values)
		if v != pair.average {
			t.Error(
				"For Average", pair.values,
				"expected", pair.average,
				"got", v,
			)
		} else if s != pair.sum {
			t.Error(
				"For Sum", pair.values,
				"expected", pair.sum,
				"got", s,
			)

		}
	}

}
