package summary

import (
    "net/http"
)

func HandleSummary(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("summary endpoint"))
}
