# mongo

**A straightforward money library for Go**

---

## Overview

Mongo is a straightforward money library for Go that makes it easy to handle the
usually bug prone arithmetic when dealing with money.

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

	m, err := mongo.MoneyGBP(1055)

	if err != nil {
		log.Fatal("Error occured creating money")
	}

	fmt.Printf("Money: %s\n", m)

	shares := m.Split(3)
	fmt.Printf("Shares: %s\n", shares)

	shares = m.Allocate(1, 2, 3)
	fmt.Printf("Allocations: %s\n", shares)

	json, _ := json.Marshal(m)
	fmt.Println(string(json))
}
```

### Output

```
Money: £10.55
Shares: [£3.52 £3.52 £3.51]
Allocations: [£1.76 £3.52 £5.27]
{"currency":"GBP","formatted":"£10.55"}
```

## Example 2

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nomad-software/mongo"
)

func main() {

	m, err := mongo.PriceGBP(1055, 17.5)

	if err != nil {
		log.Fatal("Error occured creating price")
	}

	fmt.Printf("Price: %s\n", m)

	json, _ := json.Marshal(m)
	fmt.Println(string(json))
}
```

### Output

```
Price: £10.55
{"currency":"GBP","gross":"£10.55","net":"£8.98","tax":"£1.57","taxPercent":17.500000}
```
