package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var port = os.Getenv("PORT")
var authorName = os.Getenv("AUTHOR_NAME")

var templates = template.Must(template.ParseFiles("index.html"))

func makeIndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		board, err := ioutil.ReadFile("board.map")
		if err != nil {
			log.Printf("%s ERROR: %s\n", r.RemoteAddr, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		vars := struct {
			Board string
			Name  string
		}{string(board), authorName}

		w.Header().Add("Content-Type", "text/html; charset=utf8")
		err = templates.ExecuteTemplate(w, "index.html", vars)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func makeStaticFileHandler(filename, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Printf("%s ERROR: %s\n", r.RemoteAddr, err.Error())
			http.NotFound(w, r)
			return
		}
		w.Header().Add("Content-Type", contentType)
		w.Write(body)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", makeIndexHandler())
	http.HandleFunc("/style.css", makeStaticFileHandler("./style.css", "text/css"))
	http.HandleFunc("/clarity.js", makeStaticFileHandler("./clarity.js", "application/javascript"))

	log.Printf("Server running on %s...\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, logRequest(http.DefaultServeMux)))
}
