# go_trimmed_mean

Submitted assignment for Northwestern MSDS 431

A Go package to calculate trimmed mean

## Overview

The trimmed mean is a useful approach to find the "typical" or average value in a set of numbers because it removes the influence of outliers. 

This package calculates a percentage-based trimmed mean. Given a slice of integers or floats and a percentage to trim the slice by, TrimmedMean returns an average value.

TrimmedMean also uses generics similar to what's laid out in the [tutorial documentation](https://go.dev/doc/tutorial/generics)

`evaluate.go` contains some helper functions for benchmark testing, including use of the new `math/rand/v2` to generate random numbers


## Installation

run `go get github.com/ejones77/trimmed_mean`

## Example Usage

```
package main

import (
	"fmt"

	"github.com/ejones77/trimmed_mean"
)

func main() {
	nums := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Result, err := trimmed_mean.TrimmedMean(nums, 0.2)
	if err != nil {
		fmt.Println(Result)
	}
}

```

## To run benchmarks from this repository

- clone the repository
- in the repository directory, run `go test -bench=.`

## Dependency for the R comparison

- microbenchmark

## Results of benchmarking compared to `trimmed_mean.R`

The trimmed mean is run on a sample of 500 floats and 500 integers

This operation performs 50 times, and the benchmark evaluation executes those 50 calculations in 100 iterations (a total of 5,000 means calculated).

Both programs return average execution time in nanoseconds

- Go: 995,120 nanoseconds/operation
- R: 2,693,313 nanoseconds/operation

These results show that Go achieves a 63% reduction in execution time!