package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("Value of i:", i)
    }

    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    //While loop in go
    i := 0
    for i < 5 {
        fmt.Println("Value of i:", i)
        i++
    }


    rand.Seed(time.Now().UnixNano())

    weather := rand.Intn(3)

    switch weather {
    case 0:
        fmt.Println("The weather is sunny ☀️. Time to wear shades and enjoy!")
    case 1:
        fmt.Println("The weather is cloudy ☁️. Perfect for a cozy day indoors!")
    case 2:
        fmt.Println("The weather is rainy 🌧️. Time to dance in the rain and sing!")
    default:
        fmt.Println("Hmm, seems like the weather is unpredictable today! 🤔")
    }
}
