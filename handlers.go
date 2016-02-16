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

	var apiResponse ApiMeetingResponse 
	apiResponse.ReturnCode = "ok" 
	apiResponse.MeetingId  = "meeting_id"    
	apiResponse.AttendeePW  = "ap"     
	apiResponse.ModeratorPW  = "mp"   
	apiResponse.CreateTime   = "create time"    
	apiResponse.VoiceBridge  = "voice bridge"    
	apiResponse.DialNumber   = "dial number"    
	apiResponse.CreateDate   = "create date"   
	apiResponse.HasUserJoined = "has user joined"   
	apiResponse.Duration      = "duration"   
	apiResponse.HasBeenForciblyEnded = "has been forcibly ended"
	apiResponse.MessageKey    = "message key"   
	apiResponse.Message       = "message"   

	//negotiator.Negotiate(writer, request, apiResponse)
	MarshallResult(apiResponse, "xml", writer, request, nil)
}

func JoinApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)

	
	var apiresponse ApiInvalidResponse
	apiresponse.ReturnCode = "ok"
	apiresponse.MessageKey = "message key"
	apiresponse.Message = "message message"
	negotiator.Negotiate(writer, request, apiresponse)
	//MarshallResult(apiresponse, "xml", writer, request, nil)
}

func IsMeetingRunningApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var apiResponse ApiInvalidResponse
	apiResponse.ReturnCode = "ok"	 
	apiResponse.MessageKey = "invalid param"
	apiResponse.Message    = "Invalid parameter"  
	negotiator.Negotiate(writer, request, apiResponse)
}

func GetMeetingInfoApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	var error1 ApiError 
	error1.Key = "error_1"
	error1.Message = "error 1 message"
	var error2 ApiError 
	error2.Key = "error_2"
	error2.Message = "error 2 message"

	errors := []ApiError{error1, error2}
	var apiResponse ApiErrorResponse
	apiResponse.ReturnCode = "ok"
	apiResponse.Errors = errors
	negotiator.Negotiate(writer, request, apiResponse)
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
	var apiResponse Response
	apiResponse.Status = "ok"
	apiResponse.Error = ""
	negotiator.Negotiate(writer, request, apiResponse)
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

func MarshallResult(apiresponse ApiMeetingResponse, marshall string, writer http.ResponseWriter, request *http.Request, err error) {
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


