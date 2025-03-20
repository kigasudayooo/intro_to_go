// main.go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    // /timeにアクセスが来たときのハンドラを登録
    http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now()
        fmt.Fprintf(w, currentTime.String())
    })

    // サーバー
    http.ListenAndServe(":8010", nil)
}
