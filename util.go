package main

import "strings"
import "crypto/sha1"
import "fmt"
import "sort"

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

    // To perform the opertion you want
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