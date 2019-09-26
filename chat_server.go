package main

import (
    "fmt"
    "net"
)

var chatclients map[string]*net.Conn

//ReceiveFunc Receive Client's Message
func ReceiveFunc(conn net.Conn) {
    remoteAddr := conn.RemoteAddr().String()

    fmt.Println("Accepted Connenction ", remoteAddr)

    if _, ok := chatclients[remoteAddr]; !ok {
        //Read Loop
        chatclients[remoteAddr] = &conn

        for {
            //Read Inifinit Loop
            b := make([]byte, 1024)
            nSendByte, err := conn.Read(b)

            fmt.Println("Received Byte :", nSendByte)
            fmt.Println("Received Data : ", string(b[:]))

            if err != nil {
                fmt.Println("Read Error ", err)
                fmt.Println(remoteAddr, " Client Error : Close Connnection")
                delete(chatclients, remoteAddr)
                return
            }

            if nSendByte > 0 {
                for strIP, connValue := range chatclients {
                    if strIP != remoteAddr {
                        nSendByte, err := (*connValue).Write(b)
                        if nSendByte == 0 {
                            fmt.Println("Error ", err)
                        }
                    }
                }
            }
        }
    }

}

//AcceptFunc Accept Clients
func AcceptFunc() {

    ln, err := net.Listen("tcp", ":8081")

    if err != nil {
        //hanle error
        fmt.Println("Server Listen Error")
    }

    for {
        conn, err := ln.Accept()

        if err != nil {
            fmt.Println("Server Accept Error")
        } else {
            go ReceiveFunc(conn)
        }
    }
}

func main() {
    fmt.Println("Launching server...")
    chatclients = make(map[string]*net.Conn)

    go AcceptFunc()

    fmt.Println("If U Want Exit, then Press Enter...")
    fmt.Scanln()

    for k := range chatclients {
        (*chatclients[k]).Close()
        delete(chatclients, k)
    }

    fmt.Println("Server Exit")
    fmt.Scanln()

}
