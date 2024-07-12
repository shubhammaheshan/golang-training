// package main

// import "fmt"

// // Define an interface
// type Animal interface {
// 	Speak() string
// }

// // Define a struct that implements the interface
// type Dog struct{}

// func (d Dog) Speak() string {
// 	return "Woof!"
// }

// type Cat struct{}

// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// func main() {
// 	var a Animal

// 	a = Dog{}
// 	fmt.Println(a.Speak())

// 	a = Cat{}
// 	fmt.Println(a.Speak())
// }

// package main

// import "testing"

// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
// 	}
// }

// func Add(a, b int) int {
// 	return a + b
// }

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
