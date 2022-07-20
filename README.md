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
	cache.Set("key", "value")
	value, ok := cache.Get("key")
	fmt.Println(value, ok) // [my_first_cache], <nil>

	value2, ok := cache.Get("nokey")
	fmt.Println(value2, ok) // <nil>, key doesnt exist
}

```
