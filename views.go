package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"labix.org/v2/mgo/bson"
)

func handleChannel(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		name := req.FormValue("name")
		channel := Channel{name}

		storeChannel(channel)

		resJSON, err := json.Marshal(channel)
		if err != nil {
			log.Println("Error in marshaling", err.Error())
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusCreated)
		fmt.Fprint(res, string(resJSON))
	}
}

func handlePerformer(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		name := req.FormValue("name")
		performer := Performer{Name: name}

		storePerformer(performer)

		resJSON, err := json.Marshal(performer)
		if err != nil {
			log.Println("Error in marshaling", err.Error())
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusCreated)
		fmt.Fprint(res, string(resJSON))
	}
}

func handleSong(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		performer := req.FormValue("performer")
		title := req.FormValue("title")

		song := Song{Performer: performer, Title: title}

		storeSong(song)

		resJSON, err := json.Marshal(song)
		if err != nil {
			log.Println("Error in marshaling", err.Error())
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusCreated)
		fmt.Fprint(res, string(resJSON))
	}
}

func handlePlay(res http.ResponseWriter, req *http.Request) {
	log.Println("Entering handleplay")

	switch req.Method {
	case "POST":
		channel := req.FormValue("channel")
		performer := req.FormValue("performer")
		title := req.FormValue("title")
		sstart_time := req.FormValue("start")
		send_time := req.FormValue("end")
		time_layout := "2006-01-02T15:04:05"
		start_time, err := time.Parse(time_layout, sstart_time)
		if err != nil {
			log.Println(sstart_time, err)
			return
		}

		end_time, err := time.Parse(time_layout, send_time)
		if err != nil {
			log.Println(err)
			return
		}

		play := Play{
			Channel:   channel,
			Performer: performer,
			Title:     title,
			Start:     start_time,
			End:       end_time,
		}

		storePlay(play)
		resJSON, err := json.Marshal(play)
		if err != nil {
			log.Println("Error in marshaling", err.Error())
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusCreated)
		fmt.Fprint(res, string(resJSON))
	}
}

func handleChannelPlays(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// form fields
	channel := req.FormValue("channel")
	startTime := convTime(req.FormValue("start"))
	endTime := convTime(req.FormValue("end"))

	coll := mgoSess.DB("bmat-test").C("plays")

	type Result struct {
		Performer string    `json:"performer"`
		Title     string    `json:"title"`
		Start     time.Time `json:"start"`
		End       time.Time `json:"end"`
	}
	var result []Result

	err := coll.Find(bson.M{
		"channel": channel,
		"start": bson.M{
			"$gt": startTime,
			"$lt": endTime,
		},
	}).All(&result)

	if err != nil {
		log.Println("Error on query: ", err)
	}

	type Response struct {
		Code   int      `json:"code"`
		Result []Result `json:"result"`
	}

	response := &Response{
		Code:   0,
		Result: result,
	}
	resJSON, _ := json.Marshal(response)

	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(resJSON))
}

func handleTopPlays(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// form fields
	channels := createList(req.FormValue("channels"))
	//limit := req.FormValue("limit")
	pivotTime := convTime(req.FormValue("start"))

	//startTime := pivotTime.AddDate(0, 0, -7)
	endTime := pivotTime.AddDate(0, 0, 6)

	coll := mgoSess.DB("bmat-test").C("plays")

	type Result struct {
		Title     string `json:"title"`
		Performer string `json:"performer"`
		Plays     int    `json:"plays"`
		Rank      int    `json:"rank"`
	}
	var result []Result

	err := coll.Find(bson.M{
		"channel": bson.M{
			"$in": channels,
		},
		"start": bson.M{
			"$gt": pivotTime,
			"$lt": endTime,
		},
	}).All(&result)

	if err != nil {
		log.Println("Error on query: ", err)
	}

	log.Println(result)

	//type Response struct {
	//Code   int      `json:"code"`
	//Result []Result `json:"result"`
	//}

	//response := &Response{
	//Code:   0,
	//Result: result,
	//}
	//resJSON, _ := json.Marshal(response)

	//res.WriteHeader(http.StatusCreated)
	//fmt.Fprint(res, string(resJSON))
}
func handleSongPlays(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// form fields
	title := req.FormValue("title")
	performer := req.FormValue("performer")
	startTime := convTime(req.FormValue("start"))
	endTime := convTime(req.FormValue("end"))

	coll := mgoSess.DB("bmat-test").C("plays")

	type Result struct {
		Channel string    `json:"channel"`
		Start   time.Time `json:"start"`
		End     time.Time `json:"end"`
	}
	var result []Result

	err := coll.Find(bson.M{
		"title":     title,
		"performer": performer,
		"start": bson.M{
			"$gt": startTime,
			"$lt": endTime,
		},
	}).All(&result)

	if err != nil {
		log.Println("Error on query: ", err)
	}

	type Response struct {
		Code   int      `json:"code"`
		Result []Result `json:"result"`
	}

	response := &Response{
		Code:   0,
		Result: result,
	}
	resJSON, _ := json.Marshal(response)

	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(resJSON))
}
