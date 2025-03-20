package main
import (
    "fmt"
    "net/http"
)
// トップページの処理
func handleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ようこそ、ホームページへ！")
}
// あいさつページの処理
func handleGreeting(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name") // URLパラメータから名前を取得
    if name == "" {
        name = "ゲスト"
    }
    fmt.Fprintf(w, "こんにちは、%sさん！", name)
}
func main() {
    // 各URLパスと対応する処理を設定
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/greeting", handleGreeting)
    // サーバー起動
    fmt.Println("サーバーを起動します... http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

