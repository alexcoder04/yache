
# yache

A very simple library to store data in a Go variable for a limited number of seconds.

## How to use

```sh
go get github.com/alexcoder04/yache
```

```go
import (
    "time"

    "github.com/alexcoder04/yache"
)

func main() {
    // create a store to save data only for 5 seconds
    store := yache.NewStore(5)

    // save data
    store.Set("test", "Hello World")

    // get data
    fmt.Println(store.Get("test")) // prints "Hello World"

    // wait
    time.Sleep(7 * time.Second)

    // get data again
    fmt.Println(store.Get("test")) // prints nil
}
```
