package pkggin

// Tipos pra Postman import
type Variable struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	Disabled *bool  `json:"disabled,omitempty"`
}

type Mode string

var modeDefault = Mode("formdata")

type Body struct {
	Mode     Mode       `json:"mode"`
	Formdata []Variable `json:"formdata"`
}

// type ParamRequest struct {
// 	Path string `json:"path"`
// 	Body `json:"body,omitempty"`
// }

type URL struct {
	Raw  string   `json:"raw"`
	Host []string `json:"host"`
	Path []string `json:"path"`
}

type Request struct {
	Method string `json:"method"`
	URL    `json:"url"`
	Body   *Body `json:"body,omitempty"`
}

type Info struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}
type PostmanItem struct {
	Name    string `json:"name"`
	Request `json:"request"`
}

type ColletionFolder struct {
	Name string        `json:"name"`
	Item []PostmanItem `json:"item,omitempty"`
}

type PostmanCollection struct {
	Info `json:"info"`
	Item []ColletionFolder `json:"item"`
}
