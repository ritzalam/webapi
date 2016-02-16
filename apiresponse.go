package main

import (
	"encoding/xml"
)

type Response struct {
	XMLName          xml.Name       `json:"-"     xml:"response"`
	Status           string 	    `json:"status" xml:"status"`
	Error            string         `json:"error" xml:"error"`
}

type ApiMeetingResponse struct {
	XMLName          xml.Name       `json:"-"     xml:"response"`
	ReturnCode       string 	    `json:"returncode" xml:"returncode"`
	MeetingId        string 	    `json:"meetingID" xml:"meetingID"`
	AttendeePW       string         `json:"attendeePW" xml:"attendeePW"`
	ModeratorPW      string         `json:"moderatorPW" xml:"moderatorPW"`
	CreateTime       string         `json:"createTime" xml:"createTime"`
	VoiceBridge      string         `json:"voiceBridge" xml:"voiceBridge"`
	DialNumber       string         `json:"dialNumber" xml:"dialNumber"`
	CreateDate       string         `json:"createDate" xml:"createDate"`
	HasUserJoined    string         `json:"hasUserJoined" xml:"hasUserJoined"`
	Duration         string         `json:"duration" xml:"duration"`
	HasBeenForciblyEnded string     `json:"hasBeenForciblyEnded" xml:"hasBeenForciblyEnded"`
	MessageKey       string         `json:"messageKey" xml:"messageKey"`
	Message          string         `json:"message" xml:"message"`
}

type ApiErrorResponse struct {
	XMLName          xml.Name       `json:"-"     xml:"response"`
	ReturnCode       string 	    `json:"returncode" xml:"returncode"`
	Errors           []ApiError     `json:"errors" xml:"errors"`
}

type ApiError struct {
	XMLName          xml.Name       `json:"-"     xml:"-"`
	Key              string         `json:"key" xml:"key"`
	Message          string         `json:"error" xml:"error"`	
}

type ApiInvalidResponse struct {
	XMLName         xml.Name    `json:"-" xml:"response"`
	ReturnCode      string 	    `json:"returncode" xml:"returncode"`	 
	MessageKey      string      `json:"messageKey" xml:"messageKey"`
	Message         string      `json:"message" xml:"message"`
}
