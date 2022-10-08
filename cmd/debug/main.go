package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/AzraelSec/go-i3blocks/pkg/protocol"
	"github.com/AzraelSec/go-i3blocks/pkg/utils"
)

const KEY = "value"

func main() {
	state, ok := protocol.GetState()
	var text string
	if !ok {
		text = "err 1"
	} else {
		text = state[KEY]

		num, err := strconv.Atoi(state[KEY])
		if err == nil {
			state[KEY] = strconv.Itoa(num + 1)
		}
	}
	output := &protocol.I3BlocksOutput{
		FullText:  text,
		Color:     utils.RandColorStr(),
		Separator: true,
		State:     state,
	}

	json, err := json.Marshal(output)
	if err != nil {
		output.FullText = "err 2"
	}
	fmt.Println(string(json))
}
