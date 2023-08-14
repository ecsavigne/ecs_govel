package repositories

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"

	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"
)

// Repositories System
func (sr *SystemConfigRepositories) AnalizeServiceLog(CompanyPhone string, DateInit string, DateEnd string) []migrations.WhatchDog {
	return models.WhatchDogModel.LoadFromWhatchDog(CompanyPhone, DateInit, DateEnd)
}

// Repositories System
func (sr *SystemConfigRepositories) AvailableStorage() structs.Response {
	cmd := exec.Command("df", "-h")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	stdout := strings.Split(string(out), "\n")[3]
	stdout = strings.Replace(stdout, "      ", "\t", -1)
	stdout = strings.Replace(stdout, "   ", "\t", -1)
	stdout = strings.Replace(stdout, "  ", "\t", -1)
	result := strings.Split(stdout, "\t")
	storage := strings.Replace(result[3], "G", "", -1)

	return structs.Response{
		Status:  true,
		Message: storage,
	}
}

// Repositories System
func (sr *SystemConfigRepositories) AvailableWhatsmeowProcess() structs.Response {
	whatsmeowX := make(map[string]int)
	indexes := []string{}
	indexes = append(indexes, "13371", "13372", "13373", "13374", "13375", "13376", "13377", "13378", "13379", "13310", "13311", "13312", "13313", "13314", "13315", "13316", "13317", "13318", "13319", "13320")
	indexes = append(indexes, "13321", "13322", "13323", "13324", "13325", "13326", "13327", "13328", "13329", "13330", "13331", "13332", "13333", "13334", "13335", "13336", "13337", "13338", "13339", "13340")

	for _, m := range indexes {
		whatsmeowX[m] = 0
	}

	out, err := exec.Command("bash", "-c", "ps -aux | grep main133").Output()
	if err != nil {
		panic(err)
	}

	stdout := strings.Replace(string(out), "       ", "\t", -1)
	stdout = strings.Replace(stdout, "      ", "\t", -1)
	stdout = strings.Replace(stdout, "   ", "\t", -1)
	stdout = strings.Replace(stdout, "  ", "\t", -1)
	stdout = strings.Replace(stdout, " ", "\t", -1)

	rows := strings.Split(stdout, "\n")
	for _, line := range rows {
		cols := strings.Split(line, "\t")
		if len(cols) > 10 {
			main := strings.Replace(cols[10], "/home/josesh/go/src/github.com/app-socialhub-pro/new_whatsmeow/Binary/main", "", -1)
			if sr.inArray(main, indexes) {
				whatsmeowX[main] = 1
			}
		}
	}
	js, _ := json.Marshal(whatsmeowX)
	logMessage := string(js)

	return structs.Response{
		Status:     true,
		Message:    "",
		LogMessage: logMessage,
	}
}

// Repositories System
func (sr *SystemConfigRepositories) inArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Repositories System
func (sr *SystemConfigRepositories) FreeHdSpace() structs.Response {
	path := "/home/josesh/go/src/github.com/app-socialhub-pro/new_whatsmeow/storage/"

	// 1. delete received file from laravel
	os.RemoveAll(path + "api/")
	os.Mkdir(path+"api/", 0777)

	// 2. delete received files from whatsapp
	os.RemoveAll(path + "audio/")
	os.Mkdir(path+"audio/", 0777)
	os.RemoveAll(path + "document/")
	os.Mkdir(path+"document/", 0777)
	os.RemoveAll(path + "image/")
	os.Mkdir(path+"image/", 0777)
	os.RemoveAll(path + "video/")
	os.Mkdir(path+"video/", 0777)
	os.RemoveAll(path + "video/")
	os.Mkdir(path+"video/", 0777)

	// 3. delete *logs files
	os.Remove("/var/logg/whatsappx/*logg.1")
	os.Remove("/var/logg/whatsappx/*logg.2")
	os.Remove("/var/logg/whatsappx/*logg.3")
	os.Remove("/var/logg/whatsappx/*logg.4")
	os.Remove("/var/logg/whatsappx/*logg.5")
	os.Remove("/var/logg/whatsappx/*logg.6")
	os.Remove("/var/logg/whatsappx/*logg.7")
	os.Remove("/var/logg/whatsappx/*logg.8")
	os.Remove("/var/logg/whatsappx/*logg.9")
	os.Remove("/var/logg/whatsappx/*logg.10")

	// 4. delete whatsmeow logs files
	os.Remove("/var/logg/whatsmeow")
	os.Mkdir("/var/logg/whatsmeow", 0777)

	// 5. delete bigger files
	// out, err := exec.Command("bash", "-c", "du -h storage/messageFiles/ | grep G").Output()
	// stdout := strings.Replace(string(out), "       ", "\t", -1)
	// stdout = strings.Replace(stdout, "      ", "\t", -1)
	// stdout = strings.Replace(stdout, "   ", "\t", -1)
	// stdout = strings.Replace(stdout, "  ", "\t", -1)
	// stdout = strings.Replace(stdout, " ", "\t", -1)

	// if err != nil {
	// 	logg.Log(err.Error())
	// }

	// rows := strings.Split(stdout, "\n")
	// for _, line := range rows {
	// 	if line != "" {
	// 		cols := strings.Split(line, "	")
	// 		if cols[1] != "storage/messageFiles/" {
	// 			fmt.Println(cols[1])
	// 		}
	// 	}

	// }

	return structs.Response{
		Status:  true,
		Message: "",
	}
}
