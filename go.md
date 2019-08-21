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
func main() {

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
	for i, v := range x {
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
