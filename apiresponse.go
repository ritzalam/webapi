package main

import (
	"encoding/xml"
)

type APIResponse struct {
	XMLName xml.Name `json:"-"     xml:"response"`
	Status  string   `json:"status" xml:"status"`
	Error   string   `json:"error" xml:"error"`
}
