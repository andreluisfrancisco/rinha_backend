package healthcheck

import "fmt"

func StartHealthCheck() {
    go func() {
        for {
            fmt.Println("health checking...")
        }
    }()
}
