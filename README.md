# hcache

[![Build Status](https://travis-ci.org/s2gatev/hcache.svg?branch=master)](https://travis-ci.org/s2gatev/hcache)
[![Coverage Status](https://coveralls.io/repos/github/s2gatev/hcache/badge.svg?branch=master)](https://coveralls.io/github/s2gatev/hcache?branch=master)

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
