package world

import (
	"aakimov/marsgame/server"
	"log"
	"time"
)

// max moving forward per turn
const MaxMovingLength float64 = 7

// max rotation per turn in radians
const MaxRotationValue float64 = 0.5
const MaxCannonRotationValue float64 = 0.8

type World struct {
	Server    *server.Server
	players   map[string]*Player
	objects   map[string]*Object
	changeLog *ChangeLog
	timeId    int64
}

func NewWorld(server *server.Server) World {
	return World{
		Server:    server,
		players:   make(map[string]*Player),
		objects:   make(map[string]*Object),
		changeLog: NewChangeLog(),
	}
}

func makeTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func timeStampDif(t1, t2 time.Time) int64 {
	return makeTimestamp(t2) - makeTimestamp(t1)
}

func (w *World) codeRun() {
	// players are empty at start, so this block is for future
	for _, player := range w.players {
		go player.mainProgram.Run()
		go player.listen()
	}
}

func (w *World) Run() {
	ticker := time.NewTicker(200 * time.Millisecond)
	go w.sendChangelogLoop()
	w.codeRun()

	serverStartTime := time.Now()

	// endless loop here
	for t := range ticker.C {
		w.timeId = timeStampDif(serverStartTime, t)
		//log.Printf("Game tick %v\n", timeId)

		w.listenChannels()
		changeByTime := NewChangeByTime(w.timeId)
		for _, player := range w.players {
			if ch := player.run(w); ch != nil {
				changeByTime.Add(ch)
			}
		}
		if changeByTime.IsNotEmpty() {
			w.changeLog.AddToBuffer(changeByTime)
		}
	}
}

func (w *World) listenChannels() {
	select {
	case client := <-w.Server.NewClientCh:
		player := NewPlayer(client.Id, client, NewMech(), w)
		log.Printf("New player [%s] added to the game", player.id)

		w.players[player.id] = player
		go player.mainProgram.Run()
		go player.listen()
	case saveCode := <-w.Server.SaveAstCodeCh:
		player, ok := w.players[saveCode.UserId]
		if !ok {
			log.Fatalf("Save code attempt for inexistant player [%s]", saveCode.UserId)
		}
		player.saveAstCode(saveCode.SourceCode)
	case programFlowCmd := <-w.Server.ProgramFlowCh:
		player, ok := w.players[programFlowCmd.UserId]
		if !ok {
			log.Fatalf("Save code attempt for inexistant player [%s]", programFlowCmd.UserId)
		}
		player.mainProgram.operateState(programFlowCmd.FlowCmd)
	default:
		// noop
	}
}

func (w *World) sendChangelogLoop() {
	for {
		select {
		case ch := <-w.changeLog.changesByTimeCh:
			if w.changeLog.AddAndCheckSize(ch) {
				w.changeLog.Optimize()
				command := server.PackStructToCommand("worldChanges", w.changeLog.changesByTimeLog)
				for _, player := range w.players {
					player.client.SendCommand(command)
				}
				w.changeLog.changesByTimeLog = make([]*ChangeByTime, 0, ChangelogBufferSize)
			}
		default:
			// noop
		}
	}
}