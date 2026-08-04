// Fibonacci Benchmark
import "times"

fn fib(n) {
    if n <= 1 {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}

start := times.now()
res := fib(35)
elapsed := times.since(start)

println("fib(30) =", res)
println("Time taken:", elapsed, "ms")
println("Time taken:", float(elapsed)/times.second, "sec")
