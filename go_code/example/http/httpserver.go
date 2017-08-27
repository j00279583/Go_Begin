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

	 r.ParseForm()

	 data := r.URL.Query()
	 fmt.Println("all data", data)


         var NameExist, IdExist bool
	 var Id,Name []string

	 if Id, NameExist = r.Form["Id"]; NameExist{
		 fmt.Println("Id is ",Id[0])
	 }

	 if Name, IdExist = r.Form["Name"]; IdExist{
		 fmt.Println("Name is ",Name[0])
	 }
	 if NameExist && IdExist{
		  w.Write([]byte(fmt.Sprintf("welcome %s to come in",Name[0])))
	 }

 }