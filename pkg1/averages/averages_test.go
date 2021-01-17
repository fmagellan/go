package averages

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 1 {
		t.Error("Expected 1, and got:", v)
	}
}
