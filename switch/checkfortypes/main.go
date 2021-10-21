package main

var (
	any interface{}
)

func convert(any interface{}) {
	switch t := any.(type) {
	case int:
		println("i is interger", t)
	case string:
		println("i is string", t)
	case float32:
		println("i is float32", t)
	default:
		println("sorry type not found")
	}
}

func main() {
	any = 100
	convert(any)
	any = float64(99.99)
	convert(any)
	any = "foo"
	convert(any)
	convert(float32(188.0))
}