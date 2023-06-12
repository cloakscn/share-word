package main

import (
	"testing"
)

func TestAAA(t *testing.T) {
	var (
		around         = 10
		base           = 100000
		x      float64 = float64(base)
		li             = 0.05
	)
	// for i := 0; i < 360*around; i++ {
	// 	x = x*((360+li)/360) + float64(base)
	// }

	for i := 0; i < around; i++ {
		x = x*(1+li) + float64(base)
	}
	t.Fatal(x)
	t.Fatal(x / (float64(around) * float64(base)))
}
