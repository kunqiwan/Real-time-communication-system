package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//online user list - use map to store
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//Message Channel
	Message chan string
}

// create one server interface
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//user is online - tell other user
	user := NewUser(conn)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

}

// start server interface
func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		//nobody calls, it would go to err condition,then continue
		if err != nil {
			fmt.Println()
			continue
		}
		//else,do handler
		go this.Handler(conn)
	}
}
