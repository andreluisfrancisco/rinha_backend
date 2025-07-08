package payments

import (
    "fmt"
    "net/http"
)

func PaymentWorker() {
    fmt.Println("Worker is alive")
}

func HandlePayments(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("payments endpoint"))
}
