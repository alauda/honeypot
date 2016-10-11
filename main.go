package main

import (
	"compress/gzip"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Record struct {
	Url    string
	Header map[string][]string
	Body   string
}

// Decode request body by content encoding
func decodeRequestBody(r *http.Request) (string, error) {

	ce, _ := r.Header["Content-Encoding"]

	for _, x := range ce {
		switch x {
		case "gzip":
			gr, e := gzip.NewReader(r.Body)
			defer gr.Close()
			if e != nil {
				return "", e
			}
			return readAll(gr)
		case "deflate":
			fr, e := zlib.NewReader(r.Body)
			defer fr.Close()
			if e != nil {
				return "", e
			}
			return readAll(fr)
		}
	}

	// No Content-Encoding header detected
	return readAll(r.Body)
}

// Read from reader to string
func readAll(r io.Reader) (string, error) {
	if body, err := ioutil.ReadAll(r); err != nil {
		return "", err
	} else {
		return string(body), nil
	}
}

// Handle all request.
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	// TODO should be configurable?
	w.WriteHeader(http.StatusOK)
	if status == http.StatusNotFound {
		record := Record{
			Url:    r.URL.String(),
			Header: r.Header,
		}
		if body, err := decodeRequestBody(r); err != nil {
			record.Body = fmt.Sprintf("Parse error: %s", err.Error())
		} else {
			record.Body = body
		}

		log.Println("--------------------------------------------------------------------------------")
		log.Println(record)
	}
}

// Default handler, do nothing.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r, http.StatusNotFound)
}

func main() {

	port := flag.String("port", "8080", "sets service port number")
	log_param := flag.String("log", "", "save all request to a log file")

	flag.Parse()

	// save logs to file if set.
	if *log_param != "" {
		f, err := os.OpenFile(*log_param, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			panic(fmt.Sprintf("error opening file: %s", err.Error()))
		}
		defer f.Close()

		log.SetOutput(f)

	}

	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe("0.0.0.0:"+*port, nil)
	if err != nil {
		fmt.Printf("I think something here could work, but not this")
	}
}
