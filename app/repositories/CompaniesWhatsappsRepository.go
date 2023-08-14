package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"new_whatsmeow/app/helpers/logg"
	"new_whatsmeow/app/models"
	"new_whatsmeow/app/structs"
	"new_whatsmeow/database/migrations"
)

// Repositories CompanyWhatsapp
func (cw *CompaniesWhatsappsRepositories) CreateCompanyWhatsapp(Whatsapp string, CompanyId uint64, CompanyWhatsappId uint64) structs.Response {
	CompanyWhatsapp := migrations.CompanyWhatsapps{
		Whatsapp:          Whatsapp,
		CompanyId:         CompanyId,
		CompanyWhatsappId: CompanyWhatsappId,
	}

	models.CompanyWhatsappModel.CreateCompanyWhatsapp(CompanyWhatsapp)

	return structs.Response{
		Status:  true,
		Message: "",
	}
}

// Repositories CompanyWhatsapp
func (cw *CompaniesWhatsappsRepositories) UpdateCompanyWhatsapp(Whatsapp string, CompanyId uint64, CompanyWhatsappId uint64) structs.Response {
	CompanyWhatsapp := migrations.CompanyWhatsapps{
		Whatsapp:          Whatsapp,
		CompanyId:         CompanyId,
		CompanyWhatsappId: CompanyWhatsappId,
	}

	models.CompanyWhatsappModel.UpdateCompanyWhatsapp(CompanyWhatsapp)

	return structs.Response{
		Status:  true,
		Message: "",
	}
}

// Repositories CompanyWhatsapp
func (cw *CompaniesWhatsappsRepositories) DeleteCompanyWhatsapp(Whatsapp string, CompanyId uint64, CompanyWhatsappId uint64) structs.Response {
	CompanyWhatsapp := migrations.CompanyWhatsapps{
		Whatsapp:          Whatsapp,
		CompanyId:         CompanyId,
		CompanyWhatsappId: CompanyWhatsappId,
	}

	models.CompanyWhatsappModel.DeleteCompanyWhatsapp(CompanyWhatsapp)

	return structs.Response{
		Status:  true,
		Message: "",
	}
}

// Repositories CompanyWhatsapp
func (cw *CompaniesWhatsappsRepositories) UpdateWPAccount(OldWhatsapp string, NewWhatsapp string, WhatsappDns string) structs.Response {
	whatsapp := ""
	if OldWhatsapp != "" {
		whatsapp = OldWhatsapp
	} else {
		whatsapp = NewWhatsapp
	}

	WPAccount := migrations.WPAccounts{}
	if NewWhatsapp != "" {
		WPAccount.PhoneNumber = NewWhatsapp
	}
	if WhatsappDns != "" {
		WPAccount.WhatsappDns = WhatsappDns
	}

	models.CompanyWhatsappModel.UpdateWPAccount(WPAccount, whatsapp)

	return structs.Response{
		Status:  true,
		Message: "",
	}
}

// Repositorio CompanyWhatsapp
func (cw *CompaniesWhatsappsRepositories) GetAllCompanyWhatsapp() []migrations.CompanyWhatsapps {
	var cW []migrations.CompanyWhatsapps
	result := models.CompanyWhatsappModel.DB.Find(&cW)
	if result.Error != nil {
		panic("Error reading AllCompanyWhatsapp")
	}
	return cW
}

// Repositories CompanyWhatsapp
func (cw *CompaniesWhatsappsRepositories) Verify(record migrations.CompanyWhatsapps) {
	jsonPayload, _ := json.Marshal(record)
	defer func() {
		if err := recover(); err != nil {
			logMessage := ": Exception in CompaniesWhatsappController-Verify --> " + fmt.Sprintf("%+v", err)
			logg.Log("[CompaniesWhatsappController-Verify] ", "", "", logMessage, true)
			return
		}
	}()

	response, err := http.Post("https://back1.socialhub.pro/api/RPI/verifyWhatsappNumber", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic("Error requesting line 282 - [CompaniesWhatsappController-Verify]...")
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic("Error processing response io.ReadAll(response.Body) line 290 - [CompaniesWhatsappController-Verify] ...")
	} else {
		whatsapp := string(responseBody)
		if whatsapp == "false" {
			logg.GeneralLogger.Println("\033[36m" + record.Whatsapp + "\033[0m")
			// database.Database.DeleteCompanyWhatsapp(whatsapp)
		}
	}
}

// Repositories CompanyWhatsapp
func (cw *CompaniesWhatsappsRepositories) GetCompanyWhatsapp(whatsapp string) (companyWhatsapp migrations.CompanyWhatsapps, e error) {
	models.CompanyWhatsappModel.DB.Where("whatsapp = ?", whatsapp).Order("id DESC").First(&companyWhatsapp)
	if models.CompanyWhatsappModel.DB.Error != nil {
		return companyWhatsapp, models.CompanyWhatsappModel.DB.Error
	}

	return companyWhatsapp, nil
}
