package slide

import (
	"encoding/json"
)

type (
	Message struct {
		Name `json:"name"`
		Body `json:"body"`
	}
)

func readMe() Message {
	b := []byte(`{"Name":"Ken","Body":"Hello"}`)
	var m Message
	err := json.Unmarshal(b, &m)
}

