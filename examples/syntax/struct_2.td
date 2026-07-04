import "json"

type user struct {
	name string `json:"username"`
	age  int    `json:"-"`
	city string
}

fn (u user) greet(msg) {
	return u.name + " (" + string(u.age) + ") says: " + msg
}

fn (u user) birthday() {
	u.age = u.age + 1
	return u.age
}

u := user{name: "Alice", age: 30, city: "London"}

// Verify method calls
println("greet:", u.greet("Hello!"))
println("birthday:", u.birthday())
println("new age:", u.age)

// Verify JSON encoding with tags
json_str := json.encode(u) |> string
println("json:", json_str)