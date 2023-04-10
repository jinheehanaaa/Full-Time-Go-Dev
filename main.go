package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Printf("Server starting\n")
	s.loop() // block
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) quit() {
	//close(s.quitch)
	s.quitch <- struct{}{}
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			fmt.Println("quiting server")
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		}
	}
	fmt.Println("server is shutting down gracefully")
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("We received a message: ", msg)
}

func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()

	// write code here -> will not work

}
