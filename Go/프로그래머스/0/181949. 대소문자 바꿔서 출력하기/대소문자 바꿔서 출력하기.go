package main

import "fmt"

func main() {
    var s1 string
    fmt.Scan(&s1)
    result := ""
    for i := 0; i < len(s1); i++ {
        if s1[i] >= 'a' && s1[i] <= 'z' {
            result = result + string(s1[i]-32)
        } else {
            result = result + string(s1[i]+32)
        }    
    }
    fmt.Println(result)
}