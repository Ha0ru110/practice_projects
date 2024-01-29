package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http"
	"time"
)

//func main() {
	http.HandleFunc("/haohaotest", CreateHandler)
	//receivedJSONDataCh = make(chan map[string]interface{}, 10)
	//go CompareData()
	err := http.ListenAndServe(":443", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(time.Now().UnixMilli(), string(b))
	}
}
