package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Time struct {
	Time     string `json: "time,omitemty"`
	NanoTime string `json: "nanoTime,omitemty"`
}

func main() {

	log.Print("Begin working service time-service")

	http.HandleFunc("/time", serveTime)
	http.HandleFunc("/nanoTime", serveNanoTime)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func serveTime(w http.ResponseWriter, r *http.Request) {
	log.Print("serveTime() - start")
	var serverTime Time
	serverTime.Time = time.Now().String()
	json.NewEncoder(w).Encode(serverTime)
}

func serveNanoTime(w http.ResponseWriter, r *http.Request) {
	log.Print("serveNanoTime() - start")
	var nanoTime Time
	nanoTime.NanoTime = strconv.FormatInt(time.Now().UnixNano(), 10)
	json.NewEncoder(w).Encode(nanoTime)
}
