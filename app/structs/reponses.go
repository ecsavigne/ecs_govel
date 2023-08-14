package structs

type Response struct {
	Status           bool   `json:"status"`
	Message          string `json:"message"`
	LogMessage       string `json:"log_message"`
	Code             int    `json:"code"`
	CodeMessage      string `json:"code_message"`
	ErrorCode        int    `json:"error_code"`
	ErrorCodeMessage string `json:"error_code_message"`
	Qrcodebase64     string `json:"qrcodebase64"`
	Jid              string `json:"jid"`
	PicURL           string `json:"picurl"`
	Name             string `json:"name"`
}

type ChatsApi struct {
	Id                 int    `json:"id"`
	Source             int    `json:"source"`
	Message            string `json:"message"`
	MessageId          string `json:"message_id"`
	CompanyPhone       string `json:"phonenumber"`
	ContactPhone       string `json:"contactphone"`
	TypeId             int    `json:"type_id"`
	Path               string `json:"path"`
	ClientOriginalName string `json:"client_original_name"`
	CreatedAt          string `json:"created_at"`
}

type Groups struct {
	JID  string `json:"jid"`
	Name string `json:"name"`
}
