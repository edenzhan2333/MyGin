package main

import (
	"fmt"
	"ginn"
	"net/http"
)

func main() {
	r := ginn.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	fmt.Printf("%+v", r)

	r.Run(":9999")
}
