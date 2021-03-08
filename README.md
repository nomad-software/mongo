# mongo

**A straightforward money library for Go**

---

## Overview

Mongo is a straightforward money library for Go that can make it easy to
handle bug prone arithmetic when dealing with money.

## Documentation

https://pkg.go.dev/github.com/nomad-software/mongo#section-documentation

## Example 1

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nomad-software/mongo"
)

func main() {

	cost, err := mongo.GBP(1059)

	if err != nil {
		log.Fatal("Error occured creating money")
	}

	net := cost.Div(1.2)
	tax := cost.Sub(net)
	json, _ := json.Marshal(cost)

	fmt.Printf("cost: %s (net: %s, tax: %s)\n", cost, net, tax)
	fmt.Println(string(json))
}
```

### Output

```
cost: £10.59 (net: £8.83, tax: £1.76)
{"currency":"GBP","formatted":"£10.59"}
```

## Example 2

```go
package main

import (
	"fmt"
	"log"

	"github.com/nomad-software/mongo"
)

func main() {

	pot, err := mongo.GBP(100)

	if err != nil {
		log.Fatal("Error occured creating money")
	}

	shares := pot.Split(3)
	fmt.Printf("Pot: %s, Shares: %s\n", pot, shares)

	shares = pot.Allocate(1, 2, 3)
	fmt.Printf("Pot: %s, Allocations: %s\n", pot, shares)
}
```

### Output

```
Pot: £1.00, Shares: [£0.34 £0.33 £0.33]
Pot: £1.00, Allocations: [£0.17 £0.33 £0.50]
```

## Caveat

Apart from the unit tests, this library is completely untested and written
for the fun of it. Feedback welcome!
