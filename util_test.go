package main

import (
	"testing"
)


func TestStripChecksum(t *testing.T) {
	checksum := "52998795582ce05ed3ee2ee958cf627d3726d50c"
	result := "meetingID=random-4450270"
	queryString := result + "&checksum=" + checksum
	qstr := stripChecksum(queryString, checksum)
	if qstr != result {
		t.Error("Expected 16, got ", qstr)
	}
}


func TestCalculateChecksum(t *testing.T) {
	checksum := "d68c19a0a345b7eab78d5e11e991c026ec60db63"
	queryString := "abcdefghij"
	result := calculateChecksum(queryString)
	if result != checksum {
		t.Error("Expected 16, got ", result)
	}
}
