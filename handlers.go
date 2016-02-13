package main

import (
	"fmt"
	"encoding/json"
	"encoding/xml"
	"net/http"
)


func CreateApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func JoinApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func IsMeetingRunningApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func GetMeetingInfoApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func EndApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func GetMeetingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func GetDefaultConfigXMLApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func SetConfigXMLApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func EnterApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func ConfigXMLApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func SignOutApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func GetRecordingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func PublishRecordingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func DeleteRecordingsApiHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)
	apiresponse.Status = "ok"
	MarshallResult(apiresponse, request.URL.Query().Get("marshall"), writer, request, err)
}

func MarshallResult(apiresponse APIResponse, marshall string, writer http.ResponseWriter, request *http.Request, err error) {
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
