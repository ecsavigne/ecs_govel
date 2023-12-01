package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ValidateCorreio(correo string) (string, error) {
	patron := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	exprCorreo := regexp.MustCompile(patron)

	if exprCorreo.Match([]byte(correo)) {
		return correo, nil
	} else {
		return correo, errors.New("No es un correo válido")
	}
}

func ValidateInt(num string) (int, error) {
	if num == "" {
		return 0, errors.New("Error convirtiendo cadena vacia a numero, defaul = 0")
	}
	return strconv.Atoi(num)
}

func ValidateBool(num string) (bool, error) {
	if strings.ToLower(num) != "true" && strings.ToLower(num) != "false" {
		return false, errors.New("Error convirtiendo cadena vacia a numero, defaul = 0")
	}
	return strconv.ParseBool(num)
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

// // dado un string que representa un numero devuelve el time.Time en Hora equivalente
// // Si no es un Numero entero devuelve error
// func ValidateHora(num string) (time.Time, error) {
// 	val, err := strconv.Atoi(num)
// 	if err != nil {
// 		return _, errors.New("Error convirtiendo cadena vacia a numero, defaul = 0")
// 	}
// 	horas, _ := time.ParseDuration(""+num)
// 	return , nil
// }

func ValidateFecha(fecha string) (time.Time, error) {
	const shortForm = "2006-01-02"
	t, err := time.Parse(shortForm, fecha)
	if err != nil {
		fmt.Println(fecha)
		return time.Time{}, err
	}
	return t, nil
}
