package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func execute(args string) string {
	s, err := RunString(args)
	if err != nil {
		panic(err)
	}
	return s
}

func run(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var result string
	var status int
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var parameters map[string]string
	json.Unmarshal(body, &parameters)
	if runtime.GOOS != "linux" {
		fmt.Println(`Can't Execute with non linux Machine`)
		status = http.StatusForbidden
	} else {
		result = execute(parameters["args"])
		status = http.StatusOK
	}

	log.Printf("\t| %d \t| %d ms \t| run \t| %s", status, int64(time.Since(start)/time.Millisecond), parameters["args"])
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}

func main() {
	corsOrigin := cors.AllowAll()
	router := mux.NewRouter()
	router.HandleFunc("/run", run).Methods("POST")
	http.Handle("/", corsOrigin.Handler(router))
	log.Println("services run at port :9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
