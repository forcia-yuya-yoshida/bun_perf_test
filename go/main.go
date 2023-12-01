package main

import (
	"log"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello World!!!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	x := 0
	for i := 0; i < 100000000; i++ {
		x ++
	}
	_, err := w.Write([]byte(strconv.Itoa(x)))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/calc", calcHandler)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}