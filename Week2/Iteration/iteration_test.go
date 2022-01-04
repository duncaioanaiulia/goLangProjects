package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	fmt.Println("Benchmarks are run sequentially \n The benchmark function must run the target code b.N times \n ")
	fmt.Println("Comanda: go test -bench=.")
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
