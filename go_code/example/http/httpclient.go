package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url = "http://127.0.0.1:8080/con"

type serverInfo struct {
	Ip   string
	Port string
	Dir  string
}
type clientInfo struct {

	Name string
	ID int
}

func main()  {

	fmt.Println("client con:")
	client := &http.Client{}

	rsp ,err := client.Get(url)
	if err != nil{
		fmt.Println("client Get err.",err)
		return
	}

	defer  rsp.Body.Close()

	data, err1 := ioutil.ReadAll(rsp.Body)
	if err1 == nil {
		fmt.Println(string(data))
	}

	fmt.Println("client over", err1)

}

func (s *serverInfo)GetServerInfo()  {

	return
}