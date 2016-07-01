# hcache

```go
package main

import (
	"fmt"

	"github.com/s2gatev/hcache"
)

type data struct {
  value int
}

func main() {
  cache := hcache.New()
  cache.Insert(&data{value: 42}, "important", "answer")
  
  // ...
  
  d := cache.Get("important", "answer").(*data)
  fmt.Println(d.value) // 42
}
```
