package messages

type MessageType int

const (
	INIT MessageType = iota
	JOIN
	PLAY
)

var stateMap = map[MessageType]string{
	INIT: "init",
	JOIN: "join",
	PLAY: "play",
}

func (mt MessageType) String() string {
	return stateMap[mt]
}
