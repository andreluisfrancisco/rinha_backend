package healthcheck

import (
    "net/http"
    "sync"
    "time"
)

type Status int

const (
    StatusUnknown Status = iota
    StatusHealthy
    StatusUnhealthy
)

var serviceHealth struct {
    sync.RWMutex
    statuses map[string]Status
}

var client = &http.Client{
    Timeout: 1 * time.Second
}

func init() {
    serviceHealth.statuses = make(map[string]Status)
}

func StartHealthChecks(urls ...string) {
    for _, url := range urls {
        serviceHealth.statuses[url] = StatusUnknown

        go func(u string) {
            for {
                resp, err := client.Get(u)

                serviceHealth.Lock()
                if err != nil || resp.StatusCode >= http.StatusInternalServerError {
                    serviceHealth.statuses[u] = StatusUnhealthy
                } else {
                    serviceHealth.statuses[u] = StatusHealthy
                }
                serviceHealth.Unlock()

                if resp != nil {
                    resp.Body.Close()
                }

                time.Sleep(500 * time.Millisecond)
            }
        }(url)
    }
}

func IsHealthy(url string) bool {
    serviceHealth.RLock()
    defer serviceHealth.RUnlock()
    return serviceHealth.statuses[url] == StatusHealthy
}
