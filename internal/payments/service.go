package payments

import (
    "fmt"
    "net/http"
)

func PaymentWorker() {
    fmt.Println("executando")
}

func HandlePayments(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("payments"))
}