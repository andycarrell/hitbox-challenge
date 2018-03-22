package main

import (
	"bytes"
	"fmt"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

func writeImage(w http.ResponseWriter, img draw.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	writeImage(w, getImage(increment(r.URL.Path)))
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	writeImage(w, getImage(delete(r.URL.Path)))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		handleDelete(w, r)
		return
	}
	handleCount(w, r)
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "pong") })
	http.HandleFunc("/counter/", handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
