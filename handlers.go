package main

import (
	"fmt"
	"net/http"
	"github.com/jchannon/negotiator"
	"encoding/xml"
	"encoding/json"
)

func CreateApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)

	params := request.URL.Query()
	for k := range params {
        fmt.Println("Key:", k, "Value:", params[k])
    }

	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	//negotiator.Negotiate(writer, request, apiresponse)
	MarshallResult(apiresponse, "json", writer, request, nil)
}

func JoinApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)

	
	var apiresponse InvalidResponse
	apiresponse.ReturnCode = "ok"
	apiresponse.MessageKey = "message key"
	apiresponse.Message = "message message"
	negotiator.Negotiate(writer, request, apiresponse)
	//MarshallResult(apiresponse, "xml", writer, request, nil)
}

func IsMeetingRunningApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func GetMeetingInfoApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func EndApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func GetMeetingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func GetDefaultConfigXMLApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func SetConfigXMLApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func EnterApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func ConfigXMLApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func SignOutApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func GetRecordingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func PublishRecordingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func DeleteRecordingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiresponse Response
	apiresponse.Status = "ok"
	apiresponse.Error = ""
	negotiator.Negotiate(writer, request, apiresponse)
}

func MarshallResult(apiresponse Response, marshall string, writer http.ResponseWriter, request *http.Request, err error) {
	// Marshall the result and return it
	var result []byte

	switch marshall {
	case "", "json":
		result, err = json.Marshal(apiresponse)
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	case "xml":
		result, err = xml.MarshalIndent(apiresponse, "", "  ")
		result = []byte(xml.Header + string(result))
		writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Write(result)
}

type Response struct {
	XMLName xml.Name `json:"-"     xml:"response"`
	Status  string   `json:"status" xml:"status"`
	Error   string   `json:"error" xml:"error"`
}

type ErrorResponse struct {
	XMLName xml.Name `json:"-"     xml:"response"`
	Status  string   `json:"status" xml:"status"`
	Error   string   `json:"error" xml:"error"`
}

type InvalidResponse struct {
	XMLName 		xml.Name 	`json:"-" xml:"response"`
	ReturnCode 		string 		`json:"returncode" xml:"returncode"`	 
	MessageKey  	string   	`json:"status" xml:"status"`
	Message   		string   	`json:"error" xml:"error"`
}
