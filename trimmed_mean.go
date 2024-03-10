package trimmed_mean

import (
	"errors"
	"math"
	"sort"

	"github.com/montanaflynn/stats"
)

type Number interface {
	int | float64
}

// remove a proportion of the smallest and largest values from a slice of floats
func TrimmedMean[N Number](nums []N, percent float64) (float64, error) {
	if percent < 0 || percent > 1 {
		return 0, errors.New("percent must be between 0 and 1")
	}
	if len(nums) == 0 {
		return 0, errors.New("nums cannot be empty")
	}

	floats := make([]float64, len(nums))
	for i, num := range nums {
		floats[i] = float64(num)
	}

	sort.Float64s(floats)
	n := int(math.Round(float64(len(floats)) * percent))

	// remove n elements from the beginning and end of the slice
	floats = floats[n : len(floats)-n]
	mean, _ := stats.Mean(floats)
	return mean, nil
}
