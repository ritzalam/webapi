package main

import (
	"fmt"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	"time"
	"unicode/utf8"
)

var clogger = loggo.GetLogger("client")

type Consolidator struct {
	conference string
	user       string
	client     string
	message    string
}

// DefaultFormatter provides a simple concatenation of all the components.
type ClientLogFormatter struct{}

// Format returns the parameters separated by spaces except for filename and
// line which are separated by a colon.  The timestamp is shown to second
// resolution in UTC.
func (*ClientLogFormatter) Format(level loggo.Level, module, filename string, line int, timestamp time.Time, message string) string {
	//ts := time.Now().Format(time.RFC3339)
	ts := time.Now().Format("2006-01-02T15:04:05.999999Z07:00")

	return fmt.Sprintf("%s %s", ts, message)
}

func WriteToLogFile(client string, meeting string, userid string, message string, remoteip string) (int64, error) {
	var err error
	var filesize int64 = -1

	data := meeting + " " + userid + " " + message
	//fmt.Println(data)
	clogger.Infof(data)

	return filesize, err
}

func CheckVars(client string, message string, meetingid string, userid string) bool {

	valid := (len(message) != 0)
	clientallowed := false

	// Check if the sent client name is allowed to write to logs
	for _, element := range viper.GetStringSlice("AllowedClients") {
		clientallowed = clientallowed || (element == client)
	}

	if !clientallowed {
		return false
	}

	values := [4]string{client, message, meetingid, userid}

	for _, element := range values {
		valid = valid && utf8.ValidString(element)
	}

	return valid
}
