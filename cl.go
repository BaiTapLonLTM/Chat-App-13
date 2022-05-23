package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log" //cac thu vien can su dung
	"net"
	"os"
	"strings"
)

func onMessage(conn net.Conn) { //nhan tin nhan tu server
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Print(msg)
		fmt.Print("Chat:")
	}
}

func main() {
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key") //tai 2 chung chi tu client da tao
	if err != nil {
		log.Fatal(err)
	}
	hostName := "127.0.0.1" //local host
	portNum := "6600"       // cong ket noi den server
	log.Printf("Connecting to %s\n", hostName)
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	connection, err := tls.Dial("tcp", hostName+":"+portNum, &config) //khoi tao ket noi den server
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Client is online now gogogo!!!\n")
	fmt.Print(" What your name:")
	nameReader := bufio.NewReader(os.Stdin)     //dinh danh, nhan ten dai dien cho client
	nameInput, _ := nameReader.ReadString('\n') //truyen ki tu

	fmt.Println("********** MESSAGES **********")
	go onMessage(connection) //goi ham connection de nhan tin nhan tu server tren mot goroutines
	fmt.Print("Chat:")
	for {

		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.TrimSpace(msg) == "STOP" { //lenh thoat client
			fmt.Println("Client off now!!! ")
			fmt.Println("Bye!!!")
			break
		}
		nameInput = strings.Replace(nameInput, "\r\n", "", -1) // chuan hoa lai chuoi vua nhap tu bufio
		msg = strings.Replace(msg, "\r\n", "", -1)
		msg = fmt.Sprintf("%s : %s\n", nameInput, msg)

		connection.Write([]byte(msg)) //gui giu lieu den server
	}

	connection.Close()
}
