package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	// Input form
	r.ParseMultipartForm(10 << 20)

	// input filer
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("StatusInternalServerError")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Upload File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("Header: %+v\n", handler.Header)

	// Write
	tempFile, err := ioutil.TempFile("upload-image", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write((fileBytes))

	// return
	fmt.Fprintf(w, "Succes Upload File\n")
}

func setUp() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":5000", nil)
}

func main() {
	fmt.Println("upload")
	setUp()
}
