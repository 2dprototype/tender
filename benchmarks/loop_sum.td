// Tight Loop Sum Benchmark (10,000,000 iterations)
import "times"

start := times.now()

sum := 0
for i := 0; i < 10000000; i++ {
    sum += i
}

elapsed := times.since(start)

println("10M Loop Sum =", sum)
println("Time taken:", elapsed, "ms")
println("Time taken:", float(elapsed)/times.second, "sec")