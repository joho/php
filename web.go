package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type onlyFilesFileSystem struct {
	fs http.FileSystem
}

func (fs onlyFilesFileSystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return fileOnlyReaddirFile{f}, nil
}

type fileOnlyReaddirFile struct {
	http.File
}

func (f fileOnlyReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "5000"
	}
	listenOn := fmt.Sprintf(":%v", portNumber)

	fs := onlyFilesFileSystem{http.Dir("public")}
	http.Handle("/", http.FileServer(fs))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The homepage")
	})

	log.Fatal(http.ListenAndServe(listenOn, nil))
}
