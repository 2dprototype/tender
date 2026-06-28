// tender@v0.0.5

// demo 1
for i in range(0, 10) {
    println <| `counting... ${i}`
}

// demo 2
type player struct {
    name   string
    email  string
    rank   int
    crypto bigfloat
}

const p1 = player{
    name: "Jack",
    email: "jack@email.com",
    rank: 0,
    crypto: bigfloat(104879274.348)
}

str := `
This is a template string,
    ${p1.name}
    ${p1.email}
    ${p1.rank}
    ${p1.crypto}
`

str |> println
