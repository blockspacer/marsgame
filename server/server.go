package server

import (
	"log"
	"strconv"
)

type Server struct {
	clients         map[string]*Client
	connectClientCh chan *Client
	leaveClientCh   chan *Client
	doneCh          chan bool
	errCh           chan error
	NewClientCh     chan *Client
	SaveAstCodeCh   chan *SaveAstCode
	ProgramFlowCh   chan *ProgramFlow
}

func NewServer() *Server {
	return &Server{
		clients:         make(map[string]*Client),
		connectClientCh: make(chan *Client),
		leaveClientCh:   make(chan *Client),
		NewClientCh:     make(chan *Client),
		doneCh:          make(chan bool),
		errCh:           make(chan error),
		SaveAstCodeCh:   make(chan *SaveAstCode),
		ProgramFlowCh:   make(chan *ProgramFlow),
	}
}

func (s *Server) ListenClients() {
	log.Println("Server start listening")

	for {
		select {
		case c := <-s.connectClientCh:
			log.Printf("Client [%s] registered!", c.Id)
			s.clients[c.Id] = c

		case c := <-s.leaveClientCh:
			log.Printf("Client [%s] unconnected", c.Id)
			delete(s.clients, c.Id)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}

func (s *Server) connectClient(client *Client) {
	s.connectClientCh <- client
	s.NewClientCh <- client
}

type SaveAstCode struct {
	UserId     string
	SourceCode string
}

type ProgramFlowType int

const (
	StopProgram ProgramFlowType = iota
	StartProgram
)

type ProgramFlow struct {
	UserId  string          `json:"userId"`
	FlowCmd ProgramFlowType `json:"flowCmd"`
}

func (s *Server) HandleCommand(client *Client, command *Command) {
	switch command.Type {
	case "saveCode":
		s.SaveAstCodeCh <- &SaveAstCode{
			UserId:     client.Id,
			SourceCode: command.Payload,
		}
	case "programFlow":
		flowAsInt, _ := strconv.Atoi(command.Payload)
		s.ProgramFlowCh <- &ProgramFlow{
			UserId:  client.Id,
			FlowCmd: ProgramFlowType(flowAsInt),
		}
	default:
		log.Printf("Unknown command %s", command.ToSting())
	}
}
