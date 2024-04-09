package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "初期実行")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("サーバー起動したよ① :8080")
	http.ListenAndServe(":8080", nil)
}
