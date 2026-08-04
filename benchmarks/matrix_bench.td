// Matrix Operations Benchmark
import "times"

start := times.now()

m1 := matrix([
    [1.5, 2.5, 3.5],
    [4.5, 5.5, 6.5],
    [7.5, 8.5, 9.5]
])

m2 := matrix([
    [9.0, 8.0, 7.0],
    [6.0, 5.0, 4.0],
    [3.0, 2.0, 1.0]
])

sum_det := 0.0
for i := 0; i < 50000; i++ {
    m3 := m1 * m2
    sum_det += m3.det
}

elapsed := times.since(start)

println("50k Matrix Multiplications Determinant Sum =", sum_det)
println("Time taken:", elapsed, "ms")
println("Time taken:", float(elapsed)/times.second, "sec")