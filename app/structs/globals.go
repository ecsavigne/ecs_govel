package structs

import (
	"go.mau.fi/whatsmeow"
)

type ClassClient struct {
	WAClient     *whatsmeow.Client
	CompanyPhone string
}

// var Connections map[string]*whatsmeow.Client
var Connections = make(map[string]*whatsmeow.Client, 50) //memory allocation to allocate 50 clients
var Dns = make(map[string]string, 50)                    //memory allocation to allocate 50 clients
