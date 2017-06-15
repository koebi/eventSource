package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", ServeIndex)
	http.HandleFunc("/events", ServeEvents)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func ServeIndex(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func ServeEvents(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/event-stream")

	for {
		//sending a default event with some data
		time.Sleep(time.Second)
		if _, err := fmt.Fprintf(res, "data: Hello-World-Message-Event\n\n"); err != nil {
			log.Printf("could not write: %v", err)
			return
		}
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
		// sending an "add"-event
		time.Sleep(time.Second)
		if _, err := fmt.Fprintf(res, "event:add\ndata: Hello-World-Add-Event\n\n"); err != nil {
			log.Printf("could not write: %v", err)
			return
		}
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}

		//sending an "alert"-event
		time.Sleep(time.Second)
		if _, err := fmt.Fprintf(res, "event:alert\ndata: Alert!\n\n"); err != nil {
			log.Printf("could not write: %v", err)
			return
		}
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
	}
}
