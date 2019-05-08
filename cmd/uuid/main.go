package main

import (
	"log"
	"fmt"
	"net/http"
	"flag"
	"github.com/google/uuid"
)

func main(){
	portPtr := flag.Int("port", 8000, "http port for uuid server")
	flag.Parse()
	port := *portPtr
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		id, _ := uuid.NewRandom()
		fmt.Fprintf(w, "%s", id)
	})
	log.Printf("start http server at port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
