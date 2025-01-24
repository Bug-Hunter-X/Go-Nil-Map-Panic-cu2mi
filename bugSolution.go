```go
package main

import "fmt"

func main() {
    var m map[string]int
    // Check if the map is nil before accessing it.
    if m == nil {
        m = make(map[string]int)
    }
    m["a"] = 1
    fmt.Println(m["a"])
}
```