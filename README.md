# hashring

一致性 Hash 环

```go
package main

import (
    "hashring"
    "fmt"
)

func main() {
    hashRing := hashring.NewHashRing(150)
    hashRing.AddNodes(map[string]int{
        "192.168.1.101": 1,
        "192.168.1.101": 2,
        "192.168.1.101": 3
    })
    nodeName, _ := hashRing.NodeLoadBalance(key)
    fmt.Printf("node load balance , node name %s", nodeName)
}
```
