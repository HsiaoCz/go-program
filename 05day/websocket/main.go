package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

// build a chat with websocket

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWsOrderBook(ws *websocket.Conn) {
	fmt.Println("new incoming connection form client to orderbook feed:", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("orderbook data --> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

// handler
func (s *Server) handleWs(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readloop(ws)
}

func (s *Server) readloop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read err:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("thank you for the msg!!!"))
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error:", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWs))
	http.Handle("/orderbookfeed", websocket.Handler(server.handleWsOrderBook))
	http.ListenAndServe(":3000", nil)
}