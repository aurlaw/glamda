package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"os/exec" // "path/filepath"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func RemoteCall(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	function := vars["function"]
	fmt.Fprintln(w, "RemoteCall:", function)

	cmdArgs := []string{""}
	var (
		out []byte
		err error
	)
	if out, err = exec.Command(function, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(w, "There was an error running go command: ", err)
	}
    result := fmt.Sprintf("%s", out)
	fmt.Fprintf(w, result)

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	//E:\\Temp
	cmdName := "go"
	cmdArgs := []string{"build", header.Filename}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(w, "There was an error running go command: ", err)
	}
	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)
}
