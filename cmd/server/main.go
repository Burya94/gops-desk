package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received / request")
	io.WriteString(w, "Server is running.\n")
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receiver /info request")
	io.WriteString(w, "This is a test golang server.\n")
}

func Run() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/info", getInfo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR - failed to start server")
		os.Exit(1)
	}
}
