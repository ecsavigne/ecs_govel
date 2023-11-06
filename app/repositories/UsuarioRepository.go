package repositories

import (
	"crypto/tls"
	"ecs_govel/app/models"
	"ecs_govel/config/db"
	"errors"
	"fmt"
	"net/mail"
	"net/smtp"
)

type UsuarioRepository struct {
}

// Metodos del Repositorio
func (u UsuarioRepository) RegistrarUsr(nome, correio, pass string) interface{} {
	usr := new(models.Usuario)
	usr.Nombre = nome
	usr.Password = pass
	usr.Mail = correio
	if err := db.Orm.Create(usr).Error; err != nil {
		panic("[Repositories.Usuarirepository.RegistrarUsr: Line 18] - " + err.Error())
	}
	return true
}

func (u UsuarioRepository) CambiarPass(idUsuario int, newPass string) interface{} {
	println("idUsuario:", idUsuario, "newPass:", newPass)
	if res := db.Orm.Model(new(models.Usuario)).Where("id = ?", idUsuario).Update("password", newPass); res.Error != nil {
		return errors.New("[Repositories.Usuarirepository.CambiarPass: Line 26] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (u UsuarioRepository) RecuperarPass(correio string) interface{} {
	res := db.Orm.Model(new(models.Usuario)).Where("mail = ?", correio).Update("password", "12345")
	if res.Error != nil {
		return errors.New("[Repositories.Usuarirepository.RecuperarPass: Line 37] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (u UsuarioRepository) VerificarUsr(usr *models.Usuario) interface{} {
	res := db.Orm.Model(new(models.Usuario)).Find(new(models.Usuario), usr)
	if res.Error != nil {
		return errors.New("[Repositories.Usuarirepository.RecuperarPass: Line 37] - " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return false
	}
	return true
}

func sendMail(msg string) {
	println("Enviar correo")
	// Configurar los detalles del servidor SMTP
	smtpHost := "smtp.mail.com"
	smtpPort := 465
	username := "ecsavigne@gmail.com"
	password := "SAVCoe8612..levisM"

	// Configurar el mensaje de correo electrónico
	from := mail.Address{"ecs", "ecsavigne@yahoo.com"}
	to := mail.Address{"ecs", "ecsavigne@gmail.com"}
	subject := "Asunto del correo"
	body := "TTTTTTT"

	header := make(map[string]string)
	header["from"] = from.String()
	header["to"] = to.String()
	header["subject"] = subject

	message := ""
	for k, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, value)
	}
	message += body
	// Crear el mensaje de correo

	// Autenticarse en el servidor SMTP
	auth := smtp.PlainAuth("", username, password, smtpHost)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}
	conn, err := tls.Dial("tcp", smtpHost+":"+string(smtpPort), tlsConfig)
	client, err := smtp.NewClient(conn, smtpHost)
	err = client.Auth(auth)
	if err != nil {
		println("Error client Auth")
	}

	err = client.Mail(from.String())
	if err != nil {
		println("Error en MAil")
	}
	err = client.Rcpt(to.Address)
	if err != nil {
		fmt.Println("Error Rcpt")
	}

	w, err := client.Data()
	if err != nil {
		println("Error Data")
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		println("Error Write")
	}
	err = w.Close()
	if err != nil {
		println("TTT")
	}

	println("Correo electrónico enviado correctamente")
}
