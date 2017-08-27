package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	//"net/url"
)

type ServerInfo struct {
	Ip   string
	Port string
	Dir  string
}
type ClientInfo struct {
	Name string
	ID string
}

type ConfigData struct {
	 ServerData ServerInfo
	 ClientData ClientInfo
}

func main()  {

	fmt.Println("client con:")
	client := &http.Client{}

	serveInst, clientInst, err := conInst.ServerData.GetServerInfoFromConf()
	if err != nil{
		fmt.Println(err)
		return
	}

	url := serveInst.constructUrl()
	urlpara := clientInst.constructUrl(&url)

	req, e := http.NewRequest("Get",urlpara, nil)
	if e !=nil{
		fmt.Println(e)
	}

	//v := req.Form
	//v.Set("Name",clientInst.Name)
	//v.Set("Id",clientInst.ID)

	rsp ,err1 := client.Do(req)
	if err1 != nil{
		fmt.Println("client Get err.",err1)
		return
	}

	defer  rsp.Body.Close()
	data, err2 := ioutil.ReadAll(rsp.Body)
	if err2 == nil {
		fmt.Println(string(data))
	}

	fmt.Println("client over", err2)

}

var conInst = NewConfInfo()

func NewConfInfo() *ConfigData{
	return &ConfigData{}
}


func (s *ServerInfo)GetServerInfoFromConf() (*ServerInfo, *ClientInfo,error){
	conf := &ConfigData{}
	data, err := ioutil.ReadFile("conf/info.json")
	if err != nil{
		fmt.Println("readdir failed, err is ",err)
		return nil,nil,err
	}

	err1 := json.Unmarshal(data,conf)
	if err1 != nil {
		fmt.Println("unmarshal failed, err is ",err1)
		return nil, nil,err1
	}

	fmt.Println("Get info is ",conf)
	return &conf.ServerData, &conf.ClientData,nil

}

func (s *ServerInfo)constructUrl() string {

	url := "http://" + s.Ip + ":" + s.Port + "/" + s.Dir

	fmt.Println("url is",url)
	return url
}

func (s *ClientInfo)constructUrl(url *string) string {

	//*url =*url + "?" + "Name" + "=" + s.Name
	*url =*url + "?" + "Name" + "=" + s.Name + "&"+ "Id" + "=" + s.ID

	fmt.Println("url is", *url)
	return *url
}

//
//func (s *ServerInfo)ServerInfoSetServerInfo(ip, port, addr string)  {
//	s.Dir =addr
//	s.Port = port
//	s.Ip = ip
//
//	return
//}
//
//func (s *ServerInfo)GetServerIp()  string{
//
//	return s.Ip
//}
//
//func (s *ServerInfo)GetServerPort()  string{
//
//	return s.Port
//}
//func (s *ServerInfo)GetServerIAddr() string{
//
//	return s.Dir
//}

