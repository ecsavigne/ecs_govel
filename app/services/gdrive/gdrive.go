package gdrive

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"new_whatsmeow/app/helpers/logg"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const (
	ID_FOLDER_STORE = "1SxyJg4uaalYBBIIYKk5BJuUbmHgtRQKD"
	BASE_URL        = "https://drive.google.com/uc?export=view&id="
)

type InfoFileGoogleService struct {
	MimeType    string   `json:"mimeType,omitempty"`
	Id          string   `json:"id,omitempty"`
	NameFile    string   `json:"nameFile,omitempty"`
	CreatedTime string   `json:"createdTime,omitempty"`
	Extension   string   `json:"extension,omitempty"`
	Parents     []string `json:"parents,omitempty"`
}
type File struct {
}
type ServiceGoogle struct {
	GoogleDriveService *drive.Service
}

// func (srv *ServiceGoogle) tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		logg.GeneralLogger.Printf(file + "\n")
// 		logg.GeneralLogger.Printf(os.Args[0])
// 		logg.GeneralLogger.Printf(filepath.Dir(os.Args[0]))
// 		logg.GeneralLogger.Printf("\n\033[31mError opening file of credentials\033[0m\n")
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(tok)
// 	return tok, err
// }

// func (srv *ServiceGoogle) saveToken(path string, token *oauth2.Token) {
// 	logg.GeneralLogger.Printf("\n\033[34mSaving credential file to: %s\033[0m\n", path)
// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		logg.ErrorLogger.Printf("\n\033[31mUnable to cache oauth token: %v\033[0m\n", err)
// 		return
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }

// func (srv *ServiceGoogle) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	logg.GeneralLogger.Printf("\n\033[34mGo to the following link in your browser then type the "+"authorization code: %v\n\033[0m\n", authURL)
// 	var authCode string
// 	if _, err := fmt.Scan(&authCode); err != nil {
// 		logg.ErrorLogger.Printf("\n\033[31mUnable to read authorization code %v\033[31m\n", err)
// 		return nil
// 	}

// 	tok, err := config.Exchange(context.TODO(), authCode)
// 	if err != nil {
// 		logg.ErrorLogger.Printf("\n\033[31mUnable to retrieve token from web %v\033[31m\n", err)
// 		return nil
// 	}
// 	return tok
// }

func (srv *ServiceGoogle) createFolder(name string, parentId string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{"1SxyJg4uaalYBBIIYKk5BJuUbmHgtRQKD"}, //parentId
	}

	file, err := srv.GoogleDriveService.Files.Create(d).Do()

	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mCould not create dir: %s\033[0m\n", err.Error())
		return nil, err
	}

	return file, nil
}

// func (srv *ServiceGoogle) getClient(config *oauth2.Config) *http.Client {
// 	tokFile := "./app/services/gdrive/token.json"
// 	tok, err := srv.tokenFromFile(tokFile)
// 	if err != nil {
// 		tok = srv.getTokenFromWeb(config)
// 		srv.saveToken(tokFile, tok)
// 	}
// 	return config.Client(context.Background(), tok)
// }

func (srv *ServiceGoogle) CreateGoogleService() *ServiceGoogle {
	credentials, err := os.ReadFile("./app/services/gdrive/credentials.json")
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError al leer archivo de credenciales: %v\033[0m\n", err)
		return nil
	}
	//config, err := google.ConfigFromJSON(credentials, drive.DriveScope)
	//config, err := google.JWTConfigFromJSON(credentials, drive.DriveScope)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError al crear configuración de autenticación: %v\033[0m\n", err)
		return nil
	}
	//client := srv.getClient(config)
	//client := config.Client(context.Background())

	driveService, err := drive.NewService(context.Background(), option.WithCredentialsJSON(credentials)) //WithHTTPClient(client)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError al crear cliente de la API de Google Drive: %v\033[0m\n", err)
		return nil
	}

	srv.GoogleDriveService = driveService
	return srv
}

func (srv *ServiceGoogle) CreateGoogleService1() *ServiceGoogle {
	data, err := os.ReadFile("./app/services/gdrive/credentials.json")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	conf, err := google.JWTConfigFromJSON(data, drive.DriveScope)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	// Initiate an http.Client. The following GET request will be
	// authorized and authenticated on the behalf of
	// your service account.
	client := conf.Client(context.Background())
	r, ee := client.Get("https://www.googleapis.com/drive/v3/files")
	body, err := io.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}
	if ee != nil {
		fmt.Println("Ocurrio un error")
	}
	defer r.Body.Close()
	return nil
}

func (srv *ServiceGoogle) CreateFolder(nameFolder string) string {
	file, err := srv.createFolder(nameFolder, ID_FOLDER_STORE)
	if err != nil {
		defer func() {
			logg.ErrorLogger.Printf("\n\033[31mOcurred one error creating Folder: (%s). Function: ServiceGoogle.CreateFolder line:95 Error Description: %s\033[0m\n", nameFolder, err.Error())
		}()
		return ""
	}
	return file.Id
}

func (srv *ServiceGoogle) GetFolderByName(folderName string) []*drive.File {
	fmt.Println(srv)
	query := fmt.Sprintf("name='%s' and mimeType='application/vnd.google-apps.folder'", folderName)
	folders, err := srv.GoogleDriveService.Files.List().Q(query).Fields("files(id)").Do()
	if err != nil {
		logg.GeneralLogger.Printf("\n\033[31mError al buscar la carpeta: %s\n\033[0m", err.Error())
		return nil
	}
	fmt.Println(folders.Files)
	return folders.Files
}

func (srv *ServiceGoogle) GetAllFileFromFolder(folderId string) []string {
	files, err := srv.GoogleDriveService.Files.List().Q(fmt.Sprintf("'%s' in parents and trashed=false", folderId)).Do()
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError to obtain filess of the folder: %v\033[0m\n", err)
		return nil
	}
	logg.ErrorLogger.Printf("\n\033[31mThe folder with id: %s have the filess %v\033[0m\n", folderId, files)
	arrInfoFiles := make([]*InfoFileGoogleService, 0)
	arrId := make([]string, 0)
	for i, elem := range files.Files {
		arrInfoFiles = append(arrInfoFiles, (&InfoFileGoogleService{}).CloneFileDrive(elem))
		arrId = append(arrId, elem.Id)
		logg.ErrorLogger.Printf("\n\033[31mArchivo con Id: %s\n Nombre: %s\n, Extension: %s\nTypeMime:%s\033[0m\n", elem.Id, elem.Name, arrInfoFiles[i].Extension, elem.MimeType)
	}
	return arrId
}

func (srv *ServiceGoogle) SaveGoogleDriveLocal(path string, idGoogleDrv string) error {
	resp, err := srv.GoogleDriveService.Files.Get(idGoogleDrv).Download()
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mfailed to download file %v\033[0m\n", err)
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(path)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mfailed to create file: %v\033[0m\n", err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	logg.GeneralLogger.Println("\n\033[34mCoping\n\033[0m")
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mfailed to save file on disk: %v\033[0m\n", err)
		return err
	}
	logg.GeneralLogger.Println("\n\033[34mCopied\n\033[0m")
	return nil
}

func (srv *ServiceGoogle) ReadFile(idGoogleDrv string) ([]byte, error) {
	resp, err := srv.GoogleDriveService.Files.Get(idGoogleDrv).Download()
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mFailed to read file: %v\033[0m\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mFailed to read file: %v\033[0m\n", err)
		return nil, err
	}

	return data, nil
}

func (srv *ServiceGoogle) WriteFile(pathFile, typeMime, parentId string) *drive.File {
	tempArray := strings.Split(pathFile, "/")
	tempArr2 := strings.Split(tempArray[len(tempArray)-1], ".")
	if tempArr2[len(tempArr2)-1] == "png" && typeMime == "text/plain; charset=utf-8" {
		typeMime = "image/png"
	}

	fileDrive := &drive.File{
		Name:     tempArray[len(tempArray)-1],
		Parents:  []string{parentId},
		MimeType: typeMime,
	}

	file, err := os.Open(pathFile)
	logg.GeneralLogger.Printf("\n\033[34mPATH DE FILE LOCAL: %s parentId: %s \n\033[0m\n", pathFile, parentId)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError opening file to the up: %v\033[0m\n", err)
		return nil
	}
	defer file.Close()
	servFil, err1 := srv.GoogleDriveService.Files.Create(fileDrive).Media(file).Do()
	if err1 != nil {

		logg.ErrorLogger.Printf("\n\033[31mError opening file to the up Google Drive: %v\n\033[0m\n", err1)
		return nil
	}
	//os.Remove(pathFile)
	return servFil
}

func (srv *ServiceGoogle) DeleteFile(fileId string) {
	logg.GeneralLogger.Printf("\n\033[34mFILE:%s ID:%s delete satifactory of the Google Drive \n\033[0m\n", srv.GetNameFile(fileId), fileId)
	err := srv.GoogleDriveService.Files.Delete(fileId).Do()
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mNo se pudo eliminar el archivo: %v\n\033[0m\n", err)
		return
	}
}

func (srv *ServiceGoogle) DeleteListFile(listFile []string) {
	for _, value := range listFile {
		srv.DeleteFile(value)
	}
}

func (srv *ServiceGoogle) DeleteFilesOfFolder(folderId string) {
	srv.DeleteListFile(srv.GetAllFileFromFolder(folderId))
}

func (srv *ServiceGoogle) EmptyGoogleDrive() {
	files, err := srv.GoogleDriveService.Files.List().Fields("nextPageToken, files(id)").Do()
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mErro ao deletar arquivo paso Obtener list files Error:\033[0m%v\n", err)
		return
	}
	for _, file := range files.Files {
		err := srv.GoogleDriveService.Files.Delete(file.Id).Do()
		if err != nil {
			logg.ErrorLogger.Printf("\n\033[31mErro ao deletar arquivo paso apagar arquivo Error:\n\033[0m%v", err)
			return
		}
	}
	logg.GeneralLogger.Println("\n\033[33mApagados tudos os arquivos de google Drive sem erro\n\033[0m")
}

func (srv *ServiceGoogle) GetNameFile(fileId string) string {
	file, err := srv.GoogleDriveService.Files.Get(fileId).Fields("name").Do()
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mDo not get file: %v\n\033[0m\n", err)
		return ""
	}

	return file.Name
}

func (srv *ServiceGoogle) UploadAllFilesOfDir(pathFolderRoot, idFolderDrive string) {
	if idFolderDrive == "" {
		idFolderDrive = "1obsfO-1SxyJg4uaalYBBIIYKk5BJuUbmHgtRQKD"
	}
	files, err := os.ReadDir(pathFolderRoot)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError al leer la carpeta: %v\n\033[0m\n", err)
		return
	}
	cont := 0
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
			typeMime := (&File{}).GetMimeType(pathFolderRoot + "/" + file.Name())
			logg.GeneralLogger.Printf("\n\033[31mMimeType: %s\n\033[0m\n", typeMime)
			srv.WriteFile(pathFolderRoot+"/"+file.Name(), typeMime, idFolderDrive)
			fmt.Println("")
			cont++
		}
	}
	logg.GeneralLogger.Printf("\n\033[34mArchivos subidos: %d\n\033[0m\n", cont)
}

func (f *InfoFileGoogleService) CloneFileDrive(fileGoogle *drive.File) *InfoFileGoogleService {
	f.MimeType = fileGoogle.MimeType
	f.Id = fileGoogle.Id
	f.NameFile = fileGoogle.Name
	f.CreatedTime = fileGoogle.CreatedTime
	f.Extension = strings.Split(fileGoogle.MimeType, "/")[1]
	f.Parents = fileGoogle.Parents
	return f
}

func (f *File) GetMimeType(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mOcurred one error to open File: %s Error: %s\033[0m\n", filePath, err.Error())
		return ""
	}
	defer file.Close()
	fileHeader := make([]byte, 512)
	_, err = file.Read(fileHeader)
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError to read of File: %s Error: %s\033[0m\n", filePath, err.Error())
		return ""
	}
	return http.DetectContentType(fileHeader)
}

func (srv *ServiceGoogle) GetStorageSpace() (int64, int64, int64) {
	resp, err := srv.GoogleDriveService.About.Get().Fields("storageQuota").Do()
	if err != nil {
		logg.ErrorLogger.Printf("\n\033[31mError al obtener los detalles de la cuota: %v\033[0m\n", err)
		return 0, 0, 0
	}
	// Total en bytes
	totalSpace := resp.StorageQuota.Limit
	usedSpace := resp.StorageQuota.Usage
	remainingSpace := totalSpace - usedSpace
	return totalSpace / (1024 * 1024 * 1024), usedSpace / (1024 * 1024 * 1024), remainingSpace / (1024 * 1024 * 1024)
}
