package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/andreluisfrancisco/rinha_backend/internal/payments"
    "github.com/andreluisfrancisco/rinha_backend/internal/summary"
)

func main() {
    for i := 0; i < 16; i++ {
        go payments.PaymentWorker()
    }

    http.HandleFunc("/payments", payments.HandlePayments)
    http.HandleFunc("/payments-summary", summary.HandleSummary)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Rinha Backend"))
    })

    fmt.Println("Server on :9999")
    err := http.ListenAndServe(":9999", nil)
    if err != nil {
        fmt.Println("Erro servidor:", err)
        os.Exit(1)
    }
}