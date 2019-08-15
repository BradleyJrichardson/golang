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
