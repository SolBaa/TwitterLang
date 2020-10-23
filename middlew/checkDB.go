package middlew

import (
	"net/http"

	"github.com/SolBaa/TwitterLang/bd"
)

// CheckDB jigjeriogjeroigjeiogjerioj
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
