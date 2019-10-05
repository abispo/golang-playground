package handlers

import "net/http"

// OkHandler ...
func OkHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok.\n"))
}
