package common

import (
	"fmt"
	"testing"
)

func TestChoose(t *testing.T) {
	fmt.Println(Choose(true, "true", "false"))
	fmt.Println(Choose(true, 1, 2))
}

func BenchmarkChoose(b *testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		result = Choose(true, "true", "false")
	}
	_ = result
}

func TestChooseValue(t *testing.T) {
	fmt.Println(ChooseValue(0, "true", "false"))
	fmt.Println(ChooseValue(1, 1, 2))
}
