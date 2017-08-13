package main

import (

	"net"
	"fmt"
	"encoding/json"

)

type MsgData struct {
	OperType  int
	StuInfo   StuentInfo

}
type StuentInfo struct {
	ID  uint32
	Name string
}

type MsgDataManager struct {
	MsgDataInfoList []StuentInfo
}

const (
	OperTypeAdd = iota
	OperTypeDel
	OperTypeUpdate
	OperTypeQuery
)


var globalMsgDataManager *MsgDataManager

func main() {
	fmt.Println("server running")

	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println(err)
	}

	l, err1 := net.ListenUDP("udp", localAddr)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer l.Close()

	fmt.Println("listen udp come in")

	var readMsg = make([]byte, 2048)
	for {
		fmt.Println("begin read")
		n, remoteAddr,err2 := l.ReadFromUDP(readMsg)
		if err2 != nil {
			fmt.Println(err2)
		}

		msg := make([]byte, n)
		copy(msg, readMsg)

		fmt.Println("receive msg number:",n)

		var Msgdata = MsgData{}
		err3 := json.Unmarshal(msg, &Msgdata)
		if err3 != nil{
			fmt.Println(err3)
		}

		msgProc(&Msgdata, globalMsgDataManager)

		l.WriteToUDP([]byte("hello,client"), remoteAddr)

	}
}

func msgProc(msgData *MsgData,m *MsgDataManager) {

	switch msgData.OperType{

		case OperTypeAdd:


		fmt.Println("recevie msgInfo :", msgData)

		m.MsgDataInfoList = append(m.MsgDataInfoList, msgData.StuInfo)


		case OperTypeDel:

		case OperTypeUpdate:

		case OperTypeQuery:

		default:
			fmt.Println("illegal opertype")

		}

}


func NewMsgDataManager()  *MsgDataManager{

	return &MsgDataManager{make([]StuentInfo,0)}
}

func init() {

	globalMsgDataManager = NewMsgDataManager()

}