package main

import (
	"log"
	"time"
)

type Channel struct {
	Name string `bson:"channel"`
}

type Performer struct {
	Name string `bson:"performer"`
}

type Song struct {
	Performer string `bson:"performer"`
	Title     string `bson:"title"`
	Length    int32  `bson:"length"`
}

type Play struct {
	Channel   string    `json:"channel"`
	Performer string    `json:"performer"`
	Title     string    `json:"title"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

func storeChannel(channel Channel) {
	coll := mgoSess.DB("bmat-test").C("channels")
	err := coll.Insert(channel)
	if err != nil {
		// Duplicate index, go on
		log.Println("Trying to insert duplicate channel")
	} else {
		log.Println("Adding channel: ", channel.Name)
	}
}

func storePerformer(performer Performer) {
	coll := mgoSess.DB("bmat-test").C("performers")
	err := coll.Insert(performer)
	if err != nil {
		// Duplicate index, go on
		log.Println("Trying to insert duplicate performer")
	} else {
		log.Println("Adding performer: ", performer.Name)
	}
}

func storeSong(song Song) {
	coll := mgoSess.DB("bmat-test").C("songs")
	err := coll.Insert(song)
	if err != nil {
		// Duplicate index, go on
		log.Println("Trying to insert duplicate Song")
	} else {
		log.Println("Adding song: ", song.Title)
		performer := Performer{song.Performer}
		storePerformer(performer)
	}
}

func storePlay(play Play) {
	coll := mgoSess.DB("bmat-test").C("plays")
	err := coll.Insert(play)
	if err != nil {
		// Duplicate index, go on
		log.Println("Trying to insert duplicate Song")
	} else {
		log.Println("Adding play: ", play.Title)
	}
}
