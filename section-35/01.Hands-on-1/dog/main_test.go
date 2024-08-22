package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(10)
	if y != 70 {
		t.Error("Expected", 70, "Got", y)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(20)
	if y != 140 {
		t.Error("Expected", 140, "Got", y)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))
	// Output:
	// 140
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}
}
