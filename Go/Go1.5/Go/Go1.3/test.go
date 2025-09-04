package main

import "fmt"

func main() {
mapp := map[string]int{"first": 1, "second": 2}
key := "second"
value, ok := mapp[key]
if ok {
fmt.Printf{"ключ %s найден %d", key, value}
} else {
fmt.Printf("ключа %s нет %d", key, value)
}
}

