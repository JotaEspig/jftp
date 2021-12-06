package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Server is a struct that contains information about a net.Listener
// 	Args:
// 		Host (string): The ips it is listening to. Example: 0.0.0.0, 127.0.0.1, etc.
//		Port (int): The port it is listening on. Example: 5001.
// 		serverListener (net.Listener): It is the listener that uses the previously arguments.
type Server struct {
	Host           string
	Port           int
	serverListener net.Listener
}

// It is a constructor function, that creates a Server struct
//	Args:
//		host (string): The IPs it will listening to
//		port (int): The port it will listening on
//	Returns:
//		s (Server)
func CreateServer(host string, port int) Server {
	var s Server
	s.Host = host
	s.Port = port
	s.serverListener = nil
	return s
}

// Sets the net.Listener in the struct
// 	Returns:
//		err (error)
func (s *Server) listen() error {
	var err error
	s.serverListener, err = net.Listen("tcp4", s.Host+":"+fmt.Sprint(s.Port))
	return err
}

// Accepts a connection
// 	Returns:
//		conn (net.Conn), err (error)
func (s *Server) accept() (net.Conn, error) {
	conn, err := s.serverListener.Accept()
	return conn, err
}

// Handles a connection
//	Args:
//		conn (net.Conn): Connection with the client
func (s *Server) handleConnection(conn net.Conn) {
	log.Printf("Connected with %v\n", conn.LocalAddr().String())
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("%v: %v\n", conn.LocalAddr().String(), err)
			return
		}
		log.Printf("%v sent: %v", conn.LocalAddr().String(), string(data))
		conn.Write([]byte(data))
	}
}

// Starts the server.
// 	Returns:
//		err (error): An error that occurred in other fucntions
func (s *Server) Start() error {
	err := s.listen()
	if err != nil {
		return nil
	}
	for {
		conn, err := s.accept()
		if err == nil {
			go s.handleConnection(conn)
		}

	}
}
