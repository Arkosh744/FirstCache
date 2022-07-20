Go in-memory cache
================================
Go Scheduler helps you to manage functions that should be executed every N seconds/minutes/hours etc.

See it in action:

to install module:

    go get -u github.com/Arkosh744/FirstCache

## Example #1

```go
package main

import (
	"fmt"
	"github.com/Arkosh744/FirstCache"
)

func main() {
	cache := FirstCache.NewCache()

	cache.Set("keykey", "valuetop1")
	fmt.Println(cache) // &{map[keykey:[valuetop1]]}

	value, ok := cache.Get("keykey")
	fmt.Println(value, ok) // [my_first_cache] <nil>

	value, ok = cache.Get("nokey")
	fmt.Println(value, ok) // <nil> key doesnt exist

	cache.Delete("keykey")
	value, ok = cache.Get("keykey")
	fmt.Println(value, ok) // [my_first_cache] <nil>
}
```
