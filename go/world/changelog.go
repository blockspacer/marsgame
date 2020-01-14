package world

import (
	"aakimov/marsgame/go/physics"
)

const (
	TypePlayer = "player"
	TypeObject = "object"
)

const ChangelogChannelBufferSize = 10
const ChangelogBufferSize = 8

type ChangeLog struct {
	changesByTimeCh  chan *ChangeByTime
	changesByTimeLog []*ChangeByTime
}

type ChangeByTime struct {
	TimeId          int64             `json:"timeId"`
	ChangesByObject []*ChangeByObject `json:"changesByObject"`
}

type ChangeByObject struct {
	ObjType string         `json:"objType"`
	ObjId   string         `json:"objId"`
	Pos     *physics.Point `json:"pos"`
	Angle   *float64       `json:"angle"`
	length  *float64
}

func NewChangeByTime(timeId int64) *ChangeByTime {
	return &ChangeByTime{
		TimeId:          timeId,
		ChangesByObject: make([]*ChangeByObject, 0),
	}
}

func (ch *ChangeByTime) Add(changeByObject *ChangeByObject) {
	ch.ChangesByObject = append(ch.ChangesByObject, changeByObject)
}

func (ch *ChangeByTime) IsNotEmpty() bool {
	return len(ch.ChangesByObject) > 0
}

func NewChangeLog() *ChangeLog {
	return &ChangeLog{
		changesByTimeCh:  make(chan *ChangeByTime, ChangelogChannelBufferSize),
		changesByTimeLog: make([]*ChangeByTime, 0, ChangelogBufferSize),
	}
}

func (ch *ChangeLog) AddToBuffer(changeByTime *ChangeByTime) {
	ch.changesByTimeCh <- changeByTime
}

func (ch *ChangeLog) AddAndCheckSize(changeByTime *ChangeByTime) bool {
	ch.changesByTimeLog = append(ch.changesByTimeLog, changeByTime)
	return len(ch.changesByTimeLog) >= ChangelogBufferSize
}