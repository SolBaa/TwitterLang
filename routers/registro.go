package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SolBaa/TwitterLang/bd"

	"github.com/SolBaa/TwitterLang/models"
)

// Registro es la funcion para crear la bd en el regitro de usurario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, err.Error()+" Error en los datos recibidos", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email del usuario es requerido", 400)
		return
	}
	if len(t.Password) == 0 {
		http.Error(w, "La contraseña es incorrecta", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario con ese Email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario", 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
