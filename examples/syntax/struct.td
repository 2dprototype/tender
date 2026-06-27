// tender@v0.0.5

type User struct {
	name string
	age  int
}

// 1. Basic instantiations
u1 := User{name: "Alice", age: 30}
u2 := User{"Bob", 25}
u3 := User{}

println("u1 name:", u1.name, "age:", u1.age)
println("u2 name:", u2.name, "age:", u2.age)
println("u3 name:", u3.name, "age:", u3.age) // Should print zero-values: "" and 0

// 2. Field modification
u1.age = 31
println("u1 new age:", u1.age)

// 3. Type checks
println("u1 is_struct:", is_struct(u1))
println("u1 typeof:", typeof(u1))

// 4. Nested structs
type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

l := Line{}
println("l.p1.x zero-value:", l.p1.x) // Should be 0
l.p1 = Point{x: 1, y: 2}
println("l.p1.x after set:", l.p1.x) // Should be 1
l.p1.y = 5
println("l.p1.y after set field:", l.p1.y) // Should be 5

// 5. Anonymous struct literal
a := struct {
	title string
	val   float
}{title: "Pi", val: 3.14}
println("anonymous struct title:", a.title, "val:", a.val)
println("anonymous struct typeof:", typeof(a))

// 6. Multiple fields same type
type Rect struct {
	width, height int
}
r := Rect{width: 10, height: 20}
println("rect width:", r.width, "height:", r.height)

// 7. Embedded anonymous field
type Address struct {
	city string
}
type Employee struct {
	Address
	salary int
}
emp := Employee{Address: Address{city: "London"}, salary: 5000}
println("emp city:", emp.Address.city, "salary:", emp.salary)