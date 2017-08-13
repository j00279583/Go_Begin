package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"bufio"
	"strings"
	"errors"
	"encoding/json"
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

input:
	GlobalConsoleInst.WelcomeInfo()

	ip, port, e:= GlobalConsoleInst.GetServerInformation()
	if e != nil{
		fmt.Printf("illegal inpuy")
		goto input
	}

	con, err := GlobalConsoleInst.ConnectToServer(ip, port)
	if err != nil{
		fmt.Println("Connect server failed,err",err)
		return
	}

	GlobalConsoleInst.HandleMsgPro(con)

}
func (c *console) WelcomeInfo() {

}

func (c *console) HelpInfo()  error{

	fmt.Println("input like this:  [opertype] [ID] [name]")
	fmt.Printf("operTyoe: \r\n 1 for add \r\n 2 for del \r\n 3 for uypdate \r\n 4 for query\r\n ")
	fmt.Printf("stuInfo: \r\n 1、stuID 2、stuNamer\r\n")


	return nil
}


func (c *console) GetServerInformation()  (string, int,error){

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	msg := string(data)
	msgslnce := strings.Split(msg," ")
	if len(msgslnce) != 2{
		fmt.Println("illegal param num.\r\n")
		return " ",0,errors.New("illegal param num")
	}
	ip := msgslnce[0]
	port, _ := strconv.Atoi(msgslnce[1])

	fmt.Println(ip)
	fmt.Println(port)

	return ip, port,nil
}


func (c *console) ConnectToServer(ip string, port int)  (*net.UDPConn, error){

	localAddr, err := net.ResolveUDPAddr("udp","127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("before",ip)

	ip1 := net.ParseIP(ip)

	fmt.Println("after",ip1)


	remoteAddr := net.UDPAddr{IP:   ip1,
				Port:	port,
				}

	con, err1 := net.DialUDP("udp", localAddr, &remoteAddr)
	if err1 != nil{
		fmt.Println("dailudp failed ",err1)
		return con,err1
	}


	fmt.Println("connect to server success\r\n")
	return con, nil
}


func (c *console) CheckAddr(ip net.IP, port uint32)  error{


	fmt.Println("connect to server success\r\n")
	return nil
}

func (c *console) CheckUserInput(data []byte)  (bool, bool) {

	msg := string(data)
	if len(msg) < 0 {
		fmt.Println("length illegal")
		return false,false

	}
	msgslnce := strings.Split(msg," ")

	fmt.Printf("operMsg is %v, msgLen is %d",msgslnce, len(msgslnce))

	if len(msg) == 1 && (msg == "q"|| msg == "Q"){
		return false,true
	}else if len(msg) == 1 {
		return false,false
	}

	operType, _:= strconv.Atoi(msgslnce[0])
	switch operType {

		case OperTypeAdd,OperTypeUpdate:

			if len(msgslnce) != 3{
				fmt.Println("illegal input: \r\n param is not illegal.")
				return false,false
			}

		case OperTypeDel,OperTypeQuery:
			if len(msg) != 2 {
				fmt.Println("illegal input: \r\n param is not illegal.")
				return false,false
			}

		default:
			fmt.Println("illegal opertype")
		}

	return true,false
}

func (c *console) HandleMsgPro(con net.Conn)  error{

	GlobalConsoleInst.HelpInfo()

	for{

		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()

		legalFlag, quitFlag:= GlobalConsoleInst.CheckUserInput(data)
		if legalFlag != true{
			if quitFlag != true {
				fmt.Println("input wrong")
				GlobalConsoleInst.HelpInfo()
				continue
			}
			break
		}

		err := GlobalConsoleInst.DoMsg(con, data)
		if err != nil{
			fmt.Printf("handle msg wrong err is %v",err)
		}

		GlobalConsoleInst.HelpInfo()
	}
	GlobalConsoleInst.SayBye()

	return nil
}

func (c *console) DoMsg(con net.Conn, data []byte) error {
		msg := string(data)
	        msgslnce := strings.Split(msg, " ")
		operType, _:= strconv.Atoi(msgslnce[0])

		switch operType{
		case OperTypeAdd:
			fmt.Println("Add info come in ")
			ID, _:= strconv.Atoi(msgslnce[1])

			var sendMsg = MsgData{
			OperType:operType,
			StuInfo:StuentInfo{uint32(ID),msgslnce[2]},
			}

			fmt.Println(sendMsg)

		       byteMsg, err := json.Marshal(sendMsg)
			if err != nil {
				fmt.Println(err)
				return err

			}
			nWrite, err2 := con.Write(byteMsg)
			if err2 != nil {
				fmt.Println(err2,nWrite)
			}

		case OperTypeDel:

		case OperTypeUpdate:

		case OperTypeQuery:

		default:
			fmt.Println("illegal opertype")

		}

	fmt.Println(" msg handle success\r\n")
	return nil
}


func (c *console) SayBye()  error{

	fmt.Println(" disconnect with server, goodbye \r\n")
	return nil
}


func NewConsole() *console{

	return &console{}
}

func init() {
	GlobalConsoleInst = NewConsole()
}




