package main

import(
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)

)

func add(x, y, z int) int {
	return (x+y)*z
}

func swap(x, y string) (string, string){
	return y, x
}

func testVariables(){
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

}

func otherTestVariables(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func main(){
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Nextafter(2,3))
	fmt.Println("PI ", math.Pi)
	fmt.Println(add(42, 13, 10))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	testVariables()

	otherTestVariables()
}