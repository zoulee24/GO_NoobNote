package main

import "fmt"

type test struct {
	a int
	b int
}

func (t *test) get_a() int{
	return t.a
}

func (t test) get_b() int{
	t.b = 5
	return t.b
}

func main() {
	tt := &test{1, 2}
	fmt.Println(tt)
	fmt.Println("")
	fmt.Println(tt.get_a())
	fmt.Println(tt.get_b())
}