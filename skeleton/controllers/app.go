package controllers

import "net/http"
import "fmt"

func index(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	fmt.Fprintf(w, "Home")
}
