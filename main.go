package main

import (
    "log"
    "net/http"
    "text/template"
    "path/filepath"
    "sync"
)

type templateHandler struct {
    once        sync.Once
    filename    string
    temp1       *template.Template  // 1つのテンプレート
}

// ServeHTTP: HTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.temp1 =
            template.Must(template.ParseFiles(filepath.Join("templates",
                t.filename)))
    })
    t.temp1.Execute(w, nil)
}

func main() {
    // ルート
    http.Handle("/", &templateHandler{filename: "chat.html"})
    // Webサーバーを開始します
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAddServe:", err)
    }
}
