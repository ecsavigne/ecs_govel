package helpers

import (
	"errors"
	"regexp"
	"strconv"
)

func ValidateCorreio(correo string) (string, error) {
	patron := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	exprCorreo := regexp.MustCompile(patron)

	if exprCorreo.Match([]byte(correo)) {
		return correo, nil
	} else {
		return correo, errors.New("No es un correo")
	}
}

func ValidateInt(num string) (int, error) {
	if num == "" {
		return 0, errors.New("Error convirtiendo cadena vacia a numero, defaul = 0")
	}
	return strconv.Atoi(num)
}

func ValidateCi(ci string) (string, error) {
	patron := `^\d{11}$`
	exprCi := regexp.MustCompile(patron)

	if exprCi.Match([]byte(ci)) {
		return ci, nil
	} else {
		return ci, errors.New("No es un ci valido")
	}
}
