package main

import (
	"log"
	"strings"
	"time"
)

func convTime(strTime string) time.Time {

	timePattern := "2006-01-02T15:04:05"
	convTime, err := time.Parse(timePattern, strTime)
	if err != nil {
		log.Println(strTime, err)
	}

	return convTime
}

func createList(strList string) []string {

	var lst []string
	for _, v := range strings.Split(strList[1:len(strList)-1], ",") {
		lst = append(lst, strings.Trim(v, " \""))
	}

	return lst
}
