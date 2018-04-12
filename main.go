package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

const DefaultHttpPort = "8080"

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
	//log.Printf("%s %s [%s]", r.Method, r.RequestURI, r.RemoteAddr)
	log.Printf("%s %s [%s]", r.Method, r.URL.EscapedPath(), r.RemoteAddr)
}

func main() {
	port := func() string {
		if val, ok := os.LookupEnv("HTTP_PORT"); ok {
			return val
		}
		return DefaultHttpPort
	}()
		
	http.HandleFunc("/", helloworld)
	log.Println("serving http on port", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
