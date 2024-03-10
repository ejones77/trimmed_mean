# go_trimmed_mean

Submitted assignment for Northwestern MSDS 431

A Go package to calculate trimmed mean

## Overview

The trimmed mean is a useful approach to find the "typical" or average value in a set of numbers because it removes the influence of outliers. 

This package calculates a percentage-based trimmed mean. Given a slice of integers or floats and a percentage to trim the slice by, TrimmedMean returns an average value.

TrimmedMean also uses generics similar to what's laid out in the [tutorial documentation](https://go.dev/doc/tutorial/generics)

`evaluate.go` contains some helper functions for benchmark testing, including use of the new `math/rand/v2` to generate random numbers


## Setup 


## To run benchmarks from this repository
- clone the repository
- in the repository directory, run `go test -bench=.`