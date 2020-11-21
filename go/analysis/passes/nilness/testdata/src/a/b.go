package a

func g1(x, y *X) {
	print(x.f) // want "may nil dereference in field selection"
}

func g2() (m1, m2 map[string]string) {
	m2 = map[string]string{}

	m2["a"] = "b"
	m1["a"] = "b" // want "nil dereference in map update"
	return
}

func g3(i interface{}) string {
	m := i.(map[string]string) // want "may nil assert in type assertion"
	return m["x"]
}

type T struct {
	a int
}

func (t *T) g() int {
	return t.a // we don't catch this panic because t is a receiver
}
func g4(a *int) {

	*a = 1 // want "may nil assert in type assertion"

	*a = 1 // we don't catch this panic because just checked it
}
