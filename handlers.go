package main

import (
	"fmt"
	"net/http"
	"github.com/jchannon/negotiator"
	"encoding/xml"
	"encoding/json"
	"strings"
)

const RESP_CODE_SUCCESS = "SUCCESS"
const RESP_CODE_FAILED = "FAILED"

func CreateApiHandler(writer http.ResponseWriter, request *http.Request) {
	queryStr := request.URL.RawQuery
	fmt.Println("Query string: " + queryStr)

	params := request.URL.Query()
	for k := range params {
        fmt.Println("Key:", k, "Value:", params[k])
    }

    // BEGIN - backward compatibility
    if _, ok := params["checksum"]; !ok {
		invalid("checksumError", "You did not pass the checksum security check", writer, request)
		return
	}

    if _, ok := params["meetingID"]; ok {
    	fmt.Println(params["meetingID"])
    	meetingId := params["meetingID"]
    	trimmedMeetingId := strings.TrimSpace(meetingId[0])
    	if trimmedMeetingId == "" {
 			invalid("missingParamMeetingID", "Please do not pass a space only as meetingID.", writer, request)
			return   		
		} else {
			fmt.Println("trimmedMeetingId=" + trimmedMeetingId)
		}
	} else {
 		invalid("missingParamMeetingID", "You must specify a meeting ID for the meeting.", writer, request)
		return  		
	}
    // END - backward compatibility
/*
    if (! paramsProcessorUtil.isChecksumSame(API_CALL, params.checksum, request.getQueryString())) {
      invalid("checksumError", "You did not pass the checksum security check")
      return
    }
    // END - backward compatibility
	
    ApiErrors errors = new ApiErrors();
    paramsProcessorUtil.processRequiredCreateParams(params, errors);

    if (errors.hasErrors()) {
      respondWithErrors(errors)
      return
    }
            
    // Do we agree with the checksum? If not, complain.
    if (! paramsProcessorUtil.isChecksumSame(API_CALL, params.checksum, request.getQueryString())) {
      errors.checksumError()
      respondWithErrors(errors)
      return
    }
      
    // Translate the external meeting id into an internal meeting id.
    String internalMeetingId = paramsProcessorUtil.convertToInternalMeetingId(params.meetingID);		
    Meeting existing = meetingService.getNotEndedMeetingWithId(internalMeetingId);

    if (existing != null) {
      log.debug "Existing conference found"
      Map<String, Object> updateParams = paramsProcessorUtil.processUpdateCreateParams(params);
      if (existing.getViewerPassword().equals(params.get("attendeePW")) && existing.getModeratorPassword().equals(params.get("moderatorPW"))) {
        paramsProcessorUtil.updateMeeting(updateParams, existing);
        // trying to create a conference a second time, return success, but give extra info
        // Ignore pre-uploaded presentations. We only allow uploading of presentation once.
        //uploadDocuments(existing);
        respondWithConference(existing, "duplicateWarning", "This conference was already in existence and may currently be in progress.");
      } else {
          // BEGIN - backward compatibility
          invalid("idNotUnique", "A meeting already exists with that meeting ID.  Please use a different meeting ID.");
          return;
          // END - backward compatibility

          // enforce meetingID unique-ness
          errors.nonUniqueMeetingIdError()
          respondWithErrors(errors)
      } 
      
      return;    
    }
     
    Meeting newMeeting = paramsProcessorUtil.processCreateParams(params);      

    meetingService.createMeeting(newMeeting);
    
    // See if the request came with pre-uploading of presentation.
    uploadDocuments(newMeeting);    
    respondWithConference(newMeeting, null, null)
*/
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

	negotiator.Negotiate(writer, request, apiResponse)
	//MarshallResult(apiResponse, "xml", writer, request, nil)
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

func invalid(key, msg string, writer http.ResponseWriter, request *http.Request) {
	var apiResponse ApiInvalidResponse
	apiResponse.ReturnCode = RESP_CODE_FAILED	 
	apiResponse.MessageKey = key
	apiResponse.Message    = msg
	negotiator.Negotiate(writer, request, apiResponse)
}