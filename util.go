package main

import "strings"
import "crypto/sha1"
import "fmt"

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