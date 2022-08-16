Go in-memory cache with ttl
================================

See it in action:

to install module:

    go get -u github.com/Arkosh744/FirstCache

## Example

```go
package main

import (
	"fmt"
	"github.com/Arkosh744/FirstCache"
)

func main() {
	cache := NewCache()
	ctx := context.Background()

	cache.Set("userId", 42, time.Second*5, ctx)
	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	log.Println(userId.Value) // Output: 42

	err = cache.Delete("userId")
	if err != nil {
		return
	}
	time.Sleep(time.Second * 6) // прошло 6 секунд
	userId, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}

}
```
