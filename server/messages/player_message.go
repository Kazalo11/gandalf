package messages

import (
	"encoding/json"
	"github.com/Kazalo11/gandalf/server/action"
)

type PlayerMessage struct {
	BaseMessage
	Action action.Action `json:"action"`
	Data   any           `json:"data"`
}

func parsePlayerMessage(message []byte) (PlayerMessage, error) {
	var m PlayerMessage
	err := json.Unmarshal(message, &m)
	if err != nil {
		return PlayerMessage{}, err
	}
	return m, nil
}
