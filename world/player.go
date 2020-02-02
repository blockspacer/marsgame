package world

import (
	"aakimov/marsgame/server"
	"aakimov/marslang/ast"
	"aakimov/marslang/lexer"
	"aakimov/marslang/parser"
	"log"
	"time"
)

type Player struct {
	id          int
	world       *World
	client      *server.Client
	mech        *Mech
	mainProgram *Code
	runSpeedMs  time.Duration
	outputCh    chan *MechOutputVars
	io4ClientCh chan *IO4Client
	codeSaveCh  chan *ast.StatementsBlock
	flowCh      chan ProgramState
	errorCh     chan *Error
}

func NewPlayer(id int, client *server.Client, w *World, m *Mech, runSpeedMs time.Duration) *Player {
	player := &Player{
		id:          id,
		world:       w,
		client:      client,
		mech:        m,
		mainProgram: NewCode("main", w, m, runSpeedMs),
		runSpeedMs:  runSpeedMs,
		outputCh:    make(chan *MechOutputVars, 1),
		codeSaveCh:  make(chan *ast.StatementsBlock, 1),
		io4ClientCh: make(chan *IO4Client, 1),
		flowCh:      make(chan ProgramState, 1),
		errorCh:     make(chan *Error, 10),
	}
	return player
}

func (p *Player) setBaseParams() {
	p.mech.throttle = 1
}

func (p *Player) saveAstCode(sourceCode string) {
	l := lexer.New(sourceCode)
	pr, err := parser.New(l)
	if err != nil {
		p.errorCh <- &Error{
			ErrorType: Lexing,
			Message:   err.Error(),
		}
		log.Printf("Lexing error: %s", err.Error())
		return
	}
	astProgram, err := pr.Parse()
	if err != nil {
		p.errorCh <- &Error{
			ErrorType: Parsing,
			Message:   err.Error(),
		}
		log.Printf("Parsing error: %s", err.Error())
		return
	}
	log.Println("Code parsed")
	p.codeSaveCh <- astProgram
}

func (p *Player) listen() {
	log.Printf("Player [%d] listening started", p.id)
	for {
		select {
		case codeOutputs := <-p.outputCh:
			//log.Printf("Write code runProgram result for player [%s]: mThr: %f, mrThr: %f, crThr: %f, shoot: %f",
			//	p.id, codeOutputs.MThrottle, codeOutputs.RThrottle, codeOutputs.CRThrottle, codeOutputs.Shoot)
			p.mech.Lock()
			p.mech.throttle = codeOutputs.MThrottle
			p.mech.rotateThrottle = codeOutputs.RThrottle
			p.mech.cannon.rotateThrottle = codeOutputs.CRThrottle
			if codeOutputs.Shoot != 0 {
				p.mech.cannon.shoot.state = WillShoot
				p.mech.cannon.shoot.delay = int(codeOutputs.Shoot * 1000)
			}
			p.mech.Unlock()
		case codeError := <-p.errorCh:
			p.client.PackAndSendCommand("codeError", codeError)
		case io4Client := <-p.io4ClientCh:
			p.client.PackAndSendCommand("codeInputOutput", io4Client)
		}
	}
}
