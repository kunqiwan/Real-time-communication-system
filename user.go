package main

import "net"

// create user sturt
type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// create new user interface
func NewUser(conn net.Conn) *User {
	//use user address as their username
	//every user should have their own channel to listen
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	//when a new user is created, the goroutine should be created to listen this user
	go user.ListenMessage()
	return user
}

// listen method
func (this *User) ListenMessage() {
	//for loop , if there is a message from channel, it will be sent to client
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))

	}
}
