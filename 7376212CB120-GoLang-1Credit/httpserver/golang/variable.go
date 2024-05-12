package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func isEven(num int) bool {
    if num%2 == 0 {
        return true
    }
    return false
}

func main() {
    var a int = 10
    var b float64 = 5.5
    var str string = "Hello, GoLang!"
    var isTrue bool = true

    fmt.Println("a =", a)
    fmt.Println("b =", b)
    fmt.Println("str =", str)
    fmt.Println("isTrue =", isTrue)

    if isEven(a) {
        fmt.Println(a, "is even")
    } else {
        fmt.Println(a, "is odd")
    }

    for i := 0; i < 5; i++ {
        fmt.Printf("Iteration %d\n", i+1)
    }

    sum := add(3, 4)
    fmt.Println("Sum of 3 and 4 is", sum)
}
