package main

import (
    "fmt"
)

func main() {
    array1 := [3]string{"a", "a", "a"}
    array2 := [3]string{"a", "a", "a"}

    fmt.Printf("Array 1: \n%s\n", array1)
    fmt.Printf("Array 2: \n%s\n", array2)

    count := 0

    for i := 0 ; i<len(array1); i++ {
        if array1[i] != array2[i] {
            count++;
            fmt.Printf("Index ke %d berbeda\n", i)
        }
    }

    if count == 0 {
        fmt.Printf("Kedua array sama\n")
    }

}