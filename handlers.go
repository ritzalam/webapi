package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

func LogHandler(writer http.ResponseWriter, request *http.Request) {
	var apiresponse APIResponse
	var err error

	// Route URL
	vars := mux.Vars(request)
	client := vars["client"]
	meetingid := vars["meetingid"]
	userid := vars["userid"]

	// POST variables
	message := request.FormValue("message")
	message = strings.TrimSpace(message)

	// Check variable validity
	if CheckVars(client, message, meetingid, userid) {
		_, err = WriteToLogFile(client, meetingid, userid, message, request.RemoteAddr)
		apiresponse.Error = "nil"
		apiresponse.Status = "ok"
	} else {
		apiresponse.Status = "error"
	}

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
