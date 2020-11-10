package bd

import (
	"github.com/SolBaa/TwitterLang/models"
	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin realiza el qeucheo de login de la DB
func IntentoLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	passWord := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passWord)
	if err != nil {
		return usu, false
	}
	return usu, true

}
