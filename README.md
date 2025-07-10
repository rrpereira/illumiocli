# illumiocli

Go client library for interacting with the Illumio PCE API.

## Installation

```bash
go get github.com/rrpereira/illumiocli
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/rrpereira/illumiocli"
)

func main() {
    // Create a new client
    client := illumiocli.NewClient(
        "https://your-pce-host.com",
        "your-username", 
        "your-api-key",
        "v1",
    )
    
    // Get clusters
    clusters, err := client.GetClusters()
    if err != nil {
        log.Fatal(err)
    }
    
    for _, cluster := range clusters {
        fmt.Printf("Cluster: %s - %s\n", cluster.Name, cluster.Description)
    }
}
```

## Features

- Container cluster management
- Authentication with PCE API
- Type-safe API responses

## License

See LICENSE file for details.
