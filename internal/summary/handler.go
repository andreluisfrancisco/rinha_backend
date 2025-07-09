package summary

import (
    "encoding/json"
    "net/http"
    "sync"
)

type SummaryData struct {
    TotalAmount       int `json:"total_processado"`
    TotalTransactions int `json:"total_transacoes"`
}

var globalSummary struct {
    sync.RWMutex
    data SummaryData
}

func RecordSuccessfulPayment(amount int) {
    globalSummary.Lock()
    globalSummary.data.TotalAmount += amount
    globalSummary.data.TotalTransactions++
    globalSummary.Unlock()
}

func HandleSummary(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    globalSummary.RLock()
    dataCopy := globalSummary.data
    globalSummary.RUnlock()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(dataCopy)
}