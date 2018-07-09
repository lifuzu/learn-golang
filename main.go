package main

import (
    "fmt"
    // Rename the following math package to `m`
    m "github.com/lifuzu/learn-golang/math"
    "github.com/lifuzu/learn-golang/shape"
    // Call the methods without explicitly stating the package - NOT RECOMMENDED
    . "github.com/lifuzu/learn-golang/person"
)

var (
    _ int64 = s()
)

func init() {
    fmt.Println("init in main.go")
}

func init() {
    fmt.Println("init in main.go - second time")
}

func s() int64 {
    fmt.Println("s function in main.go")
    return 54
}

func main() {

    fmt.Println("main begins")

    var a [5]float64
    a[4] = 100
    fmt.Println(a)

    var b = [5]int{0}
    fmt.Println(b)
    fmt.Printf("Hello world!\n")

    var oc = [5]float64{12, 23, 34, 45, 56}

    fmt.Println("Average float64{12, 23, 34, 45, 56}: ", m.Average(oc[:]))

    c := shape.Circle{X: 4, Y: 5, R: 6}
    fmt.Println(c.Area())

    r := shape.Rectangle{2, 3, 4, 5}

    fmt.Println(shape.TotalArea(&c, &r))

        // Using embedded types
    andr := Android{Person: Person{Name: "Android"}}
    andr.Talk()


    fmt.Println("main ends")
}





