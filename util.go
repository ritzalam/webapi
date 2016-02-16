package main

import "strings"
import "crypto/sha1"
import "fmt"
import "sort"
import "strconv"
import "time"
import "regexp"

func stripChecksum(queryString, checksum string) string {
	if (queryString != "" ) {
	    // handle either checksum as first or middle / end parameter
	    queryString = strings.Replace(queryString, "&checksum=" + checksum, "", -1);
	    queryString = strings.Replace(queryString, "checksum=" + checksum + "&", "", -1);
	    queryString = strings.Replace(queryString, "checksum=" + checksum, "", -1);
	}

	return queryString
}

func calculateChecksum(str string) string {
	checksum := sha1.Sum([]byte(str))
	fmt.Printf("%x\n", checksum)
	s := fmt.Sprintf("%x", checksum)
	return s
}

func sortMap(data map[string]string) map[string]string {
	// To store the keys in slice in sorted order
    var keys []string
    for k := range data {
        keys = append(keys, k)
    }
    sort.Strings(keys)

	sorted := make(map[string]string)

    for _, k := range keys {
        fmt.Println("Key:", k, "Value:", data[k])
        sorted[k] = data[k]
    }

    return sorted
}

func replaceKeyword(message, keyword, newValue string) string {
	return strings.Replace(message, keyword, newValue, -1)
}

func substituteKeywords(message, dialNumber, confNum, meetingName string) string {
	DIAL_NUM := "%%DIALNUM%%"
	CONF_NUM := "%%CONFNUM%%"
	CONF_NAME := "%%CONFNAME%%"

	message = replaceKeyword(message, DIAL_NUM, dialNumber)
	message = replaceKeyword(message, CONF_NUM, confNum)
	message = replaceKeyword(message, CONF_NAME, meetingName)

	return message
}

func generateInternalMeetingId(extMeetingId string) string {
	now := Int64ToString(currentTimeMillis())
	return convertToInternalMeetingId(extMeetingId) + "-" + now
}

func convertToInternalMeetingId(extMeetingId string) string {
	extId := sha1.Sum([]byte(extMeetingId))
	return fmt.Sprintf("%x", extId)
}

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func currentTimeMillis() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func stripMeta(param string) string {
	return strings.Replace(param, "meta_", "", 1)
}

func isValid(meta string) bool {
	match, err := regexp.MatchString("meta_[a-zA-Z][a-zA-Z0-9-]*$", meta)
	if err != nil {
		return false
	}
	return match
}

func processMetaParam(params map[string]string) map[string]string {
    metas := make(map[string]string)
    for key, _ := range params {
        fmt.Println("Key:", key, "Value:", params[key])
        if isValid(key) {
        	// Need to lowercase to maintain backward compatibility with 0.81
    		metaKey := strings.ToLower(stripMeta(key))
    		fmt.Println("metaKey:", metaKey, "Value:", params[key])
    		metas[metaKey] = params[key]
        }
    }

    return metas
}
