package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// init db as global handler
var mgoSess, mgoErr = mgo.Dial("localhost")

func main() {
	if mgoErr != nil {
		panic(mgoErr)
	}
	mgoSess.SetMode(mgo.Monotonic, true)
	defer mgoSess.Close()

	// set time to UTC due to mongo conversion
	time.Local = time.UTC

	initDB()

	router := mux.NewRouter()
	router.HandleFunc("/add_channel", handleChannel).Methods("POST")
	router.HandleFunc("/add_performer", handlePerformer).Methods("POST")
	router.HandleFunc("/add_song", handleSong).Methods("POST")
	router.HandleFunc("/add_play", handlePlay).Methods("POST")
	router.HandleFunc("/get_song_plays", handleSongPlays).Methods("GET")
	router.HandleFunc("/get_channel_plays", handleChannelPlays).Methods("GET")
	router.HandleFunc("/get_top", handleTopPlays).Methods("GET")

	log.Println("Starting API at port 5000...")
	http.ListenAndServe(":5000", router)
}

func initDB() {
	// drop database
	err := mgoSess.DB("bmat-test").DropDatabase()
	if err != nil {
		panic(err)
	}

	// create channels
	c := mgoSess.DB("bmat-test").C("channels")
	index := mgo.Index{
		Key:        []string{"channel"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// create performers
	c = mgoSess.DB("bmat-test").C("performers")
	index = mgo.Index{
		Key:        []string{"performer"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// create songs
	c = mgoSess.DB("bmat-test").C("songs")
	index = mgo.Index{
		Key:        []string{"performer", "title"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// create songs
	c = mgoSess.DB("bmat-test").C("plays")
	index = mgo.Index{
		Key:        []string{"channel", "start"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
