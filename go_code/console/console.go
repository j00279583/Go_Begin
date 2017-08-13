package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"bufio"
	//"encoding/json"
	"strings"
	//"encoding/json"
)

const (
	OperTypeAdd = iota
	OperTypeDel
	OperTypeUpdate
	OperTypeQuery
)

type console struct {

}

type MsgData struct {
	OperType  int
	StuInfo   StuentInfo

}
type StuentInfo struct {
	ID  uint32
	Name string
}

type ServerAddr struct {
	ip net.IP
	port uint32
}

var GlobalConsoleInst *console

func main()  {
	fmt.Println("Welcome to console:")
	fmt.Printf("Input server IP and port:\r\n")

	GlobalConsoleInst = NewConsole()

	GlobalConsoleInst.GetServerInformation()

	GlobalConsoleInst.ConnectToServer([]byte("127.0.0.1"),8889)

	con, err := net.ListenUDP("tcp",&net.UDPAddr{})
	if err != nil {
		fmt.Println(err)

	}

	GlobalConsoleInst.HandleMsgPro(con)

}

func (c *console) HelpInfo()  error{

	fmt.Println("input like this:  [opertype] [stuInfo]")
	fmt.Printf("operTyoe: \r\n 1 for add \r\n 2 for del \r\n 3 for uypdate \r\n 4 for query\r\n ")
	fmt.Printf("stuInfo: \r\n 1、stuID 2、stuName\r\n")


	return nil
}


func (c *console) GetServerInformation()  (string, uint32){

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	msg := string(data)
	msgslnce := strings.Split(msg," ")

	ip := msgslnce[0]
	port,_ := strconv.Atoi(msgslnce[1])

	return ip , uint32(port)
}


func (c *console) ConnectToServer(ip net.IP, port uint32)  error{




	fmt.Println("connect to server success\r\n")
	return nil
}


func (c *console) CheckAddr(ip net.IP, port uint32)  error{


	fmt.Println("connect to server success\r\n")
	return nil
}

func (c *console) CheckUserInput(data []byte)  bool {

	msg := string(data)
	if len(msg) < 0 {
		fmt.Println("illegal input:")
		return false

	}
	msgslnce := strings.Split(msg," ")

	operType, _:= strconv.Atoi(msgslnce[0])
	switch operType {

		case OperTypeAdd,OperTypeUpdate:
			fmt.Println("Add info come in ")
			if len(msg) != 3{
				fmt.Println("illegal input:")
				return false
			}

		case OperTypeDel,OperTypeQuery:
			if len(msg) != 2 {
				fmt.Println("illegal input:")
				return false
			}

		default:
			fmt.Println("illegal opertype")
		}

	return true
}

func (c *console) HandleMsgPro(con net.Conn)  error{

	GlobalConsoleInst.HelpInfo()


	for{

		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()

		illegalflag := GlobalConsoleInst.CheckUserInput(data)
		if illegalflag != true{
			fmt.Println("input wrong")
			continue
		}


		flag, err := GlobalConsoleInst.DoMsg(con,data)
		if flag == true{
			break
		}

		fmt.Println(err)



	}


	GlobalConsoleInst.SayBye()

	return nil
}

func (c *console) DoMsg(con net.Conn, data []byte)  (bool ,error) {

	msg := string(data)
	msgslnce := strings.Split(msg," ")

	operType, _:= strconv.Atoi(msgslnce[0])
	if len(msg) == 1 && (msg == "q"|| msg == "Q"){
		return true,nil
	}else if len(msg) == 1 {
		return false,nil
	}else {
		switch operType{

		case OperTypeAdd:
			fmt.Println("Add info come in ")

			ID, _:= strconv.Atoi(msgslnce[1])

			var sendMsg = MsgData{
			OperType:operType,
			StuInfo:StuentInfo{uint32(ID),msgslnce[2]},
			}

			fmt.Println(sendMsg)

			//byteMsg, err := json.Marshal(sendMsg)
			//if err != nil {
			//	fmt.Println(err)
			//
			//}
			//
			//con.Write(byteMsg)

		case OperTypeDel:

		case OperTypeUpdate:

		case OperTypeQuery:

		default:
			fmt.Println("illegal opertype")

		}


		fmt.Println(msg)

	}


	fmt.Println(" msg handle success\r\n")
	return false,nil
}


func (c *console) SayBye()  error{


	fmt.Println(" disconnect with server, goodbye \r\n")
	return nil
}


func NewConsole() *console{

	return &console{}
}





