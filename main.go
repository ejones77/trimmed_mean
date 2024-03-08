package trimmed_mean

import (
	"errors"
	"math"
	"sort"

	"github.com/montanaflynn/stats"
)

// remove a proportion of the smallest and largest values from a slice of floats
func TrimmedMean(nums []float64, percent float64) (float64, error) {
	// Quick error handling for invalid inputs
	if percent < 0 || percent > 1 {
		return 0, errors.New("percent must be between 0 and 1")
	}
	if len(nums) == 0 {
		return 0, errors.New("numsay cannot be empty")
	}

	sort.Float64s(nums)
	n := int(math.Round(float64(len(nums)) * percent))

	// remove n elements from the beginning and end of the slice
	nums = nums[n : len(nums)-n]
	mean, _ := stats.Mean(nums)
	return mean, nil

}
