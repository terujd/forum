package server

import (
	"fmt"
	//"log"
	"net/http"
)

var Port = "1337"

func StartServer() {
	http.Handle("/css/", http.StripPrefix(http.FileServer(http.Dir("./front-end/css/main.css"))))
	http.HandleFunc("/", StartPage)
	fmt.Println("Server started on port", Port)

	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		panic(err)
		//log.Fatal(err)
	}
}
