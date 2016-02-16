package main

import (
	"testing"
	"fmt"
	"sort"
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

func TestSortMap(t *testing.T) {
	elements := make(map[string]string)
  	elements["H"] = "Hydrogen"
  	elements["He"] = "Helium"
  	elements["Li"] = "Lithium"
  	elements["Be"] = "Beryllium"
  	elements["B"] = "Boron"
  	elements["C"] = "Carbon"
  	elements["N"] = "Nitrogen"
  	elements["O"] = "Oxygen"
  	elements["F"] = "Fluorine"
  	elements["Ne"] = "Neon"

    var keys []string
    for k := range elements {
    	fmt.Println("Key:", k)
        keys = append(keys, k)
    }
    sort.Strings(keys)

	sorted := make(map[string]string)

    for _, k := range keys {
        fmt.Println("Key:", k, "Value:", elements[k])
        sorted[k] = elements[k]
    }

//  	data := sortMap(elements)
//	var keys [10]string
//    for k := range data {
//        keys = append(keys, k)
//    }

//  	result := [10]string {"B", "Be", "C", "F", "H", "He", "Li", "N", "Ne", "O"}

//	if keys != result {
//		t.Error("Expected 16, got ", keys)
//	}
}

func TestSubstituteKeywords(t *testing.T) {
	dialNum := "6135551234"
	confNum := "85115"
	confName := "Demo Meeting"
	welcome := "<br>Welcome to <b>%%CONFNAME%%</b>! DialNum=%%DIALNUM%% ConfNum=%%CONFNUM%%"
	result := "<br>Welcome to <b>" + confName + "</b>! DialNum=" + dialNum + " ConfNum=" + confNum

	newWelcome := substituteKeywords(welcome, dialNum, confNum, confName)
	fmt.Println(newWelcome)
	if result != newWelcome {
		t.Error("Expected 16, got ", result)
	}
}

// "685ccde7fb1b853825eee3cf73c5d6205d3a0191-1455569115186"
func TestConvertToInternalMeetingId(t *testing.T) {
	extId := "685ccde7fb1b853825eee3cf73c5d6205d3a0191"
	result := convertToInternalMeetingId("random-7603468")
	fmt.Println(result)
	if result != extId {
		t.Error("Expected 16, got ", result)
	}	
}

func TestInt64ToString(t *testing.T) {
	v := int64(1257894000000)
	now := Int64ToString(v)
	result := "1257894000000"
	fmt.Println(now)
	if result != now {
		t.Error("Expected 16, got ", now)
	}	
}

func TestProcessMetaParam(t *testing.T) {
	metas := make(map[string]string)
  	metas["meta_H"] = "Hydrogen"
  	metas["meta_He"] = "Helium"
  	metas["meta_Li"] = "Lithium"
  	metas["meta_Be"] = "Beryllium"
  	metas["meta_B"] = "Boron"
  	metas["meta_C"] = "Carbon"
  	metas["meta_N"] = "Nitrogen"
  	metas["meta_O"] = "Oxygen"
  	metas["meta_F"] = "Fluorine"
  	metas["meta_Ne"] = "Neon"

  	result := processMetaParam(metas)

    for key, value := range result {
        fmt.Println("rKey:", key, "Value:", value)
    }
}