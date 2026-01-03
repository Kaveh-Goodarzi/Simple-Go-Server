package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("OK"))
	})
}

