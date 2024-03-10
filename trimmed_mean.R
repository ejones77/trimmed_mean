library(microbenchmark)

generate_samples <- function() {
  floats <- runif(500, min = 0, max = 100)
  ints <- sample(1:100, 500, replace = TRUE)
  list(floats = floats, ints = ints)
}

trimmed_mean <- function(x, trim = 0.05) {
  mean(x, trim = trim)
}

evaluate <- function() {
  samples <- generate_samples()
  for (i in 1:50) {
    trimmed_mean(samples$floats, 0.05)
    trimmed_mean(samples$ints, 0.05)
  }
}

benchmark_evaluate <- function(n) {
  times <- microbenchmark(evaluate(), times = n)
  avg_time <- mean(times$time)
  cat("BenchmarkEvaluate\n")
  cat(sprintf("Number of iterations: %d\n", n))
  cat(sprintf("Average time per iteration: %.2f ns\n", avg_time))
}

benchmark_evaluate(100)