## Control Flow

```go
for init; condition; post {}

for i := 0; i <= 100; i++ {
    fmt.Println(i)
    }

j := 7

for j <= 7 {
  fmt.Println("for while loop")
  j++
}

for {
  fmt.Println("infinite loop")
  if true {
    break
  }
}
```

```go
func main() {
	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)

		}
	}
}

```

```go
func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)

	}
	fmt.Println("done.")
}
```

### Initialization statement

allows for reduction of scope of variables, as the assignment is inside of the if statement it's contained!

```go
func main() {o

	if x := 42; x == 42 {
		fmt.Println("001")
	}
}

```

```go

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

```go
func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
	case (4 == 4):
		fmt.Println("also true, does it print?")
	}
}
```

```go
func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
	}
}
```

```go
func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	default:
		fmt.Println("this is default")
	}
}

```

```go
func main() {
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}
}
```

```go
func main() {
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}
}
```

## Grouping Data

```go
package main

import "fmt"

func main() {
	// x := type{values} composite literal
	//  a SLICE allows you to group together values of the same type
	x := []int{4, 3, 2, 1, 0}
	length := len(x)

	if length == 5 {
		fmt.Println("length is 5")
	}

	fmt.Printf("the array we have constructed with a composite literal is %v\n", x)
	fmt.Printf("which has a length of %v\n", length)
	fmt.Printf("the first element of the array is %v\n", x[0])

	// the range keyword in the below loop is similar to a forEach in javaScript
	for index, value := range x {
		fmt.Printf("index: %v\t value: %v\n", index, value)
	}

	// slicing a slice
	slice1 := x[1:]
	slice2 := x[5:]
	slice3 := x[1:3]

	fmt.Printf("slice1: %v\t slice2: %v\t slice3: %v\t\n", slice1, slice2, slice3)

	// appending to a slice
	// append is a built in function of the langauge
	x = append(x, 1, 2, 3, 4, 5)
	fmt.Println(x)

	x = append(x, slice1...)
	fmt.Println(x)

	// delete from a slice
	x1 := []int{4, 5, 6, 7, 8, 9}
	x2 := []int{10, 11, 12, 13, 14, 15, 16}

	fmt.Printf("x1 = %v\nx2 = %v\n", x1, x2)

	x1 = append(x1, x2...)
	fmt.Printf("x1 is now %v\n", x1)

	// to delete from a slice we need to slice out the values we desire and then append them back together
	fmt.Printf("x1[:2] = %v\nx1[4:] = %v\n", x1[:2], x1[4:])
	// the below code slices out the [6,7]
	x1 = append(x1[:2], x1[4:]...)
	fmt.Printf("x1 is now %v\n", x1)

	// built in function MAKE
	// so a slice is built atop of an array slices and dynamic and can change size... when they do change size they need to reconstruct the array that they are built upon... therefore using processing
	// MAKE removes this processing
	fmt.Println("making shit")
	x3 := make([]int, 10, 11)
	fmt.Println(x3)
	fmt.Println(len(x3))
	fmt.Println(cap(x3))

	x3[0] = 1
	x3 = append(x3, 11)
	fmt.Printf("x3 = %v\t len = %v\t capacity = %v\n", x3, len(x3), cap(x3))
	x3 = append(x3, 12)
	fmt.Printf("x3 = %v\t len = %v\t capacity = %v\n", x3, len(x3), cap(x3))

	// when we exceed the capicity the go runtime creates a new array and copies all the numbers over!
	// so there's no downside of using make?

	// a slice of a slice of TYPE string
	// multidimensional slice
	jb := []string{"James", "Bond", "Chocolate", "Martinini"}
	mp := []string{"Miss", "March", "Vanilla", "Red Wine"}

	fmt.Printf("jb = %v\n mp = %v\n", jb, mp)
	// below is a slice of a slice string
	xp := [][]string{jb, mp}

	fmt.Printf("slice of slices %v", xp)
}

```

# Maps

```go
package main

import "fmt"

func main() {
	// the type of this data structure is a "map[string]int{}"

	x := map[string]int{
		"James":            32,
		"Miss Moneypenney": 21,
	}
	fmt.Println(x)

	fmt.Printf("accessing a value from a key x[James] = %v\n", x["James"])
	fmt.Printf("wrong key returns default value which is false = %v\n", x["Stu"])

	v, ok := x["Stu"]
	fmt.Printf("v = %v\t ok = %v\n", v, ok)

	// a very common idiomatic chunk of code lies below
	if v, ok := x["James"]; ok {
		fmt.Printf("entered conditional v = %v\n%v\n", v, x)
	}

	// adding new item to map
	x["Todd"] = 33
	fmt.Println("our map is now\n\n", x)

	// for range of map syntax
	fmt.Println("--------------- range loop below -------------")
	for k, v := range x {
		fmt.Printf("key: %v\t value: %v\n", k, v)
	}

	fmt.Println("--------------- range loop below -------------")
	y := []int{1, 2, 3, 4, 5, 6}
	for i, v := range y {
		fmt.Printf("index: %v\t value: %v\n", i, v)
	}

	// deleting something from a map
	// delete(map, "key")
	delete(x, "James")
	fmt.Println(x)

	if v, ok := x["Miss Moneypenney"]; ok {
		fmt.Println(v)
		delete(x, "Miss Moneypenney")
	}
}

```

```go
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range x {
		fmt.Printf("index: %v\t value: %v\n", i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	fmt.Printf("type: %T\nvalue: %v\n", x, x)

	var y [11]int
	for i := 0; i <= 10; i++ {
		y[i] = i
		fmt.Println(y)
	}

	// slicing slices
	slice1 := append(x[:1], y[:1]...)
	fmt.Println(slice1)

	// appending slices
	slice2 := []int{42, 43, 45, 46, 47, 48, 49}
	slice2 = append(slice2, 50)
	fmt.Println(slice2)
	slice2 = append(slice2, 51, 52, 53)
	fmt.Println(slice2)
	slice3 := []int{54, 55, 56, 57}
	slice2 = append(slice2, slice3...)

	// deleting from a slice
	slice4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice5 := append(slice4[:3], slice4[8:]...)
	fmt.Println(slice5)

	// big slice
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(states)
	fmt.Printf("the length of the slice is %v\nthe capcity of the slice is %v\n", len(states), cap(states))

	for j := 0; j <= 49; j++ {
		fmt.Printf("index: %v\tstate: %v\n", j, states[j])
	}

	// slice of a slice string
	slice6 := []string{"James", "Bond", "Shaken, not stirred"}
	slice7 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	slice8 := [][]string{slice6, slice7}
	fmt.Printf("slice of a slice of type string %v\n", slice8)

	// maps
	map1 := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for k, v := range map1 {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}


```

# Structs

```go
package main

import "fmt"

// the type is person and the underlying type is struct
type person struct {
	first string
	last  string
	age   int
}

// a secretAent is everything a person is and more..
type secretAgent struct {
	person
	ltk bool
}

func main() {
	// p1 is a variable that stores a value of type person
	// known as a composite/aggregate/complex data structure as it contain different types
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   24,
	}
	fmt.Printf("p1: %v\tp2: %v\n", p1, p2)
	// we can access values with dot notation like javascript!
	fmt.Println(p1.first, p2.first)

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32},
		ltk: true,
	}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(sa1.person.first, sa1.person.last, sa1.age, sa1.ltk)
	// both work! you dont need to specify that the key first is part of the embedded type person but will need to if there is a naming collision

	// if you only have to use a struct in one little area the below method of creating an anonmymous struct is preffered for cleaner code
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p3)
}


```

```go
package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

type vehicle struct {
	doors  int
	colors []string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

	// Complex structs
	vehicle1 := truck{
		vehicle: vehicle{
			doors:  6,
			colors: []string{"black", "grey", "silver"},
		},
		fourWheel: true,
	}

	vehicle2 := sedan{
		vehicle: vehicle{
			doors:  4,
			colors: []string{"white", "red", "blue"},
		},
		luxury: false,
	}

	vehicle3 := sedan{
		vehicle: vehicle{
			doors:  4,
			colors: []string{"silver", "black"},
		},
		luxury: true,
	}

	// anonoymous Structs
	fmt.Println(vehicle1, vehicle2, vehicle3)

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}

```

# Functions

```go
package main

import "fmt"

func main() {
	foo()
	bar("Brad")
	x := woo("Brad")
	y, j := tang("Brad", "Richardson")
	fmt.Println(x)
	fmt.Println(y, j)
	sum := sum(4, 4, 6, 8)
	fmt.Println("the sum is", sum)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello from bar, ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func tang(x string, y string) (string, bool) {
	a := fmt.Sprint(x, " ", y, " ", "says Hello")
	b := true
	return a, b
}

// an example of a function with variadic parameters
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0 // var sum int
	for i, v := range x {
		sum += v
		fmt.Println("for item in position,", i, "we are now adding,", v, "to the total", sum)
	}
	return sum

}
```

## Unfurling a Slice and multiple parameters with one variadic

```go
package main

import (
	"fmt"
)

func main() {
	xi := []int{12,12,12}
	x := sum("james", xi...)
	fmt.Println("The total is", x)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}
```

## Defer keyword

```go
func main() {
	defer foo()
	bar()
}
// right when func main exists the defered functions then run
// this is used for opening and closing files...

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}

```

## Methods

```go
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("i am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"bond",
			"james",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"miss",
			"moneypenny",
		},
		ltk: true,
	}

	sa1.speak()
	sa2.speak()
}

```

# Interfaces and Polymorhism

[https://www.ardanlabs.com/blog/2015/09/composition-with-go.html]

```go
// a value can be of more than one type
package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// interfaces allow a value to be of more than one type and thus used polymorphically, which is when a function has different outputs based on the secondary type of the struct in this example
type human interface {
	speak()
}
// an interface says hey baby if you have these methods then your my type...

// we can pass both instances of the structs as they have the attached function speak which satifys the interface

func bar(h human) {
	switch h.(type) { // special switch statement for types
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}


```

# Anonymous Functions

```go
func main() {
	func() {
		fmt.Println("anonymous func")
	}()

	func(x int) {
		fmt.Println("my age:", x)
	}(23)
}
```

# Function Expressions

```go
func main() {
	// in go functions are first class citizens and it can be treated as all other types can be ...

	// a func expression is simple assigning a function to a variable.. like is JS we do const foo = () => console.log("my func expression")

	f := func(x int) {
		fmt.Println("my func expression", x)
	}

	f(42)
}
```

# Returning a function

```go
func main() {

	s1 := foo()
	fmt.Println(s1)

	fmt.Println(bar()())
	x := bar()
	fmt.Printf("%T\n", x)
	fmt.Println(x())
	i := x()
	fmt.Println(i)
}

// returning a string
func foo() string {
	return "hello world"

}

func bar() func() int {
	return func() int {
		return 451
	}
}

```

# Callback

```go
package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)
}

func sum(xi ...int) int {
	// fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}


func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

```

,

# Closure

```go
package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	{
		y := 35
		fmt.Println(y)
	}

	a := incrementor()
	fmt.Println(a())
}

func foo() {
	fmt.Println("hello")
	x++
}

func incrementor() func() int {
	var xi int
	return func() int {
		xi++
		return xi
	}
}

```

# Playing with functions

```go

package main

import (
	"fmt"
	"math"
)

// Exercise Four
type person struct {
	first string
	last  string
	age   int
}

func (s person) describe() {
	fmt.Println("i am", s.first, s.last, s.age)
}

// Exerise Five
type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	// Exercise One
	fmt.Println("Exercise 1 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("foo returns:", foo())
	x, y := bar()
	fmt.Println("bar returns:", x, y)

	// Exercise Two
	fmt.Println("Exercise 2 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo2(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar2(ii2)
	fmt.Println(n2)

	// Exercise Three
	fmt.Println("Exercise 3 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	defer foo3()

	// Exercise Four
	fmt.Println("Exercise 4 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	p1 := person{
		"Brad",
		"Richardson",
		23,
	}
	p2 := person{
		"Adam",
		"Smith",
		31,
	}
	foo4(p1, p2)

	// Exercise Five
	fmt.Println("Exercise 5 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}
	info(circ)
	info(squa)

	// Exercise Six
	fmt.Println("Exercise 6 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	x6 := func(x int) int {
		return x * 10
	}(10)
	fmt.Println(x6)


	// Exercise Seven
	fmt.Println("Exercise 7 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	func foo7() {
		return foo8(){

		}
	}

}

func foo() int {
	return 57
}
func bar() (int, string) {
	return 75, "string"
}

func foo2(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func foo3() {
	fmt.Println("This is a defered function")
}

func foo4(x person, y person) {
	x.describe()
	y.describe()
}
```

# Introduction to Pointers

```go
package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	// the & operator shows us the address of the value
	// and what is the address?
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // type is a pointer to an int which is a differnt type! *int != int

	// var b *int = &a
	b := &a
	fmt.Println(b)  // b now points to that memory address
	fmt.Println(*b) // the value at that address!!
	// the * operator shows us the value at the address de-referencing the address
	fmt.Println(*&a)

	// changing the value at the address
	*b = 43

	fmt.Println(a)

	// when to use pointers?
	// pointers are performant to pass around big chunks of data, just passing an address is better than passing the chunk

	// de-referance an address and change the value that is at that address

	fmt.Println("------------------------------")
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}


```

```go
package main

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Brad Richardson"
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}


// p *person = value of type pointer to person
func changeMe(p *person){
 p.name = "Steven Heinz"
}
```

# My first go Program!

taking data from a struct and marshal/unmarshal it to json/struct

```go
// https://godoc.org/encoding/json#Marshal

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	j := encodePersonToJSON(p1, p2)
	p := decodeJSONtoPerson(j)

	fmt.Println(j)
	fmt.Println(p)
}

// func Marshal(v interface{}) ([]byte, error)
func encodePersonToJSON(p1 person, p2 person) string {
	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	return string(bs) // string(bs) - convert byte slice to a string
}

// func Unmarshal(data []byte, v interface{}) error
func decodeJSONtoPerson(j string) []person {
	var people []person
	bs := []byte(j)
	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
}
	return people
}


```

# Another method and example

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

}

```

# Digging into things

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground")
	os.Stdout.Write([]byte("Hello, playground"))
}


```

# Sort Package

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

```

# Custom Sort

```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

// sort uses these three methods...
func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}

```

# Using bcrypt package

```go
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}

```

# Concurrency

take advantage of those multiple cores

# Wait Group

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

```

```go

```

```go

```

```go

```

```go

```

```go

```
