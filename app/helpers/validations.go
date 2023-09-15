package helpers

import (
	"app/models"
	"strconv"
)

func ValidateApiHash(ApiHash string) {
	if ApiHash == "" {
		panic("Missing apihash in request")
	}
	if !models.ApplicationModel.AppIsAuth(ApiHash) {
		panic("Unauthorized apihash")
	}
}

func ValidateContactPhone(ContactPhone string) {
	if ContactPhone == "" {
		panic("Missing RemoteJid in request")
	}
}

func ValidateCompanyPhone(CompanyPhone string) {
	if CompanyPhone == "" {
		panic("Missing CompanyPhone in request")
	}
}

func ValidatePerPage(PerPage string) {
	if PerPage == "" {
		panic("Missing PerPage in request")
	}
}

func ValidatePage(Page string) {
	if Page == "" {
		panic("Missing Page in request")
	}
}

func ValidateAction(Page string) {
	if Page == "" {
		panic("Missing Action in request")
	}
}

func ValidateInitDate(InitDate string) {
	if InitDate == "" {
		panic("Missing InitDate in request")
	}
}

func ValidateEndDate(EndDate string) {
	if EndDate == "" {
		panic("Missing EndDate in request")
	}
}

func ValidateSource(Source string) int {
	var source int
	if Source == "" {
		source = 0 // panic("Missing Source in request")
	}
	source, _ = strconv.Atoi(Source)
	return source
}

func ValidateUserId(UserId string) uint64 {
	var val uint64
	if UserId == "" {
		val = 0 // panic("Missing UserId in request")
	}
	val, _ = strconv.ParseUint(UserId, 10, 64)
	return val
}

func ValidateCompanyId(CompanyId string) uint64 {
	var val uint64
	if CompanyId == "" {
		val = 0 // panic("Missing CompanyId in request")
	}
	val, _ = strconv.ParseUint(CompanyId, 10, 64)
	return val
}

func ValidateCompanyWhatsappId(CompanyWhatsappId string) uint64 {
	if CompanyWhatsappId == "" {
		panic("Missing CompanyWhatsappId in request")
	}
	companyWhatsappId, err := strconv.ParseUint(CompanyWhatsappId, 10, 64)
	if err != nil {
		panic("Invalid CompanyWhatsappId: " + CompanyWhatsappId)
	}
	return companyWhatsappId
}

func ValidateCompanyIdAndCompanyWhatsappId(CompanyId string, CompanyWhatsappId string, CompanyPhone string) (uint64, uint64) {
	var companyId, companyWhatsappId uint64
	if CompanyId == "" || CompanyWhatsappId == "" {
		company, err := models.CompanyWhatsappModel.GetCompanyWhatsapp(CompanyPhone)
		if err != nil {
			panic("Client " + CompanyPhone + ": CompanyWhatsap " + CompanyPhone + " does not exist in database. Error is: " + err.Error() + "")
		} else {
			companyId = company.CompanyId
			companyWhatsappId = company.CompanyWhatsappId
		}
	} else {
		companyId, _ = strconv.ParseUint(CompanyId, 10, 64)
		companyWhatsappId, _ = strconv.ParseUint(CompanyWhatsappId, 10, 64)
	}
	return companyId, companyWhatsappId
}

func ValidatePath(Path string) {
	if Path == "" {
		panic("Missing Path in request")
	}
}

func ValidateWebHook(WebHook string) {
	if WebHook == "" {
		panic("Missing WebHook in request")
	}
}

func ValidateMessageId(MsgId string) {
	if MsgId == "" {
		panic("Missing MessageId in request")
	}
}

func ValidateFirstMessageId(FirstMessageId string) {
	if FirstMessageId == "" {
		panic("Missing FirstMessageId in request")
	}
}

func ValidateWhatsapp(Whatsapp string) {
	if Whatsapp == "" {
		panic("Missing Whatsapp in request")
	}
}

func ValidateWhatsappDns(WhatsappDns string) {
	if WhatsappDns == "" {
		panic("Missing WhatsappDns in request")
	}
}
func ValidateStatus(Status string) {
	if Status == "" {
		panic("Missing Status in request")
	}
}

func ValidateFilneame(Filneame string) {
	if Filneame == "" {
		panic("Missing Filneame in request")
	}
}
