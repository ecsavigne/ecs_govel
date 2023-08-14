package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/services/gdrive"
	"strconv"

	"os"
	"strings"

	mt "github.com/gabriel-vasile/mimetype"
)

func StoreMessageFile(file multipart.File, clientOriginalName string, contactPhone string) (string, string, error) {
	filename := strings.Split(clientOriginalName, ".")
	ext := filename[len(filename)-1]
	mimetype := APIExts[ext]
	if mimetype == "" {
		fileHeader := make([]byte, 512)
		file.Read(fileHeader)
		mimetype = mt.Detect(fileHeader).String()
	}

	CreateDirIfNotExist(os.Getenv("MESSAJE_FILES") + "/api/" + contactPhone)
	filepath := os.Getenv("MESSAJE_FILES") + "/api/" + contactPhone + "/" + clientOriginalName
	storedFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", "", err
	}
	defer storedFile.Close()
	io.Copy(storedFile, file)
	return filepath, mimetype, nil
}

func CreateDirIfNotExist(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			logg.ErrorLogger.Printf("\033[31m %v \033[0m\n", err)
		}
	}
}

func DownloadProfilePicture(url string, ContactPhone string) {
	response, e := http.Get(url)
	if e != nil {
		fmt.Println(e)
	}
	defer response.Body.Close()

	CreateDirIfNotExist(os.Getenv("MESSAJE_FILES") + "/messageFiles/" + ContactPhone)
	file, err := os.Create(os.Getenv("MESSAJE_FILES") + "/messageFiles/" + ContactPhone + "/pic_url.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println(err)
	}
}

/*
/////////////////////////////////////////////////////

	 @filePath -> path del fichero a subir en dive
	 @mimeType -> Tipo mime
	 @contactId-> Id del contacto
	 return (
		@idFolderParent -> Id de la carpeta padre
		@fileUploadId
	 )

/////////////////////////////////////////////////////
*/
func SaveInGoogleDrive(filePath, mimeType, contactId string) (idFolderParent, fileUploadId string) {
	fmt.Println("Debug 5.8.1")
	srv := new(gdrive.ServiceGoogle)
	fmt.Println(srv)
	srv.CreateGoogleService()
	fmt.Println(srv)
	fmt.Printf("\nServicio: %v\n", srv.GoogleDriveService)
	fmt.Println("Debug 5.8.2")
	_, _, remainingSpace := srv.GetStorageSpace()
	fmt.Println("Debug 5.8.3")
	limitSpaceDrive, _ := strconv.ParseInt(os.Getenv("LIMIT_SPACE_DRIVE"), 10, 64)
	if remainingSpace > limitSpaceDrive {
		if len(srv.GetFolderByName(contactId)) == 0 {
			idFolderParent = srv.CreateFolder(contactId)
		} else {
			idFolderParent = srv.GetFolderByName(contactId)[0].Id
		}
		//3 Escribrir en Drive el File y guardar idFile de google drive En la tabla Chats
		fileUpload := srv.WriteFile(filePath, mimeType, idFolderParent)
		fileUploadId = fileUpload.Id
		logg.GeneralLogger.Printf("[helpers.SaveInGoogleDrive] - Archivo: %s escrito en Drive", fileUpload.Id)
		// os.Remove(filePath)
	} else {
		idFolderParent, fileUploadId = "", ""
		logg.ErrorLogger.Printf("\033[31m[helpers.SaveInGoogleDrive] - Quedam menos de %d Gb por segurança não colara mais \033[0m\n", limitSpaceDrive)
	}
	fmt.Println("Debug 5.8.4")
	return idFolderParent, fileUploadId
}

// ej: //go helpers.EmptyGoogleDrive()
func EmptyGoogleDrive() {
	srv.CreateGoogleService()
	srv.EmptyGoogleDrive()
}
