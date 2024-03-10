package trimmed_mean

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func GenerateSamples() ([]float64, []int) {
	floats := make([]float64, 500)
	ints := make([]int, 500)
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < 500; i++ {
		floats[i] = r.Float64() * 100
		ints[i] = r.IntN(100)
	}

	return floats, ints
}

func Evaluate() {
	floatSample, intSample := GenerateSamples()
	start := time.Now()
	for i := 0; i < 100; i++ {
		TrimmedMean(floatSample, 0.05)
		TrimmedMean(intSample, 0.05)
	}
	end := time.Since(start)
	fmt.Println("Time taken:", end)
}
