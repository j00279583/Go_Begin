package main

import (
	"fmt"
	"net/http"
)

func main()  {

	fmt.Println("server listening:")

	http.HandleFunc("/con",conncet)
	http.ListenAndServe(":8080",nil)

}

 func conncet(w http.ResponseWriter, r *http.Request) {

	 w.Write([]byte("welcome to come in\r\n"))
 }