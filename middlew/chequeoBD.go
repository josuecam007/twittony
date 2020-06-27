package middlew

import (
	"net/http"

	"github.com/josuecam007/twittony/bd"
)

/*ChequeoBD es el middleware que me permite conocer el estado de la base de datos, recibe http.HandlerFunc y revisa contra bd.ChequeoConexion() si hay conexion exitosa devuelve http.HandlerFunc para sederlo al siguiente eslabon */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == false {
			http.Error(w, "Conexion de base de datos perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
