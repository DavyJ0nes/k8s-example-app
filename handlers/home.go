package handlers

import "net/http"
import "fmt"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! Your Request Was Processed")
}
