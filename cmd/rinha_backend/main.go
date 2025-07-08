package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"

    "github.com/andreluisfrancisco/rinha_backend/internal/payments"
    "github.com/andreluisfrancisco/rinha_backend/internal/summary"
)

func main() {
    port := getEnv("HTTP_PORT", "9999")
    numWorkers := getEnvAsInt("PAYMENT_WORKERS", 16)

    fmt.Printf("Iniciando %d workers de pagamento...\n", numWorkers)

    for i := 0; i < numWorkers; i++ {
        go payments.PaymentWorker()
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/payments", payments.HandlePayments)
    mux.HandleFunc("/payments-summary", summary.HandleSummary)
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Rinha Backend"))
    })

    fmt.Println("Servidor on :" + port)

    if err := http.ListenAndServe(":"+port, mux); err != nil {
        fmt.Println("Erro ao iniciar servidor:", err)
        os.Exit(1)
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func getEnvAsInt(key string, fallback int) int {
    if valueStr, ok := os.LookupEnv(key); ok {
        if value, err := strconv.Atoi(valueStr); err == nil {
            return value
        }
    }
    return fallback
}


