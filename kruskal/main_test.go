package kruskal

import (
	"testing"
)


func TestCorrectness(t *testing.T) {
	expectedValues := []int{12, 6359060}
	tests := 2
	for i := 0; i < tests; i++ {
		expected := expectedValues[i]
		actual := Kruskal(i)

		if actual != expected{
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	}
}

func BenchmarkSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 2; j++ {
			Kruskal(j)
		}
	}
}

