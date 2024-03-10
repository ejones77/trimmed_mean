package trimmed_mean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimmedMean(t *testing.T) {
	arr := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	percent := 0.2
	expected := 5.5
	result, err := TrimmedMean(arr, percent)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func BenchmarkEvaluate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Evaluate()
	}
}
