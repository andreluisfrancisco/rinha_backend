package payments

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Payment struct {
    Valor int `json:"valor"`
    Descricao string `json:"descricao"`
}

var PaymentChannel = make(chan Payment, 10000)
func HandlePayments(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var p Payment
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if p.Valor <= 0 || len(p.Descricao) == 0 || len(p.Descricao) > 10 {
        w.WriteHeader(http.StatusUnprocessableEntity)
        return
    }

    PaymentChannel <- p

    w.WriteHeader(http.StatusAccepted)
}

func PaymentWorker() {
    for payment := range PaymentChannel {

        fmt.Printf("Processando pagamento: Valor=%d, Descrição=%s\n", payment.Valor, payment.Descricao)

        time.Sleep(100 * time.Millisecond)
    }
}