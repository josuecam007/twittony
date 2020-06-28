package routers

import (
	"encoding/json"
	"net/http"

	"github.com/josuecam007/twittony/bd"
	"github.com/josuecam007/twittony/models"
)

/*Registro es la funcion para crear en la base de datos el registro de usuarios*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error()+"tony", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, usuarioExistente, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if usuarioExistente == true {
		http.Error(w, "El usuario ya existe en la base de datos", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
